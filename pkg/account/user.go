package account

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/repository/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

type userService struct {
	userRepository repository.UserStorage
}

func NewUser() comm.UserDescriptor {
	return userService{
		userRepository: repository.NewUser(),
	}
}

func (s userService) Create(ctx context.Context, u *types.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.userRepository.Create(ctx, u)
}

func (s userService) Update(ctx context.Context, u *types.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.userRepository.Update(ctx, u)
}

func (s userService) Delete(ctx context.Context, id string) error {
	return s.userRepository.Delete(ctx, id)
}

func (s userService) Get(ctx context.Context, r *filters.User) (types.User, error) {
	return s.userRepository.Get(ctx, r)
}

func (s userService) GetAll(ctx context.Context, r *filters.User) (types.User, error) {
	return s.userRepository.GetAll(ctx, r)
}
