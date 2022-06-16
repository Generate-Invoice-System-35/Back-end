package model

type SendCustomer struct {
	To      string `json:"to" form:"to"`
	Subject string `json:"subject" form:"subject"`
	Body    string `json:"body" form:"body"`
}
