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

type UserRepository struct {
	db  *sqlx.DB
	log echo.Logger
}

func NewUserRepository(db *sqlx.DB, log echo.Logger) *UserRepository {
	return &UserRepository{db, log}
}

func (ur *UserRepository) Insert(ctx context.Context, user *model.User) (*model.User, error) {
	tx, err := ur.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		ur.log.Error(err)
		return nil, err
	}

	param := []any{user.Email, user.Username, user.Password}
	query := `INSERT INTO "chat"."user" ("email", "username", "password") VALUES ($1, $2, $3) RETURNING "id"`

	row := tx.QueryRowx(query, param...)
	defer func() { _ = tx.Rollback() }()

	if err := row.Scan(&user.ID); err != nil {
		ur.log.Errorj(log.JSON{"error": err, "query": query, "param": param})
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		ur.log.Errorj(log.JSON{"error": err, "query": query, "param": param})
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) SelectWhereId(ctx context.Context, userId uint64) (*model.User, error) {
	tx, err := ur.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: true})
	if err != nil {
		ur.log.Error(err)
		return nil, err
	}

	param := []any{userId}
	query := `SELECT * FROM "chat"."user" WHERE "id" = $1`
	row := tx.QueryRowx(query, param...)
	defer func() { _ = tx.Rollback() }()

	var dest model.User
	if err := row.StructScan(&dest); err != nil {
		ur.log.Errorj(log.JSON{"error": err, "query": query, "param": param})
		return nil, helper.ParseError(err)
	}

	return &dest, nil
}

func (ur *UserRepository) SelectWhereUsername(ctx context.Context, username string) (*model.User, error) {
	tx, err := ur.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: true})
	if err != nil {
		ur.log.Error(err)
		return nil, err
	}

	param := []any{username}
	query := `SELECT * FROM "chat"."user" WHERE "username" = $1`
	row := tx.QueryRowx(query, param...)
	defer func() { _ = tx.Rollback() }()

	var dest model.User
	if err := row.StructScan(&dest); err != nil {
		ur.log.Errorj(log.JSON{"error": err, "query": query, "param": param})
		return nil, helper.ParseError(err)
	}

	return &dest, nil
}
