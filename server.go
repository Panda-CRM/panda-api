package main

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/routers"
	"log"
	"os"
)

//ENV_RUN_PORT variavel ambiente da porta de inicialização
const ENV_RUN_PORT = "PORT"

//RUN_PORT porta de inicialização padrão
var RUN_PORT = "8080"

func init() {
	getEnvPort()
	database.RebuildDataBase(true)
}

func getEnvPort() {
	log.Print("[CONFIG] Lendo configurações de inicialização")
	port := os.Getenv(ENV_RUN_PORT)
	if len(port) > 0 {
		RUN_PORT = port
	}
}

func main() {
	routers.InitRoutes().Run(":" + RUN_PORT)
}
