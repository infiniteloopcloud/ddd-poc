package payment

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type transactionService struct {
	paymentStorage repository.PaymentStorage
}

func NewTransaction() comm.TransactionDescriptor {
	return transactionService{
		paymentStorage: repository.New(),
	}
}

func (t transactionService) Create(ctx context.Context, r *types.Transaction) error {
	if err := r.Validate(); err != nil {
		return err
	}

	return t.paymentStorage.Transaction.Create(ctx, r)
}

func (t transactionService) Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error) {
	return t.paymentStorage.Transaction.Get(ctx, f)
}

func (t transactionService) GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error) {
	return t.paymentStorage.Transaction.GetAll(ctx, f)
}
