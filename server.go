package main

import (
	"os"
	"github.com/wilsontamarozzi/panda-api/routers"
)

const ENV_PORT = "PORT"
var PORT_DEFAULT string = "8080"

func init() {
	port := os.Getenv(ENV_PORT)

	if len(port) > 0 {
		PORT_DEFAULT = port
	}
}

func main() {
	//Inicia todas as rotas
	router := routers.InitRoutes()

	//Inicia o Server
    router.Run(":" + PORT_DEFAULT)
}