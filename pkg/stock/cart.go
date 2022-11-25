package stock

import (
	"context"
	"errors"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

type cartService struct {
	payment        comm.PaymentDescriptor
	cartStorage    repository.CartStorage
	productStorage repository.ProductStorage
}

func NewCart(
	payment comm.PaymentDescriptor,
	cartStorage repository.CartStorage,
	productStorage repository.ProductStorage,
) comm.CartDescriptor {
	return cartService{
		payment:        payment,
		cartStorage:    cartStorage,
		productStorage: productStorage,
	}
}

func (cs cartService) Create(ctx context.Context, r *types.Cart) (*types.Cart, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	return cs.cartStorage.Create(ctx, r)
}

func (cs cartService) Update(ctx context.Context, r *types.Cart) (*types.Cart, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	return cs.cartStorage.Update(ctx, r)
}

func (cs cartService) Delete(ctx context.Context, id string) error {
	if _, err := cs.cartStorage.Get(ctx, &filters.Cart{
		ID: id,
	}); err != nil {
		return err
	}
	return cs.cartStorage.Delete(ctx, id)
}

func (cs cartService) Get(ctx context.Context, f *filters.Cart) (*types.Cart, error) {
	return cs.cartStorage.Get(ctx, f)
}

func (cs cartService) GetAll(ctx context.Context, f *filters.Cart) ([]types.Cart, error) {
	return cs.cartStorage.GetAll(ctx, f)
}

func (cs cartService) AddProductToCart(ctx context.Context, productID, cartID string) error {
	p, err := cs.productStorage.Get(ctx, &filters.Product{
		ID: productID,
	})
	if err != nil {
		return err
	}
	c, err := cs.cartStorage.Get(ctx, &filters.Cart{
		ID: cartID,
	})
	if err != nil {
		return err
	}
	c.Products = append(c.Products, *p)

	_, err = cs.cartStorage.Update(ctx, c)
	return err
}

func (cs cartService) DeleteProductFromCart(ctx context.Context, productID, cartID string) error {
	c, err := cs.cartStorage.Get(ctx, &filters.Cart{
		ID: cartID,
	})
	if err != nil {
		return err
	}
	var foundIndex = -1
	for i, p := range c.Products {
		if p.ID == productID {
			foundIndex = i
			break
		}
	}
	if foundIndex < 0 {
		return errors.New("product is not in the cart")
	}

	c.Products = append(c.Products[:foundIndex], c.Products[foundIndex+1:]...)

	_, err = cs.cartStorage.Update(ctx, c)
	return err
}

func (cs cartService) Pay(ctx context.Context, r *types.PayCart) error {
	if err := r.Validate(); err != nil {
		return err
	}
	return cs.payment.Create()
}
