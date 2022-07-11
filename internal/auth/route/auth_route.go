package route

import (
	"net/http"

	"Back-end/config"
	d "Back-end/database"
	h "Back-end/internal/auth/handler"
	r "Back-end/internal/auth/repository"
	u "Back-end/internal/auth/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterAuthGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlAuthRepository(db)
	service := u.NewServiceAuth(repo, conf)
	hand := h.EchoAuthController{Service: service}
	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}

	e.POST("/register", hand.RegisterController, middleware.Logger(), middleware.CORSWithConfig(DefaultCORSConfig))
	e.POST("/login", hand.LoginController, middleware.Logger(), middleware.CORSWithConfig(DefaultCORSConfig))
}
