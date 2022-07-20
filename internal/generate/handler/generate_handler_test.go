package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/generate/handler"
	"Back-end/internal/generate/handler/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateFileController(t *testing.T) {
	service := mocks.MockGenerateService{}
	generateController := handler.EchoUploadCSVController{
		Service: &service,
	}
	e := echo.New()

	// t.Run("Success", func(t *testing.T) {
	// 	service.On("GenerateFileService", mock.Anything).Return(nil).Once()

	// 	r := httptest.NewRequest("POST", "/generate/file", nil)
	// 	w := httptest.NewRecorder()
	// 	echoContext := e.NewContext(r, w)

	// 	err := generateController.GenerateFileController(echoContext)
	// 	if err != nil {
	// 		return
	// 	}
	// 	assert.Equal(t, 201, w.Result().StatusCode)
	// })

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("GenerateFileService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("POST", "/generate/file", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := generateController.GenerateFileController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGenerateInvoicesController(t *testing.T) {
	service := mocks.MockGenerateService{}
	generateController := handler.EchoUploadCSVController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GenerateInvoiceService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("POST", "/generate/invoices", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := generateController.GenerateInvoicesController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("GenerateInvoiceService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("POST", "/generate/invoices", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := generateController.GenerateInvoicesController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
