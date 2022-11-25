package merchant

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type Storage interface {
	Create(ctx context.Context, m *types.Merchant) error
	Update(ctx context.Context, m *types.Merchant) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters.Merchant) (types.Merchant, error)
	GetAll(ctx context.Context, r *filters.Merchant) (types.Merchant, error)
}

type merchantStorage struct{}

func NewMerchant() Storage {
	return merchantStorage{}
}

func (s merchantStorage) Create(ctx context.Context, m *types.Merchant) error {
	return newSQL().Create(ctx, m)
}

func (s merchantStorage) Update(ctx context.Context, m *types.Merchant) error {
	return newSQL().Update(ctx, m)
}

func (s merchantStorage) Delete(ctx context.Context, id string) error {
	return newSQL().Delete(ctx, id)
}

func (s merchantStorage) Get(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	return newSQL().Get(ctx, r)
}

func (s merchantStorage) GetAll(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	return newSQL().GetAll(ctx, r)
}
