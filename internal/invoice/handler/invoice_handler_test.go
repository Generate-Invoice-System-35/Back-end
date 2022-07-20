package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/invoice/handler"
	"Back-end/internal/invoice/handler/mocks"
	"Back-end/internal/invoice/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInvoiceController(t *testing.T) {
	service := mocks.MockInvoiceService{}
	invoiceController := handler.EchoInvoiceController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("CreateInvoiceService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("POST", "/invoice", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.CreateInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Error Internal Server Error", func(t *testing.T) {
		service.On("CreateInvoiceService", mock.Anything).Return(errors.New("Error Internal Server Error")).Once()

		r := httptest.NewRequest("POST", "/invoice", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.CreateInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetInvoicesController(t *testing.T) {
	service := mocks.MockInvoiceService{}
	invoiceController := handler.EchoInvoiceController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetAllInvoicesService").Return([]model.Invoice{}).Once()

		r := httptest.NewRequest("GET", "/invoice", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoicesController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetInvoicesPaginationController(t *testing.T) {
	service := mocks.MockInvoiceService{}
	invoiceController := handler.EchoInvoiceController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetInvoicesPaginationService", mock.Anything).Return([]model.Invoice{}, nil).Once()

		r := httptest.NewRequest("POST", "/invoice/pagination", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoicesPaginationController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetInvoicesPaginationService", mock.Anything).Return([]model.Invoice{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("POST", "/invoice/pagination", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoicesPaginationController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestGetInvoiceController(t *testing.T) {
	service := mocks.MockInvoiceService{}
	invoiceController := handler.EchoInvoiceController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetInvoiceByIDService", mock.Anything).Return(model.Invoice{}, nil).Once()

		r := httptest.NewRequest("GET", "/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetInvoiceByIDService", mock.Anything).Return(model.Invoice{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestGetInvoicesByPaymentStatusController(t *testing.T) {
	service := mocks.MockInvoiceService{}
	invoiceController := handler.EchoInvoiceController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetInovicesByPaymentStatusService", mock.Anything).Return([]model.Invoice{}, nil).Once()

		r := httptest.NewRequest("GET", "/invoice/status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoicesByPaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetInovicesByPaymentStatusService", mock.Anything).Return([]model.Invoice{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/invoice/status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoicesByPaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestGetInvoicesByNameCustomerController(t *testing.T) {
	service := mocks.MockInvoiceService{}
	invoiceController := handler.EchoInvoiceController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetInvoicesByNameCustomerService", mock.Anything).Return([]model.Invoice{}, nil).Once()

		r := httptest.NewRequest("POST", "/invoice/search", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoicesByNameCustomerController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetInvoicesByNameCustomerService", mock.Anything).Return([]model.Invoice{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("POST", "/invoice/search", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetInvoicesByNameCustomerController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestUpdateInvoiceController(t *testing.T) {
	service := mocks.MockInvoiceService{}
	invoiceController := handler.EchoInvoiceController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("UpdateInvoiceByIDService", mock.Anything, mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("PUT", "/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.UpdateInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("UpdateInvoiceByIDService", mock.Anything, mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("PUT", "/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.UpdateInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestDeleteInvoiceController(t *testing.T) {
	service := mocks.MockInvoiceService{}
	invoiceController := handler.EchoInvoiceController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("DeleteInvoiceByIDService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.DeleteInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server Error", func(t *testing.T) {
		service.On("DeleteInvoiceByIDService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("DELETE", "/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.DeleteInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
