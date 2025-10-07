package db

import (
	"context"

	"github.com/AttendEase-ERP/erp-backend/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB(connString string) {
	var err error
	Pool, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("unable to connect to database")
	}
}
