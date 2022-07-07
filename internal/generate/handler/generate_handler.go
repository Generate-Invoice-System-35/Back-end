package handler

import (
	"encoding/csv"
	"net/http"

	"Back-end/internal/generate/adapter"

	"github.com/labstack/echo/v4"
)

type EchoUploadCSVController struct {
	Service adapter.AdapterGenerateInvoiceService
}

// GenerateFileController godoc
// @Summary      Generate File Invoices
// @Description  User can generate invoice file format csv of excel for sent to the client
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /generate/file [post]
// @param        data  body      model.Invoice  true  "required"
// @Success      201   {object}  model.Invoice
// @Failure      500   {object}  model.Invoice
func (ce *EchoUploadCSVController) GenerateFileController(c echo.Context) error {
	//-----------
	// Read file
	//-----------

	// Source
	file, err1 := c.FormFile("file")
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "upload failed",
			"error":   err1,
		})
	}

	src, err2 := file.Open()
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "open failed",
			"error":   err2,
		})
	}
	defer src.Close()

	// Read CSV Values using csv.Reader
	csvReader := csv.NewReader(src)
	data, err3 := csvReader.ReadAll()
	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "read failed",
			"error":   err3,
		})
	}

	// Convert Records to Array of Struct
	err4 := ce.Service.GenerateFileService(data)
	if err4 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "convert failed",
			"error":   err4,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "generate file success",
	})
}

// GenerateInvoicesController godoc
// @Summary      Generate Invoices
// @Description  User can generate invoice for send to the client
// @Tags         Invoice
// @accept       json
// @Produce      json
// @Router       /generate/invoices [post]
// @param        data  body      model.Invoice  true  "required"
// @Success      201   {object}  model.Invoice
// @Failure      500   {object}  model.Invoice
func (ce *EchoUploadCSVController) GenerateInvoicesController(c echo.Context) error {
	type Data struct {
		IDS []int `json:"ids" form:"ids"`
	}
	datas := new(Data)
	c.Bind(&datas)

	err := ce.Service.GenerateInvoiceService(datas.IDS)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "generate invoices success",
	})
}
