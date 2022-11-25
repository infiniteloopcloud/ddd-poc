package comm

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	filters2 "github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type UserDescriptor interface {
	Create(ctx context.Context, u *types.User) error
	Update(ctx context.Context, u *types.User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters2.User) (types.User, error)
	GetAll(ctx context.Context, r *filters2.User) (types.User, error)
}

type MerchantDescriptor interface {
	Create(ctx context.Context, u *types.Merchant) error
	Update(ctx context.Context, u *types.Merchant) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters2.Merchant) (types.Merchant, error)
	GetAll(ctx context.Context, r *filters2.Merchant) (types.Merchant, error)
}
