package repo

import (
	"context"
	"database/sql"

	"app/src/core/model"
	"github.com/jmoiron/sqlx"

	"github.com/rs/zerolog/log"
)

type ChatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(pool *sqlx.DB) *ChatRepository {
	return &ChatRepository{
		db: pool,
	}
}

func (r *ChatRepository) Insert(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	txx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		log.Error().Err(err).Msg("ChatRepository:Insert cannot begin transaction")
		return nil, err
	}

	qry := `INSERT INTO "chat"."chat" ("name") VALUES ($1) RETURNING "id"`
	row := txx.QueryRow(qry)
	defer func() { _ = txx.Rollback() }()

	if err := row.Scan(chat.ID); err != nil {
		log.Error().Err(err).Str("query", qry).Msg("ChatRepository:Insert cannot scan")
		return nil, err
	}

	return chat, nil
}

func (r *ChatRepository) SelectWhereId(ctx context.Context, id uint64) (*model.Chat, error) {
	txx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		log.Error().Err(err).Msg("ChatRepository:SelectWhereId cannot begin transaction")
		return nil, err
	}

	var dest model.Chat
	qry := `SELECT * FROM "chat"."chat" WHERE "id" = $1`
	row := txx.QueryRowx(qry, id)
	defer func() { _ = txx.Rollback() }()

	if err := row.StructScan(&dest); err != nil {
		log.Error().Err(err).Str("query", qry).Msg("ChatRepository:SelectWhereId cannot scan")
		return nil, err
	}

	return &dest, nil
}
