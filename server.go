package main

import (
	"os"
	"panda-api/routers"
)

func GetENVPort() string {
	env := os.Getenv("PORT")

	if env == "" {
		return "8080"
	}
	
	return env
}

func main() {
	//Recebe a porta que irá abrir conexão
	port := GetENVPort()

	//Inicia todas as rotas
	router := routers.InitRoutes()

	//Inicia o Server
    router.Run(":" + port)
}