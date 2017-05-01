package main

import (
	"github.com/wilsontamarozzi/panda-api/routers"
	"os"
)

const ENV_RUN_PORT = "PORT"

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
	routers.InitRoutes().Run(":" + RUN_PORT)
}
