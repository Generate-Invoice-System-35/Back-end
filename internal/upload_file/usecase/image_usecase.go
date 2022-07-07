package usecase

import (
	"mime/multipart"
	"time"

	"Back-end/config"
	"Back-end/internal/upload_file/adapter"
	"Back-end/internal/upload_file/model"
)

type serviceUpload struct {
	c    config.Config
	repo adapter.AdapterUploadImageRepository
}

func (s *serviceUpload) CreateImageService(image model.File, file *multipart.FileHeader) error {
	image.File_Name = file.Filename
	image.File_Size = int(file.Size)
	image.Created_At = time.Now()
	image.Updated_At = time.Now()

	err := s.repo.CreateImage(image)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceUpload) GetAllImagesService() []model.File {
	return s.repo.GetAllImages()
}

func (s *serviceUpload) GetImageByIDService(id int) (model.File, error) {
	return s.repo.GetImageByID(id)
}

func (s *serviceUpload) UpdateImageService(id int, image model.File, file *multipart.FileHeader) error {
	image.File_Name = file.Filename
	image.File_Size = int(file.Size)
	image.Updated_At = time.Now()

	err := s.repo.UpdateImage(id, image)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceUpload) DeleteImageService(id int) error {
	return s.repo.DeleteImage(id)
}

func NewServiceUpload(repo adapter.AdapterUploadImageRepository, c config.Config) adapter.AdapterUploadImageService {
	return &serviceUpload{
		repo: repo,
		c:    c,
	}
}
