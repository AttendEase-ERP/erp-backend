package handlers

import (
	"net/http"

	"github.com/AttendEase-ERP/erp-backend/internal/auth"
)

func FetchEmail(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.GetUserEmail(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Write([]byte(claims))
}
