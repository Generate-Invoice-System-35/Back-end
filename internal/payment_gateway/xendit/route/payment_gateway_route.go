package route

import (
	"Back-end/config"
	d "Back-end/database"
	h "Back-end/internal/payment_gateway/xendit/handler/http"
	r "Back-end/internal/payment_gateway/xendit/repository/mysql"
	u "Back-end/internal/payment_gateway/xendit/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaymentGatewayGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlPaymentGatewayRepository(db)
	service := u.NewServicePaymentGateway(repo, conf)
	hand := h.EchoPaymentGatewayController{Service: service}

	apiInvoice := e.Group("/payment/xendit/invoice",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiInvoice.POST("/:id", hand.CreateXenditPaymentInvoiceController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoice.GET("/:id", hand.GetXenditPaymentInvoiceController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoice.GET("", hand.GetAllXenditPaymentInvoiceController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoice.GET("/expire/:id", hand.ExpireXenditPaymentInvoiceController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoice.POST("/callback", hand.CallbackXenditPaymentInvoiceController)
}
