package service

import "goCourseService/internal/repository"

type CourseService struct {
	repos repository.Course
}

func NewCourseService(repos repository.Course) *CourseService {
	return &CourseService{
		repos: repos,
	}
}

