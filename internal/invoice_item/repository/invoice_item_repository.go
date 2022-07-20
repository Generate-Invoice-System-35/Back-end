package repository

import (
	"fmt"

	invoice "Back-end/internal/invoice/model"
	"Back-end/internal/invoice_item/adapter"
	"Back-end/internal/invoice_item/model"

	"gorm.io/gorm"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *RepositoryMysqlLayer) CreateInvoiceItem(item model.InvoiceItem) error {
	res := r.DB.Create(&item)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllInvoiceItems() []model.InvoiceItem {
	items := []model.InvoiceItem{}
	r.DB.Find(&items)

	return items
}

func (r *RepositoryMysqlLayer) GetInvoiceItemByID(id int) (item model.InvoiceItem, err error) {
	res := r.DB.Where("id = ?", id).Find(&item)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) GetInvoiceItemByNumber(number string) (items []model.InvoiceItem, err error) {
	var inv invoice.Invoice
	res1 := r.DB.Where("number = ?", number).Find(&inv)
	if res1.RowsAffected < 1 {
		err = fmt.Errorf("number invoice not found")
	}

	res2 := r.DB.Where("id_invoice = ?", inv.ID).Find(&items)
	if res2.RowsAffected < 1 {
		err = fmt.Errorf("id invoice not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateInvoiceItemByID(id int, item model.InvoiceItem) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&item)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteInvoiceItemByID(id int) error {
	res := r.DB.Delete(&model.InvoiceItem{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlInvoiceItemRepository(db *gorm.DB) adapter.AdapterInvoiceItemRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
