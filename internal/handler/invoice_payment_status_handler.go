package handler

import (
	"net/http"
	"strconv"

	"Back-end/internal/adapter"
	"Back-end/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoInvoicePaymentStatusController struct {
	Service adapter.AdapterPaymentStatusService
}

// CreateInvoicePaymentStatusController godoc
// @Summary      Create Invoice Payment Status
// @Description  Admin can create invoice payment for table invoice
// @Tags         Invoice Payment Status
// @accept       json
// @Produce      json
// @Router       /invoice-payment-status [post]
// @param        data  body      model.InvoicePaymentStatus  true  "required"
// @Success      201   {object}  model.InvoicePaymentStatus
// @Failure      500   {object}  model.InvoicePaymentStatus
// @Security     JWT
func (ce *EchoInvoicePaymentStatusController) CreateInvoicePaymentStatusController(c echo.Context) error {
	invoice_payment_status := model.InvoicePaymentStatus{}
	c.Bind(&invoice_payment_status)

	err := ce.Service.CreateInvoicePaymentStatusService(invoice_payment_status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages":               "success",
		"invoice payment status": invoice_payment_status,
	})
}

// GetInvoicesPaymentStatusController godoc
// @Summary      Get All Invoice Payment Status Information
// @Description  Admin can get all invoices payment status information
// @Tags         Invoice Payment Status
// @accept       json
// @Produce      json
// @Router       /invoice-payment-status [get]
// @Success      200   {object}  model.InvoicePaymentStatus
// @Security     JWT
func (ce *EchoInvoicePaymentStatusController) GetInvoicesPaymentStatusController(c echo.Context) error {
	invoices_payment_status := ce.Service.GetAllInvoicesPaymentStatusService()

	return c.JSONPretty(http.StatusOK, invoices_payment_status, " ")
}

// GetInvoicePaymentStatusController godoc
// @Summary      Get Invoice Payment Status Information by Id
// @Description  Admin can get invoice payment status information by id
// @Tags         Invoice Payment Status
// @accept       json
// @Produce      json
// @Router       /invoice-payment-status/{id} [get]
// @param        id    path      int                         true  "id"
// @Success      200  {object}  model.InvoicePaymentStatus
// @Failure      404  {object}  model.InvoicePaymentStatus
// @Security     JWT
func (ce *EchoInvoicePaymentStatusController) GetInvoicePaymentStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	invoice_payment_status, err := ce.Service.GetInvoicePaymentStatusByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSONPretty(http.StatusOK, invoice_payment_status, " ")
}

// UpdateInvoicePaymentStatusController godoc
// @Summary      Update Invoice Payment Status Information
// @Description  User can update invoice payment status information
// @Tags         Invoice Payment Status
// @accept       json
// @Produce      json
// @Router       /invoice-payment-status/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.InvoicePaymentStatus  true  "required"
// @Success      200  {object}  model.InvoicePaymentStatus
// @Failure      500   {object}  model.InvoicePaymentStatus
// @Security     JWT
func (ce *EchoInvoicePaymentStatusController) UpdateInvoicePaymentStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	invoice_payment_status := model.InvoicePaymentStatus{}
	c.Bind(&invoice_payment_status)

	err := ce.Service.UpdateInvoicePaymentStatusByIDService(intID, invoice_payment_status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":               "edited",
		"id":                     intID,
		"invoice payment status": invoice_payment_status,
	})
}

// DeleteInvoicePaymentStatusController godoc
// @Summary      Delete Invoice Payment Status Information
// @Description  Admin can delete invoice payment status information
// @Tags         Invoice Payment Status
// @accept       json
// @Produce      json
// @Router       /invoice-payment-status/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.InvoicePaymentStatus
// @Failure      500  {object}  model.InvoicePaymentStatus
// @Security     JWT
func (ce *EchoInvoicePaymentStatusController) DeleteInvoicePaymentStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteInvoicePaymentStatusByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "deleted",
		"id":       intID,
	})
}
