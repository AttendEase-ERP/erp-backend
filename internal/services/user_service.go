package services

import (
	"context"

	"github.com/AttendEase-ERP/erp-backend/internal/models"
	"github.com/AttendEase-ERP/erp-backend/internal/repository"
)

func GetUserByEmail(ctx context.Context, email string) (*models.UserDetails, error) {
	return repository.GetUserDetailsByEmail(ctx, email)
}
