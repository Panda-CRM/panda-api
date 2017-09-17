package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Panda-CRM/panda-api/controllers"
	"github.com/Panda-CRM/panda-api/repositories"
)

func AddRoutesSaleProducts(r *gin.RouterGroup) {
	controller := controllers.SaleProductController{Repository: repositories.NewSaleProductRepository()}
	routes := r.Group("/sale_products")
	{
		routes.GET("", controller.List)
	}
}
