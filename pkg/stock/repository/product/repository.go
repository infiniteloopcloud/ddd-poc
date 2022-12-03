package product

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type Storage interface {
	Create(ctx context.Context, r *proto.Product) (*proto.Product, error)
	Update(ctx context.Context, r *proto.Product) (*proto.Product, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Product) (*proto.Product, error)
	GetAll(ctx context.Context, f *filters.Product) ([]proto.Product, error)
}

type productStorage struct{}

func New() Storage {
	return productStorage{}
}

func (p productStorage) Create(ctx context.Context, r *proto.Product) (*proto.Product, error) {
	return newSQL().Create(ctx, r)
}

func (p productStorage) Update(ctx context.Context, r *proto.Product) (*proto.Product, error) {
	return newSQL().Update(ctx, r)
}

func (p productStorage) Delete(ctx context.Context, id string) error {
	return newSQL().Delete(ctx, id)
}

func (p productStorage) Get(ctx context.Context, f *filters.Product) (*proto.Product, error) {
	return newSQL().Get(ctx, f)
}

func (p productStorage) GetAll(ctx context.Context, f *filters.Product) ([]proto.Product, error) {
	return newSQL().GetAll(ctx, f)
}
