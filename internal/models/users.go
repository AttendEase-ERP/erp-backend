package models

type UserRole struct {
	Role string `json:"role"` // teacher, admin, student
}

type UserDetails struct {
	Email            string   `json:"email"`
	Role             UserRole `json:"role"`
	Name             string   `json:"name"`
	Section          []string `json:"section"`
	Subject          string   `json:"subject"`
	Semester         string   `json:"semester"`
	CourseName       string   `json:"course_name"`
	CourseDuration   string   `json:"course_duration"`
	EnrollmentNumber string   `json:"enrollment_number"`
}
/* 
type AttendanceStatus string

const (
	AttendancePresent AttendanceStatus = "present"
	AttendanceAbsent  AttendanceStatus = "absent"
)

type Student struct {
	ID               string `db:"id"`
	Name             string `db:"name"`
	Email            string `db:"email"`
	EnrollmentNumber string `db:"enrollment_number"`
	CurrentSemester  int    `db:"current_semester"`
	SectionID        string `db:"section_id"`
}

type Section struct {
	ID          string `db:"id"`
	SectionName string `db:"section_name"`
	CourseID    string `db:"course_id"`
}

type Course struct {
	ID             string `db:"id"`
	CourseName     string `db:"course_name"`
	CourseDuration string `db:"course_duration"`
}

type Attendance struct {
	ID        string           `db:"id"`
	StudentID string           `db:"student_id"`
	TeacherID string           `db:"teacher_id"`
	SubjectID string           `db:"subject_id"`
	Date      time.Time        `db:"date"`
	Status    AttendanceStatus `db:"attendance_status"`
}
 */
type StudentsAttendanceList struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	EnrollmentNumber string `json:"enrollment_number"`
	Status           string `json:"attendance_status"`
}
