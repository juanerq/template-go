package main

import (
	_ "app/docs"
	routes "app/internal"
	config "app/internal/shared/config/env"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title MS API
// @version 1.0
// @description Microservicio

// @host localhost:8080
// @BasePath /
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

	router.Run(fmt.Sprintf(":%s", envConfig.Port))
}
