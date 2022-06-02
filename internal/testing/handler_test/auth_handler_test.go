package mocks

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/handler"
	"Back-end/internal/testing/handler_test/mocks"

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
		service.On("RegisterService", mock.Anything).
			Return(nil, 201).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.RegisterController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Failed Expectation Failed", func(t *testing.T) {
		service.On("RegisterService", mock.Anything).
			Return(errors.New("Failed Register Controller"), 417).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.RegisterController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 417, w.Result().StatusCode)
	})

	t.Run("Failed Internal Server Error", func(t *testing.T) {
		service.On("RegisterService", mock.Anything).
			Return(errors.New("Failed Register Controller"), 500).Once()

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
		service.On("LoginService", mock.Anything, mock.Anything).
			Return("token", 200).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.LoginController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Failed Unauthorized", func(t *testing.T) {
		service.On("LoginService", mock.Anything, mock.Anything).
			Return("token", 401).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.LoginController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 401, w.Result().StatusCode)
	})

	t.Run("Failed Internal Server", func(t *testing.T) {
		service.On("LoginService", mock.Anything, mock.Anything).
			Return("token", 500).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := authController.LoginController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
