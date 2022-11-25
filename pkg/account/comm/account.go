package comm

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type UserDescriptor interface {
	Create(ctx context.Context, u *proto.User) error
	Update(ctx context.Context, u *proto.User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters.User) (proto.User, error)
	GetAll(ctx context.Context, r *filters.User) (proto.User, error)
}

type MerchantDescriptor interface {
	Create(ctx context.Context, u *proto.Merchant) error
	Update(ctx context.Context, u *proto.Merchant) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters.Merchant) (proto.Merchant, error)
	GetAll(ctx context.Context, r *filters.Merchant) (proto.Merchant, error)
}
