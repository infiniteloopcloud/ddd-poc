package comm

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type TransactionDescriptor interface {
	Create(ctx context.Context, r *proto.Transaction) error
	Get(ctx context.Context, f *filters.Transaction) (*proto.Transaction, error)
	GetAll(ctx context.Context, f *filters.Transaction) ([]proto.Transaction, error)
}

type GatewayDescriptor interface {
	Setup(ctx context.Context, r *proto.Gateway) error
	Update(ctx context.Context, r *proto.Gateway) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Gateway) (*proto.Gateway, error)
	GetAll(ctx context.Context, f *filters.Gateway) ([]proto.Gateway, error)
}
