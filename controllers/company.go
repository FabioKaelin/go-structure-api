package controllers

import (
	"api1-new/pkg/company"

	"github.com/gin-gonic/gin"
)

// GetCompanies godoc
// @Summary      Get all companies
// @Tags         companies
// @Produce      json
// @Success      200  {array}  company.Company
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

// PostCompanies godoc
// @Summary      Post a company
// @Tags         companies
// @Produce      json
// @Accept       json
// @Param        company  body  company.Company  true  "Company"
// @Success      201  {string}  string
// @Failure      500  {string}  string
// @Router       /companies [post]
func PostCompanies(c *gin.Context) {
	// bind the request body to the struct
	var companyData company.Company
	err := c.BindJSON(&companyData)
	if err != nil {
		c.IndentedJSON(500, err)
		return
	}
	err = company.CreateCompany(companyData.Name)
	if err != nil {
		c.IndentedJSON(500, err)
		return
	}
	c.IndentedJSON(201, "Company created")
}
