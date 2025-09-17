package main

import (
	"fmt"

	"github.com/AttendEase-ERP/erp-backend/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println(cfg.DatabaseURL)
}
