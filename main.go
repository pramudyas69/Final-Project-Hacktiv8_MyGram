package main

import (
	"MyGram/configs"
	"MyGram/routes"
	"MyGram/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := SetupRouter()

	log.Fatal(router.Run(":" + utils.GoDotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	db := configs.Connection()

	router := gin.Default()

	routes.InitAuthRoute(db, router)
	return router
}
