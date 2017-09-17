package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

type SaleController struct {
	Repository repositories.SaleRepository
}

func (controller SaleController) List(c *gin.Context) {
	params := c.Request.URL.Query()
	sales := controller.Repository.List(params)

	c.JSON(200, sales)
}
