package transaction

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type sqlHandler struct {
}

func newSQL() Storage {
	return sqlHandler{}
}

func (s sqlHandler) Create(ctx context.Context, r *types.Transaction) error {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error) {
	// TODO implement me
	panic("implement me")
}

func (s sqlHandler) GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error) {
	// TODO implement me
	panic("implement me")
}
