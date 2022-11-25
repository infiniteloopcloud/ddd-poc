package account

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type userService struct {
	storage repository.AccountStorage
}

func NewUser() comm.UserDescriptor {
	return userService{
		storage: repository.New(),
	}
}

func (s userService) Create(ctx context.Context, u *types.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.storage.User.Create(ctx, u)
}

func (s userService) Update(ctx context.Context, u *types.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.storage.User.Update(ctx, u)
}

func (s userService) Delete(ctx context.Context, id string) error {
	return s.storage.User.Delete(ctx, id)
}

func (s userService) Get(ctx context.Context, r *filters.User) (types.User, error) {
	return s.storage.User.Get(ctx, r)
}

func (s userService) GetAll(ctx context.Context, r *filters.User) (types.User, error) {
	return s.storage.User.GetAll(ctx, r)
}
