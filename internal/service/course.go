package service

import (
	"goCourseService/internal/repository"
	"goCourseService/models"
)

type CourseService struct {
	repos repository.Course
}

func NewCourseService(repos repository.Course) *CourseService {
	return &CourseService{
		repos: repos,
	}
}

func (s *CourseService) CreateCourse(input models.Course, teacherId int) error {
	return s.repos.CreateCourse(input, teacherId)
}

func (s *CourseService) GetCourses() ([]models.Course, error) {
	return s.repos.GetCourses()
}

func (s *CourseService) GetCourseById(courseId int) (models.Course, error) {
	return s.repos.GetCourseById(courseId)
}

func (s *CourseService) UpdateCourse(input models.Course, courseId int) error {
	return s.repos.UpdateCourse(input, courseId)
}

func (s *CourseService) DeleteCourse(courseId int) error {
	return s.repos.DeleteCourse(courseId)
}
