package controllers

import (
	"api1-new/pkg/person"

	"github.com/gin-gonic/gin"
)

// GetPersons godoc
//	@Summary	Get all persons
//	@Tags		persons
//	@Produce	json
//	@Success	200	{array}		person.Person
//	@Failure	500	{string}	string
//	@Router		/persons [get]
func GetPersons(c *gin.Context) {
	data, err := person.GetPerson()
	if err != nil {
		c.IndentedJSON(500, err)
		return
	}
	c.IndentedJSON(200, data)
}

// PostPersons godoc
//	@Summary	Post a person
//	@Tags		persons
//	@Produce	json
//	@Accept		json
//	@Param		person	body		person.PersonPost	true	"Person"
//	@Success	201		{string}	string
//	@Failure	400		{string}	string
//	@Failure	500		{string}	string
//	@Router		/persons [post]
func PostPersons(c *gin.Context) {
	// bind the request body to the struct
	var personData person.PersonPost
	err := c.BindJSON(&personData)
	if err != nil {
		c.IndentedJSON(400, err)
		return
	}
	err = person.CreatePerson(personData.Name, personData.Age, personData.CompanyId)
	if err != nil {
		c.IndentedJSON(500, err)
		return
	}
	c.IndentedJSON(201, "Person created")
}
