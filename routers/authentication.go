package routers

import (
    "github.com/gin-gonic/gin"
    "panda-api/middleware"
)

func AddRoutesAuthentication(r *gin.RouterGroup) {
    r.POST("/auth/auth_token", middleware.SetToken)
}