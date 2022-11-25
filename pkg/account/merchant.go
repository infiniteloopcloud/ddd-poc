package account

import (
	"context"
	"time"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type merchantService struct {
	storage repository.AccountStorage
}

func NewMerchant() comm.MerchantDescriptor {
	return merchantService{
		storage: repository.New(),
	}
}

func (s merchantService) Create(ctx context.Context, u *types.Merchant) error {
	if err := s.storage.Merchant.Create(ctx, u); err != nil {
		return err
	}
	// assume that we want to create a user as well
	return s.storage.User.Create(ctx, &types.User{
		AccountTypeID: u.ID,
		Name:          "Test",
		Email:         "test@test.com",
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	})
}

func (s merchantService) Update(ctx context.Context, u *types.Merchant) error {
	return s.storage.Merchant.Update(ctx, u)
}

func (s merchantService) Delete(ctx context.Context, id string) error {
	return s.storage.Merchant.Delete(ctx, id)
}

func (s merchantService) Get(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	return s.storage.Merchant.Get(ctx, r)
}

func (s merchantService) GetAll(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	return s.storage.Merchant.GetAll(ctx, r)
}
