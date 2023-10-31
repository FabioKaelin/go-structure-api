package main

import (
	"api1-new/controllers"
	"api1-new/pkg/logger"

	"api1-new/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go Example API
// @version         1.0
// @description     Example API for demonstrating how to use Swagger with Gin and its folder structure.

// @host      localhost:8000
// @BasePath  /api

func main() {
	logger.InitializeZapCustomLogger()
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	apiGroup := server.Group("/api")
	personGroup := apiGroup.Group("/persons")
	{
		personGroup.GET("", controllers.GetPersons)
		personGroup.POST("", controllers.PostPersons)
	}
	companyGroup := apiGroup.Group("/companies")
	{
		companyGroup.GET("", controllers.GetCompanies)
	}
	docs.SwaggerInfo.BasePath = "/api"
	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.Run("0.0.0.0:8000")
}
