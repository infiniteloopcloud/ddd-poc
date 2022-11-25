package merchant

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type sql struct{}

func newSQL() Storage {
	return sql{}
}

func (s sql) Create(ctx context.Context, m *types.Merchant) error {
	//TODO implement me
	panic("implement me")
}

func (s sql) Update(ctx context.Context, m *types.Merchant) error {
	//TODO implement me
	panic("implement me")
}

func (s sql) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s sql) Get(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	//TODO implement me
	panic("implement me")
}

func (s sql) GetAll(ctx context.Context, r *filters.Merchant) (types.Merchant, error) {
	//TODO implement me
	panic("implement me")
}
