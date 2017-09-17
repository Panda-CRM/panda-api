package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Panda-CRM/panda-api/controllers"
	"github.com/Panda-CRM/panda-api/repositories"
)

func AddRoutesUser(r *gin.RouterGroup) {
	controller := controllers.UserController{Repository: repositories.NewUserRepository()}
	routes := r.Group("/users")
	{
		routes.GET("/me", controller.Me)

	}
}
