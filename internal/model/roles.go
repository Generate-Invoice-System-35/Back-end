package model

import "time"

type Role struct {
	ID         int       `json:"id" form:"id"`
	Name       string    `json:"name" form:"name"`
	Created_At time.Time `json:"created_at" form:"created_at"`
	Updated_at time.Time `json:"updated_at" form:"updated_at"`
}
