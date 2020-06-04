package infrastructure

import "github.com/labstack/echo/v4"

var router *echo.Echo

func init() {
	router = echo.New()
}

// Run HTTP server
func Run() {
	router.Logger.Fatal(router.Start(":1323"))
}
