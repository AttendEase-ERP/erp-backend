package main

import (
	"net/http"

	"github.com/AttendEase-ERP/erp-backend/internal/auth"
	"github.com/AttendEase-ERP/erp-backend/internal/db"
	"github.com/AttendEase-ERP/erp-backend/internal/router"
	"github.com/AttendEase-ERP/erp-backend/pkg/config"
	"github.com/AttendEase-ERP/erp-backend/pkg/logger"
)

func main() {
	// Initialize logger
	logger.Init()

	// load env variables
	cfg := config.LoadConfig()
	auth.InitClerk(cfg.ClerkSecretKey)

	// connect to db
	db.InitDB(cfg.DatabaseURL)

	// setup router
	r := router.NewRouter()

	// start server
	logger.Log.Info().Msgf("server running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		logger.Log.Fatal().Err(err).Msg("server crashed")
	}
}
