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

func RegisterUploadGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlUploadRepository(db)
	service := u.NewServiceUpload(repo, conf)
	hand := h.EchoUploadController{Service: service}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "storage")
	e.POST("/upload-image", hand.UploadImageController)
}
