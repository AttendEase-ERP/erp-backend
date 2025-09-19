package main

import (
	"fmt"
	"net/http"

	"github.com/AttendEase-ERP/erp-backend/internal/auth"
	"github.com/AttendEase-ERP/erp-backend/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	// initialize resources
	cfg := config.LoadConfig()
	auth.InitClerk(cfg.ClerkSecretKey)
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		// MaxAge:           300,
	}))
	r.Use(auth.RequireAuth)

	r.Get("/auth-check", authHandlerFunc)

	http.ListenAndServe(cfg.Port, r)
}

func authHandlerFunc(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.GetUserEmail(r.Context())
	if err != nil {
		fmt.Println(claims)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Println(claims)
	w.Write([]byte(claims))
}
