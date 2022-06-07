package repository

import (
	"fmt"

	"Back-end/internal/adapter"
	"Back-end/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateInvoice(invoice model.Invoice) error {
	res := r.DB.Create(&invoice)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllInvoices() []model.Invoice {
	invoices := []model.Invoice{}
	r.DB.Find(&invoices)

	return invoices
}

func (r *RepositoryMysqlLayer) GetInvoiceByID(id int) (invoice model.Invoice, err error) {
	res := r.DB.Where("id = ?", id).Find(&invoice)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateInvoiceByID(id int, invoice model.Invoice) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&invoice)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteInvoiceByID(id int) error {
	res := r.DB.Delete(&model.Invoice{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlInvoiceRepository(db *gorm.DB) adapter.AdapterInvoiceRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
