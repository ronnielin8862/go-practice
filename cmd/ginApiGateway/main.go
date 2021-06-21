package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Print("要開始了喔")
}

func GetMethod(c *gin.Context) {
	fmt.Print("進入Get")

	// c.Params()
	c.String(http.StatusOK, "Get完成")

}

func main() {
	server := gin.Default()

	server.GET("/", GetMethod)

	server.Run(":3000")

}
