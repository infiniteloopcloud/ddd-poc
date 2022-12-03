package comm

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type ProductDescriptor interface {
	Create(ctx context.Context, r *proto.Product) (*proto.Product, error)
	Update(ctx context.Context, r *proto.Product) (*proto.Product, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Product) (*proto.Product, error)
	GetAll(ctx context.Context, f *filters.Product) ([]proto.Product, error)
}

type CartDescriptor interface {
	Create(ctx context.Context, r *proto.Cart) (*proto.Cart, error)
	Update(ctx context.Context, r *proto.Cart) (*proto.Cart, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, f *filters.Cart) (*proto.Cart, error)
	GetAll(ctx context.Context, f *filters.Cart) ([]proto.Cart, error)

	AddProductToCart(ctx context.Context, productID, cartID string) error
	DeleteProductFromCart(ctx context.Context, productID, cartID string) error
	Pay(ctx context.Context, r *proto.PayCart) error
}
