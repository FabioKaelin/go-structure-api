package controllers

import (
	"api1-new/pkg/db"

	"github.com/gin-gonic/gin"
)

// GetCompanies godoc
// @Summary      Get all companies
// @Tags         companies
// @Produce      json
// @Success      200  {array}  db.company
// @Router       /companies [get]
func GetCompanies(c *gin.Context) {
	data := db.GetCompanies()
	c.IndentedJSON(200, data)
}
