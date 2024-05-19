package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourse(c *gin.Context) {


}

func (h *Handler) GetCourses(c *gin.Context) {

}

func (h *Handler) GetCourseById(c *gin.Context) {
	_, err := h.validator.ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

}

func (h *Handler) DeleteCourse(c *gin.Context) {

}

func (h *Handler) UpdateCourse(c *gin.Context) {

}