package usecase

import (
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"time"

	"Back-end/config"
	"Back-end/internal/adapter"
	"Back-end/internal/model"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

const XENDIT_SECRET_KEY = "xnd_development_4aoNFVAmHyBXgOeFJ01dl8a3S5s1snPh02uMVEqKpiXD4NaxUe7xCClxqcCW6"

type serviceGenerate struct {
	c    config.Config
	repo adapter.AdapterGenerateInvoiceRepository
}

func (s *serviceGenerate) GenerateFileService(data [][]string) error {
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

func (s *serviceGenerate) GenerateInvoiceService(ids []int) error {
	for i := 0; i < len(ids); i++ {
		/* <========== CREATE INVOICE PAYMENT ==========> */
		inv, invItem, errRepo := s.repo.GetInvoices(ids[i])
		if errRepo != nil {
			log.Print(errRepo)
			return errRepo
		}

		if inv.ID_Payment_Status == 3 {
			return fmt.Errorf("invoice already paid")
		}

		xendit.Opt.SecretKey = XENDIT_SECRET_KEY

		customer := xendit.InvoiceCustomer{
			GivenNames:   inv.Name,
			Email:        inv.Email,
			MobileNumber: inv.Number,
			Address:      inv.Address,
		}

		var items []xendit.InvoiceItem
		for i := 0; i < len(invItem); i++ {
			item := xendit.InvoiceItem{
				Name:     invItem[i].Product,
				Quantity: invItem[i].Qty,
				Price:    float64(invItem[i].Price),
				Category: invItem[i].Category,
				Url:      "",
			}
			items = append(items, item)
		}

		var fees []xendit.InvoiceFee
		fee := xendit.InvoiceFee{
			Type:  "ADMIN",
			Value: 5000,
		}
		fees = append(fees, fee)

		NotificationType := []string{"whatsapp", "email", "sms"}

		customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  NotificationType,
			InvoiceReminder: NotificationType,
			InvoicePaid:     NotificationType,
			InvoiceExpired:  NotificationType,
		}

		total, error := s.repo.GetTotalAmount(inv.ID)
		if error != nil {
			return error
		}
		for i := 0; i < len(fees); i++ {
			total += float32(fees[i].Value)
		}

		data := invoice.CreateParams{
			ExternalID:                     strconv.Itoa(inv.ID),
			Amount:                         float64(total),
			Description:                    inv.Description,
			InvoiceDuration:                86400,
			Customer:                       customer,
			CustomerNotificationPreference: customerNotificationPreference,
			SuccessRedirectURL:             "https://http.cat/200",
			FailureRedirectURL:             "https://http.cat/406",
			Currency:                       "IDR",
			Items:                          items,
			Fees:                           fees,
		}

		resp, err := invoice.Create(&data)
		if err != nil {
			log.Print(err)
			return err
		}

		var statusInvoice model.Invoice
		statusInvoice.ID_Payment_Status = 2
		s.repo.UpdateStatusInvoice(inv.ID, statusInvoice)

		transaction := model.TransactionRecord{
			ID_Invoice:         inv.ID,
			ID_Invoice_Payment: resp.ID,
			ID_User_Payment:    resp.UserID,
			Created_At:         time.Now(),
			Updated_At:         time.Now(),
		}

		errTransaction := s.repo.CreateTransactionRecord(inv.ID, transaction)
		if errTransaction != nil {
			log.Print(err)
			return err
		}

		/* <========== SEND EMAIL TO CUSTOMER ==========> */
		server := "smtp-mail.outlook.com"
		port := 587
		user := "fattureinvoices35@outlook.com"
		from := user
		pass := "FattureInvoices123456789"
		dest := inv.Email

		auth := LoginAuth(user, pass)

		to := []string{dest}

		subject := "Invoice (Ref INV/2022/" + inv.Number
		body := "Hi " + inv.Name + ",\n\n" +
			"We hope you're well. Please see attached invoice number " +
			inv.Number + ", due on " + inv.Due_Date.String() +
			". Don't hesitate to reach out if you have any questions.\n" +
			"Here the link link payment : " + resp.InvoiceURL +
			"\n\n Kind regards,\n\nFatture Invoices"

		message := []byte("From: " + from + "\n" +
			"To: " + dest + "\n" +
			"Subject: " + subject + "\n\n" +
			body)

		endpoint := fmt.Sprintf("%v:%v", server, port)
		errSendEmail := smtp.SendMail(endpoint, auth, from, to, message)
		if errSendEmail != nil {
			log.Fatal(err)
		}

		msg := model.SendCustomer{
			To:         dest,
			Subject:    subject,
			Body:       body,
			Created_At: time.Now(),
			Updated_At: time.Now(),
		}
		s.repo.SendEmail(msg)
	}

	return nil
}

func NewServiceGenerate(repo adapter.AdapterGenerateInvoiceRepository, c config.Config) adapter.AdapterGenerateInvoiceService {
	return &serviceGenerate{
		repo: repo,
		c:    c,
	}
}
