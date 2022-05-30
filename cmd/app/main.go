package main

import (
	// "os"

	"github.com/labstack/echo/v4"

	conf "Back-end/config"
	// docs "Back-end/docs"
	rest "Back-end/internal/route"
	// echoSwag "github.com/swaggo/echo-swagger"
)

func main() {
	config := conf.InitConfig()
	e := echo.New()

	rest.RegisterRoleGroupAPI(e, config)

	// e.GET("/swagger/*", echoSwag.WrapHandler)
	// docs.SwaggerInfo.Host = os.Getenv("APP_HOST")

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
