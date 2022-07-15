package handler

import (
	"net/http"
	"strconv"

	"Back-end/internal/invoice/adapter"
	"Back-end/internal/invoice/model"

	"github.com/labstack/echo/v4"
)

type EchoInvoiceController struct {
	Service adapter.AdapterInvoiceService
}

// CreateInvoiceController godoc
// @Summary      Create Invoice
// @Description  User can create invoice for sent to the client
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /invoice [post]
// @param        data  body      model.Invoice  true  "required"
// @Success      201   {object}  model.Invoice
// @Failure      500   {object}  model.Invoice
// @Security     JWT
func (ce *EchoInvoiceController) CreateInvoiceController(c echo.Context) error {
	invoice := model.Invoice{}
	c.Bind(&invoice)

	err := ce.Service.CreateInvoiceService(invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"invoice": invoice,
	})
}

// GetInvoicesController godoc
// @Summary      Get All Invoices Information
// @Description  User can get all invoices information
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /invoice [get]
// @Success      200   {object}  model.Invoice
// @Security     JWT
func (ce *EchoInvoiceController) GetInvoicesController(c echo.Context) error {
	invoices := ce.Service.GetAllInvoicesService()

	return c.JSONPretty(http.StatusOK, invoices, " ")
}

// GetInvoicesPaginationController godoc
// @Summary      Get Invoices Information By Pagination
// @Description  User can get invoices information by limit, page, and sort settings
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /invoice/pagination [post]
// @Success      200   {object}  model.Invoice
// @Failure      404   {object}  model.Invoice
// @Security     JWT
func (ce *EchoInvoiceController) GetInvoicesPaginationController(c echo.Context) error {
	// Initializing Default
	pagination := model.Pagination{
		Limit: 5,
		Page:  1,
		Sort:  "created_at asc",
	}

	c.Bind(&pagination)
	invoices, err := ce.Service.GetInvoicesPaginationService(pagination)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no id",
		})
	}

	return c.JSONPretty(http.StatusOK, invoices, " ")
}

func (ce *EchoInvoiceController) GetTotalPagesPaginationController(c echo.Context) error {
	pages, err := ce.Service.GetTotalPagesPaginationService()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no invoices",
		})
	}

	return c.JSONPretty(http.StatusOK, pages, " ")
}

// GetInvoiceController godoc
// @Summary      Get Invoice Information by Id
// @Description  User can get invoice information by id
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /invoice/{id} [get]
// @param        id    path      int            true  "id"
// @Success      200  {object}  model.Invoice
// @Failure      404  {object}  model.Invoice
// @Security     JWT
func (ce *EchoInvoiceController) GetInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	invoice, err := ce.Service.GetInvoiceByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no id",
		})
	}

	return c.JSONPretty(http.StatusOK, invoice, " ")
}

// GetInvoicesByPaymentStatusController godoc
// @Summary      Get Invoice Information by Payment Status
// @Description  User can get invoice information by payment status
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /invoice/status/{id} [get]
// @param        id    path      int            true  "id"
// @Success      200  {object}  model.Invoice
// @Security     JWT
func (ce *EchoInvoiceController) GetInvoicesByPaymentStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	// Initializing Default
	pagination := model.Pagination{
		Limit: 5,
		Page:  1,
		Sort:  "created_at asc",
	}
	c.Bind(&pagination)

	invoices := ce.Service.GetInovicesByPaymentStatusService(intID, pagination)
	return c.JSONPretty(http.StatusOK, invoices, " ")
}

// GetInvoicesByNameCustomerController godoc
// @Summary      Get Invoice Information by Search Name Customer
// @Description  User can get invoice information by search name customer
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /invoice/search [post]
// @param        data body      model.Invoice  true  "required"
// @Success      200  {object}  model.Invoice
// @Failure      404  {object}  model.Invoice
// @Security     JWT
func (ce *EchoInvoiceController) GetInvoicesByNameCustomerController(c echo.Context) error {
	type Keyword struct {
		Name string `json:"name" form:"name"`
	}
	key := new(Keyword)
	c.Bind(key)

	invoices, err := ce.Service.GetInvoicesByNameCustomerService(key.Name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "invoices not found",
		})
	}

	return c.JSONPretty(http.StatusOK, invoices, " ")
}

// UpdateInvoiceController godoc
// @Summary      Update Invoice Information
// @Description  User can update invoice information
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /invoice/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      model.Invoice  true  "required"
// @Success      200  {object}  model.Invoice
// @Failure      500   {object}  model.Invoice
// @Security     JWT
func (ce *EchoInvoiceController) UpdateInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	invoice := model.Invoice{}
	c.Bind(&invoice)

	err := ce.Service.UpdateInvoiceByIDService(intID, invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "edited",
		"id":      intID,
		"invoice": invoice,
	})
}

// DeleteInvoiceController godoc
// @Summary      Delete Invoice Information
// @Description  User can delete invoice information if they want it
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /invoice/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.Invoice
// @Failure      500  {object}  model.Invoice
// @Security     JWT
func (ce *EchoInvoiceController) DeleteInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteInvoiceByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
		"id":      intID,
	})
}
