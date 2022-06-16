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

func RegisterSendCustomerGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlSendCustomerRepository(db)
	service := u.NewServiceSendCustomer(repo, conf)
	hand := h.EchoSendCustomerController{Service: service}

	apiSend := e.Group("/send",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiSend.POST("/email", hand.SendEmailController)
	apiSend.POST("/whatsapp", hand.SendWhatsappController)
}
