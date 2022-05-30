package chat

import (
	"context"
	"errors"

	"app/src/core/dep"
	"app/src/core/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type Service struct {
	db *pgxpool.Pool
}

func NewChatService() *Service {
	c := new(Service)
	c.db = dep.DB
	return c
}

func (s Service) GetChat(ctx context.Context, id uint64) (*model.Chat, error) {
	var chat model.Chat
	sql := "select * from chat where id = $1"
	opt := pgx.TxOptions{
		IsoLevel:       pgx.ReadUncommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: pgx.NotDeferrable,
	}

	err := s.db.BeginTxFunc(ctx, opt, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, sql, id)
		return row.Scan(&chat.ID, &chat.Name)
	})

	if err != nil {
		log.Error().Caller().Err(err).Uint64("id", id).Send()
		switch {
		default:
			return nil, err
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrIDNotFound
		}
	}

	return &chat, nil
}

func (s Service) GetChatMessages(ctx context.Context, id, page, size uint64) ([]model.ChatMessage, error) {
	errLog := log.Error().Caller().Uint64("chat", id).Uint64("page", page).Uint64("size", size)
	sql := "select * from chat_message where chat = $1 limit $2 offset $3"
	opt := pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: pgx.NotDeferrable,
	}
	res := make([]model.ChatMessage, size)
	err := s.db.BeginTxFunc(ctx, opt, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, id, page*size-size, size)
		if err != nil {
			errLog.Err(err).Msg("cannot execute query")
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var msg model.ChatMessage
			err := rows.Scan(&msg.ID, &msg.Text, &msg.Chat, &msg.Sender, &msg.Timestamp)
			if err != nil {
				errLog.Err(err).Msg("cannot scan row")
				return err
			}

			res = append(res, msg)
		}

		return nil
	})

	if err != nil {
		errLog.Err(err).Send()
		switch {
		default:
			return nil, err

		case errors.Is(err, pgx.ErrNoRows):
			return res, nil
		}
	}

	return res, err
}
