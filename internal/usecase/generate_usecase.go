package usecase

import (
	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"
	"strconv"
	"time"
)

type serviceGenerate struct {
	c    config.Config
	repo adapter.AdapterGenerateInvoiceRepository
}

func (s *serviceGenerate) CreateInvoiceGenerateService(data [][]string) error {
	var err error = nil

	for i, line := range data {
		if i > 0 {
			var flag bool
			var recInv model.Invoice
			var recItm model.InvoiceItem

			// INVOICE
			recInv, flag = s.repo.NumberInvoiceExists(line[0])
			if !flag {
				recInv.ID_Payment_Status = 1

				// Number Row
				recInv.Number = line[0]
				// Customer Name Row
				recInv.Name = line[1]
				// Email Row
				recInv.Email = line[2]
				// Phone Number Row
				recInv.Phone_Number = line[3]
				// Address Row
				recInv.Address = line[4]
				// Description Row
				recInv.Description = line[5]
				// Invoice Date Row
				inv_date, _ := time.Parse("2006-01-02", line[6])
				recInv.Invoice_Date = inv_date
				// Due Date Row
				due_date, _ := time.Parse("2006-01-02", line[7])
				recInv.Due_Date = due_date

				recInv.Created_At = time.Now()
				recInv.Updated_At = time.Now()
				err = s.repo.CreateInvoiceGenerate(recInv)
				if err != nil {
					break
				}
			}

			// INVOICE ITEM
			var qty int
			var price float64

			// ID Invoice Item
			recInv, _ = s.repo.NumberInvoiceExists(line[0])
			recItm.ID_Invoice = recInv.ID
			// Invoice Item Product
			recItm.Product = line[8]
			// Invoice Item Label
			recItm.Category = line[9]
			// Invoice Item Quantity
			qty, _ = strconv.Atoi(line[10])
			recItm.Qty = qty
			// Invoice Item Rate
			price, _ = strconv.ParseFloat(line[11], 32)
			recItm.Price = float32(price)
			// Invoice Item SubTotal
			recItm.Subtotal = (recItm.Price * float32(recItm.Qty))

			recItm.Created_At = time.Now()
			recItm.Updated_At = time.Now()

			err = s.repo.CreateInvoiceItemsGenerate(recItm)
			if err != nil {
				break
			}
		}
	}
	return err
}

func NewServiceGenerate(repo adapter.AdapterGenerateInvoiceRepository, c config.Config) adapter.AdapterGenerateInvoiceService {
	return &serviceGenerate{
		repo: repo,
		c:    c,
	}
}
