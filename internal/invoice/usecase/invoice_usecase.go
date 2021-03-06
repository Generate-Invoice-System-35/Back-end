package usecase

import (
	"time"

	"Back-end/config"
	"Back-end/internal/invoice/adapter"
	"Back-end/internal/invoice/model"
)

type serviceInvoice struct {
	c    config.Config
	repo adapter.AdapterInvoiceRepository
}

func (s *serviceInvoice) CreateInvoiceService(invoice model.Invoice) error {
	invoice.Created_At = time.Now()
	invoice.Updated_At = time.Now()
	return s.repo.CreateInvoice(invoice)
}

func (s *serviceInvoice) GetAllInvoicesService() []model.Invoice {
	return s.repo.GetAllInvoices()
}

func (s *serviceInvoice) GetInvoiceByIDService(id int) (model.Invoice, error) {
	return s.repo.GetInvoiceByID(id)
}

func (s *serviceInvoice) GetInvoicesPaginationService(pagination model.Pagination) ([]model.Invoice, error) {
	return s.repo.GetInvoicesPagination(pagination)
}

func (s *serviceInvoice) GetTotalPagesPaginationService() (int, error) {
	invoices, err := s.repo.GetTotalPagesPagination()

	div := invoices / 5
	mod := invoices % 5

	if mod > 0 {
		div++
	}

	return div, err
}

func (s *serviceInvoice) GetInovicesByPaymentStatusService(status int, pagination model.Pagination) []model.Invoice {
	return s.repo.GetInvoicesByPaymentStatus(status, pagination)
}

func (s *serviceInvoice) GetInvoicesByNameCustomerService(name string) ([]model.Invoice, error) {
	return s.repo.GetInvoicesByNameCustomer(name)
}

func (s *serviceInvoice) UpdateInvoiceByIDService(id int, invoice model.Invoice) error {
	invoice.Updated_At = time.Now()
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
