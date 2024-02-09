package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

var persons []Person

func HelloWorld(ctx *gin.Context) {
	name := ctx.Param("name")
	msg := fmt.Sprintf("Hello World This is %s", name)
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"messsage":    msg,
	})
}

func AddPerson(ctx *gin.Context) {
	var person Person
	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing body: " + err.Error(),
		})
	}
	persons = append(persons, person)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "person added successfully",
	})

}

func GetAllPersons(ctx *gin.Context) {
	byteData, err := json.Marshal(persons)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing body: " + err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": string(byteData),
	})
}

func main() {
	fmt.Println("Welcome to go web development framwork gin")

	// initialize the gin
	r := gin.Default()
	
	r.GET("/:name", HelloWorld)
	r.POST("/person", AddPerson)
	r.GET("/person", GetAllPersons)

	r.Run(":5000")
}
