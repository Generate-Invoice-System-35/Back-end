package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

func (r *RepositoryMysqlLayer) CreateInvoicesGenerate(invoice []model.Invoice, item []model.InvoiceItem) error {
	resInv := r.DB.Create(&invoice)
	if resInv.RowsAffected < 1 {
		return fmt.Errorf("error insert invoice")
	}

	resItm := r.DB.Create(&item)
	if resItm.RowsAffected < 1 {
		return fmt.Errorf("error insert invoice item")
	}

	return nil
}

func (r *RepositoryMysqlLayer) NumberInvoiceExists(number string) (invoice model.Invoice, flag bool) {
	res := r.DB.Where("number = ?", number).Find(&invoice)
	flag = false
	if res.RowsAffected > 0 {
		flag = true
	}

	return
}

func (r *RepositoryMysqlLayer) CreateInvoiceGenerate(invoice model.Invoice) error {
	resInv := r.DB.Create(&invoice)
	if resInv.RowsAffected < 1 {
		return fmt.Errorf("error insert invoice")
	}

	return nil
}

func (r *RepositoryMysqlLayer) CreateInvoiceItemsGenerate(item model.InvoiceItem) error {
	resItm := r.DB.Create(&item)
	if resItm.RowsAffected < 1 {
		return fmt.Errorf("error insert invoice item")
	}

	return nil
}

func NewMysqlGenerateRepository(db *gorm.DB) adapter.AdapterGenerateInvoiceRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
