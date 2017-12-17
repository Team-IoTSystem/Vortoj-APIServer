package main

import (
	_ "strconv"

	"./controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return controllers.ConnectDB(c)
	})
	e.GET("/Packet/:id", func(c echo.Context) error {
		return controllers.SelectPacketData(c)
	})
	e.GET("/Packet/new", func(c echo.Context) error {
		return controllers.NewPacketData(c)
	})
	e.POST("/InsertDB", func(c echo.Context) error {
		return controllers.InsertDB(c)
	})
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
