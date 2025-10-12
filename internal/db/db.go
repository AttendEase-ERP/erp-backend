package db

import (
	"context"
	"time"

	"github.com/AttendEase-ERP/erp-backend/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB(connString string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	Pool, err = pgxpool.New(ctx, connString)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to initialize DB pool")
	}

	if err = Pool.Ping(ctx); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to ping DB — check your Neon branch or credentials")
	}

	logger.Log.Info().Msg("Successfully connected to Neon database")
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
		logger.Log.Info().Msg("Database connection pool closed")
	}
}
