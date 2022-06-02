package repo

import (
	"context"
	"database/sql"
	"github.com/alexeymyakinin/ruck/app/src/core/helper"
	"github.com/alexeymyakinin/ruck/app/src/core/model"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ChatRepository struct {
	db  *sqlx.DB
	log echo.Logger
}

func NewChatRepository(db *sqlx.DB, log echo.Logger) *ChatRepository {
	return &ChatRepository{db, log}
}

func (cr *ChatRepository) Insert(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	tx, err := cr.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		cr.log.Error(err)
		return nil, err
	}

	param := []any{chat.Name}
	query := `INSERT INTO "chat"."chat" ("name") VALUES ($1) RETURNING "id"`

	row := tx.QueryRow(query, param...)
	defer func() { _ = tx.Rollback() }()

	if err := row.Scan(chat.ID); err != nil {
		cr.log.Errorj(log.JSON{"error": err, "query": query, "param": param})
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		cr.log.Errorj(log.JSON{"error": err, "query": query, "param": param})
		return nil, err
	}

	return chat, nil
}

func (cr *ChatRepository) SelectWhereId(ctx context.Context, chatId uint64) (*model.Chat, error) {
	tx, err := cr.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: true})
	if err != nil {
		cr.log.Error(err)
		return nil, err
	}

	param := []any{chatId}
	query := `SELECT * FROM "chat"."chat" WHERE "id" = $1`
	row := tx.QueryRowx(query, param...)
	defer func() { _ = tx.Rollback() }()

	var dest model.Chat
	if err := row.StructScan(&dest); err != nil {
		cr.log.Errorj(log.JSON{"error": err, "query": query, "param": param})
		return nil, helper.ParseError(err)
	}

	return &dest, nil
}
