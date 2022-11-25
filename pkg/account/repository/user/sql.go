package user

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
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

func (s userSQL) Get(ctx context.Context, r *filters.User) (types.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s userSQL) GetAll(ctx context.Context, r *filters.User) (types.User, error) {
	//TODO implement me
	panic("implement me")
}
