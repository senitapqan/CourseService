package repository

import (
	"fmt"
	"goCourseService/consts"
	"goCourseService/dtos"
	"goCourseService/models"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
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

func (r *CourseRepository) CreateCourse(input models.Course, teacherId int) (int, error) {
	var courseId int
	query := fmt.Sprintf(`insert into %s (title, subtitle, description, teacher_id, teacher_full_name) 
		values($1, $2, $3, $4, $5) returning id`, consts.CourseTable)

	err := r.db.Get(&courseId, query, input.Title, input.Subtitle, input.Description, input.TeacherId, input.TeacherFullName)
	if err != nil {
		return -1, err
	}

	return courseId, nil
}

func (r *CourseRepository) GetCourses(page, limit int) ([]dtos.Course, dtos.Pagination, error) {
	var courses []dtos.Course
	query := fmt.Sprintf(`select c.title, c.subtitle, c.description, c.teacher_full_name from %s c
			where c.draft = true
			order by c.created_time desc
			limit %d offset %d`, consts.CourseTable, limit, (page-1)*limit)

	log.Info().Msg(query)
	err := r.db.Select(&courses, query)
	if err != nil {
		return nil, dtos.Pagination{}, err
	}

	query = fmt.Sprintf(`select count(*) from %s c
			where c.draft = true`, consts.CourseTable)

	var totalRows int
	err = r.db.Get(&totalRows, query)
	if err != nil {
		return nil, dtos.Pagination{}, err
	}

	return courses, dtos.Pagination{
		PageElementsCount: limit,
		CurrentPage:       page,
		TotalPage:         (totalRows + limit - 1) / limit,
	}, nil
}

func (r *CourseRepository) GetCourseById(courseId int) (dtos.Course, error) {
	var course dtos.Course
	query := fmt.Sprintf(`select c.title, c.subtitle, c.description, c.teacher_id, c.teacher_full_name from %s c
			where c.draft = true and c.id = $1`, consts.CourseTable)
	err := r.db.Get(&course, query, courseId)
	return course, err
}

func (r *CourseRepository) UpdateCourse(input models.Course, courseId int) error {
	return nil
}

func (r *CourseRepository) DeleteCourse(courseId int) error {
	return nil
}

func (r *CourseRepository) RegisterForCourse(clientId, courseId int, email string) error {
	query := fmt.Sprintf("insert into %s (course_id, student_id, student_email) values($1, $2, $3)", consts.CourseStudentsTable)
	_, err := r.db.Exec(query, courseId, clientId, email)

	return err
}

func (r *CourseRepository) GetCourseFollowers(courseId int) ([]string, error) {
	var result []string
	query := fmt.Sprintf("select distinct student_email from %s where course_id = $1", consts.CourseStudentsTable)
	rows, err := r.db.Query(query, courseId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			return nil, err
		}
		result = append(result, email)
	}

	return result, nil
}
