package repository

import (
	"fmt"
	"log"

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

func (r *RepositoryMysqlLayer) CreateTransactionRecord(id int, record model.TransactionRecord) error {
	var transaction model.TransactionRecord
	res1 := r.DB.Where("id_invoice = ?", id).Find(&transaction)
	if res1.RowsAffected < 1 {
		res2 := r.DB.Create(&record)
		if res2.RowsAffected < 1 {
			log.Print("error insert")
			return fmt.Errorf("error insert")
		}
	} else {
		res3 := r.DB.Where("id_invoice = ?", id).UpdateColumns(&record)
		if res3.RowsAffected < 1 {
			log.Print("error update")
			return fmt.Errorf("error update")
		}
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetInvoices(id int) (inv model.Invoice, items []model.InvoiceItem, err error) {
	res1 := r.DB.Where("id = ?", id).Find(&inv)
	if res1.RowsAffected < 1 {
		log.Printf("not found invoice")
		err = fmt.Errorf("not found invoice")
	}

	res2 := r.DB.Where("id_invoice = ?", id).Find(&items)
	if res2.RowsAffected < 1 {
		log.Printf("not found invoice items")
		err = fmt.Errorf("not found invoice items")
	}

	return
}

func (r *RepositoryMysqlLayer) GetTotalAmount(id int) (float32, error) {
	var err error = nil
	var total float32 = 0
	var items []model.InvoiceItem

	res := r.DB.Where("id_invoice = ?", id).Find(&items)
	if res.RowsAffected < 1 {
		log.Printf("not found invoice items")
		err = fmt.Errorf("not found invoice items")
	}

	for i := 0; i < len(items); i++ {
		total += items[i].Subtotal
	}

	return total, err
}

func (r *RepositoryMysqlLayer) UpdateStatusInvoice(id int, invoice model.Invoice) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&invoice)
	if res.RowsAffected < 1 {
		log.Print("Error Update")
		return fmt.Errorf("error update")
	}

	return nil
}

func NewMysqlGenerateRepository(db *gorm.DB) adapter.AdapterGenerateInvoiceRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
