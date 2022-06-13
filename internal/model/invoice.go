package model

import "time"

type Invoice struct {
	ID                int       `json:"id" form:"id"`
	ID_Payment_Status int       `json:"id_payment_status" form:"id_payment_status"`
	Number            string    `json:"number" form:"number"`
	Buyer_Name        string    `json:"buyer_name" form:"buyer_name"`
	Invoice_Date      time.Time `json:"invoice_date" form:"invoice_date"`
	Due_Date          time.Time `json:"due_date" form:"due_date"`
}
