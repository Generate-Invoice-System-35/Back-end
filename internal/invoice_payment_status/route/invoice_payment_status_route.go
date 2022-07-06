package route

import (
	"Back-end/config"
	d "Back-end/database"
	h "Back-end/internal/handler"
	r "Back-end/internal/repository"
	u "Back-end/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterInvoicePaymentStatusGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewmYsqlInvoicePaymentStatusRepository(db)
	service := u.NewServiceInvoicePaymentStatus(repo, conf)
	hand := h.EchoInvoicePaymentStatusController{Service: service}

	apiInvoicePaymentStatus := e.Group("/invoice-payment-status",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiInvoicePaymentStatus.POST("", hand.CreateInvoicePaymentStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoicePaymentStatus.GET("", hand.GetInvoicesPaymentStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoicePaymentStatus.GET("/:id", hand.GetInvoicePaymentStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoicePaymentStatus.PUT("/:id", hand.UpdateInvoicePaymentStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoicePaymentStatus.DELETE("/:id", hand.DeleteInvoicePaymentStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
}
