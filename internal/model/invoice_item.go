package model

type InvoiceItem struct {
	ID         int     `json:"id" form:"id"`
	ID_Invoice int     `json:"id_invoice" form:"id_invoice"`
	Product    string  `json:"product" form:"product"`
	Label      string  `json:"label" form:"label"`
	Qty        int     `json:"qty" form:"qty"`
	Price      float32 `json:"price" form:"price"`
	Tax        float32 `json:"tax" form:"tax"`
	Subtotal   float32 `json:"subtotal" form:"subtotal"`
}
