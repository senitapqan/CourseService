package dtos

type Course struct {
	Id              int    `json:"id" db:"id"`
	Title           string `json:"title" db:"title"`
	Subtitle        string `json:"subtitle" db:"subtitle"`
	Description     string `json:"description" db:"description"`
	TeacherId       int    `json:"teacher_id" db:"teacher_id"`
	TeacherFullName string `json:"teahcer_full_name" db:"teacher_full_name"`
	Price           int    `json:"price" db:"price"`
}

type GetAllCoursesResponse struct {
	Data       []Course   `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type GetCourseById struct {
	Data       Course     `json:"data"`
	Pagination Pagination `json:"pagination"`
}
