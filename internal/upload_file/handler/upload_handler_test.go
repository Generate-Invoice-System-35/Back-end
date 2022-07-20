package handler_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"Back-end/internal/upload_file/handler"
	"Back-end/internal/upload_file/handler/mocks"
	"Back-end/internal/upload_file/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUploadImageController(t *testing.T) {
	service := mocks.MockUploadService{}
	uploadImageController := handler.EchoUploadImageController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("CreateImageService", mock.Anything, mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("POST", "/upload-image", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.UploadImageController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("CreateImageService", mock.Anything, mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("POST", "/upload-image", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.UploadImageController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetImagesController(t *testing.T) {
	service := mocks.MockUploadService{}
	uploadImageController := handler.EchoUploadImageController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetAllImagesService").Return([]model.File{}).Once()

		r := httptest.NewRequest("GET", "/upload-image", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.GetImagesController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetImageController(t *testing.T) {
	service := mocks.MockUploadService{}
	uploadImageController := handler.EchoUploadImageController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("GetImageByIDService", mock.Anything).Return(model.File{}, nil).Once()

		r := httptest.NewRequest("GET", "/upload-image/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.GetImageController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("GetImageByIDService", mock.Anything).Return(model.File{}, errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("GET", "/upload-image/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.GetImageController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestUpdateImageController(t *testing.T) {
	service := mocks.MockUploadService{}
	uploadImageController := handler.EchoUploadImageController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("UpdateImageService", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("PUT", "/upload-image/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.UpdateImageController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("UpdateImageService", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("PUT", "/upload-image/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.UpdateImageController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestDeleteImageController(t *testing.T) {
	service := mocks.MockUploadService{}
	uploadImageController := handler.EchoUploadImageController{
		Service: &service,
	}
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		service.On("DeleteImageService", mock.Anything).Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/upload-image/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.DeleteImageController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 200, w.Result().StatusCode)
	})

	t.Run("Error Internal Server", func(t *testing.T) {
		service.On("DeleteImageService", mock.Anything).Return(errors.New("Error Internal Server")).Once()

		r := httptest.NewRequest("DELETE", "/upload-image/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := uploadImageController.DeleteImageController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
