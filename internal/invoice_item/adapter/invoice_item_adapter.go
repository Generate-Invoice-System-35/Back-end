package adapter

import "Back-end/internal/invoice_item/model"

type AdapterInvoiceItemRepository interface {
	CreateInvoiceItem(item model.InvoiceItem) error
	GetAllInvoiceItems() []model.InvoiceItem
	GetInvoiceItemByID(id int) (invItem model.InvoiceItem, err error)
	GetInvoiceItemByNumber(number string) (invItem []model.InvoiceItem, err error)
	UpdateInvoiceItemByID(id int, invItem model.InvoiceItem) error
	DeleteInvoiceItemByID(id int) error
}

type AdapterInvoiceItemService interface {
	CreateInvoiceItemService(item model.InvoiceItem) error
	GetAllInvoiceItemsService() []model.InvoiceItem
	GetInvoiceItemByIDService(id int) (invItem model.InvoiceItem, err error)
	GetInvoiceItemByNumberService(number string) (invItem []model.InvoiceItem, err error)
	UpdateInvoiceItemByIDService(id int, invItem model.InvoiceItem) error
	DeleteInvoiceItemByIDService(id int) error
}
