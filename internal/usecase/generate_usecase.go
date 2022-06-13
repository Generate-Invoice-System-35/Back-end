package usecase

import (
	"strconv"
	"time"

	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

type serviceGenerate struct {
	c    config.Config
	repo adapter.AdapterGenerateInvoiceRepository
}

func (s *serviceGenerate) CreateInvoiceGenerateService(data [][]string) error {
	var invoiceList []model.Invoice
	var itemList []model.InvoiceItem

	for i, line := range data {
		if i > 0 {
			var recInv model.Invoice
			var recItm model.InvoiceItem
			for j, field := range line {
				var qty int
				var rate, tax float64
				recInv.ID_Payment_Status = 2
				switch j {
				case 0:
					recInv.Number = field
				case 1:
					recInv.Buyer_Name = field
				case 2:
					inv_date, _ := time.Parse("2006-01-02", field)
					recInv.Invoice_Date = inv_date
				case 3:
					due_date, _ := time.Parse("2006-01-02", field)
					recInv.Due_Date = due_date
				case 4:
					recItm.Product = field
				case 5:
					recItm.Label = field
				case 6:
					qty, _ = strconv.Atoi(field)
					recItm.Qty = qty
				case 7:
					rate, _ = strconv.ParseFloat(field, 32)
					recItm.Rate = float32(rate)
				case 8:
					tax, _ = strconv.ParseFloat(field, 32)
					recItm.Tax = float32(tax)
				}
				tempTotal := (recItm.Rate * float32(recItm.Qty))
				recItm.Subtotal = tempTotal + (tempTotal * (recItm.Tax / 100))
			}
			invoiceList = append(invoiceList, recInv)
			itemList = append(itemList, recItm)
		}
	}
	return s.repo.CreateInvoiceGenerate(invoiceList, itemList)
}

func NewServiceGenerate(repo adapter.AdapterGenerateInvoiceRepository, c config.Config) adapter.AdapterGenerateInvoiceService {
	return &serviceGenerate{
		repo: repo,
		c:    c,
	}
}
