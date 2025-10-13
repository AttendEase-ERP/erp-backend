package repository

import (
	"context"

	"github.com/AttendEase-ERP/erp-backend/internal/db"
	"github.com/AttendEase-ERP/erp-backend/internal/models"
	"github.com/AttendEase-ERP/erp-backend/pkg/logger"
)

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	row := db.Pool.QueryRow(ctx, `SELECT id, email, role FROM "Users" WHERE "email" ILIKE $1`, "%"+email+"%")

	var u models.User
	if err := row.Scan(&u.ID, &u.Email, &u.Role); err != nil {
		logger.Log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return &u, nil
}
