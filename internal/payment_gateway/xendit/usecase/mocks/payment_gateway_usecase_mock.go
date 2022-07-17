package mocks

import (
	invoices "Back-end/internal/invoice/model"
	items "Back-end/internal/invoice_item/model"
	transactions "Back-end/internal/payment_gateway/xendit/model"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
)

type MockPaymentGatewayRepository struct {
	mock.Mock
}

func (r *MockPaymentGatewayRepository) CreateTransactionRecord(id int, transaction transactions.TransactionRecord) error {
	ret := r.Called(id, transaction)

	var err error
	if res, ok := ret.Get(0).(func(int, transactions.TransactionRecord) error); ok {
		err = res(id, transaction)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockPaymentGatewayRepository) GetIDInvoicePayment(id int) (record transactions.TransactionRecord, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) transactions.TransactionRecord); ok {
		record = res(id)
	} else {
		record = ret.Get(0).(transactions.TransactionRecord)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockPaymentGatewayRepository) GetInvoices(id int) (invoice invoices.Invoice, item []items.InvoiceItem, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) invoices.Invoice); ok {
		invoice = res(id)
	} else {
		invoice = ret.Get(0).(invoices.Invoice)
	}

	if res, ok := ret.Get(1).(func(int) []items.InvoiceItem); ok {
		item = res(id)
	} else {
		item = ret.Get(1).([]items.InvoiceItem)
	}

	if res, ok := ret.Get(2).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(2)
	}

	return
}

func (r *MockPaymentGatewayRepository) GetTotalAmount(id int) (total float32, err error) {
	ret := r.Called(id)

	if res, ok := ret.Get(0).(func(int) float32); ok {
		total = res(id)
	} else {
		total = ret.Get(0).(float32)
	}

	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return
}

func (r *MockPaymentGatewayRepository) UpdateStatusInvoice(id int, invoice invoices.Invoice) error {
	ret := r.Called(id, invoice)

	var err error
	if res, ok := ret.Get(0).(func(int, invoices.Invoice) error); ok {
		err = res(id, invoice)
	} else {
		err = ret.Error(0)
	}

	return err
}

