package handler

import (
	"net/http"
	"strconv"

	"Back-end/internal/adapter"
	"Back-end/internal/model"

	"github.com/labstack/echo/v4"
)

type EchoInvoiceItemController struct {
	Service adapter.AdapterInvoiceItemService
}

// CreateInvoiceItemController godoc
// @Summary      Create Invoice Item
// @Description  User can create invoice item for the detail of invoice
// @Tags         Invoice Item
// @accept       json
// @Produce      json
// @Router       /invoice-item [post]
// @param        data  body      model.InvoiceItem  true  "required"
// @Success      201   {object}  model.InvoiceItem
// @Failure      500   {object}  model.InvoiceItem
// @Security     JWT
func (ce *EchoInvoiceItemController) CreateInvoiceItemController(c echo.Context) error {
	item := model.InvoiceItem{}
	c.Bind(&item)

	err := ce.Service.CreateInvoiceItemService(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages":     "success",
		"invoice item": item,
	})
}

// GetInvoiceItemsController godoc
// @Summary      Get All Invoice Item Information
// @Description  User can get all invoices item information that has been created from user itself
// @Tags         Invoice Item
// @accept       json
// @Produce      json
// @Router       /invoice-item [get]
// @Success      200  {object}  model.InvoiceItem
// @Security     JWT
func (ce *EchoInvoiceItemController) GetInvoiceItemsController(c echo.Context) error {
	items := ce.Service.GetAllInvoiceItemsService()

	return c.JSONPretty(http.StatusOK, items, " ")
}

// GetInvoiceItemController godoc
// @Summary      Get Invoice Item Information by Id
// @Description  User can get invoice item information by id
// @Tags         Invoice Item
// @accept       json
// @Produce      json
// @Router       /invoice-item/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.InvoiceItem
// @Failure      404  {object}  model.InvoiceItem
// @Security     JWT
func (ce *EchoInvoiceItemController) GetInvoiceItemController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	item, err := ce.Service.GetInvoiceItemByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSONPretty(http.StatusOK, item, " ")
}

// UpdateInvoiceItemController godoc
// @Summary      Update Invoice Item Information
// @Description  User can update invoice item information
// @Tags         Invoice Item
// @accept       json
// @Produce      json
// @Router       /invoice-item/{id} [put]
// @param        id   path      int  true  "id"
// @param        data body      model.InvoiceItem  true  "required"
// @Success      200  {object}  model.InvoiceItem
// @Failure      500  {object}  model.InvoiceItem
// @Security     JWT
func (ce *EchoInvoiceItemController) UpdateInvoiceItemController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	item := model.InvoiceItem{}
	c.Bind(&item)

	err := ce.Service.UpdateInvoiceItemByIDService(intID, item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":     "edited",
		"id":           intID,
		"invoice item": item,
	})
}

// DeleteInvoiceItemController godoc
// @Summary      Delete Invoice Item Information
// @Description  User can delete invoice item information if they want it
// @Tags         Invoice Item
// @accept       json
// @Produce      json
// @Router       /invoice-item/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.InvoiceItem
// @Failure      500  {object}  model.InvoiceItem
// @Security     JWT
func (ce *EchoInvoiceItemController) DeleteInvoiceItemController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteInvoiceItemByIDService(intID)
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
