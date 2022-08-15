package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return StartWebSocketClient(c)
	})
	log.Fatal(e.Start("0.0.0.0:8080"))
}

func StartWebSocketClient(ctx echo.Context) error {
	rq := ctx.Request()
	addr := rq.RemoteAddr
	fmt.Println("addr:", addr)
	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := upgrader.Upgrade(ctx.Response(), rq, nil)
	if err != nil {
		fmt.Println("err = ", err)
	}

	for {
		m, c, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("err = ", err)
			return err
		}
		fmt.Println("m = ", m, "c = ", string(c))

		err = conn.WriteMessage(m, c)
		if err != nil {
			fmt.Println("err = ", err)
			return err
		}
	}

}
