package repository

import (
	//"goCourseService/models"

	"goCourseService/dtos"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Course interface {

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
		User: NewUserRepository(redis),
	}
}