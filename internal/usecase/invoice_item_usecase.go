package usecase

import (
	"time"

	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

type serviceInvoiceItem struct {
	c    config.Config
	repo adapter.AdapterInvoiceItemRepository
}

func (s *serviceInvoiceItem) CreateInvoiceItemService(item model.InvoiceItem) error {
	item.Created_At = time.Now()
	item.Updated_At = time.Now()

	return s.repo.CreateInvoiceItem(item)
}

func (s *serviceInvoiceItem) GetAllInvoiceItemsService() []model.InvoiceItem {
	return s.repo.GetAllInvoiceItems()
}

func (s *serviceInvoiceItem) GetInvoiceItemByIDService(id int) (model.InvoiceItem, error) {
	return s.repo.GetInvoiceItemByID(id)
}

func (s *serviceInvoiceItem) UpdateInvoiceItemByIDService(id int, item model.InvoiceItem) error {
	item.Updated_At = time.Now()

	return s.repo.UpdateInvoiceItemByID(id, item)
}

func (s *serviceInvoiceItem) DeleteInvoiceItemByIDService(id int) error {
	return s.repo.DeleteInvoiceItemByID(id)
}

func NewServiceInvoiceItem(repo adapter.AdapterInvoiceItemRepository, c config.Config) adapter.AdapterInvoiceItemService {
	return &serviceInvoiceItem{
		repo: repo,
		c:    c,
	}
}
