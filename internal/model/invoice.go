package model

import "time"

type Invoice struct {
	ID                int       `json:"id" form:"id"`
	ID_Payment_Status int       `json:"id_payment_status" form:"id_payment_status"`
	Number            string    `json:"number" form:"number"`
	Date              time.Time `json:"date" form:"date"`
	Due_Date          time.Time `json:"due_date" form:"due_date"`
	Issuer_Name       string    `json:"issuer_name" form:"issuer_name"`
	Issuer_Street     string    `json:"issuer_street" form:"issuer_street"`
	Issuer_City       string    `json:"issuer_city" form:"issuer_city"`
	Issuer_Zip        int       `json:"issuer_zip" form:"issuer_zip"`
	Buyer_Name        string    `json:"buyer_name" form:"buyer_name"`
	Buyer_Street      string    `json:"buyer_street" form:"buyer_street"`
	Buyer_City        string    `json:"buyer_city" form:"buyer_city"`
	Buyer_Zip         int       `json:"buyer_zip" form:"buyer_zip"`
	Tax               float32   `json:"tax" form:"tax"`
	Total             float32   `json:"total" form:"total"`
}
