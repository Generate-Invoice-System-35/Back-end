package handler

import (
	"net/http"
	"strconv"

	"Back-end/internal/payment_gateway/xendit/adapter"
	"Back-end/internal/payment_gateway/xendit/model"

	"github.com/labstack/echo/v4"
)

const CALLBACK_PUBLIC_KEY = "xnd_public_development_HBTjAJMAM0TbLDQ5mMoITLNUYJi8b3bb4Uge7xtN2zDuu7L8uZycmMikwFf1W"

type EchoPaymentGatewayController struct {
	Service adapter.AdapterPaymentGatewayService
}

// CreateXenditPaymentInvoiceController godoc
// @Summary      Create Payment Invoice Using Xendit
// @Description  User can create payment invoice by using xendit
// @Tags         TransactionRecord
// @accept       json
// @Produce      json
// @Router       /payment/xendit/invoice/{id} [post]
// @param        id   path      int  true  "id"
// @Success      200   {object}  model.TransactionRecord
// @Failure      404   {object}  model.TransactionRecord
// @Security     JWT
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

	return c.JSONPretty(http.StatusCreated, resp, " ")
}

// GetXenditPaymentInvoiceController godoc
// @Summary      Get Xendit Payment Invoice By ID
// @Description  User can get xendit payment invoice by id
// @Tags         TransactionRecord
// @accept       json
// @Produce      json
// @Router       /payment/xendit/invoice/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200   {object}  model.TransactionRecord
// @Failure      404   {object}  model.TransactionRecord
// @Security     JWT
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

// GetAllXenditPaymentInvoiceController godoc
// @Summary      Get All Xendit Payment Invoice
// @Description  User can get all xendit payment invoice
// @Tags         TransactionRecord
// @accept       json
// @Produce      json
// @Router       /payment/xendit/invoice [get]
// @param        data  body      model.TransactionRecord  true  "required"
// @Success      200   {object}  model.TransactionRecord
// @Failure      404   {object}  model.TransactionRecord
// @Security     JWT
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

// ExpireXenditPaymentInvoiceController godoc
// @Summary      Expired Xendit Payment Invoice
// @Description  User can expired xendit payment invoice
// @Tags         TransactionRecord
// @accept       json
// @Produce      json
// @Router       /payment/xendit/invoice/expire/{id} [get]
// @param        id    path      int  true  "id"
// @Success      200   {object}  model.TransactionRecord
// @Failure      404   {object}  model.TransactionRecord
// @Security     JWT
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

// CallbackXenditPaymentInvoiceController godoc
// @Summary      Xendit can Callback By Using This Route
// @Description  Xendit can callback by this route if customer is have been paying or the invoice is expired
// @Tags         CallbackInvoice
// @accept       json
// @Produce      json
// @Router       /payment/xendit/invoice/callback [post]
// @param        data  body      model.CallbackInvoice  true  "required"
// @Success      200   {object}  model.CallbackInvoice
// @Failure      404   {object}  model.CallbackInvoice
// @Security     JWT
func (ce *EchoPaymentGatewayController) CallbackXenditPaymentInvoiceController(c echo.Context) error {
	invoiceCallback := model.CallbackInvoice{}
	c.Bind(&invoiceCallback)

	err := ce.Service.CallbackXenditPaymentInvoiceService(invoiceCallback)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
			"error":    err,
		})
	}

	return c.JSONPretty(http.StatusOK, invoiceCallback, " ")
}
