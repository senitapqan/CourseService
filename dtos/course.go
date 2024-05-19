package dtos

type Course struct {
	Id              int
	Title           string
	Subtitle        string
	Description     string
	TeacherId       int
	TeacherFullName string
	Price           int
}

type GetAllCoursesResponse struct {
	Data       []Course   `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type GetCourseById struct {
	Data Course `json:"data"`
}
