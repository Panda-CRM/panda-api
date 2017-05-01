package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/controllers"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

func AddRoutesPeople(r *gin.RouterGroup) {
	controller := controllers.PersonController{Repository: repositories.NewPersonRepository()}
	routes := r.Group("/people")
	{
		routes.GET("/", controller.GetAll)
		routes.GET("/:id", controller.Get)
		routes.DELETE("/:id", controller.Delete)
		routes.POST("", controller.Create)
		routes.PUT("/:id", controller.Update)
	}
}