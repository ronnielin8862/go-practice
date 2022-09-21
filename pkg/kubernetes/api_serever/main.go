package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/get", GetMethod)
	server.GET("/start_mark/*filepath", StartMark)

	err := server.Run(":3000")
	if err != nil {
		fmt.Println("start gin err : ", err)
	}
}

func StartMark(c *gin.Context) {
	fmt.Println("Request URI = ", c.Request.RequestURI)
	fmt.Println("Request Param = ", c.Param("filepath"))
	c.String(200, "complete")
}

func GetMethod(c *gin.Context) {
	fmt.Println("Get entrance")
	type Person struct {
		Name string `json:"name" binding:"required"`
	}
	var person Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.String(400, "must have name")
		return
	}
	c.String(200, fmt.Sprint("Get finish,", person.Name))
}
