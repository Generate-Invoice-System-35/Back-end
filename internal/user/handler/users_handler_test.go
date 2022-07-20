package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/user/handler"
	"Back-end/internal/user/handler/mocks"
	"Back-end/internal/user/model"

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

	t.Run("Success", func(t *testing.T) {
		service.On("GetAllUsersService").Return([]model.User{}).Once()

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

	t.Run("Success", func(t *testing.T) {
		service.On("GetUserByIDService", mock.Anything).Return(model.User{}, nil).Once()

		r := httptest.NewRequest("GET", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.GetUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("GetUserByIDService", mock.Anything).Return(model.User{}, errors.New("Error Not Found")).Once()

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
		service.On("UpdateUserByIDService", mock.Anything, mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("PUT", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.UpdateUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("UpdateUserByIDService", mock.Anything, mock.Anything).Return(errors.New("Error Not Found")).Once()

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
		service.On("DeleteUserByIDService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/user/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.DeleteUsercontroller(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("DeleteUserByIDService", mock.Anything).Return(errors.New("Error Not Found")).Once()

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

func TestChangeUsernameController(t *testing.T) {
	service := mocks.MockUserService{}
	userController := handler.EchoUserController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("UpdateUsernameService", mock.Anything, mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("PUT", "/user/update/username/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.ChangeUsernameController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("UpdateUsernameService", mock.Anything, mock.Anything).Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("PUT", "/user/update/username/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.ChangeUsernameController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})
}

func TestChangePasswordController(t *testing.T) {
	service := mocks.MockUserService{}
	userController := handler.EchoUserController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("UpdatePasswordService", mock.Anything, mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("PUT", "/user/update/password/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.ChangePasswordController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Not Found", func(t *testing.T) {
		service.On("UpdatePasswordService", mock.Anything, mock.Anything).Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("PUT", "/user/update/password/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.ChangePasswordController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})
}
