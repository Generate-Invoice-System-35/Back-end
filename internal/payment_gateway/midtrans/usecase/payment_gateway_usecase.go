package usecase

import (
	"Back-end/config"
	"Back-end/internal/payment_gateway/midtrans/adapter"
)

type servicePaymentGateway struct {
	c    config.Config
	repo adapter.AdapterPaymentGatewayRepository
}

func (s *servicePaymentGateway) ChargeTransactionService(id int) (int, error) {
	// // 1. Set you ServerKey with globally
	// midtrans.ServerKey = "YOUR-SERVER-KEY"
	// midtrans.Environment = midtrans.Sandbox

	// // 2. Initiate charge request
	// chargeReq := &coreapi.ChargeReq{
	// 	PaymentType: midtrans.SourceCreditCard,
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  "12345",
	// 		GrossAmt: 200000,
	// 	},
	// 	CreditCard: &coreapi.CreditCardDetails{
	// 		TokenID:        "YOUR-CC-TOKEN",
	// 		Authentication: true,
	// 	},
	// 	Items: &[]midtrans.ItemDetail{
	// 		coreapi.ItemDetail{
	// 			ID:    "ITEM1",
	// 			Price: 200000,
	// 			Qty:   1,
	// 			Name:  "Someitem",
	// 		},
	// 	},
	// }

	// // 3. Request to Midtrans using global config
	// coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
	// fmt.Println("Response :", coreApiRes)

	return 0, nil
}

func NewServicePaymentGateway(repo adapter.AdapterPaymentGatewayRepository, c config.Config) adapter.AdapterPaymentGatewayService {
	return &servicePaymentGateway{
		repo: repo,
		c:    c,
	}
}
