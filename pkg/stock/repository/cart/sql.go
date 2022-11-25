package cart

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type sqlHandler struct{}

func newSQL() Storage {
	return sqlHandler{}
}

func (s sqlHandler) Create(ctx context.Context, r *proto.Cart) (*proto.Cart, error) {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) Update(ctx context.Context, r *proto.Cart) (*proto.Cart, error) {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) Delete(ctx context.Context, id string) error {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) Get(ctx context.Context, f *filters.Cart) (*proto.Cart, error) {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) GetAll(ctx context.Context, f *filters.Cart) ([]proto.Cart, error) {
	// TODO implement me
	panic("implement me")
}
