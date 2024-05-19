package handler

import (
	"goCourseService/models"
	"goCourseService/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourse(c *gin.Context) {
	teacherId, _ := getId(c, teacherCtx)
	
	var input models.Course
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	courseId, err := h.service.Course.CreateCourse(input, teacherId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, courseId)
}

func (h *Handler) GetCourses(c *gin.Context) {
	page, limit, err := h.validator.ValidatePage(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	courses, pagination, err := h.service.Course.GetCourses(page, limit)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	
	c.JSON(http.StatusOK, dtos.GetAllCoursesResponse{
		Data:       courses,
		Pagination: pagination,
	  })
}

func (h *Handler) GetCourseById(c *gin.Context) {
	courseId, err := h.validator.ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

	course, err := h.service.Course.GetCourseById(courseId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	
	c.JSON(http.StatusOK, dtos.GetCourseById) {
		Data: course,
		Pagination: pagination
	}
}

func (h *Handler) DeleteCourse(c *gin.Context) {

}

func (h *Handler) UpdateCourse(c *gin.Context) {

}