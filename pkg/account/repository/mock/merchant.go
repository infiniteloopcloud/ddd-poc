package mock

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type Merchant struct {
	CreateFunc func(ctx context.Context, m *proto.Merchant) error
	UpdateFunc func(ctx context.Context, m *proto.Merchant) error
	DeleteFunc func(ctx context.Context, id string) error
	GetFunc    func(ctx context.Context, r *filters.Merchant) (proto.Merchant, error)
	GetAllFunc func(ctx context.Context, r *filters.Merchant) (proto.Merchant, error)
}

func (s Merchant) Create(ctx context.Context, m *proto.Merchant) error {
	return s.CreateFunc(ctx, m)
}

func (s Merchant) Update(ctx context.Context, m *proto.Merchant) error {
	return s.UpdateFunc(ctx, m)
}

func (s Merchant) Delete(ctx context.Context, id string) error {
	return s.DeleteFunc(ctx, id)
}

func (s Merchant) Get(ctx context.Context, r *filters.Merchant) (proto.Merchant, error) {
	return s.GetFunc(ctx, r)
}

func (s Merchant) GetAll(ctx context.Context, r *filters.Merchant) (proto.Merchant, error) {
	return s.GetAllFunc(ctx, r)
}
