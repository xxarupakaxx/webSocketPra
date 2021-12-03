package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)
import "golang.org/x/net/websocket"

func handlewebSocket(c echo.Context) error {
	websocket.Handler(func(conn *websocket.Conn) {
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				log.Fatalf("err:%v", err)
			}
		}(conn)

		err := websocket.Message.Send(conn, "Server:Hello, Client!")
		if err != nil {
			c.Logger().Error(err)
		}

		for true {
			msg := ""
			err = websocket.Message.Receive(conn, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			err := websocket.Message.Send(conn, fmt.Sprintf("Server: \"%s\" received!", msg))

			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Static("/","public")
	e.GET("/ws",handlewebSocket)
	e.Logger.Fatal(e.Start(":8000"))
}
