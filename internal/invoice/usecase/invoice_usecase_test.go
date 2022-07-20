package usecase_test

import (
	"errors"
	"testing"
	"time"

	"Back-end/config"
	"Back-end/internal/invoice/model"
	"Back-end/internal/invoice/usecase"
	"Back-end/internal/invoice/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInvoiceService(t *testing.T) {
	repo := mocks.MockInvoiceRepository{}
	invoice := model.Invoice{
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
		repo.On("CreateInvoice", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		err := svc.CreateInvoiceService(invoice)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("CreateInvoice", mock.Anything).Return(errors.New("Failed Create Invoices")).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		err := svc.CreateInvoiceService(invoice)

		assert.Error(t, err)
	})
}

func TestGetAllInvoicesService(t *testing.T) {
	repo := mocks.MockInvoiceRepository{}
	data := []model.Invoice{
		{
			ID:           1,
			Number:       "INV/2020/0001",
			Name:         "issuer name test",
			Email:        "testgmail@gmail.com",
			Phone_Number: "12315415",
			Address:      "address test",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetAllInvoices").Return(data).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice := svc.GetAllInvoicesService()

		assert.Equal(t, invoice, data)
	})
}

func TestGetInvoiceByIDService(t *testing.T) {
	repo := mocks.MockInvoiceRepository{}
	data := model.Invoice{
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
	id := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoiceByID", mock.Anything).Return(data, nil).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice, err := svc.GetInvoiceByIDService(id)

		assert.Equal(t, id, invoice.ID)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoiceByID", mock.Anything).Return(data, errors.New("Failed Get Invoice By ID")).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice, err := svc.GetInvoiceByIDService(id)

		assert.Equal(t, id, invoice.ID)
		assert.Error(t, err)
	})
}

func TestGetInvoicesPaginationService(t *testing.T) {
	repo := mocks.MockInvoiceRepository{}
	data := []model.Invoice{
		{
			ID:           1,
			Number:       "INV/2020/0001",
			Name:         "issuer name test",
			Email:        "testgmail@gmail.com",
			Phone_Number: "12315415",
			Address:      "address test",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
		{
			ID:           2,
			Number:       "INV/2020/0002",
			Name:         "issuer name test2",
			Email:        "testgmail2@gmail.com",
			Phone_Number: "123154152",
			Address:      "address test2",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
		{
			ID:           3,
			Number:       "INV/2020/0003",
			Name:         "issuer name test3",
			Email:        "testgmail3@gmail.com",
			Phone_Number: "123154153",
			Address:      "address test3",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
	}
	pagination := model.Pagination{
		Page: 1,
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoicesPagination", mock.Anything).Return(data, nil).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice, err := svc.GetInvoicesPaginationService(pagination)

		assert.Equal(t, data, invoice)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoicesPagination", mock.Anything).Return(data, errors.New("Failed Get Invoices By Pagination")).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice, err := svc.GetInvoicesPaginationService(pagination)

		assert.Equal(t, data, invoice)
		assert.Error(t, err)
	})
}

func TestGetInovicesByPaymentStatusService(t *testing.T) {
	repo := mocks.MockInvoiceRepository{}
	data := []model.Invoice{
		{
			ID:           1,
			Number:       "INV/2020/0001",
			Name:         "issuer name test",
			Email:        "testgmail@gmail.com",
			Phone_Number: "12315415",
			Address:      "address test",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
		{
			ID:           2,
			Number:       "INV/2020/0002",
			Name:         "issuer name test2",
			Email:        "testgmail2@gmail.com",
			Phone_Number: "123154152",
			Address:      "address test2",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
		{
			ID:           3,
			Number:       "INV/2020/0003",
			Name:         "issuer name test3",
			Email:        "testgmail3@gmail.com",
			Phone_Number: "123154153",
			Address:      "address test3",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
	}
	page := model.Pagination{
		Page: 1,
	}
	status := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoicesByPaymentStatus", mock.Anything).Return(data, nil).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice := svc.GetInovicesByPaymentStatusService(status, page)

		assert.Equal(t, data, invoice)
		// assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoicesByPaymentStatus", mock.Anything).Return(data, errors.New("Failed Get Invoices By Payment Status")).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice := svc.GetInovicesByPaymentStatusService(status, page)

		assert.Equal(t, data, invoice)
		// assert.Error(t, err)
	})
}

func TestGetInvoicesByNameCustomerService(t *testing.T) {
	repo := mocks.MockInvoiceRepository{}
	data := []model.Invoice{
		{
			ID:           1,
			Number:       "INV/2020/0001",
			Name:         "issuer name test",
			Email:        "testgmail@gmail.com",
			Phone_Number: "12315415",
			Address:      "address test",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
		{
			ID:           2,
			Number:       "INV/2020/0002",
			Name:         "issuer name test2",
			Email:        "testgmail2@gmail.com",
			Phone_Number: "123154152",
			Address:      "address test2",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
		{
			ID:           3,
			Number:       "INV/2020/0003",
			Name:         "issuer name test3",
			Email:        "testgmail3@gmail.com",
			Phone_Number: "123154153",
			Address:      "address test3",
			Description:  "-",
			Invoice_Date: time.Now(),
			Due_Date:     time.Now(),
		},
	}
	search := "name"

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoicesByNameCustomer", mock.Anything).Return(data, nil).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice, err := svc.GetInvoicesByNameCustomerService(search)

		assert.Equal(t, invoice, data)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoicesByNameCustomer", mock.Anything).Return(data, errors.New("Failed Get Invoices By Search Name Customer")).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		invoice, err := svc.GetInvoicesByNameCustomerService(search)

		assert.Equal(t, invoice, data)
		assert.Error(t, err)
	})
}

func TestUpdateInvoiceByIDService(t *testing.T) {
	repo := mocks.MockInvoiceRepository{}
	data := model.Invoice{
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
	id := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("UpdateInvoiceByID", mock.Anything, mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		err := svc.UpdateInvoiceByIDService(id, data)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("UpdateInvoiceByID", mock.Anything, mock.Anything).Return(errors.New("Failed Update Invoice By ID")).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		err := svc.UpdateInvoiceByIDService(id, data)

		assert.Error(t, err)
	})
}

func TestDeleteInvoiceByIDService(t *testing.T) {
	repo := mocks.MockInvoiceRepository{}
	id := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("DeleteInvoiceByID", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		err := svc.DeleteInvoiceByIDService(id)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("DeleteInvoiceByID", mock.Anything).Return(errors.New("Failed Update Invoice By ID")).Once()

		svc := usecase.NewServiceInvoice(&repo, config.Config{})
		err := svc.DeleteInvoiceByIDService(id)

		assert.Error(t, err)
	})
}
