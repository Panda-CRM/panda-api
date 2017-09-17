package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Panda-CRM/panda-api/controllers"
	"github.com/Panda-CRM/panda-api/repositories"
)

func AddRoutesPeople(r *gin.RouterGroup) {
	controller := controllers.PersonController{Repository: repositories.NewPersonRepository()}
	routes := r.Group("/people")
	{
		routes.GET("", controller.List)
		routes.GET("/:id", controller.Get)
		routes.DELETE("/:id", controller.Delete)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
	}
}
