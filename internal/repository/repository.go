package repository

import (
	//"goCourseService/models"

	"goCourseService/dtos"
	"goCourseService/models"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Course interface {
	CreateCourse(input models.Course, teacherId int) (int, error)
	GetCourseById(courseId int) (dtos.Course, error)
	GetCourses(page, limit int) ([]dtos.Course, dtos.Pagination, error)
	UpdateCourse(input models.Course, courseId int) error
	DeleteCourse(courseId int) error

	RegisterForCourse(clientId, courseId int, email string) error
}

type User interface {
	GetUserFromCacheWithToken(token string) (dtos.User, error)
}

type Repository struct {
	Course
	User
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{
		Course: NewCourseRepository(db, redis),
		User:   NewUserRepository(redis),
	}
}
