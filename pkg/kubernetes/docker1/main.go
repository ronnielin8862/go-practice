package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/get", GetMethod)
	server.GET("/start_mark/*filepath", StartMark)

	server.Run(":3000")
}

func StartMark(c *gin.Context) {
	fmt.Println("Request URI = ", c.Request.RequestURI)
	fmt.Println("Request Param = ", c.Param("filepath"))
	c.String(200, "完成")
}

func GetMethod(c *gin.Context) {
	fmt.Println("進入Get")
	type Person struct {
		Name string `form:"name"`
	}
	var person Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Println("body = ", person.Name)
	c.String(200, "Get完成")
}
