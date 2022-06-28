package usecase

import (
	"log"
	"strconv"

	"Back-end/config"
	"Back-end/internal/payment_gateway/adapter"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type servicePaymentGateway struct {
	c    config.Config
	repo adapter.AdapterPaymentGatewayRepository
}

func (s *servicePaymentGateway) CreateInvoiceService(id int) (*xendit.Invoice, error) {
	inv, invItem, errRepo := s.repo.GetInvoices(id)
	if errRepo != nil {
		log.Print(errRepo)
		return nil, errRepo
	}

	xendit.Opt.SecretKey = ""
	// customerAddress := xendit.CustomerAddress{
	// 	Country:     "Indonesia",
	// 	StreetLine1: "Jalan Makan",
	// 	StreetLine2: "Kecamatan Kebayoran Baru",
	// 	City:        "Jakarta Selatan",
	// 	State:       "Daerah Khusus Ibukota Jakarta",
	// 	PostalCode:  "12345",
	// }

	customer := xendit.InvoiceCustomer{
		GivenNames:   inv.Buyer_Name,
		Email:        "johndoe@example.com",
		MobileNumber: "+6287774441111",
		Address:      "",
	}

	item := xendit.InvoiceItem{
		Name:     invItem.Product,
		Quantity: invItem.Qty,
		Price:    float64(invItem.Rate),
		Category: invItem.Label,
		Url:      "",
	}

	// items := xendit.Invoice{
	// 	Items: []xendit.InvoiceItem{item},
	// }

	fee := xendit.InvoiceFee{
		Type:  "ADMIN",
		Value: 5000,
	}

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

	data := invoice.CreateParams{
		ExternalID:                     strconv.Itoa(inv.ID),
		Amount:                         50000,
		Description:                    "Invoice Demo #123",
		InvoiceDuration:                86400,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		SuccessRedirectURL:             "https://www.google.com",
		FailureRedirectURL:             "https://www.google.com",
		Currency:                       "IDR",
		Items:                          []xendit.InvoiceItem{item},
		Fees:                           []xendit.InvoiceFee{fee},
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Print(err)
		return resp, err
	}

	// fmt.Printf("created invoice: %+v\n", resp)
	return resp, nil
}

func NewServicePaymentGateway(repo adapter.AdapterPaymentGatewayRepository, c config.Config) adapter.AdapterPaymentGatewayService {
	return &servicePaymentGateway{
		repo: repo,
		c:    c,
	}
}
