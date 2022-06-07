package adapter

import (
	"Back-end/internal/model"
	"mime/multipart"
)

type AdapterUploadRepository interface {
	UploadImage(image model.File) error
}

type AdapterUploadService interface {
	UploadImageService(image model.File, file *multipart.FileHeader) error
}
