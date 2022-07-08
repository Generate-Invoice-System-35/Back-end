package model

import "time"

type Invoice struct {
	ID                int       `json:"id" form:"id"`
	ID_Payment_Status int       `json:"id_payment_status" form:"id_payment_status"`
	Number            string    `json:"number" form:"number"`
	Name              string    `json:"name" form:"name"`
	Email             string    `json:"email" form:"email"`
	Phone_Number      string    `json:"phone_number" form:"phone_number"`
	Address           string    `json:"address" form:"address"`
	Description       string    `json:"description" form:"description"`
	Invoice_Date      time.Time `json:"invoice_date" form:"invoice_date"`
	Due_Date          time.Time `json:"due_date" form:"due_date"`
	Created_At        time.Time `json:"created_at" form:"created_at"`
	Updated_At        time.Time `json:"updated_at" form:"updated_at"`
}

type Pagination struct {
	Limit int    `json:"limit" form:"limit"`
	Page  int    `json:"page" form:"page"`
	Sort  string `json:"sort" form:"sort"`
}
