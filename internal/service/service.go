package service

import (
	"goCourseService/dtos"
	"goCourseService/internal/repository"
	"goCourseService/models"
)

type Course interface {
	CreateCourse(input models.Course, teacherId int) (int, error)
	GetCourseById(courseId int) (dtos.Course, error)
	GetCourses(page, limit int) ([]dtos.Course, dtos.Pagination, error)
	UpdateCourse(input models.Course, courseId int) error
	DeleteCourse(courseId int) error

	RegisterForCourse(clientId, courseId int, email string) error
}

type Plan interface {
	CreatePlan(input models.Plan) (int, error)
}

type User interface {
	GetUserFromCacheWithToken(token string) (dtos.User, error)
}

type Service struct {
	Course
	User
	Plan
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Course: NewCourseService(repos.Course),
		User:   NewUserService(repos.User),
		Plan:   NewPlanService(repos.Plan, repos.Course),
	}
}
