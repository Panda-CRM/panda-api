package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Panda-CRM/panda-api/controllers"
	"github.com/Panda-CRM/panda-api/repositories"
)

func AddRoutesTaskCategories(r *gin.RouterGroup) {
	controller := controllers.TaskCategoryController{Repository: repositories.NewTaskCategoryRepository()}
	routes := r.Group("/task_categories")
	{
		routes.GET("", controller.List)
		routes.GET("/:id", controller.Get)
		routes.DELETE("/:id", controller.Delete)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
	}
}