func MockData() (*xendit.Invoice, []xendit.Invoice) {
	expireDate, _ := time.Parse("Mon Jan 02 2006 15:04:05 GMT-0700", "2022-07-17T09:10:39.736Z")
	createdDate, _ := time.Parse("Mon Jan 02 2006 15:04:05 GMT-0700", "2022-07-17T09:10:39.736Z")
	updatedDate, _ := time.Parse("Mon Jan 02 2006 15:04:05 GMT-0700", "2022-07-17T09:10:39.736Z")
	data := &xendit.Invoice{
		ID:                        "62d3d20e36992aedcd007a6e",
		InvoiceURL:                "https://checkout-staging.xendit.co/web/62d3d20e36992aedcd007a6e",
		UserID:                    "62ac09b74c8f5874ad68dede",
		ExternalID:                "2",
		Status:                    "PENDING",
		MerchantName:              "Made in Bali",
		MerchantProfilePictureURL: "https://xnd-merchant-logos.s3.amazonaws.com/business/production/62ac09b74c8f5874ad68dede-1655442253521.png",
		Amount:                    1005000,
		Items: []xendit.InvoiceItem{
			{
				Name:     "Baju",
				Price:    100000,
				Quantity: 5,
				Category: "Clothe",
			},
			{
				Name:     "Baju",
				Price:    100000,
				Quantity: 5,
				Category: "Clothe",
			},
		},
		Fees: []xendit.InvoiceFee{
			{
				Type:  "ADMIN",
				Value: 5000,
			},
		},
		Description: "-",
		ExpiryDate:  &expireDate,
		Customer: xendit.InvoiceCustomer{
			GivenNames:   "Angga Aditya",
			Email:        "agungangga2001@gmail.com",
			MobileNumber: "00002",
		},
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated: []string{
				"whatsapp",
				"email",
				"sms",
			},
			InvoiceReminder: []string{
				"whatsapp",
				"email",
				"sms",
			},
			InvoicePaid: []string{
				"whatsapp",
				"email",
				"sms",
			},
			InvoiceExpired: []string{
				"whatsapp",
				"email",
				"sms",
			},
		},
		AvailableBanks: []xendit.InvoiceBank{
			{
				BankCode:          "MANDIRI",
				CollectionType:    "POOL",
				BankAccountNumber: "",
				TransferAmount:    1005000,
				BankBranch:        "Virtual Account",
				AccountHolderName: "MADE IN BALI",
				IdentityAmount:    0,
			},
			{
				BankCode:          "BRI",
				CollectionType:    "POOL",
				BankAccountNumber: "",
				TransferAmount:    1005000,
				BankBranch:        "Virtual Account",
				AccountHolderName: "MADE IN BALI",
				IdentityAmount:    0,
			},
			{
				BankCode:          "BNI",
				CollectionType:    "POOL",
				BankAccountNumber: "",
				TransferAmount:    1005000,
				BankBranch:        "Virtual Account",
				AccountHolderName: "MADE IN BALI",
				IdentityAmount:    0,
			},
			{
				BankCode:          "PERMATA",
				CollectionType:    "POOL",
				BankAccountNumber: "",
				TransferAmount:    1005000,
				BankBranch:        "Virtual Account",
				AccountHolderName: "MADE IN BALI",
				IdentityAmount:    0,
			},
			{
				BankCode:          "BCA",
				CollectionType:    "POOL",
				BankAccountNumber: "",
				TransferAmount:    1005000,
				BankBranch:        "Virtual Account",
				AccountHolderName: "MADE IN BALI",
				IdentityAmount:    0,
			},
		},
		AvailableEWallets: []xendit.InvoiceEWallet{
			{
				EWalletType: "OVO",
			},
			{
				EWalletType: "DANA",
			},
			{
				EWalletType: "SHOPEEPAY",
			},
			{
				EWalletType: "LINKAJA",
			},
		},
		AvailableRetailOutlets: []xendit.InvoiceRetailOutlet{
			{
				RetailOutletName: "ALFAMART",
			},
			{
				RetailOutletName: "INDOMARET",
			},
		},
		ShouldSendEmail:    false,
		Created:            &createdDate,
		Updated:            &updatedDate,
		Currency:           "IDR",
		PaymentDetail:      xendit.InvoicePaymentDetail{},
		SuccessRedirectURL: "https://http.cat/200",
		FailureRedirectURL: "https://http.cat/406",
	}

	datas := []xendit.Invoice{
		{
			ID:                        "62d3d20e36992aedcd007a6e",
			InvoiceURL:                "https://checkout-staging.xendit.co/web/62d3d20e36992aedcd007a6e",
			UserID:                    "62ac09b74c8f5874ad68dede",
			ExternalID:                "2",
			Status:                    "PENDING",
			MerchantName:              "Made in Bali",
			MerchantProfilePictureURL: "https://xnd-merchant-logos.s3.amazonaws.com/business/production/62ac09b74c8f5874ad68dede-1655442253521.png",
			Amount:                    1005000,
			Items: []xendit.InvoiceItem{
				{
					Name:     "Baju",
					Price:    100000,
					Quantity: 5,
					Category: "Clothe",
				},
				{
					Name:     "Baju",
					Price:    100000,
					Quantity: 5,
					Category: "Clothe",
				},
			},
			Fees: []xendit.InvoiceFee{
				{
					Type:  "ADMIN",
					Value: 5000,
				},
			},
			Description: "-",
			ExpiryDate:  &expireDate,
			Customer: xendit.InvoiceCustomer{
				GivenNames:   "Angga Aditya",
				Email:        "agungangga2001@gmail.com",
				MobileNumber: "00002",
			},
			CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
				InvoiceCreated: []string{
					"whatsapp",
					"email",
					"sms",
				},
				InvoiceReminder: []string{
					"whatsapp",
					"email",
					"sms",
				},
				InvoicePaid: []string{
					"whatsapp",
					"email",
					"sms",
				},
				InvoiceExpired: []string{
					"whatsapp",
					"email",
					"sms",
				},
			},
			AvailableBanks: []xendit.InvoiceBank{
				{
					BankCode:          "MANDIRI",
					CollectionType:    "POOL",
					BankAccountNumber: "",
					TransferAmount:    1005000,
					BankBranch:        "Virtual Account",
					AccountHolderName: "MADE IN BALI",
					IdentityAmount:    0,
				},
				{
					BankCode:          "BRI",
					CollectionType:    "POOL",
					BankAccountNumber: "",
					TransferAmount:    1005000,
					BankBranch:        "Virtual Account",
					AccountHolderName: "MADE IN BALI",
					IdentityAmount:    0,
				},
				{
					BankCode:          "BNI",
					CollectionType:    "POOL",
					BankAccountNumber: "",
					TransferAmount:    1005000,
					BankBranch:        "Virtual Account",
					AccountHolderName: "MADE IN BALI",
					IdentityAmount:    0,
				},
				{
					BankCode:          "PERMATA",
					CollectionType:    "POOL",
					BankAccountNumber: "",
					TransferAmount:    1005000,
					BankBranch:        "Virtual Account",
					AccountHolderName: "MADE IN BALI",
					IdentityAmount:    0,
				},
				{
					BankCode:          "BCA",
					CollectionType:    "POOL",
					BankAccountNumber: "",
					TransferAmount:    1005000,
					BankBranch:        "Virtual Account",
					AccountHolderName: "MADE IN BALI",
					IdentityAmount:    0,
				},
			},
			AvailableEWallets: []xendit.InvoiceEWallet{
				{
					EWalletType: "OVO",
				},
				{
					EWalletType: "DANA",
				},
				{
					EWalletType: "SHOPEEPAY",
				},
				{
					EWalletType: "LINKAJA",
				},
			},
			AvailableRetailOutlets: []xendit.InvoiceRetailOutlet{
				{
					RetailOutletName: "ALFAMART",
				},
				{
					RetailOutletName: "INDOMARET",
				},
			},
			ShouldSendEmail:    false,
			Created:            &createdDate,
			Updated:            &updatedDate,
			Currency:           "IDR",
			PaymentDetail:      xendit.InvoicePaymentDetail{},
			SuccessRedirectURL: "https://http.cat/200",
			FailureRedirectURL: "https://http.cat/406",
		},
	}

	return data, datas
}
