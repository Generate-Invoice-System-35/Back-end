package repository

import (
	"fmt"

	"Back-end/internal/invoice/adapter"
	"Back-end/internal/invoice/model"

	"gorm.io/gorm"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}

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

func (r *RepositoryMysqlLayer) GetInvoicesPagination(pagination model.Pagination) (invoice []model.Invoice, err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := r.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	res := queryBuilder.Model(&model.Invoice{}).Where(invoice).Find(&invoice)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) GetInvoicesByPaymentStatus(status int) (invoice []model.Invoice, err error) {
	res := r.DB.Where("id_payment_status = ?", status).Find(&invoice)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("status not found")
	}

	return
}

func (r *RepositoryMysqlLayer) GetInvoicesByNameCustomer(name string) (invoice []model.Invoice, err error) {
	res := r.DB.Where("name like ?", "%"+name+"%").Find(&invoice)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("invoices not found")
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
