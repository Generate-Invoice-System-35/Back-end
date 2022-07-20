package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	handler "Back-end/internal/payment_gateway/xendit/handler/http"
	"Back-end/internal/payment_gateway/xendit/handler/http/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
)

func TestCreateXenditPaymentInvoiceController(t *testing.T) {
	service := mocks.MockPaymentGatewayXenditService{}
	invoiceController := handler.EchoPaymentGatewayController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("CreateXenditPaymentInvoiceService", mock.Anything).Return(&xendit.Invoice{}, nil).Once()

		r := httptest.NewRequest("POST", "/payment/xendit/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.CreateXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("CreateXenditPaymentInvoiceService", mock.Anything).Return(&xendit.Invoice{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("POST", "/payment/xendit/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.CreateXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestGetXenditPaymentInvoiceController(t *testing.T) {
	service := mocks.MockPaymentGatewayXenditService{}
	invoiceController := handler.EchoPaymentGatewayController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetXenditPaymentInvoiceService", mock.Anything).Return(&xendit.Invoice{}, nil).Once()

		r := httptest.NewRequest("GET", "/payment/xendit/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetXenditPaymentInvoiceService", mock.Anything).Return(&xendit.Invoice{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/payment/xendit/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestGetAllXenditPaymentInvoiceController(t *testing.T) {
	service := mocks.MockPaymentGatewayXenditService{}
	invoiceController := handler.EchoPaymentGatewayController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetAllXenditPaymentInvoiceService", mock.Anything).Return([]xendit.Invoice{}, nil).Once()

		r := httptest.NewRequest("GET", "/payment/xendit/invoice", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetAllXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		service.On("GetAllXenditPaymentInvoiceService", mock.Anything).Return([]xendit.Invoice{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/payment/xendit/invoice", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.GetAllXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestExpireXenditPaymentInvoiceController(t *testing.T) {
	service := mocks.MockPaymentGatewayXenditService{}
	invoiceController := handler.EchoPaymentGatewayController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("ExpireXenditPaymentInvoiceService", mock.Anything).Return(&xendit.Invoice{}, nil).Once()

		r := httptest.NewRequest("GET", "/payment/xendit/invoice/expire/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.ExpireXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("ExpireXenditPaymentInvoiceService", mock.Anything).Return(&xendit.Invoice{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/payment/xendit/invoice/expire/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.ExpireXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestCallbackXenditPaymentInvoiceController(t *testing.T) {
	service := mocks.MockPaymentGatewayXenditService{}
	invoiceController := handler.EchoPaymentGatewayController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("CallbackXenditPaymentInvoiceService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("POST", "/payment/xendit/invoice/callback", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.CallbackXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("CallbackXenditPaymentInvoiceService", mock.Anything).Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("POST", "/payment/xendit/invoice/callback", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := invoiceController.CallbackXenditPaymentInvoiceController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 404, w.Result().StatusCode)
	})
}
