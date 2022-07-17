package usecase_test

import (
	"errors"
	"mime/multipart"
	"net/textproto"
	"testing"

	"Back-end/config"
	"Back-end/internal/upload_file/model"
	"Back-end/internal/upload_file/usecase"
	"Back-end/internal/upload_file/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateImageService(t *testing.T) {
	repo := mocks.MockImageRepository{}
	dataImage := model.File{
		File_Name: "Testing Image",
	}
	dataFile := multipart.FileHeader{
		Filename: "main-qimg-55949c34563b674c37450e20ce61fa9b-lq.jpg",
		Header: textproto.MIMEHeader{
			"Content-Disposition": []string{
				"form-data; name=\"file\"; filename=\"main-qimg-55949c34563b674c37450e20ce61fa9b-lq.jpg\"",
			},
			"Content-Type": []string{
				"image/jpeg",
			},
		},
		Size: 36121,
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("CreateImage", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		err := svc.CreateImageService(dataImage, &dataFile)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("CreateImage", mock.Anything).Return(errors.New("Failed Upload Image")).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		err := svc.CreateImageService(dataImage, &dataFile)

		assert.Error(t, err)
	})
}

func TestGetAllImagesService(t *testing.T) {
	repo := mocks.MockImageRepository{}
	data := []model.File{
		{
			ID:        1,
			Name:      "Perocbaan File 8",
			File_Name: "987aaa900530affcf9b4b56bf5180ed874ac25c409ab00dff77911e1d28d6c15.jpg",
			File_Size: 94710,
		},
		{
			ID:        2,
			Name:      "Perocbaan File 10",
			File_Name: "main-qimg-55949c34563b674c37450e20ce61fa9b-lq.jpg",
			File_Size: 36121,
		},
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetAllImages").Return(data).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		images := svc.GetAllImagesService()

		assert.Equal(t, images[0], data[0])
		assert.Equal(t, images[1], data[1])
	})
}

func TestGetImageByIDService(t *testing.T) {
	repo := mocks.MockImageRepository{}
	data := model.File{

		ID:        1,
		Name:      "Perocbaan File 8",
		File_Name: "987aaa900530affcf9b4b56bf5180ed874ac25c409ab00dff77911e1d28d6c15.jpg",
		File_Size: 94710,
	}
	id := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("GetImageByID", mock.Anything).Return(data, nil).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		image, err := svc.GetImageByIDService(id)

		assert.Equal(t, data, image)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetImageByID", mock.Anything).Return(data, errors.New("Failed Get Image By ID")).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		image, err := svc.GetImageByIDService(id)

		assert.Equal(t, data, image)
		assert.Error(t, err)
	})
}

func TestUpdateImageService(t *testing.T) {
	repo := mocks.MockImageRepository{}
	id := 1
	dataImage := model.File{
		File_Name: "Testing Image",
	}
	dataFile := multipart.FileHeader{
		Filename: "main-qimg-55949c34563b674c37450e20ce61fa9b-lq.jpg",
		Header: textproto.MIMEHeader{
			"Content-Disposition": []string{
				"form-data; name=\"file\"; filename=\"main-qimg-55949c34563b674c37450e20ce61fa9b-lq.jpg\"",
			},
			"Content-Type": []string{
				"image/jpeg",
			},
		},
		Size: 36121,
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("UpdateImage", mock.Anything, mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		err := svc.UpdateImageService(id, dataImage, &dataFile)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("UpdateImage", mock.Anything, mock.Anything).Return(errors.New("Failed Update Image")).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		err := svc.UpdateImageService(id, dataImage, &dataFile)

		assert.Error(t, err)
	})
}

func TestDeleteImageService(t *testing.T) {
	repo := mocks.MockImageRepository{}
	id := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("DeleteImage", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		err := svc.DeleteImageService(id)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("DeleteImage", mock.Anything).Return(errors.New("Failed Delete Image")).Once()

		svc := usecase.NewServiceUpload(&repo, config.Config{})
		err := svc.DeleteImageService(id)

		assert.Error(t, err)
	})
}
