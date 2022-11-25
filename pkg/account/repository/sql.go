package repository

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	filters2 "github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type userSQL struct{}

func newUserSQL() UserStorage {
	return userSQL{}
}

func (s userSQL) Create(ctx context.Context, u *types.User) error {
	//TODO implement me
	panic("implement me")
}

func (s userSQL) Update(ctx context.Context, u *types.User) error {
	//TODO implement me
	panic("implement me")
}

func (s userSQL) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s userSQL) Get(ctx context.Context, r *filters2.User) (types.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s userSQL) GetAll(ctx context.Context, r *filters2.User) (types.User, error) {
	//TODO implement me
	panic("implement me")
}

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

func (s merchantSQL) Get(ctx context.Context, r *filters2.Merchant) (types.Merchant, error) {
	//TODO implement me
	panic("implement me")
}

func (s merchantSQL) GetAll(ctx context.Context, r *filters2.Merchant) (types.Merchant, error) {
	//TODO implement me
	panic("implement me")
}
