package route

import (
	"Back-end/config"
	d "Back-end/database"
	h "Back-end/internal/generate/handler"
	r "Back-end/internal/generate/repository"
	u "Back-end/internal/generate/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterGenerateInvoiceGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlGenerateRepository(db)
	service := u.NewServiceGenerate(repo, conf)
	hand := h.EchoUploadCSVController{Service: service}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/generate/file", hand.GenerateFileController, middleware.JWT([]byte(conf.JWT_KEY)))
	e.POST("/generate/invoices", hand.GenerateInvoicesController, middleware.JWT([]byte(conf.JWT_KEY)))
}
