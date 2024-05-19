package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
	rdb *redis.Client
}

func NewCourseRepository(db *sqlx.DB, rdb *redis.Client) *CourseRepository {
	return &CourseRepository{
		db: db,
		rdb: rdb,
	}
}