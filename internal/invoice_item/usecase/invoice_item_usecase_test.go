package usecase_test

import (
	"errors"
	"testing"

	"Back-end/config"
	"Back-end/internal/invoice_item/model"
	"Back-end/internal/invoice_item/usecase"
	"Back-end/internal/invoice_item/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInvoiceItemService(t *testing.T) {
	repo := mocks.MockInvoiceItemRepository{}
	item := model.InvoiceItem{
		ID_Invoice: 1,
		Product:    "HP Iphone 77",
		Category:   "Technology",
		Qty:        1,
		Price:      3000000,
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("CreateInvoiceItem", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		err := svc.CreateInvoiceItemService(item)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("CreateInvoiceItem", mock.Anything).Return(errors.New("Failed Create Invoices Item")).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		err := svc.CreateInvoiceItemService(item)

		assert.Error(t, err)
	})
}

func TestGetAllInvoiceItemsService(t *testing.T) {
	repo := mocks.MockInvoiceItemRepository{}
	data := []model.InvoiceItem{
		{
			ID_Invoice: 1,
			Product:    "HP Iphone 77",
			Category:   "Technology",
			Qty:        1,
			Price:      3000000,
		},
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetAllInvoiceItems").Return(data).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		item := svc.GetAllInvoiceItemsService()

		assert.Equal(t, item, data)
	})
}

func TestGetInvoiceItemByIDService(t *testing.T) {
	repo := mocks.MockInvoiceItemRepository{}
	data := model.InvoiceItem{
		ID_Invoice: 1,
		Product:    "HP Iphone 77",
		Category:   "Technology",
		Qty:        1,
		Price:      3000000,
	}
	id := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoiceItemByID", mock.Anything).Return(data, nil).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		item, err := svc.GetInvoiceItemByIDService(id)

		assert.Equal(t, item, data)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoiceItemByID", mock.Anything).Return(data, errors.New("Failed Get Invoice Item By ID")).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		item, err := svc.GetInvoiceItemByIDService(id)

		assert.Equal(t, item, data)
		assert.Error(t, err)
	})
}

func TestGetInvoiceItemByNumberService(t *testing.T) {
	repo := mocks.MockInvoiceItemRepository{}
	data := []model.InvoiceItem{
		{
			ID_Invoice: 1,
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
	number := "00001"

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoiceItemByNumber", mock.Anything).Return(data, nil).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		item, err := svc.GetInvoiceItemByNumberService(number)

		assert.Equal(t, item, data)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoiceItemByNumber", mock.Anything).Return(data, errors.New("Failed Get Invoice Item By Number Invoice")).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		item, err := svc.GetInvoiceItemByNumberService(number)

		assert.Equal(t, item, data)
		assert.Error(t, err)
	})
}

func TestUpdateInvoiceItemByIDService(t *testing.T) {
	repo := mocks.MockInvoiceItemRepository{}
	data := model.InvoiceItem{
		ID_Invoice: 1,
		Product:    "HP Iphone 77",
		Category:   "Technology",
		Qty:        1,
		Price:      3000000,
	}
	id := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("UpdateInvoiceItemByID", mock.Anything, mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		err := svc.UpdateInvoiceItemByIDService(id, data)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("UpdateInvoiceItemByID", mock.Anything, mock.Anything).Return(errors.New("Failed Update Invoice Item By ID")).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		err := svc.UpdateInvoiceItemByIDService(id, data)

		assert.Error(t, err)
	})
}

func TestDeleteInvoiceItemByIDService(t *testing.T) {
	repo := mocks.MockInvoiceItemRepository{}
	id := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("DeleteInvoiceItemByID", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		err := svc.DeleteInvoiceItemByIDService(id)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("DeleteInvoiceItemByID", mock.Anything).Return(errors.New("Failed Update Invoice Item By ID")).Once()

		svc := usecase.NewServiceInvoiceItem(&repo, config.Config{})
		err := svc.DeleteInvoiceItemByIDService(id)

		assert.Error(t, err)
	})
}
