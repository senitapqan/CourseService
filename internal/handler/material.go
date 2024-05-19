package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMaterial(c *gin.Context) {

}

func (h *Handler) GetMaterials(c *gin.Context) {

}

func (h *Handler) GetMaterialById(c *gin.Context) {
	_, err := h.validator.ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
	}

}

func (h *Handler) DeleteMaterial(c *gin.Context) {

}

func (h *Handler) UpdateMaterial(c *gin.Context) {

}