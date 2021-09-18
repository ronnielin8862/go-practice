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
	fmt.Println("進入Get")

	name := c.DefaultQuery("name", "沒輸入姓名")
	old := c.DefaultQuery("old", "沒輸入年齡")

	fmt.Println("name ============ " + name)
	fmt.Println("old ============ " + old)

	c.String(http.StatusOK, "Get完成, name = "+name+", old = "+old)

}

func main() {
	server := gin.Default()

	server.GET("/get", GetMethod)

	server.Run(":3000")

}
