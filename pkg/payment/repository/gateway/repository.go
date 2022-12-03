package gateway

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type Storage interface {
	Create(ctx context.Context, r *proto.Gateway) error
	Update(ctx context.Context, r *proto.Gateway) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Gateway) (*proto.Gateway, error)
	GetAll(ctx context.Context, f *filters.Gateway) ([]proto.Gateway, error)
}

type gatewayStorage struct{}

func New() Storage {
	return gatewayStorage{}
}

func (g gatewayStorage) Create(ctx context.Context, r *proto.Gateway) error {
	return newSQL().Create(ctx, r)
}

func (g gatewayStorage) Update(ctx context.Context, r *proto.Gateway) error {
	return newSQL().Update(ctx, r)
}

func (g gatewayStorage) Delete(ctx context.Context, id string) error {
	return newSQL().Delete(ctx, id)
}

func (g gatewayStorage) Get(ctx context.Context, f *filters.Gateway) (*proto.Gateway, error) {
	return newSQL().Get(ctx, f)
}

func (g gatewayStorage) GetAll(ctx context.Context, f *filters.Gateway) ([]proto.Gateway, error) {
	return newSQL().GetAll(ctx, f)
}
