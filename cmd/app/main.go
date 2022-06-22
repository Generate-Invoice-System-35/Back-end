package main

import (
	"os"

	"github.com/labstack/echo/v4"

	conf "Back-end/config"
	docs "Back-end/docs"
	pmgt "Back-end/internal/payment_gateway"
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

	e.Static("storage", "storage")

	pmgt.PaymentGateway(config)
	rest.RegisterAuthGroupAPI(e, config)
	rest.RegisterUserGroupAPI(e, config)
	rest.RegisterUploadImageGroupAPI(e, config)
	rest.RegisterGenerateInvoiceGroupAPI(e, config)
	rest.RegisterInvoiceGroupAPI(e, config)
	rest.RegisterInvoiceItemGroupAPI(e, config)
	rest.RegisterInvoicePaymentStatusGroupAPI(e, config)

	e.GET("/swagger/*", echoSwag.WrapHandler)
	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
