package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/controllers"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

func AddRoutesSales(r *gin.RouterGroup) {
	controller := controllers.SaleController{Repository: repositories.NewSaleRepository()}
	routes := r.Group("/sales")
	{
		routes.GET("/", controller.List)
	}
}
