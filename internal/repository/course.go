package repository

import (
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
	
}

func (s *CourseRepository) GetCourses() ([]models.Course, error) {

}

func (s *CourseRepository) GetCourseById(courseId int) (models.Course, error) {

}

func (s *CourseRepository) UpdateCourse(input models.Course, courseId int) error {

}

func (s *CourseRepository) DeleteCourse(courseId int) error {

}
