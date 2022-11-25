package user

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type Storage interface {
	Create(ctx context.Context, u *proto.User) error
	Update(ctx context.Context, u *proto.User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters.User) (proto.User, error)
	GetAll(ctx context.Context, r *filters.User) (proto.User, error)
}

type userStorage struct{}

func NewUser() Storage {
	return userStorage{}
}

func (s userStorage) Create(ctx context.Context, u *proto.User) error {
	return newSQL().Create(ctx, u)
}

func (s userStorage) Update(ctx context.Context, u *proto.User) error {
	return newSQL().Update(ctx, u)
}

func (s userStorage) Delete(ctx context.Context, id string) error {
	return newSQL().Delete(ctx, id)
}

func (s userStorage) Get(ctx context.Context, r *filters.User) (proto.User, error) {
	return newSQL().Get(ctx, r)
}

func (s userStorage) GetAll(ctx context.Context, r *filters.User) (proto.User, error) {
	return newSQL().GetAll(ctx, r)
}
