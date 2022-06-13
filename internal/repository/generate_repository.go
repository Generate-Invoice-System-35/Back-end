package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

func (r *RepositoryMysqlLayer) CreateInvoiceGenerate(invoice []model.Invoice, item []model.InvoiceItem) error {
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

func NewMysqlGenerateRepository(db *gorm.DB) adapter.AdapterGenerateInvoiceRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
