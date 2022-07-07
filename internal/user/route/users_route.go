package route

import (
	"Back-end/config"
	d "Back-end/database"
	h "Back-end/internal/user/handler"
	r "Back-end/internal/user/repository"
	u "Back-end/internal/user/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlUserRepository(db)
	service := u.NewServiceUser(repo, conf)
	hand := h.EchoUserController{Service: service}

	apiUser := e.Group("/user",
		middleware.Logger(),
		middleware.CORS(),
	)

	// apiUser.GET("", hand.GetUsersController, middleware.JWTWithConfig(
	// 	middleware.JWTConfig{
	// 		SigningKey: []byte(conf.JWT_KEY),
	// 		ErrorHandlerWithContext: func(err error, c echo.Context) error {
	// 			return c.JSONPretty(404, map[string]interface{}{
	// 				"messages": "token tidak valid",
	// 			}, " ")
	// 		},
	// 		SuccessHandler: func(c echo.Context) {},
	// 	},
	// ))
	apiUser.GET("", hand.GetUsersController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.GET("/:id", hand.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/:id", hand.UpdateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/update/username/:id", hand.ChangeUsernameController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/update/password/:id", hand.ChangePasswordController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/:id", hand.DeleteUsercontroller, middleware.JWT([]byte(conf.JWT_KEY)))
}
