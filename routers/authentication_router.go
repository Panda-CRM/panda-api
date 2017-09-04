package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/middleware"
)

func AddRoutesAuthentication(r *gin.RouterGroup) {
	r.POST("/auth/auth_token", middleware.SetToken)
}
