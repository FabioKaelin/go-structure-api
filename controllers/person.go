package controllers

import (
	"api1-new/pkg/person"

	"github.com/gin-gonic/gin"
)

// GetPersons godoc
// @Summary      Get all persons
// @Tags         persons
// @Produce      json
// @Success      200  {array}  person.Person
// @Router       /persons [get]
func GetPersons(c *gin.Context) {
	data := person.GetPerson()
	c.IndentedJSON(200, data)
}

// PostPersons godoc
// @Summary      Post a person
// @Tags         persons
// @Produce      json
// @Success      200  {string}  string
// @Router       /persons [post]
func PostPersons(c *gin.Context) {
	// placeholder
	c.JSON(200, "Placeholder")
}
