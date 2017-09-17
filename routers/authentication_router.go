package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/middleware"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

func AddRoutesAuthentication(r *gin.RouterGroup) {
	//controller := controllers.AccountController{Repository: repositories.NewUserRepository()}
	controller := middleware.AccountController{Repository: repositories.NewUserRepository()}
	r.POST("/accounts/token", middleware.MY_JWT.LoginHandler)
	r.POST("/accounts/forgot-password", controller.ForgotPassword)
	r.GET("/accounts/reset-password", controller.ResetPassword)
	r.POST("/accounts/reset-password", controller.CheckTokenValid)

	//pedi requisição - forgot password
	//cliquei no link - reset password get com token
	//alterei a senha - reset password post

	//r.POST("/auth/auth_token", middleware.SetToken)
	//auth.GET("/auth/refresh_token", authMiddleware.RefreshHandler)
}
