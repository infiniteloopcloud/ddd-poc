package transaction

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type Storage interface {
	Create(ctx context.Context, r *proto.Transaction) error
	Update(ctx context.Context, r *proto.Transaction) error
	Get(ctx context.Context, f *filters.Transaction) (*proto.Transaction, error)
	GetAll(ctx context.Context, f *filters.Transaction) ([]proto.Transaction, error)
}

type transactionStorage struct{}

func New() Storage {
	return transactionStorage{}
}

func (t transactionStorage) Create(ctx context.Context, r *proto.Transaction) error {
	return newSQL().Create(ctx, r)
}

func (t transactionStorage) Update(ctx context.Context, r *proto.Transaction) error {
	return newSQL().Update(ctx, r)
}

func (t transactionStorage) Get(ctx context.Context, f *filters.Transaction) (*proto.Transaction, error) {
	return newSQL().Get(ctx, f)
}

func (t transactionStorage) GetAll(ctx context.Context, f *filters.Transaction) ([]proto.Transaction, error) {
	return newSQL().GetAll(ctx, f)
}
