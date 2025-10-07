package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AttendEase-ERP/erp-backend/internal/services"
)

func GetUserRole(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"role": user.Role,
	})
}
