package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

type SaleProductController struct {
	Repository repositories.SaleProductRepository
}

func (controller SaleProductController) List(c *gin.Context) {
	params := c.Request.URL.Query()
	sales := controller.Repository.List(params)

	c.JSON(200, sales)
}
