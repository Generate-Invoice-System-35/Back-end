package usecase

import (
	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

type serviceInvoice struct {
	c    config.Config
	repo adapter.AdapterInvoiceRepository
}

func (s *serviceInvoice) CreateInvoiceService(invoice model.Invoice) error {
	return s.repo.CreateInvoice(invoice)
}

func (s *serviceInvoice) GetAllInvoicesService() []model.Invoice {
	return s.repo.GetAllInvoices()
}

func (s *serviceInvoice) GetInvoiceByIDService(id int) (model.Invoice, error) {
	return s.repo.GetInvoiceByID(id)
}

func (s *serviceInvoice) UpdateInvoiceByIDService(id int, invoice model.Invoice) error {
	return s.repo.UpdateInvoiceByID(id, invoice)
}

func (s *serviceInvoice) DeleteInvoiceByIDService(id int) error {
	return s.repo.DeleteInvoiceByID(id)
}

func NewServiceInvoice(repo adapter.AdapterInvoiceRepository, c config.Config) adapter.AdapterInvoiceService {
	return &serviceInvoice{
		repo: repo,
		c:    c,
	}
}
