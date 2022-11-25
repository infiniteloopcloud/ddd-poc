package stock

import (
	"context"
	"errors"

	comm2 "github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type cartService struct {
	transaction comm2.TransactionDescriptor
	storage     repository.StockStorage
}

func NewCart(payment comm2.TransactionDescriptor) comm.CartDescriptor {
	return cartService{
		transaction: payment,
		storage:     repository.New(),
	}
}

func (cs cartService) Create(ctx context.Context, r *proto.Cart) (*proto.Cart, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	return cs.storage.Cart.Create(ctx, r)
}

func (cs cartService) Update(ctx context.Context, r *proto.Cart) (*proto.Cart, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	return cs.storage.Cart.Update(ctx, r)
}

func (cs cartService) Delete(ctx context.Context, id string) error {
	if _, err := cs.storage.Cart.Get(ctx, &filters.Cart{
		ID: id,
	}); err != nil {
		return err
	}
	return cs.storage.Cart.Delete(ctx, id)
}

func (cs cartService) Get(ctx context.Context, f *filters.Cart) (*proto.Cart, error) {
	return cs.storage.Cart.Get(ctx, f)
}

func (cs cartService) GetAll(ctx context.Context, f *filters.Cart) ([]proto.Cart, error) {
	return cs.storage.Cart.GetAll(ctx, f)
}

func (cs cartService) AddProductToCart(ctx context.Context, productID, cartID string) error {
	p, err := cs.storage.Product.Get(ctx, &filters.Product{
		ID: productID,
	})
	if err != nil {
		return err
	}
	c, err := cs.storage.Cart.Get(ctx, &filters.Cart{
		ID: cartID,
	})
	if err != nil {
		return err
	}
	c.Products = append(c.Products, *p)

	_, err = cs.storage.Cart.Update(ctx, c)
	return err
}

func (cs cartService) DeleteProductFromCart(ctx context.Context, productID, cartID string) error {
	c, err := cs.storage.Cart.Get(ctx, &filters.Cart{
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

	_, err = cs.storage.Cart.Update(ctx, c)
	return err
}

func (cs cartService) Pay(ctx context.Context, r *proto.PayCart) error {
	if err := r.Validate(); err != nil {
		return err
	}
	return cs.transaction.Create(ctx, &r.Transaction)
}
