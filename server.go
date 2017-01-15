package main

import (
	"os"
	"github.com/wilsontamarozzi/panda-api/routers"
)

const ENV_RUN_PORT = "RUN_PORT"
var RUN_PORT string = "8080"

func init() {
	getEnvPort()
}

func getEnvPort() {
	port := os.Getenv(ENV_RUN_PORT)

	if len(port) > 0 {
		RUN_PORT = port
	}
}

func main() {
	//Inicia todas as rotas
	router := routers.InitRoutes()

	//Inicia o Server
    router.Run(":" + RUN_PORT)
}