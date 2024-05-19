package repository

import (
	"fmt"
	"goCourseService/dtos"
	"goCourseService/models"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewCourseRepository(db *sqlx.DB, rdb *redis.Client) *CourseRepository {
	return &CourseRepository{
		db:  db,
		rdb: rdb,
	}
}

func (s *CourseRepository) CreateCourse(input models.Course, teacherId int) (int, error) {
	query := fmt.Sprintf(`insert into %s (title, subtitle, teacher_id, teacher_full_name) 
		values($1, $2, $3, $4, $5)`)
}

func (s *CourseRepository) GetCourses(page, limit int) ([]dtos.Course, dtos.Pagination, error) {
	return nil, dtos.Pagination{}, nil
}

func (s *CourseRepository) GetCourseById(courseId int) (dtos.Course, error) {
	return dtos.Course{}, nil
}

func (s *CourseRepository) UpdateCourse(input models.Course, courseId int) error {
	return nil
}

func (s *CourseRepository) DeleteCourse(courseId int) error {
	return nil
}
