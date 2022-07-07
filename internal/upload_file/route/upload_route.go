package route

import (
	"Back-end/config"
	d "Back-end/database"
	h "Back-end/internal/upload_file/handler"
	r "Back-end/internal/upload_file/repository"
	u "Back-end/internal/upload_file/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
