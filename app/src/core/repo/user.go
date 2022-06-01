package repo

import (
	"context"
	"database/sql"
	"strconv"

	"app/src/core/model"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: conn,
	}
}

func (r *UserRepository) Insert(ctx context.Context, user *model.User) (*model.User, error) {
	txx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		log.Error().Err(err).Msg("UserRepository:Insert cannot begin transaction")
		return nil, err
	}
	qry := `INSERT INTO "chat"."user" ("email", "username", "password") VALUES ($1, $2, $3) RETURNING "id"`
	row := txx.QueryRowx(qry, user.Email, user.Username, user.Password)
	defer func() { _ = txx.Rollback() }()

	if err := row.Scan(&user.ID); err != nil {
		log.Error().Err(err).Str("query", qry).Msg("UserRepository:Insert cannot perform operation")
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		log.Error().Err(err).Str("query", qry).Msg("UserRepository:Insert cannot commit transaction")
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) SelectWhereId(ctx context.Context, userID uint64) (*model.User, error) {
	txx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	})
	if err != nil {
		log.Error().Err(err).Msg("UserRepository:SelectWhereId cannot begin transaction")
	}

	var dest model.User
	qry := `SELECT * FROM "chat"."user" WHERE "id" = $1`
	row := r.db.QueryRowx(qry, userID)
	defer func() { _ = txx.Rollback() }()

	if err := row.StructScan(&dest); err != nil {
		log.Error().
			Err(err).
			Str("id", strconv.FormatUint(userID, 10)).
			Str("query", qry).
			Msg("UserRepository:SelectWhereId cannot perform operation")
		return nil, err
	}

	return &dest, nil
}
