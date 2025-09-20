package main

import (
	"net/http"

	"github.com/AttendEase-ERP/erp-backend/internal/auth"
	"github.com/AttendEase-ERP/erp-backend/internal/router"
	"github.com/AttendEase-ERP/erp-backend/pkg/config"
)

func main() {
	// initialize resources
	cfg := config.LoadConfig()
	auth.InitClerk(cfg.ClerkSecretKey)
	r := router.NewRouter()

	http.ListenAndServe(cfg.Port, r)
}
