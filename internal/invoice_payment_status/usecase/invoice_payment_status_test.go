package usecase_test

import (
	"errors"
	"testing"

	"Back-end/config"
	"Back-end/internal/invoice_payment_status/model"
	"Back-end/internal/invoice_payment_status/usecase"
	"Back-end/internal/invoice_payment_status/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateInvoicePaymentStatusService(t *testing.T) {
	repo := mocks.MockInvoicePaymentStatusRepository{}
	status := model.InvoicePaymentStatus{
		Name: "Testing",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("CreateInvoicePaymentStatus", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		err := svc.CreateInvoicePaymentStatusService(status)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("CreateInvoicePaymentStatus", mock.Anything).Return(errors.New("Failed Crate Invoice Payment Status")).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		err := svc.CreateInvoicePaymentStatusService(status)

		assert.Error(t, err)
	})
}

func TestGetAllInvoicesPaymentStatusService(t *testing.T) {
	repo := mocks.MockInvoicePaymentStatusRepository{}
	data := []model.InvoicePaymentStatus{
		{
			ID:   1,
			Name: "Pending",
		},
		{
			ID:   2,
			Name: "Paid",
		},
		{
			ID:   3,
			Name: "Expired",
		},
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetAllInvoicesPaymentStatus").Return(data).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		status := svc.GetAllInvoicesPaymentStatusService()

		assert.Equal(t, status, data)
	})
}

func TestGetInvoicePaymentStatusByIDService(t *testing.T) {
	repo := mocks.MockInvoicePaymentStatusRepository{}
	ID := 1
	data := model.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("GetInvoicePaymentStatusByID", mock.Anything).Return(data, nil).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		status, err := svc.GetInvoicePaymentStatusByIDService(ID)

		assert.NoError(t, err)
		assert.Equal(t, status, data)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("GetInvoicePaymentStatusByID", mock.Anything).Return(data, errors.New("Failed Get Invoice Payment Status By ID")).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		_, err := svc.GetInvoicePaymentStatusByIDService(ID)

		assert.Error(t, err)
	})
}

func TestUpdateInvoicePaymentStatusByIDService(t *testing.T) {
	repo := mocks.MockInvoicePaymentStatusRepository{}
	ID := 1
	data := model.InvoicePaymentStatus{
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("UpdateInvoicePaymentStatusByID", mock.Anything, mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		err := svc.UpdateInvoicePaymentStatusByIDService(ID, data)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("UpdateInvoicePaymentStatusByID", mock.Anything, mock.Anything).Return(errors.New("Failed Update Invoice Payment Status By ID")).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		err := svc.UpdateInvoicePaymentStatusByIDService(ID, data)

		assert.Error(t, err)
	})
}

func TestDeleteInvoicePaymentStatusByIDService(t *testing.T) {
	repo := mocks.MockInvoicePaymentStatusRepository{}
	ID := 1

	t.Run("Success", func(t *testing.T) {
		repo.On("DeleteInvoicePaymentStatusByID", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		err := svc.DeleteInvoicePaymentStatusByIDService(ID)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("DeleteInvoicePaymentStatusByID", mock.Anything).Return(errors.New("Failed Delete Invoice Payment Status By ID")).Once()

		svc := usecase.NewServiceInvoicePaymentStatus(&repo, config.Config{})
		err := svc.DeleteInvoicePaymentStatusByIDService(ID)

		assert.Error(t, err)
	})
}
