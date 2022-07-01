package mysql

import (
	"fmt"
	"log"

	"Back-end/internal/model"
	"Back-end/internal/payment_gateway/xendit/adapter"

	"gorm.io/gorm"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
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

func (r *RepositoryMysqlLayer) GetIDInvoicePayment(id int) (record model.TransactionRecord, err error) {
	res := r.DB.Where("id_invoice = ?", id).Find(&record)
	if res.RowsAffected < 1 {
		log.Printf("not found invoice")
		err = fmt.Errorf("not found invoice")
	}

	return
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

func NewMysqlPaymentGatewayRepository(db *gorm.DB) adapter.AdapterPaymentGatewayRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
