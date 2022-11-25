package user

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type Storage interface {
	Create(ctx context.Context, u *types.User) error
	Update(ctx context.Context, u *types.User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters.User) (types.User, error)
	GetAll(ctx context.Context, r *filters.User) (types.User, error)
}

type userStorage struct{}

func NewUser() Storage {
	return userStorage{}
}

func (s userStorage) Create(ctx context.Context, u *types.User) error {
	return newSQL().Create(ctx, u)
}

func (s userStorage) Update(ctx context.Context, u *types.User) error {
	return newSQL().Update(ctx, u)
}

func (s userStorage) Delete(ctx context.Context, id string) error {
	return newSQL().Delete(ctx, id)
}

func (s userStorage) Get(ctx context.Context, r *filters.User) (types.User, error) {
	return newSQL().Get(ctx, r)
}

func (s userStorage) GetAll(ctx context.Context, r *filters.User) (types.User, error) {
	return newSQL().GetAll(ctx, r)
}
