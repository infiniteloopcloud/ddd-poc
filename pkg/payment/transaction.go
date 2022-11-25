package payment

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type transactionService struct {
	transactionStorage repository.TransactionStorage
}

func NewTransaction() comm.TransactionDescriptor {
	return transactionService{
		transactionStorage: repository.NewTransaction(),
	}
}

func (t transactionService) Create(ctx context.Context, r *types.Transaction) error {
	if err := r.Validate(); err != nil {
		return err
	}

	return t.transactionStorage.Create(ctx, r)
}

func (t transactionService) Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error) {
	return t.transactionStorage.Get(ctx, f)
}

func (t transactionService) GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error) {
	return t.transactionStorage.GetAll(ctx, f)
}
