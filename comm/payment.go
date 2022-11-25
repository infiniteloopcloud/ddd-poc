package comm

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type TransactionDescriptor interface {
	Create(ctx context.Context, r *types.Transaction) error
	Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error)
	GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error)
}

type GatewayDescriptor interface {
	Setup(ctx context.Context, r *types.Gateway) (*types.Gateway, error)
}
