package handler

import (
	"net/http"
	"strconv"

	"Back-end/internal/payment_gateway/adapter"

	"github.com/labstack/echo/v4"
)

type EchoPaymentGatewayController struct {
	Service adapter.AdapterPaymentGatewayService
}

func (ce *EchoPaymentGatewayController) CreatePaymentInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resp, err := ce.Service.CreateInvoiceService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
			"error":    err,
		})
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}
