package routers

import (
    "os"
    "github.com/wilsontamarozzi/panda-api/middleware"
    "github.com/bugsnag/bugsnag-go"
    "github.com/bugsnag/bugsnag-go/gin"
    "github.com/gin-gonic/gin"
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

	// Rota Global
    r := gin.New()
    
    if API_KEY_BUGSNAG != "" {
        r.Use(bugsnaggin.AutoNotify(bugsnag.Configuration{
            APIKey: API_KEY_BUGSNAG,
        }))
    }

    // Logs das rotas
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // CORS
    r.Use(middleware.CORS())

    // Grupo de rotas v1
    v1 := r.Group("api/v1")

    // Rotas autorizadas: Adicionar aqui as rotas autorizadas
    AddRoutesAuthentication(v1)

    v1.Use(middleware.AuthRequired()) 
    {    	
    	// Rotas não autorizadas: Adicionar aqui as rotas não autorizadas
    	AddRoutesPeople(v1)
        AddRoutesTaskCategories(v1)
        AddRoutesTasks(v1)
	}    

	return r
}