package dep

import (
	"context"

	"app/src/core/cfg"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

var DB *pgxpool.Pool

func init() {
	db, err := pgxpool.Connect(context.Background(), cfg.Config.DBConnectionURL)
	if err != nil {
		log.Fatal().Str("dbConnectionURL", cfg.Config.DBConnectionURL).Err(err).Msg("cannot connect to db")
	}

	DB = db
}
