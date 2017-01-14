package routers

import (
    "github.com/wilsontamarozzi/panda-api/middleware"
    "github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	// Rota Global
    r := gin.New()
    
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