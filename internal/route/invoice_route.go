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

func RegisterInvoiceGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlInvoiceRepository(db)
	service := u.NewServiceInvoice(repo, conf)
	hand := h.EchoInvoiceController{Service: service}

	apiInvoice := e.Group("/invoice",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiInvoice.POST("", hand.CreateInvoiceController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoice.GET("", hand.GetInvoicesController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoice.GET("/:id", hand.GetInvoiceController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoice.PUT("/:id", hand.UpdateInvoiceController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoice.DELETE("/:id", hand.DeleteInvoiceController, middleware.JWT([]byte(conf.JWT_KEY)))
}
