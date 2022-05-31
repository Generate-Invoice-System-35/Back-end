package mocks

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/handler"
	"Back-end/internal/model"
	"Back-end/internal/testing/handler_test/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUsersController(t *testing.T) {
	service := mocks.MockUserService{}
	userController := handler.EchoUserController{
		Service: &service,
	}
	e := echo.New()

	userData := []model.User{
		{
			ID:       1,
			Username: "username testing 1",
			Password: "password testing 1",
		},
		{
			ID:       2,
			Username: "username testing 2",
			Password: "password testing 2",
		},
	}

	t.Run("Success", func(t *testing.T) {
		service.On("GetAllUsersService").
			Return(userData).Once()

		r := httptest.NewRequest("GET", "/user", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.GetUsersController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetUserController(t *testing.T) {
	service := mocks.MockUserService{}
	userController := handler.EchoUserController{
		Service: &service,
	}
	e := echo.New()

	userData := model.User{
		ID:       1,
		Username: "username testing 1",
		Password: "password testing 1",
	}

	t.Run("Success", func(t *testing.T) {
		service.On("GetUserByIDService", mock.Anything).
			Return(userData, nil).Once()

		r := httptest.NewRequest("GET", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.GetUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Failed", func(t *testing.T) {
		service.On("GetUserByIDService", mock.Anything).
			Return(userData, errors.New("Failed Get User Controller")).Once()

		r := httptest.NewRequest("GET", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.GetUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestUpdateUserController(t *testing.T) {
	service := mocks.MockUserService{}
	userController := handler.EchoUserController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("UpdateUserByIDService", mock.Anything, mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("PUT", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.UpdateUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Failed", func(t *testing.T) {
		service.On("UpdateUserByIDService", mock.Anything, mock.Anything).
			Return(errors.New("Failed Update User Controller")).Once()

		r := httptest.NewRequest("PUT", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.UpdateUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestDeleteUsercontroller(t *testing.T) {
	service := mocks.MockUserService{}
	userController := handler.EchoUserController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("DeleteUserByIDService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.DeleteUsercontroller(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Failed", func(t *testing.T) {
		service.On("DeleteUserByIDService", mock.Anything).
			Return(errors.New("Failed Delete User Controller")).Once()

		r := httptest.NewRequest("DELETE", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.DeleteUsercontroller(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})
}
