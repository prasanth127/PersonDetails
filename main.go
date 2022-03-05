package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID           string `json :"id"`
	NAME         string `json :"name"`
	PHONE_NUMBER string `json :"phone_number"`
	CITY         string `json :"city"`
	STATE        string `json :"state"`
	STREET1      string `json :"street1"`
	ZIP_CODE     string `json :"zip_code"`
}

var persons = []person{
	{ID: "01", NAME: "golang", PHONE_NUMBER: "99777777", CITY: "hyd", STATE: "TS", STREET1: "Abcd", ZIP_CODE: "555111"},
	{ID: "01", NAME: "developer", PHONE_NUMBER: "123456", CITY: "Bangolare", STATE: "KA", STREET1: "xyz", ZIP_CODE: "000777"},
}

func main() {
	router := gin.Default()
	router.GET("/persons", getpersons)
	router.GET("/persons/:id", getpersonByID)
	router.POST("/persons", postpersons)

	router.Run("localhost:8080")
}
func getpersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)

}
func postpersons(c *gin.Context) {
	var newperson person

	if err := c.BindJSON(&newperson); err != nil {
		return
	}

	persons = append(persons, newperson)
	c.IndentedJSON(http.StatusCreated, newperson)
}

func getpersonByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range persons {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
