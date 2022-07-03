package model

import "time"

type SendCustomer struct {
	ID         int       `json:"id" form:"id"`
	To         string    `json:"to" form:"to"`
	Subject    string    `json:"subject" form:"subject"`
	Body       string    `json:"body" form:"body"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	Updated_At time.Time `json:"updated_at" form:"updated_at"`
}
