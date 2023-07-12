package main

import (
	"api1-new/controllers"
	"api1-new/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitializeZapCustomLogger()
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	apiGroup := server.Group("/api")
	personGroup := apiGroup.Group("/person")
	{
		personGroup.GET("", controllers.GetPerson)
		personGroup.POST("", controllers.PostPerson)
	}
	companyGroup := apiGroup.Group("/company")
	{
		companyGroup.GET("", controllers.GetCompany)
	}
	server.Run("0.0.0.0:8080")
}
