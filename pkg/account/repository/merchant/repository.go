package merchant

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type MerchantStorage interface {
	Create(ctx context.Context, m *types.Merchant) error
	Update(ctx context.Context, m *types.Merchant) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters.Merchant) (types.Merchant, error)
	GetAll(ctx context.Context, r *filters.Merchant) (types.Merchant, error)
}

type merchantStorage struct{}

func NewMerchant() MerchantStorage {
	return merchantStorage{}
}

func (s merchantStorage) Create(ctx context.Context, m *types.Merchant) error {
	return newMerchantSQL().Create(ctx, m)
}

func (s merchantStorage) Update(ctx context.Context, m *types.Merchant) error {
	return newMerchantSQL().Update(ctx, m)
}

func (s merchantStorage) Delete(ctx context.Context, id string) error {
	return newMerchantSQL().Delete(ctx, id)
}

func (s merchantStorage) Get(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	return newMerchantSQL().Get(ctx, r)
}

func (s merchantStorage) GetAll(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	return newMerchantSQL().GetAll(ctx, r)
}
