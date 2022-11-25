package stock

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

type productService struct {
	productStorage repository.ProductStorage
}

func NewProduct(productStorage repository.ProductStorage) comm.ProductDescriptor {
	return productService{
		productStorage: productStorage,
	}
}

func (p productService) Create(ctx context.Context, r *types.Product) (*types.Product, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return p.productStorage.Create(ctx, r)
}

func (p productService) Update(ctx context.Context, r *types.Product) (*types.Product, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return p.productStorage.Update(ctx, r)
}

func (p productService) Delete(ctx context.Context, id string) error {
	_, err := p.productStorage.Get(ctx, &filters.Product{
		ID: id,
	})
	if err != nil {
		return err
	}
	return p.productStorage.Delete(ctx, id)
}

func (p productService) Get(ctx context.Context, f *filters.Product) (*types.Product, error) {
	return p.productStorage.Get(ctx, f)
}

func (p productService) GetAll(ctx context.Context, f *filters.Product) ([]types.Product, error) {
	return p.productStorage.GetAll(ctx, f)
}