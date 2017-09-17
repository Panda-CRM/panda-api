package controllers

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

type UserController struct {
	Repository repositories.UserRepository
}

func (u UserController) Me(gc *gin.Context) {
	claims := jwt.ExtractClaims(gc)
	userRequest := claims["id"].(string)
	user := u.Repository.Get(userRequest)
	if user.IsEmpty() {
		gc.JSON(404, gin.H{"errors": "Registro não encontrado."})
		return
	}
	gc.JSON(200, gin.H{"user": user})
}
