package service

import (
	"goCourseService/consts"
	"goCourseService/internal/repository"
	"goCourseService/models"
)

type PlanService struct {
	reposPlan repository.Plan
	reposCourse repository.Course
}

func NewPlanService(reposPlan repository.Plan, reposCourse repository.Course) *PlanService {
	return &PlanService{
		reposPlan: reposPlan,
		reposCourse: reposCourse,
	}
}

func (s *PlanService) CreatePlan(input models.Plan) (int, error) {
	planId, err := s.reposPlan.CreatePlan(input)

	if err != nil {
		return -1, err
	}

	followers, err:= s.reposCourse.GetCourseFollowers(input.CourseId)

	if err != nil {
		return planId, err
	}

	err = SendNotificationRequest(consts.NotificationServiceHost, consts.NotificationServicePort, input.Title, followers)
	if err != nil {
		return planId, err
	}

	return planId, nil
}