package handler

import (
	"goCourseService/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	validator RequestValidator
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
		validator: &requestValidator{},
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	
	teach := router.Group("/teach") 
	{
		teach.Use(h.userIdentify())
		teach.Use(h.roleIdentify(teacherCtx))
		teach.POST("")
	} 	

	return router
}