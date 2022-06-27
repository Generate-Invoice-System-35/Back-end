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
				// Buyer Row
				recInv.Buyer_Name = line[1]
				// Invoice Date Row
				inv_date, _ := time.Parse("2006-01-02", line[2])
				recInv.Invoice_Date = inv_date
				// Due Date Row
				due_date, _ := time.Parse("2006-01-02", line[3])
				recInv.Due_Date = due_date

				err = s.repo.CreateInvoiceGenerate(recInv)
				if err != nil {
					break
				}
			}

			// INVOICE ITEM
			recInv, _ = s.repo.NumberInvoiceExists(line[0])
			var qty int
			var rate, tax float64

			// ID Invoice Item
			recItm.ID_Invoice = recInv.ID
			// Invoice Item Product
			recItm.Product = line[4]
			// Invoice Item Label
			recItm.Label = line[5]
			// Invoice Item Quantity
			qty, _ = strconv.Atoi(line[6])
			recItm.Qty = qty
			// Invoice Item Rate
			rate, _ = strconv.ParseFloat(line[7], 32)
			recItm.Rate = float32(rate)
			// Invoice Item Tax
			tax, _ = strconv.ParseFloat(line[8], 32)
			recItm.Tax = float32(tax)
			// Invoice Item SubTotal
			tempTotal := (recItm.Rate * float32(recItm.Qty))
			recItm.Subtotal = tempTotal + (tempTotal * (recItm.Tax / 100))

			err = s.repo.CreateInvoiceItemsGenerate(recItm)
			if err != nil {
				break
			}
		}
	}
	return err
}

// func (s *serviceGenerate) CreateInvoiceGenerateService(data [][]string) error {
// 	var err error = nil

// 	for i, line := range data {
// 		if i > 0 {
// 			var flag bool
// 			var recInv model.Invoice
// 			var recItm model.InvoiceItem

// 			recInv, flag = s.repo.NumberInvoiceExists(line[0])
// 			log.Print(flag)
// 			if !flag {
// 				recInv.ID_Payment_Status = 1

// 				// Number Row
// 				recInv.Number = line[0]
// 				// Buyer Row
// 				recInv.Buyer_Name = line[1]
// 				// Invoice Date Row
// 				inv_date, _ := time.Parse("2006-01-02", line[2])
// 				recInv.Invoice_Date = inv_date
// 				// Due Date Row
// 				due_date, _ := time.Parse("2006-01-02", line[3])
// 				recInv.Due_Date = due_date

// 				err = s.repo.CreateInvoiceGenerate(recInv)
// 				if err != nil {
// 					break
// 				}
// 			}

// 			for j, field := range line {
// 				recInv, _ = s.repo.NumberInvoiceExists(line[0])
// 				var qty int
// 				var rate, tax float64

// 				recItm.ID_Invoice = recInv.ID
// 				switch j {
// 				case 4:
// 					recItm.Product = field
// 				case 5:
// 					recItm.Label = field
// 				case 6:
// 					qty, _ = strconv.Atoi(field)
// 					recItm.Qty = qty
// 				case 7:
// 					rate, _ = strconv.ParseFloat(field, 32)
// 					recItm.Rate = float32(rate)
// 				case 8:
// 					tax, _ = strconv.ParseFloat(field, 32)
// 					recItm.Tax = float32(tax)
// 				}

// 				tempTotal := (recItm.Rate * float32(recItm.Qty))
// 				recItm.Subtotal = tempTotal + (tempTotal * (recItm.Tax / 100))
// 			}
// 			err = s.repo.CreateInvoiceItemsGenerate(recItm)
// 			if err != nil {
// 				break
// 			}
// 		}
// 	}
// 	return err
// }

func NewServiceGenerate(repo adapter.AdapterGenerateInvoiceRepository, c config.Config) adapter.AdapterGenerateInvoiceService {
	return &serviceGenerate{
		repo: repo,
		c:    c,
	}
}
