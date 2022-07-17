package usecase_test

import (
	"errors"
	"testing"
	"time"

	"Back-end/config"
	"Back-end/internal/generate/usecase"
	"Back-end/internal/generate/usecase/mocks"
	invoices "Back-end/internal/invoice/model"
	items "Back-end/internal/invoice_item/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateFileService(t *testing.T) {
	repo := mocks.MockGenerateRepository{}
	data := [][]string{
		{"00001", "Made Rahano", "mdrahano12@gmail.com", "082144948550", "Jalan Gubeng Kertajaya", "-", "2022-01-15", "2022-02-15", "Iphone ProMAX", "Technology", "1", "30000000"},
		// {"00002", "Angga Aditya", "agungangga2001@gmail.com", "0824124312", "Jalan Bali", "-", "2022-01-17", "2022-02-17", "Baju", "Clothe", "5", "100000"},
		// {"00003", "Jovin Lidan", "jovinlidan2@gmail.com", "0824124312", "Jalan Lombok", "-", "2022-01-21", "2022-02-21", "Rolex", "Watch", "1", "45000000"},
	}
	invoice := invoices.Invoice{
		ID:           1,
		Number:       "INV/2020/0001",
		Name:         "issuer name test",
		Email:        "testgmail@gmail.com",
		Phone_Number: "12315415",
		Address:      "address test",
		Description:  "-",
		Invoice_Date: time.Now(),
		Due_Date:     time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("NumberInvoiceExists", mock.Anything).Return(invoice, false).Once()
		repo.On("CreateInvoiceGenerate", mock.Anything).Return(nil).Once()
		repo.On("CreateInvoiceItemsGenerate", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceGenerate(&repo, config.Config{})
		err := svc.GenerateFileService(data)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("NumberInvoiceExists", mock.Anything).Return(invoice, true).Once()
		repo.On("CreateInvoiceGenerate", mock.Anything).Return(errors.New("Failed Create Invoice")).Once()
		repo.On("CreateInvoiceItemsGenerate", mock.Anything).Return(errors.New("Failed Create Invoice Items")).Once()

		svc := usecase.NewServiceGenerate(&repo, config.Config{})
		err := svc.GenerateFileService(data)

		assert.NoError(t, err)
	})
}

func TestGenerateInvoiceService(t *testing.T) {
	invoice := invoices.Invoice{
		ID:           1,
		Number:       "INV/2020/0001",
		Name:         "issuer name test",
		Email:        "testgmail@gmail.com",
		Phone_Number: "12315415",
		Address:      "address test",
		Description:  "-",
		Invoice_Date: time.Now(),
		Due_Date:     time.Now(),
	}
	item := []items.InvoiceItem{
		{
			ID_Invoice: 1,
			Product:    "HP Iphone 77",
			Category:   "Technology",
			Qty:        1,
			Price:      3000000,
		},
	}

	repo := mocks.MockGenerateRepository{}
	data := []int{1}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoices", mock.Anything).Return(invoice, item, nil).Once()
		repo.On("GetTotalAmount", mock.Anything).Return(float32(100000.00), nil).Once()
		repo.On("UpdateStatusInvoice", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("CreateTransactionRecord", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("SendEmail", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceGenerate(&repo, config.Config{})
		err := svc.GenerateInvoiceService(data)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoices", mock.Anything).Return(invoice, item, errors.New("Failed Get Invoices")).Once()
		repo.On("GetTotalAmount", mock.Anything).Return(float32(100000.00), errors.New("Failed Get Total Amount")).Once()
		repo.On("UpdateStatusInvoice", mock.Anything, mock.Anything).Return(errors.New("Failed Update Status Invoice")).Once()
		repo.On("CreateTransactionRecord", mock.Anything, mock.Anything).Return(errors.New("Failed Create Transaction Record")).Once()
		repo.On("SendEmail", mock.Anything).Return(errors.New("Failed Send Email to Customer")).Once()

		svc := usecase.NewServiceGenerate(&repo, config.Config{})
		err := svc.GenerateInvoiceService(data)

		assert.Error(t, err)
	})
}
