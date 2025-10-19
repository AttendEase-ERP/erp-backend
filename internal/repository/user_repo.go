package repository

import (
	"context"
	"fmt"

	"github.com/AttendEase-ERP/erp-backend/internal/db"
	"github.com/AttendEase-ERP/erp-backend/internal/models"
	"github.com/AttendEase-ERP/erp-backend/pkg/logger"
)

func GetUserDetailsByEmail(ctx context.Context, email string) (*models.UserDetails, error) {
	var role string
	err := db.Pool.QueryRow(
		ctx,
		`SELECT role FROM "Users" WHERE "email" ILIKE $1`,
		"%"+email+"%",
	).Scan(&role)
	if err != nil {
		logger.Log.Error().Err(err).Msg("error fetching user role")
		return nil, err
	}

	u := &models.UserDetails{
		Email: email,
		Role:  models.UserRole{Role: role},
	}

	switch role {
	case "teacher":
		query := `
		SELECT 
			t.name, 
			s.section_name, 
			sub.subject_name,
			sub.semester,
			course.course_name,
			course.course_duration
		FROM "Teachers" t
		LEFT JOIN "TeacherSectionAssignment" tsa ON t.id = tsa.teacher_id
		LEFT JOIN "Sections" s ON tsa.section_id = s.id
		LEFT JOIN "Subjects" sub ON tsa.subject_id = sub.id
		LEFT JOIN "courses" course ON s.course_id = course.id
		WHERE t.email ILIKE $1
	`
		err = db.Pool.QueryRow(ctx, query, "%"+email+"%").Scan(&u.Name, &u.Section, &u.Subject, &u.Semester, &u.CourseName, &u.CourseDuration)
		if err != nil {
			logger.Log.Error().Err(err).Msg("error fetching teacher details")
			return nil, err
		}

	case "student":
		query := `
			SELECT s.name, s.enrollment_number, s.current_semester, s.section_id
			FROM "Students" s
			WHERE s.email ILIKE $1
		`
		var enrollment, semester string
		err = db.Pool.QueryRow(ctx, query, "%"+email+"%").Scan(&u.Name, &enrollment, &semester, &u.Section)
		if err != nil {
			logger.Log.Error().Err(err).Msg("error fetching student details")
			return nil, err
		}

	case "admin":
		query := `
			SELECT a.name, a.course_id
			FROM "Admins" a
			WHERE a.email ILIKE $1
		`
		err = db.Pool.QueryRow(ctx, query, "%"+email+"%").Scan(&u.Name, &u.Section)
		if err != nil {
			logger.Log.Error().Err(err).Msg("error fetching admin details")
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unknown role: %s", role)
	}

	return u, nil
}
