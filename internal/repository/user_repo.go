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
			c.course_name,
			c.course_duration
		FROM "Teachers" t
		LEFT JOIN "TeacherSectionAssignment" tsa ON t.id = tsa.teacher_id
		LEFT JOIN "Sections" s ON tsa.section_id = s.id
		LEFT JOIN "Subjects" sub ON tsa.subject_id = sub.id
		LEFT JOIN "courses" c ON s.course_id = c.id
		WHERE t.email ILIKE $1
	`
		rows, err := db.Pool.Query(ctx, query, "%"+email+"%")
		if err != nil {
			logger.Log.Error().Err(err).Msg("error executing query for teacher details")
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var (
				name           string
				section        string
				subject        string
				semester       string
				courseName     string
				courseDuration string
			)

			err := rows.Scan(&name, &section, &subject, &semester, &courseName, &courseDuration)
			if err != nil {
				logger.Log.Error().Err(err).Msg("error scanning teacher details")
				return nil, err
			}

			u.Name = name
			u.Subject = subject
			u.Semester = semester
			u.CourseName = courseName
			u.CourseDuration = courseDuration

			if section != "" {
				u.Section = append(u.Section, section)
			}
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

	fmt.Println(u)
	return u, nil
}
