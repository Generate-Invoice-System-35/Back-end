package usecase_test

import (
	"errors"
	"testing"

	"Back-end/config"
	"Back-end/internal/send_customer/model"
	"Back-end/internal/send_customer/usecase"
	"Back-end/internal/send_customer/usecase/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendEmailService(t *testing.T) {
	repo := mocks.MockSendCustomerRepository{}
	data := model.SendCustomer{
		To:      "zereftheblackmage33@gmail.com",
		Subject: "Testing Subject",
		Body:    "Testing Body, I hope this is working",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("SendEmail", mock.Anything).Return(nil).Once()

		svc := usecase.NewServiceSendCustomer(&repo, config.Config{})
		err := svc.SendEmailService(data)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		repo.On("SendEmail", mock.Anything).Return(errors.New("Failed Send Email")).Once()

		svc := usecase.NewServiceSendCustomer(&repo, config.Config{})
		err := svc.SendEmailService(data)

		assert.Error(t, err)
	})
}
