package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	name := ctx.Param("name")
	msg := fmt.Sprintf("Hello World This is %s", name)
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 200,
		"messsage":    msg,
	})
}

func main() {
	fmt.Println("Welcome to go web development framwork gin")

	// initialize the gin
	r := gin.Default()

	r.GET("/:name", HelloWorld)

	r.Run(":5000")
}
