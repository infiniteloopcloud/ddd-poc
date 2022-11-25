package repository

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

type GatewayStorage interface {
	Create(ctx context.Context, r *types.Gateway) (*types.Gateway, error)
}

type gatewayStorage struct{}

func NewGateway() GatewayStorage {
	return gatewayStorage{}
}

func (g gatewayStorage) Create(ctx context.Context, r *types.Gateway) (*types.Gateway, error) {
	// TODO implement me
	panic("implement me")
}
