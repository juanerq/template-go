package main

import (
	"fmt"
	"template-go/config"
	"template-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	envConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes.RegisterRoutes(router, envConfig)

	router.Run(fmt.Sprintf(":%d", envConfig.Port))
}
