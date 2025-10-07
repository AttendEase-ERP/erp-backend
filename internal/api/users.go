package api

import (
	"encoding/json"
	"net/http"

	"github.com/AttendEase-ERP/erp-backend/internal/services"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	user, err := services.GetUserByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
