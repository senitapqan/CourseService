package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePlan(c *gin.Context) {

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