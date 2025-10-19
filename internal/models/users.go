package models

type UserRole struct {
	Role string `json:"role"` // teacher, admin, student
}

type UserDetails struct {
	Email            string   `json:"email"`
	Role             UserRole `json:"role"`
	Name             string   `json:"name"`
	Section          string   `json:"section"`
	Subject          string   `json:"subject"`
	Semester         string   `json:"semester"`
	CourseName       string   `json:"course_name"`
	CourseDuration   string   `json:"course_duration"`
	EnrollmentNumber string   `json:"enrollment_number"`
}
