package repository

import (
	"goCourseService/dtos"

	"github.com/go-redis/redis/v8"
)

type UserRepository struct {
	rdb *redis.Client
}

func NewUserRepository(rdb *redis.Client) *UserRepository {
	return &UserRepository{
		rdb: rdb,
	}
}

func (r *UserRepository) GetUserFromCacheWithToken(token string) (dtos.User, error) {
	return dtos.User{}, nil
}

