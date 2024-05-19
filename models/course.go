package models

type Course struct {
	Id              int
	Title           string
	Subtitle        string
	Description     string
	TeacherId       int
	TeacherFullName string
	Price           int
}

type Plan struct {
	Id       int
	Title    string
	CourseId int
}

type Material struct {
	Id        int
	Title     string
	PlanId    int
	VideoLink []string
	Context   string
}
