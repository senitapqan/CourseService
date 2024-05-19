package service

import (
	"goCourseService/dtos"
	"goCourseService/internal/repository"
	"goCourseService/models"
)

type Course interface {
	CreateCourse(input models.Course, teacherId int) error
	GetCourseById(courseId int) (models.Course, error)
	GetCourses() ([]models.Course, error)
	UpdateCourse(input models.Course, courseId int) error
	DeleteCourse(courseId int) error
}

type User interface {
	GetUserFromCacheWithToken(token string) (dtos.User, error)
}

type Service struct {
	Course
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Course: NewCourseService(repos.Course),
		User:   NewUserService(repos.User),
	}
}
