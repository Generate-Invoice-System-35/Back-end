package repository

import (
	"fmt"
	"log"

	"Back-end/internal/generate/adapter"
	invoice "Back-end/internal/invoice/model"
	item "Back-end/internal/invoice_item/model"
	transaction "Back-end/internal/payment_gateway/xendit/model"
	send "Back-end/internal/send_customer/model"

	"gorm.io/gorm"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *RepositoryMysqlLayer) CreateInvoicesGenerate(invoice []invoice.Invoice, item []item.InvoiceItem) error {
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

func (r *RepositoryMysqlLayer) NumberInvoiceExists(number string) (invoice invoice.Invoice, flag bool) {
	res := r.DB.Where("number = ?", number).Find(&invoice)
	flag = false
	if res.RowsAffected > 0 {
		flag = true
	}

	return
}

func (r *RepositoryMysqlLayer) CreateInvoiceGenerate(invoice invoice.Invoice) error {
	resInv := r.DB.Create(&invoice)
	if resInv.RowsAffected < 1 {
		return fmt.Errorf("error insert invoice")
	}

	return nil
}

func (r *RepositoryMysqlLayer) CreateInvoiceItemsGenerate(item item.InvoiceItem) error {
	resItm := r.DB.Create(&item)
	if resItm.RowsAffected < 1 {
		return fmt.Errorf("error insert invoice item")
	}

	return nil
}

func (r *RepositoryMysqlLayer) CreateTransactionRecord(id int, record transaction.TransactionRecord) error {
	var transaction transaction.TransactionRecord
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

func (r *RepositoryMysqlLayer) GetInvoices(id int) (inv invoice.Invoice, items []item.InvoiceItem, err error) {
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
	var items []item.InvoiceItem

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

func (r *RepositoryMysqlLayer) UpdateStatusInvoice(id int, invoice invoice.Invoice) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&invoice)
	if res.RowsAffected < 1 {
		log.Print("Error Update")
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) SendEmail(message send.SendCustomer) error {
	res := r.DB.Create(&message)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func NewMysqlGenerateRepository(db *gorm.DB) adapter.AdapterGenerateInvoiceRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
