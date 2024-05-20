package models

type Course struct {
	Id              int
	Title           string `json:"title"`
	Subtitle        string `json:"subtitle"`
	Description     string `json:"description"`
	TeacherId       int    `json:"teacher_id"`
	TeacherFullName string `json:"teacher_full_name"`
	Price           int    `json:"price"`
}

type Plan struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title" binding:"required"`
	CourseId int    `json:"course_id" db:"course_id"`
}

type Material struct {
	Id        int
	Title     string
	PlanId    int
	VideoLink []string
	Context   string
}
