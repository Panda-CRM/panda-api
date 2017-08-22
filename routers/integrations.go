package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/integrations"
	"github.com/wilsontamarozzi/panda-api/repositories"
)

func AddRoutesIntegrations(r *gin.RouterGroup) {
	cvc := integrations.IntegrationCVC{
		PersonRepository:      repositories.NewPersonRepository(),
		SaleRepository:        repositories.NewSaleRepository(),
		SaleProductRepository: repositories.NewSaleProductRepository(),
		ProductRepository:     repositories.NewProductRepository(),
	}
	routes := r.Group("/integrations")
	{
		routes.GET("/cvc/import", cvc.Import)
	}
}
