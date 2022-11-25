package stock

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type productService struct {
	storage repository.StockStorage
}

func NewProduct() comm.ProductDescriptor {
	return productService{
		storage: repository.New(),
	}
}

func (p productService) Create(ctx context.Context, r *types.Product) (*types.Product, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return p.storage.Product.Create(ctx, r)
}

func (p productService) Update(ctx context.Context, r *types.Product) (*types.Product, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	return p.storage.Product.Update(ctx, r)
}

func (p productService) Delete(ctx context.Context, id string) error {
	_, err := p.storage.Product.Get(ctx, &filters.Product{
		ID: id,
	})
	if err != nil {
		return err
	}
	return p.storage.Product.Delete(ctx, id)
}

func (p productService) Get(ctx context.Context, f *filters.Product) (*types.Product, error) {
	return p.storage.Product.Get(ctx, f)
}

func (p productService) GetAll(ctx context.Context, f *filters.Product) ([]types.Product, error) {
	return p.storage.Product.GetAll(ctx, f)
}
