package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AttendEase-ERP/erp-backend/internal/repository"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserDetailsByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"email": user.Email,
		"role":  user.Role.Role,
		"name":  user.Name,
	}

	switch user.Role.Role {
	case "teacher":
		response["section"] = user.Section
		response["subject"] = user.Subject
		response["course_name"] = user.CourseName
		response["course_duration"] = user.CourseDuration
		response["Semester"] = user.Semester

	case "student":
		response["section"] = user.Section
		response["Semester"] = user.Semester
		response["enrollment_number"] = user.EnrollmentNumber
		response["course_name"] = user.CourseName

	case "admin":
		response["course_name"] = user.CourseName
		response["course_duration"] = user.CourseDuration
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetStudentsOfTeacher(w http.ResponseWriter, r *http.Request) {
	semester := r.URL.Query().Get("semester")
	section := r.URL.Query().Get("section")
	date := r.URL.Query().Get("date")

	if semester == "" || section == "" {
		http.Error(w, "semester and section are required", http.StatusBadRequest)
		return
	}

	students, err := repository.GetStudentsOfTeacher(r.Context(), semester, section, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
