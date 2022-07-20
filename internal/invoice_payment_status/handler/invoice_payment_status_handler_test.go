package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/invoice_payment_status/handler"
	"Back-end/internal/invoice_payment_status/handler/mocks"
	"Back-end/internal/invoice_payment_status/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInvoicePaymentStatusController(t *testing.T) {
	service := mocks.MockInvoicePaymentStatusService{}
	invoicePaymentStatusController := handler.EchoInvoicePaymentStatusController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("CreateInvoicePaymentStatusService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("POST", "/invoice-payment-status", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.CreateInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("CreateInvoicePaymentStatusService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("POST", "/invoice-payment-status", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.CreateInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetInvoicesPaymentStatusController(t *testing.T) {
	service := mocks.MockInvoicePaymentStatusService{}
	invoicePaymentStatusController := handler.EchoInvoicePaymentStatusController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetAllInvoicesPaymentStatusService").Return([]model.InvoicePaymentStatus{}).Once()

		r := httptest.NewRequest("GET", "/invoice-payment-status", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.GetInvoicesPaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetInvoicePaymentStatusController(t *testing.T) {
	service := mocks.MockInvoicePaymentStatusService{}
	invoicePaymentStatusController := handler.EchoInvoicePaymentStatusController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetInvoicePaymentStatusByIDService", mock.Anything).Return(model.InvoicePaymentStatus{}, nil).Once()

		r := httptest.NewRequest("GET", "/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.GetInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetInvoicePaymentStatusByIDService", mock.Anything).Return(model.InvoicePaymentStatus{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.GetInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestUpdateInvoicePaymentStatusController(t *testing.T) {
	service := mocks.MockInvoicePaymentStatusService{}
	invoicePaymentStatusController := handler.EchoInvoicePaymentStatusController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("UpdateInvoicePaymentStatusByIDService", mock.Anything, mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("PUT", "/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.UpdateInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("UpdateInvoicePaymentStatusByIDService", mock.Anything, mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("PUT", "/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.UpdateInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestDeleteInvoicePaymentStatusController(t *testing.T) {
	service := mocks.MockInvoicePaymentStatusService{}
	invoicePaymentStatusController := handler.EchoInvoicePaymentStatusController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("DeleteInvoicePaymentStatusByIDService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.DeleteInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("DeleteInvoicePaymentStatusByIDService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("DELETE", "/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoicePaymentStatusController.DeleteInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
