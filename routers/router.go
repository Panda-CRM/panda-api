package routers

import (
	"github.com/bugsnag/bugsnag-go"
	"github.com/bugsnag/bugsnag-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/wilsontamarozzi/panda-api/middleware"
	"os"
)

const ENV_API_KEY_BUGSNAG = "API_KEY_BUGSNAG"

var API_KEY_BUGSNAG string = ""

func init() {
	getEnvAPIKeyBugsnag()
}

func getEnvAPIKeyBugsnag() {
	apiKey := os.Getenv(ENV_API_KEY_BUGSNAG)
	if len(apiKey) > 0 {
		API_KEY_BUGSNAG = apiKey
	}
}

func InitRoutes() *gin.Engine {
	router := gin.New()

	if API_KEY_BUGSNAG != "" {
		router.Use(bugsnaggin.AutoNotify(bugsnag.Configuration{
			APIKey: API_KEY_BUGSNAG,
		}))
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	v1 := router.Group("api/v1")
	{
		// Rotas autorizadas: Adicionar aqui as rotas autorizadas
		AddRoutesAuthentication(v1)

		//v1.Use(middleware.AuthRequired())
		v1.Use(middleware.AuthRequired())
		{
			// Rotas não autorizadas: Adicionar aqui as rotas não autorizadas
			AddRoutesPeople(v1)
			AddRoutesUser(v1)
			AddRoutesTaskCategories(v1)
			AddRoutesTasks(v1)
			AddRoutesSales(v1)
			AddRoutesSaleProducts(v1)
			AddRoutesIntegrations(v1)
		}
	}
	return router
}
