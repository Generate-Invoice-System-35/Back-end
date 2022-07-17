package usecase_test

import (
	"errors"
	"testing"
	"time"

	"Back-end/config"
	inv "Back-end/internal/invoice/model"
	invItem "Back-end/internal/invoice_item/model"
	"Back-end/internal/payment_gateway/xendit/model"
	"Back-end/internal/payment_gateway/xendit/usecase"
	"Back-end/internal/payment_gateway/xendit/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateXenditPaymentInvoiceService(t *testing.T) {
	repo := mocks.MockPaymentGatewayRepository{}
	invoice := inv.Invoice{
		ID:           2,
		Number:       "INV/2020/0001",
		Name:         "issuer name test",
		Email:        "testgmail@gmail.com",
		Phone_Number: "12315415",
		Address:      "address test",
		Description:  "-",
		Invoice_Date: time.Now(),
		Due_Date:     time.Now(),
	}
	item := []invItem.InvoiceItem{
		{
			ID_Invoice: 2,
			Product:    "HP Iphone 77",
			Category:   "Technology",
			Qty:        1,
			Price:      3000000,
		},
		{
			ID_Invoice: 2,
			Product:    "HP Iphone 77",
			Category:   "Technology",
			Qty:        1,
			Price:      3000000,
		},
	}
	total := float32(6000000)
	id := 2

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoices", mock.Anything).Return(invoice, item, nil).Once()
		repo.On("GetTotalAmount", mock.Anything).Return(total, nil).Once()
		repo.On("UpdateStatusInvoice", mock.Anything, mock.Anything).Return(nil).Once()
		repo.On("CreateTransactionRecord", mock.Anything, mock.Anything).Return(nil).Once()

		svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
		_, err := svc.CreateXenditPaymentInvoiceService(id)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoices", mock.Anything).Return(invoice, item, errors.New("Failed Get Invoices")).Once()
		repo.On("GetTotalAmount", mock.Anything).Return(total, errors.New("Failed Get Total Amount")).Once()
		repo.On("UpdateStatusInvoice", mock.Anything, mock.Anything).Return(errors.New("Failed Update Invoces Payment Status")).Once()
		repo.On("CreateTransactionRecord", mock.Anything, mock.Anything).Return(errors.New("Failed Record Transaction")).Once()

		svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
		_, err := svc.CreateXenditPaymentInvoiceService(id)

		assert.Error(t, err)
	})
}

func TestGetXenditPaymentInvoiceService(t *testing.T) {
	repo := mocks.MockPaymentGatewayRepository{}
	record := model.TransactionRecord{
		ID_Invoice:         2,
		ID_Invoice_Payment: "62d3d20e36992aedcd007a6e",
		ID_User_Payment:    "62ac09b74c8f5874ad68dede",
	}
	id := 2

	t.Run("Success", func(t *testing.T) {
		repo.On("GetIDInvoicePayment", mock.Anything).Return(record, nil)

		svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
		_, err := svc.GetXenditPaymentInvoiceService(id)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetIDInvoicePayment", mock.Anything).Return(record, errors.New("Failed Get Invoice By ID Invoice Payment"))

		svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
		_, err := svc.GetXenditPaymentInvoiceService(id)

		assert.NoError(t, err)
	})
}

func TestGetAllXenditPaymentInvoiceService(t *testing.T) {
	repo := mocks.MockPaymentGatewayRepository{}

	t.Run("Success", func(t *testing.T) {
		svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
		_, err := svc.GetAllXenditPaymentInvoiceService()

		assert.NoError(t, err)
	})
}

func TestExpireXenditPaymentInvoiceService(t *testing.T) {
	// repo := mocks.MockPaymentGatewayRepository{}
	// record := model.TransactionRecord{
	// 	ID_Invoice:         2,
	// 	ID_Invoice_Payment: "62d3d20e36992aedcd007a6e",
	// 	ID_User_Payment:    "62ac09b74c8f5874ad68dede",
	// }
	// id := 2

	// t.Run("Success", func(t *testing.T) {
	// 	repo.On("GetIDInvoicePayment", mock.Anything).Return(record, nil).Once()
	// 	repo.On("UpdateStatusInvoice", mock.Anything, mock.Anything).Return(nil).Once()

	// 	svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
	// 	_, err := svc.ExpireXenditPaymentInvoiceService(id)

	// 	assert.NoError(t, err)
	// })

	// t.Run("Failed", func(t *testing.T) {
	// 	repo.On("GetIDInvoicePayment", mock.Anything).Return(record, errors.New("Failed Get ID")).Once()
	// 	repo.On("UpdateStatusInvoice", mock.Anything, mock.Anything).Return(errors.New("Failed Update Status Invoice")).Once()

	// 	svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
	// 	_, err := svc.ExpireXenditPaymentInvoiceService(id)

	// 	assert.Error(t, err)
	// })
}

func TestCallbackXenditPaymentInvoiceService(t *testing.T) {
	repo := mocks.MockPaymentGatewayRepository{}
	data := model.CallbackInvoice{
		ExternalID: "1",
		Status:     "PENDING",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("UpdateStatusInvoice", mock.Anything, mock.Anything).Return(nil).Once()

		svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
		err := svc.CallbackXenditPaymentInvoiceService(data)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("UpdateStatusInvoice", mock.Anything, mock.Anything).Return(errors.New("Failed Callback Invoices")).Once()

		svc := usecase.NewServicePaymentGateway(&repo, config.Config{})
		err := svc.CallbackXenditPaymentInvoiceService(data)

		assert.Error(t, err)
	})
}
