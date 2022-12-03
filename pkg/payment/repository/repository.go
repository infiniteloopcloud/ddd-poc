package repository

import (
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository/gateway"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository/transaction"
)

type PaymentStorage struct {
	Gateway     gateway.Storage
	Transaction transaction.Storage
}

func New() PaymentStorage {
	return PaymentStorage{
		Gateway:     gateway.New(),
		Transaction: transaction.New(),
	}
}
