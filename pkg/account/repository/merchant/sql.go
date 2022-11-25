package merchant

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type merchantSQL struct{}

func newMerchantSQL() MerchantStorage {
	return merchantSQL{}
}

func (s merchantSQL) Create(ctx context.Context, m *types.Merchant) error {
	//TODO implement me
	panic("implement me")
}

func (s merchantSQL) Update(ctx context.Context, m *types.Merchant) error {
	//TODO implement me
	panic("implement me")
}

func (s merchantSQL) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s merchantSQL) Get(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	//TODO implement me
	panic("implement me")
}

func (s merchantSQL) GetAll(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	//TODO implement me
	panic("implement me")
}
