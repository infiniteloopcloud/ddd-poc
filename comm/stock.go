package comm

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type ProductDescriptor interface {
	Create(ctx context.Context, r *types.Product) (*types.Product, error)
	Update(ctx context.Context, r *types.Product) (*types.Product, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Product) (*types.Product, error)
	GetAll(ctx context.Context, f *filters.Product) ([]types.Product, error)
}

type CartDescriptor interface {
	Create(ctx context.Context, r *types.Cart) (*types.Cart, error)
	Update(ctx context.Context, r *types.Cart) (*types.Cart, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Cart) (*types.Cart, error)
	GetAll(ctx context.Context, f *filters.Cart) ([]types.Cart, error)

	AddProductToCart(ctx context.Context, productID, cartID string) error
	DeleteProductFromCart(ctx context.Context, productID, cartID string) error
	Pay(ctx context.Context, r *types.PayCart) error
}
