package repository

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

type ProductStorage interface {
	Create(ctx context.Context, r *types.Product) (*types.Product, error)
	Update(ctx context.Context, r *types.Product) (*types.Product, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Product) (*types.Product, error)
	GetAll(ctx context.Context, f *filters.Product) ([]types.Product, error)
}

type productStorage struct{}

func NewProduct() ProductStorage {
	return productStorage{}
}

func (p productStorage) Create(ctx context.Context, r *types.Product) (*types.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (p productStorage) Update(ctx context.Context, r *types.Product) (*types.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (p productStorage) Delete(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}

func (p productStorage) Get(ctx context.Context, f *filters.Product) (*types.Product, error) {
	// TODO implement me
	panic("implement me")
}

func (p productStorage) GetAll(ctx context.Context, f *filters.Product) ([]types.Product, error) {
	// TODO implement me
	panic("implement me")
}
