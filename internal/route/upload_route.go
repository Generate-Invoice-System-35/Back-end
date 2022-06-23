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

func RegisterGenerateInvoiceGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlGenerateRepository(db)
	service := u.NewServiceGenerate(repo, conf)
	hand := h.EchoUploadCSVController{Service: service}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/generate", hand.GenerateInvoicesController)
}

func RegisterUploadImageGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlUploadRepository(db)
	service := u.NewServiceUpload(repo, conf)
	hand := h.EchoUploadImageController{Service: service}

	apiImage := e.Group("/upload-image",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiImage.POST("", hand.UploadImageController)
	apiImage.GET("", hand.GetImagesController)
	apiImage.GET("/:id", hand.GetImageController)
	apiImage.PUT("/:id", hand.UpdateImageController)
	apiImage.DELETE("/:id", hand.DeleteImageController)
}
