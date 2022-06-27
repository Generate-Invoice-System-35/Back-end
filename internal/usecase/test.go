/*
Invoice id,Customer name,Invoice date,Item name,Qty,Price
inv-001,Fahmy,09/06/2022,,,
,,,Macbook M1,2,35000000
,,,Mesin kopi,1,2000000
inv-002,Go Frendi,09/06/2022,,,
,,,Kursi gaming,1,10000000
inv-003,Budi,09/06/2022,,,
*/

package usecase

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Invoice struct {
	Id           string
	CustomerName string
	InvoiceDate  string
	Items        []InvoiceItem
}

type InvoiceItem struct {
	ItemName string
	Qty      int
	Price    int
}

func main() {
	file, err := os.Open("invoice.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	invoices := []*Invoice{}
	var invoice *Invoice
	for index, row := range records {
		if index == 0 {
			continue
		}
		// invoice
		if row[0] != "" {
			if invoice != nil {
				invoices = append(invoices, invoice)
			}
			invoice = &Invoice{
				Id:           row[0],
				CustomerName: row[1],
				InvoiceDate:  row[2],
			}
			continue
		}
		// item
		qty, _ := strconv.Atoi(row[4])
		price, _ := strconv.Atoi(row[5])
		item := InvoiceItem{
			ItemName: row[3],
			Qty:      qty,
			Price:    price,
		}
		invoice.Items = append(invoice.Items, item)
	}
	// if invoice != nil {
	// 	invoices = append(invoices, invoice)
	// }
	// tampilkan
	for _, invoice := range invoices {
		fmt.Printf("%#v\n", invoice)
	}
}
