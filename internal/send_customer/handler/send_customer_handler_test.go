package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/send_customer/handler"
	"Back-end/internal/send_customer/handler/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendEmailController(t *testing.T) {
	service := mocks.MockSendCustomerService{}
	sendCustomerController := handler.EchoSendCustomerController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("SendEmailService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("POST", "/send/email", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := sendCustomerController.SendEmailController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("SendEmailService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("POST", "/send/email", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := sendCustomerController.SendEmailController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
