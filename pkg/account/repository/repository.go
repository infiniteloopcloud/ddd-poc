package repository

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/repository/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

type UserStorage interface {
	Create(ctx context.Context, u *types.User) error
	Update(ctx context.Context, u *types.User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, r *filters.User) (types.User, error)
	GetAll(ctx context.Context, r *filters.User) (types.User, error)
}

type userStorage struct{}

func NewUser() UserStorage {
	return userStorage{}
}

func (s userStorage) Create(ctx context.Context, u *types.User) error {
	return newUserSQL().Create(ctx, u)
}

func (s userStorage) Update(ctx context.Context, u *types.User) error {
	return newUserSQL().Update(ctx, u)
}

func (s userStorage) Delete(ctx context.Context, id string) error {
	return newUserSQL().Delete(ctx, id)
}

func (s userStorage) Get(ctx context.Context, r *filters.User) (types.User, error) {
	return newUserSQL().Get(ctx, r)
}

func (s userStorage) GetAll(ctx context.Context, r *filters.User) (types.User, error) {
	return newUserSQL().GetAll(ctx, r)
}

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
