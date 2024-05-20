package service

import (
	"goCourseService/internal/repository"
	"goCourseService/models"
	"goCourseService/dtos"
)

type CourseService struct {
	repos repository.Course
}

func NewCourseService(repos repository.Course) *CourseService {
	return &CourseService{
		repos: repos,
	}
}

func (s *CourseService) CreateCourse(input models.Course, teacherId int) (int, error) {
	return s.repos.CreateCourse(input, teacherId)
}

func (s *CourseService) GetCourses(page, limit int) ([]dtos.Course, dtos.Pagination, error) {
	return s.repos.GetCourses(page, limit)
}

func (s *CourseService) GetCourseById(courseId int) (dtos.Course, error) {
	return s.repos.GetCourseById(courseId)
}

func (s *CourseService) UpdateCourse(input models.Course, courseId int) error {
	return s.repos.UpdateCourse(input, courseId)
}

func (s *CourseService) DeleteCourse(courseId int) error {
	return s.repos.DeleteCourse(courseId)
}

func (s *CourseService) RegisterForCourse(clientId, courseId int, email string) error {
	return s.repos.RegisterForCourse(clientId, courseId, email)
}
