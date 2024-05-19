package service

import (
	"goCourseService/dtos"
	"goCourseService/internal/repository"
)

type UserService struct {
	repos repository.User
}

func NewUserService(repos repository.User) *UserService {
	return &UserService{
		repos: repos,
	}
}

func (s *UserService) GetUserFromCacheWithToken(token string) (dtos.User, error) {
	return s.repos.GetUserFromCacheWithToken(token)
}