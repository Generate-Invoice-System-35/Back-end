package route

import (
	"Back-end/config"
	d "Back-end/database"
	h "Back-end/internal/payment_gateway/midtrans/handler/http"
	r "Back-end/internal/payment_gateway/midtrans/repository/mysql"
	u "Back-end/internal/payment_gateway/midtrans/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaymentGatewayGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlPaymentGatewayRepository(db)
	service := u.NewServicePaymentGateway(repo, conf)
	hand := h.EchoPaymentGatewayController{Service: service}

	apiInvoice := e.Group("/payment/midtrans/invoice",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiInvoice.POST("/:id", hand.ChargeTransactionController, middleware.JWT([]byte(conf.JWT_KEY)))
}
