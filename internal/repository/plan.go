package repository

import (
	"fmt"
	"goCourseService/consts"
	"goCourseService/models"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type PlanRepository struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewPlanRepository(db *sqlx.DB, rdb *redis.Client) *PlanRepository {
	return &PlanRepository{
		db:  db,
		rdb: rdb,
	}
}

func (r *PlanRepository) CreatePlan(input models.Plan) (int, error) {
	var planId int
	query := fmt.Sprintf("insert into %s (title, course_id) values($1, $2) returning id", consts.PlanTable)

	err := r.db.Get(&planId, query, input.Title, input.CourseId)
	if err != nil {
		return -1, err
	}
	return planId, nil
}