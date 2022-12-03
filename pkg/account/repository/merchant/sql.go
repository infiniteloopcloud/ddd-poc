package merchant

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type sql struct{}

func newSQL() Storage {
	return sql{}
}

func (s sql) Create(ctx context.Context, m *proto.Merchant) error {
	// TODO implement me
	panic("implement me")
}

func (s sql) Update(ctx context.Context, m *proto.Merchant) error {
	// TODO implement me
	panic("implement me")
}

func (s sql) Delete(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}

func (s sql) Get(ctx context.Context, r *filters.Merchant) (proto.Merchant, error) {
	// TODO implement me
	panic("implement me")
}

func (s sql) GetAll(ctx context.Context, r *filters.Merchant) (proto.Merchant, error) {
	// TODO implement me
	panic("implement me")
}
