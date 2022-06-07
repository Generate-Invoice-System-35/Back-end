package repository

import (
	"fmt"

	"Back-end/internal/adapter"
	"Back-end/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateInvoicePaymentStatus(status model.InvoicePaymentStatus) error {
	res := r.DB.Create(&status)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllInvoicesPaymentStatus() []model.InvoicePaymentStatus {
	invoice_payment_status := []model.InvoicePaymentStatus{}
	r.DB.Find(&invoice_payment_status)

	return invoice_payment_status
}

func (r *RepositoryMysqlLayer) GetInvoicePaymentStatusByID(id int) (status model.InvoicePaymentStatus, err error) {
	res := r.DB.Where("id = ?", id).Find(&status)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateInvoicePaymentStatusByID(id int, status model.InvoicePaymentStatus) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&status)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteInvoicePaymentStatusByID(id int) error {
	res := r.DB.Delete(&model.InvoicePaymentStatus{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewmYsqlInvoicePaymentStatusRepository(db *gorm.DB) adapter.AdapterPaymentStatusRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
