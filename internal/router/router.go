package router

import (
	"github.com/AttendEase-ERP/erp-backend/internal/auth"
	"github.com/AttendEase-ERP/erp-backend/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter() *chi.Mux {
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

	r.Get("/fetch-email", handlers.FetchEmail)

	return r
}
