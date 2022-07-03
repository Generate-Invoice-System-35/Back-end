package usecase

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"Back-end/config"
	"Back-end/internal/model"
	"Back-end/internal/payment_gateway/xendit/adapter"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

const XENDIT_SECRET_KEY = "xnd_development_4aoNFVAmHyBXgOeFJ01dl8a3S5s1snPh02uMVEqKpiXD4NaxUe7xCClxqcCW6"

type servicePaymentGateway struct {
	c    config.Config
	repo adapter.AdapterPaymentGatewayRepository
}

func (s *servicePaymentGateway) CreateXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	inv, invItem, errRepo := s.repo.GetInvoices(id)
	if errRepo != nil {
		log.Print(errRepo)
		return nil, errRepo
	}

	if inv.ID_Payment_Status == 3 {
		return nil, fmt.Errorf("invoice already paid")
	}

	xendit.Opt.SecretKey = XENDIT_SECRET_KEY
	// customerAddress := xendit.CustomerAddress{
	// 	Country:     "Indonesia",
	// 	StreetLine1: "Jalan Makan",
	// 	StreetLine2: "Kecamatan Kebayoran Baru",
	// 	City:        "Jakarta Selatan",
	// 	State:       "Daerah Khusus Ibukota Jakarta",
	// 	PostalCode:  "12345",
	// }

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

	// items := xendit.Invoice{
	// 	Items: []xendit.InvoiceItem{item},
	// }

	var fees []xendit.InvoiceFee
	fee := xendit.InvoiceFee{
		Type:  "ADMIN",
		Value: 5000,
	}
	fees = append(fees, fee)

	// fees := xendit.Invoice{
	// 	Fees: []xendit.InvoiceFee{fee},
	// }

	NotificationType := []string{"whatsapp", "email", "sms"}

	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:  NotificationType,
		InvoiceReminder: NotificationType,
		InvoicePaid:     NotificationType,
		InvoiceExpired:  NotificationType,
	}

	total, error := s.repo.GetTotalAmount(inv.ID)
	if error != nil {
		return nil, error
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
		return resp, err
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

	log.Print(transaction)

	errTransaction := s.repo.CreateTransactionRecord(inv.ID, transaction)
	if errTransaction != nil {
		log.Print(err)
		return resp, err
	}

	// fmt.Printf("created invoice: %+v\n", resp)
	return resp, nil
}

func (s *servicePaymentGateway) GetXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	xendit.Opt.SecretKey = XENDIT_SECRET_KEY

	ID, errRepo := s.repo.GetIDInvoicePayment(id)
	if errRepo != nil {
		log.Fatal(errRepo)
		return nil, errRepo
	}

	data := invoice.GetParams{
		ID: ID.ID_Invoice_Payment,
	}

	resp, err := invoice.Get(&data)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	// fmt.Printf("retrieved invoice: %+v\n", resp)
	return resp, nil
}

func (s *servicePaymentGateway) GetAllXenditPaymentInvoiceService() ([]xendit.Invoice, error) {
	xendit.Opt.SecretKey = XENDIT_SECRET_KEY

	// createdAfter, _ := time.Parse(time.RFC3339, "2016-02-24T23:48:36.697Z")
	data := invoice.GetAllParams{
		Statuses: []string{"EXPIRED", "PENDING", "SETTLED"},
		// Limit:        5,
		// CreatedAfter: createdAfter,
	}

	resps, err := invoice.GetAll(&data)
	if err != nil {
		log.Fatal(err)
		return resps, err
	}

	// fmt.Printf("invoices: %+v\n", resps)
	return resps, nil
}

func (s *servicePaymentGateway) ExpireXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	var statusInvoice model.Invoice

	xendit.Opt.SecretKey = XENDIT_SECRET_KEY

	ID, errRepo := s.repo.GetIDInvoicePayment(id)
	if errRepo != nil {
		log.Fatal(errRepo)
		return nil, errRepo
	}

	data := invoice.ExpireParams{
		ID: ID.ID_Invoice_Payment,
	}

	resp, err := invoice.Expire(&data)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	statusInvoice.ID_Payment_Status = 4
	s.repo.UpdateStatusInvoice(id, statusInvoice)

	// fmt.Printf("expired invoice: %+v\n", resp)
	return resp, nil
}

func (s *servicePaymentGateway) CallbackXenditPaymentInvoiceService(callback model.CallbackInvoice) error {
	var statusInvoice model.Invoice

	StringID := callback.ExternalID
	ID, _ := strconv.Atoi(StringID)
	if callback.Status == "PENDING" {
		statusInvoice.ID_Payment_Status = 2
	} else if callback.Status == "PAID" {
		statusInvoice.ID_Payment_Status = 3
	} else if callback.Status == "EXPIRED" {
		statusInvoice.ID_Payment_Status = 4
	} else {
		return fmt.Errorf("status not found")
	}

	return s.repo.UpdateStatusInvoice(ID, statusInvoice)
}

func NewServicePaymentGateway(repo adapter.AdapterPaymentGatewayRepository, c config.Config) adapter.AdapterPaymentGatewayService {
	return &servicePaymentGateway{
		repo: repo,
		c:    c,
	}
}
