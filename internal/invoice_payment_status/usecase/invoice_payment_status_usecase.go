package usecase

import (
	"Back-end/config"
	"Back-end/internal/invoice_payment_status/adapter"
	"Back-end/internal/invoice_payment_status/model"
)

type serviceIPStatus struct {
	c    config.Config
	repo adapter.AdapterPaymentStatusRepository
}

func (s *serviceIPStatus) CreateInvoicePaymentStatusService(status model.InvoicePaymentStatus) error {
	return s.repo.CreateInvoicePaymentStatus(status)
}

func (s *serviceIPStatus) GetAllInvoicesPaymentStatusService() []model.InvoicePaymentStatus {
	return s.repo.GetAllInvoicesPaymentStatus()
}

func (s *serviceIPStatus) GetInvoicePaymentStatusByIDService(id int) (model.InvoicePaymentStatus, error) {
	return s.repo.GetInvoicePaymentStatusByID(id)
}

func (s *serviceIPStatus) UpdateInvoicePaymentStatusByIDService(id int, status model.InvoicePaymentStatus) error {
	return s.repo.UpdateInvoicePaymentStatusByID(id, status)
}

func (s *serviceIPStatus) DeleteInvoicePaymentStatusByIDService(id int) error {
	return s.repo.DeleteInvoicePaymentStatusByID(id)
}

func NewServiceInvoicePaymentStatus(repo adapter.AdapterPaymentStatusRepository, c config.Config) adapter.AdapterPaymentStatusService {
	return &serviceIPStatus{
		repo: repo,
		c:    c,
	}
}
