package gateway

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type Storage interface {
	Create(ctx context.Context, r *types.Gateway) error
	Update(ctx context.Context, r *types.Gateway) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Gateway) (*types.Gateway, error)
	GetAll(ctx context.Context, f *filters.Gateway) ([]types.Gateway, error)
}

type gatewayStorage struct{}

func New() Storage {
	return gatewayStorage{}
}

func (g gatewayStorage) Create(ctx context.Context, r *types.Gateway) error {
	return newSQL().Create(ctx, r)
}

func (g gatewayStorage) Update(ctx context.Context, r *types.Gateway) error {
	return newSQL().Update(ctx, r)
}

func (g gatewayStorage) Delete(ctx context.Context, id string) error {
	return newSQL().Delete(ctx, id)
}

func (g gatewayStorage) Get(ctx context.Context, f *filters.Gateway) (*types.Gateway, error) {
	return newSQL().Get(ctx, f)
}

func (g gatewayStorage) GetAll(ctx context.Context, f *filters.Gateway) ([]types.Gateway, error) {
	return newSQL().GetAll(ctx, f)
}
