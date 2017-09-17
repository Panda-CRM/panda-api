package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/controllers"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

func AddRoutesUser(r *gin.RouterGroup) {
	controller := controllers.UserController{Repository: repositories.NewUserRepository()}
	routes := r.Group("/users")
	{
		routes.GET("/me", controller.Me)

	}
}
