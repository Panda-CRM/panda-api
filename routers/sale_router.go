package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Panda-CRM/panda-api/controllers"
	"github.com/Panda-CRM/panda-api/repositories"
)

func AddRoutesSales(r *gin.RouterGroup) {
	controller := controllers.SaleController{Repository: repositories.NewSaleRepository()}
	routes := r.Group("/sales")
	{
		routes.GET("/", controller.List)
	}
}
