package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"Back-end/internal/model"
	"Back-end/internal/payment_gateway/xendit/adapter"

	"github.com/labstack/echo/v4"
)

const CALLBACK_PUBLIC_KEY = "xnd_public_development_HBTjAJMAM0TbLDQ5mMoITLNUYJi8b3bb4Uge7xtN2zDuu7L8uZycmMikwFf1W"

type EchoPaymentGatewayController struct {
	Service adapter.AdapterPaymentGatewayService
}

func (ce *EchoPaymentGatewayController) CreateXenditPaymentInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resp, err := ce.Service.CreateXenditPaymentInvoiceService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
			"error":    err,
		})
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}

func (ce *EchoPaymentGatewayController) GetXenditPaymentInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resp, err := ce.Service.GetXenditPaymentInvoiceService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
			"error":    err,
		})
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}

func (ce *EchoPaymentGatewayController) GetAllXenditPaymentInvoiceController(c echo.Context) error {
	resp, err := ce.Service.GetAllXenditPaymentInvoiceService()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
			"error":    err,
		})
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}

func (ce *EchoPaymentGatewayController) ExpireXenditPaymentInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resp, err := ce.Service.ExpireXenditPaymentInvoiceService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
			"error":    err,
		})
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}

func (ce *EchoPaymentGatewayController) CallbackXenditPaymentInvoiceController(c echo.Context) error {
	invoiceCallback := model.CallbackInvoice{}
	c.Bind(&invoiceCallback)
	fmt.Println("Invoice Callback : ", invoiceCallback)

	err := ce.Service.CallbackXenditPaymentInvoiceService(invoiceCallback)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
			"error":    err,
		})
	}

	return c.JSONPretty(http.StatusOK, invoiceCallback, " ")
}
