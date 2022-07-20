package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/invoice_item/handler"
	"Back-end/internal/invoice_item/handler/mocks"
	"Back-end/internal/invoice_item/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInvoiceItemController(t *testing.T) {
	service := mocks.MockInvoiceItemService{}
	invoiceitemController := handler.EchoInvoiceItemController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("CreateInvoiceItemService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("POST", "/invoice-item", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.CreateInvoiceItemController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("CreateInvoiceItemService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("POST", "/invoice-item", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.CreateInvoiceItemController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetInvoiceItemsController(t *testing.T) {
	service := mocks.MockInvoiceItemService{}
	invoiceitemController := handler.EchoInvoiceItemController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetAllInvoiceItemsService").Return([]model.InvoiceItem{}).Once()

		r := httptest.NewRequest("GET", "/invoice-item", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.GetInvoiceItemsController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetInvoiceItemController(t *testing.T) {
	service := mocks.MockInvoiceItemService{}
	invoiceitemController := handler.EchoInvoiceItemController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetInvoiceItemByIDService", mock.Anything).Return(model.InvoiceItem{}, nil).Once()

		r := httptest.NewRequest("GET", "/invoice-item/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.GetInvoiceItemController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetInvoiceItemByIDService", mock.Anything).Return(model.InvoiceItem{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/invoice-item/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.GetInvoiceItemController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestGetInvoiceItemsByNumberController(t *testing.T) {
	service := mocks.MockInvoiceItemService{}
	invoiceitemController := handler.EchoInvoiceItemController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetInvoiceItemByNumberService", mock.Anything).Return([]model.InvoiceItem{}, nil).Once()

		r := httptest.NewRequest("GET", "/invoice-item/number/:number", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.GetInvoiceItemsByNumberController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetInvoiceItemByNumberService", mock.Anything).Return([]model.InvoiceItem{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/invoice-item/number/:number", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.GetInvoiceItemsByNumberController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestUpdateInvoiceItemController(t *testing.T) {
	service := mocks.MockInvoiceItemService{}
	invoiceitemController := handler.EchoInvoiceItemController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("UpdateInvoiceItemByIDService", mock.Anything, mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("PUT", "/invoice-item/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.UpdateInvoiceItemController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server Error", func(t *testing.T) {
		service.On("UpdateInvoiceItemByIDService", mock.Anything, mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("PUT", "/invoice-item/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.UpdateInvoiceItemController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestDeleteInvoiceItemController(t *testing.T) {
	service := mocks.MockInvoiceItemService{}
	invoiceitemController := handler.EchoInvoiceItemController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("DeleteInvoiceItemByIDService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/invoice-item/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.DeleteInvoiceItemController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("DeleteInvoiceItemByIDService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("DELETE", "/invoice-item/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceitemController.DeleteInvoiceItemController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
