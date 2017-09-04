package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/controllers"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

func AddRoutesTasks(r *gin.RouterGroup) {
	controller := controllers.TaskController{Repository: repositories.NewTaskRepository()}
	routes := r.Group("/tasks")
	{
		routes.GET("", controller.List)
		routes.GET("/:id", controller.Get)
		routes.DELETE("/:id", controller.Delete)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
	}

	routesReports := r.Group("/reports/tasks")
	{
		routesReports.GET("", controller.ReportGeneral)
		routesReports.GET("/assignee", controller.ReportByAssignees)
		routesReports.GET("/assignee/category", controller.ReportByAssigneesAndCategory)
		routesReports.GET("/category", controller.ReportByCategories)
	}
}
