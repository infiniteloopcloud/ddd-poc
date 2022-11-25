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
	userRepository     repository.UserStorage
	merchantRepository repository.MerchantStorage
}

func NewMerchant() comm.MerchantDescriptor {
	return merchantService{
		userRepository:     NewUser(),
		merchantRepository: NewMerchant(),
	}
}

func (s merchantService) Create(ctx context.Context, u *types.Merchant) error {
	if err := s.merchantRepository.Create(ctx, u); err != nil {
		return err
	}
	return s.userRepository.Create(ctx, &types.User{
		AccountTypeID: u.ID,
		Name:          "Test",
		Email:         "test@test.com",
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	})
}

func (s merchantService) Update(ctx context.Context, u *types.Merchant) error {
	return s.merchantRepository.Update(ctx, u)
}

func (s merchantService) Delete(ctx context.Context, id string) error {
	return s.merchantRepository.Delete(ctx, id)
}

func (s merchantService) Get(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	return s.merchantRepository.Get(ctx, r)
}

func (s merchantService) GetAll(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	return s.merchantRepository.GetAll(ctx, r)
}
