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

func RegisterRoleGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlRoleRepository(db)
	service := u.NewServiceRole(repo, conf)
	hand := h.EchoRoleController{Service: service}

	apiRole := e.Group("/role",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiRole.POST("", hand.CreateRoleController)
	apiRole.GET("", hand.GetRolesController)
	apiRole.GET("/:id", hand.GetRoleController)
	apiRole.PUT("/:id", hand.UpdateRoleController)
	apiRole.DELETE("/:id", hand.DeleteRoleController)
}
