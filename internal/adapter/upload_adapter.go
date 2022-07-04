package adapter

import (
	"mime/multipart"

	"Back-end/internal/model"
)

type AdapterGenerateInvoiceRepository interface {
	CreateInvoiceGenerate(invoice model.Invoice) error
	CreateInvoiceItemsGenerate(item model.InvoiceItem) error
	CreateTransactionRecord(int, model.TransactionRecord) error
	NumberInvoiceExists(number string) (invoice model.Invoice, flag bool)
	CreateInvoicesGenerate(invoice []model.Invoice, item []model.InvoiceItem) error
	GetInvoices(int) (model.Invoice, []model.InvoiceItem, error)
	GetTotalAmount(int) (float32, error)
	UpdateStatusInvoice(int, model.Invoice) error
	SendEmail(message model.SendCustomer) error
}

type AdapterGenerateInvoiceService interface {
	GenerateFileService(data [][]string) error
	GenerateInvoiceService(data []int) error
}

type AdapterUploadImageRepository interface {
	CreateImage(image model.File) error
	GetAllImages() []model.File
	GetImageByID(id int) (image model.File, err error)
	UpdateImage(id int, image model.File) error
	DeleteImage(id int) error
}

type AdapterUploadImageService interface {
	CreateImageService(image model.File, file *multipart.FileHeader) error
	GetAllImagesService() []model.File
	GetImageByIDService(id int) (image model.File, err error)
	UpdateImageService(id int, image model.File, file *multipart.FileHeader) error
	DeleteImageService(id int) error
}
