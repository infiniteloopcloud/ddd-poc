package repository

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

type TransactionStorage interface {
	Create(ctx context.Context, r *types.Transaction) error
	Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error)
	GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error)
}

type transactionStorage struct{}

func NewTransaction() TransactionStorage {
	return transactionStorage{}
}

func (t transactionStorage) Create(ctx context.Context, r *types.Transaction) error {
	// TODO implement me
	panic("implement me")
}

func (t transactionStorage) Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error) {
	// TODO implement me
	panic("implement me")
}

func (t transactionStorage) GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error) {
	// TODO implement me
	panic("implement me")
}
