package model

import "time"

type Card struct {
	ID            int       `json:"id" form:"id"`
	Credit_Number int       `json:"credit_number" form:"credit_number"`
	Month         int       `json:"month" form:"month"`
	Year          int       `json:"year" form:"year"`
	Created_At    time.Time `json:"created_at" form:"created_at"`
	Updated_At    time.Time `json:"updated_at" form:"updated_at"`
}
