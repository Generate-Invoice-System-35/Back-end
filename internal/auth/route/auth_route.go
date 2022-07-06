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

func RegisterAuthGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlAuthRepository(db)
	service := u.NewServiceAuth(repo, conf)
	hand := h.EchoAuthController{Service: service}

	e.POST("/register", hand.RegisterController, middleware.Logger(), middleware.CORS())
	e.POST("/login", hand.LoginController, middleware.Logger(), middleware.CORS())
}
