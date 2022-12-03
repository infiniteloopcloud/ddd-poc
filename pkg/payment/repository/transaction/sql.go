package transaction

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type sqlHandler struct {
}

func newSQL() Storage {
	return sqlHandler{}
}

func (s sqlHandler) Create(ctx context.Context, r *proto.Transaction) error {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) Update(ctx context.Context, r *proto.Transaction) error {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) Get(ctx context.Context, f *filters.Transaction) (*proto.Transaction, error) {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) GetAll(ctx context.Context, f *filters.Transaction) ([]proto.Transaction, error) {
	// TODO implement me
	panic("implement me")
}
