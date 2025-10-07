package repository

import (
	"context"

	"github.com/AttendEase-ERP/erp-backend/internal/db"
	"github.com/AttendEase-ERP/erp-backend/internal/models"
)

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	row := db.Pool.QueryRow(ctx, "SELECT id, email, role FROM Users WHERE email=$1", email)
	var u models.User
	if err := row.Scan(&u.ID, &u.Email, &u.Role); err != nil {
		return nil, err
	}
	return &u, nil
}
