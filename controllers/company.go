package controllers

import (
	"api1-new/pkg/db"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	data := db.GetCompanies()
	c.JSON(200, data)
}
