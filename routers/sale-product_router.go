package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/controllers"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

func AddRoutesSaleProducts(r *gin.RouterGroup) {
	controller := controllers.SaleProductController{Repository: repositories.NewSaleProductRepository()}
	routes := r.Group("/sale_products")
	{
		routes.GET("/", controller.List)
	}
}
