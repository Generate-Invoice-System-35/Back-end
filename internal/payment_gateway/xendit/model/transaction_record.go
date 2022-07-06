package model

import "time"

type TransactionRecord struct {
	ID_Invoice         int       `json:"id_invoice" form:"id_invoice"`
	ID_Invoice_Payment string    `json:"id_invoice_payment" form:"id_invoice_payment"`
	ID_User_Payment    string    `json:"id_user_payment" form:"id_user_payment"`
	Created_At         time.Time `json:"created_at" form:"created_at"`
	Updated_At         time.Time `json:"updated_at" form:"updated_at"`
}

type CallbackInvoice struct {
	ID                 string    `json:"id"`
	ExternalID         string    `json:"external_id"`
	UserID             string    `json:"user_id"`
	IsHigh             bool      `json:"is_high"`
	PaymentMethod      string    `json:"payment_method"`
	Status             string    `json:"status"`
	MerchantName       string    `json:"merchant_name"`
	Amount             int       `json:"amount"`
	PaidAmount         int       `json:"paid_amount"`
	BankCode           string    `json:"bank_code"`
	PaidAt             time.Time `json:"paid_at"`
	PayerEmail         string    `json:"payer_email"`
	Description        string    `json:"description"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated"`
	Currency           string    `json:"currency"`
	PaymentChannel     string    `json:"payment_channel"`
	PaymentDestination string    `json:"payment_destination"`
}
