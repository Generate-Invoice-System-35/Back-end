package handler

import (
	"Back-end/internal/adapter"
	"Back-end/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoSendCustomerController struct {
	Service adapter.AdapterSendCustomerService
}

func (ce *EchoSendCustomerController) SendEmailController(c echo.Context) error {
	messages := model.SendCustomer{}
	c.Bind(&messages)

	err := ce.Service.SendEmailService(messages)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
			"error":    err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"send":     messages,
	})
}

func (ce *EchoSendCustomerController) SendWhatsappController(c echo.Context) error {
	messages := model.SendCustomer{}
	c.Bind(&messages)

	err := ce.Service.SendWhatsappService(messages)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
			"error":    err,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"send":     messages,
	})
}
