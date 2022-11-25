package cart

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type Storage interface {
	Create(ctx context.Context, r *types.Cart) (*types.Cart, error)
	Update(ctx context.Context, r *types.Cart) (*types.Cart, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Cart) (*types.Cart, error)
	GetAll(ctx context.Context, f *filters.Cart) ([]types.Cart, error)
}

type cartStorage struct{}

func New() Storage {
	return cartStorage{}
}

func (c cartStorage) Create(ctx context.Context, r *types.Cart) (*types.Cart, error) {
	return newSQL().Create(ctx, r)
}

func (c cartStorage) Update(ctx context.Context, r *types.Cart) (*types.Cart, error) {
	return newSQL().Update(ctx, r)
}

func (c cartStorage) Delete(ctx context.Context, id string) error {
	return newSQL().Delete(ctx, id)
}

func (c cartStorage) Get(ctx context.Context, f *filters.Cart) (*types.Cart, error) {
	return newSQL().Get(ctx, f)
}

func (c cartStorage) GetAll(ctx context.Context, f *filters.Cart) ([]types.Cart, error) {
	return newSQL().GetAll(ctx, f)
}
