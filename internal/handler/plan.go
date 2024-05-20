package handler

import (
	"goCourseService/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePlan(c *gin.Context) {
	var input models.Plan
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	courseId, _ := h.validator.ValidateId(c)
	input.CourseId = courseId

	planId, err := h.service.Plan.CreatePlan(input)

	if err != nil {
		if planId == -1 {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"message": "План успешно создался, но с рассылкой какието проблемы",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "План успешно создался, всем отправились уведомления",
	})
}

func (h *Handler) GetPlans(c *gin.Context) {

}

func (h *Handler) GetPlanById(c *gin.Context) {
	_, err := h.validator.ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

}

func (h *Handler) DeletePlan(c *gin.Context) {

}

func (h *Handler) UpdatePlan(c *gin.Context) {

}
