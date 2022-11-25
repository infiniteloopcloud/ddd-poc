package transaction

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type Storage interface {
	Create(ctx context.Context, r *types.Transaction) error
	Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error)
	GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error)
}

type transactionStorage struct{}

func New() Storage {
	return transactionStorage{}
}

func (t transactionStorage) Create(ctx context.Context, r *types.Transaction) error {
	return newSQL().Create(ctx, r)
}

func (t transactionStorage) Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error) {
	return newSQL().Get(ctx, f)
}

func (t transactionStorage) GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error) {
	return newSQL().GetAll(ctx, f)
}
