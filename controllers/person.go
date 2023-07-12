package controllers

import (
	"api1-new/pkg/person"

	"github.com/gin-gonic/gin"
)

func GetPerson(c *gin.Context) {
	data := person.GetPerson()
	c.JSON(200, data)
}

func PostPerson(c *gin.Context) {
	// placeholder
	c.JSON(200, "Placeholder")
}
