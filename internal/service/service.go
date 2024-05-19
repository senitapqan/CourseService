package service

import (
	"goCourseService/dtos"
	"goCourseService/internal/repository"
)

type Course interface {
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
		User: NewUserService(repos.User),
	}
}