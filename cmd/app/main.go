package main

import (
	"os"

	"github.com/labstack/echo/v4"

	"Back-end/config"
	"Back-end/docs"
	auth "Back-end/internal/auth/route"
	generate "Back-end/internal/generate/route"
	invoice "Back-end/internal/invoice/route"
	item "Back-end/internal/invoice_item/route"
	status "Back-end/internal/invoice_payment_status/route"
	payment "Back-end/internal/payment_gateway/xendit/route"
	send "Back-end/internal/send_customer/route"
	image "Back-end/internal/upload_file/route"
	user "Back-end/internal/user/route"

	echoSwag "github.com/swaggo/echo-swagger"
)

// @title        Generate Invoice System API Documentation
// @description  This is Generate Invoice System API
// @version      2.0
// @host         api.calorilin.me:8888
// @BasePath
// @schemes                     http https
// @securityDefinitions.apiKey  JWT
// @in                          header
// @name                        Authorization
func main() {
	config := config.InitConfig()
	e := echo.New()

	e.Static("images", "public")

	auth.RegisterAuthGroupAPI(e, config)
	user.RegisterUserGroupAPI(e, config)
	image.RegisterUploadImageGroupAPI(e, config)
	generate.RegisterGenerateInvoiceGroupAPI(e, config)
	invoice.RegisterInvoiceGroupAPI(e, config)
	item.RegisterInvoiceItemGroupAPI(e, config)
	send.RegisterSendCustomerGroupAPI(e, config)
	status.RegisterInvoicePaymentStatusGroupAPI(e, config)
	payment.RegisterPaymentGatewayGroupAPI(e, config)

	e.GET("/swagger/*", echoSwag.WrapHandler)
	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
