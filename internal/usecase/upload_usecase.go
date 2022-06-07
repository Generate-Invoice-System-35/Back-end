package usecase

import (
	"mime/multipart"

	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

type serviceUpload struct {
	c    config.Config
	repo adapter.AdapterUploadRepository
}

func (s *serviceUpload) UploadImageService(image model.File, file *multipart.FileHeader) error {
	image.File_Name = file.Filename
	image.File_Size = int(file.Size)

	err := s.repo.UploadImage(image)
	if err != nil {
		return err
	}

	return nil
}

func NewServiceUpload(repo adapter.AdapterUploadRepository, c config.Config) adapter.AdapterUploadService {
	return &serviceUpload{
		repo: repo,
		c:    c,
	}
}
