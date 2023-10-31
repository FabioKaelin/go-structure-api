package controllers

import (
	"api1-new/pkg/company"

	"github.com/gin-gonic/gin"
)

// GetCompanies godoc
// @Summary      Get all companies
// @Tags         companies
// @Produce      json
// @Success      200  {array}  company.company
// @Failure      500  {string}  string
// @Router       /companies [get]
func GetCompanies(c *gin.Context) {
	data, err := company.GetCompany()
	if err != nil {
		c.IndentedJSON(500, err)
		return
	}
	c.IndentedJSON(200, data)
}
