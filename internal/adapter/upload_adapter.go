package adapter

import (
	"mime/multipart"

	"Back-end/internal/model"
)

type AdapterGenerateInvoiceRepository interface {
	CreateInvoiceGenerate(invoice []model.Invoice, item []model.InvoiceItem) error
}

type AdapterGenerateInvoiceService interface {
	CreateInvoiceGenerateService(data [][]string) error
}

type AdapterUploadImageRepository interface {
	CreateImage(image model.File) error
	GetAllImages() []model.File
	GetImageByID(id int) (image model.File, urlImage string, err error)
	UpdateImage(id int, image model.File) error
	DeleteImage(id int) error
}

type AdapterUploadImageService interface {
	CreateImageService(image model.File, file *multipart.FileHeader) error
	GetAllImagesService() []model.File
	GetImageByIDService(id int) (image model.File, urlImage string, err error)
	UpdateImageService(id int, image model.File, file *multipart.FileHeader) error
	DeleteImageService(id int) error
}
