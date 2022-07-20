package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/auth/handler"
	"Back-end/internal/auth/handler/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterController(t *testing.T) {
	service := mocks.MockAuthService{}
	authController := handler.EchoAuthController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("RegisterService", mock.Anything).Return(201, nil).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.RegisterController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Error Expectation Failed", func(t *testing.T) {
		service.On("RegisterService", mock.Anything).Return(417, errors.New("Error Expectation Failed")).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.RegisterController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 417, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("RegisterService", mock.Anything).Return(500, errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.RegisterController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestLoginController(t *testing.T) {
	service := mocks.MockAuthService{}
	authController := handler.EchoAuthController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("LoginService", mock.Anything, mock.Anything).Return("token testing", 201).Once()

		r := httptest.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.LoginController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Unauthorized", func(t *testing.T) {
		service.On("LoginService", mock.Anything, mock.Anything).Return("token testing", 401).Once()

		r := httptest.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.LoginController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 401, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("LoginService", mock.Anything, mock.Anything).Return("token testing", 500).Once()

		r := httptest.NewRequest("POST", "/login", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.LoginController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
