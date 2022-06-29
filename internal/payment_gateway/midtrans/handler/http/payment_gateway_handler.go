package handler

import (
	"net/http"
	"strconv"

	"Back-end/internal/payment_gateway/midtrans/adapter"

	"github.com/labstack/echo/v4"
)

type EchoPaymentGatewayController struct {
	Service adapter.AdapterPaymentGatewayService
}

func (ce *EchoPaymentGatewayController) ChargeTransactionController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resp, err := ce.Service.ChargeTransactionService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
			"error":    err,
		})
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}
