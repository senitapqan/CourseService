package handler

import (
	"goCourseService/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service   *service.Service
	validator RequestValidator
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service:   service,
		validator: &requestValidator{},
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	teach := router.Group("/teach")
	{
		teach.Use(h.userIdentify())
		teach.Use(h.roleIdentify(teacherCtx))
		course := teach.Group("/course") 
		{
			course.POST("", h.CreateCourse)
			course.DELETE("/:id", h.DeleteCourse)
			course.PUT("/:id", h.UpdateCourse)
		}
		plan := teach.Group("/plan")
		{
			plan.Use(h.ownerIdentify())
			plan.POST("/:id", h.CreatePlan)
			plan.DELETE("/:id", h.DeleteCourse)
			plan.PUT("/:id", h.UpdatePlan)
		}
		material := teach.Group("/matrial")
		{
			material.POST("/:id", h.CreateMaterial)
			material.DELETE("/:id", h.DeleteMaterial)
			material.PUT("/:id", h.UpdateMaterial)
		}
	} 
	router.GET("/", h.GetCourses)
	router.GET("/:id", h.GetCourseById)	

	course := router.Group("/course")
	{
		course.Use(h.userIdentify())
		course.GET("/:id", h.GetPlans)
		course.GET("/plan/:id", h.GetPlanById)
		course.GET("/matrials/:id", h.GetMaterials)
		course.POST("/register/:id", h.RegisterForCourse)
	}



	return router
}
