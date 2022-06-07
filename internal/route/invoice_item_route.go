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

func RegisterInvoiceItemGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlInvoiceItemRepository(db)
	service := u.NewServiceInvoiceItem(repo, conf)
	hand := h.EchoInvoiceItemController{Service: service}

	apiInvoiceItem := e.Group("/invoice-item",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiInvoiceItem.POST("", hand.CreateInvoiceItemController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoiceItem.GET("", hand.GetInvoiceItemsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoiceItem.GET("/:id", hand.GetInvoiceItemController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoiceItem.PUT("/:id", hand.UpdateInvoiceItemController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiInvoiceItem.DELETE("/:id", hand.DeleteInvoiceItemController, middleware.JWT([]byte(conf.JWT_KEY)))
}
