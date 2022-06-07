package main

import (
	"os"

	"github.com/labstack/echo/v4"

	conf "Back-end/config"
	docs "Back-end/docs"
	rest "Back-end/internal/route"

	echoSwag "github.com/swaggo/echo-swagger"
)

// @title        Generate Invoice System API Documentation
// @description  This is Generate Invoice System API
// @version      2.0
// @host         localhost:8888
// @BasePath
// @schemes                     http https
// @securityDefinitions.apiKey  JWT
// @in                          header
// @name                        Authorization
func main() {
	config := conf.InitConfig()
	e := echo.New()

	rest.RegisterAuthGroupAPI(e, config)
	rest.RegisterUserGroupAPI(e, config)
	rest.RegisterUploadGroupAPI(e, config)

	e.GET("/swagger/*", echoSwag.WrapHandler)
	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
