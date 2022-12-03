package mock

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type Merchant struct {
	CreateFunc func(ctx context.Context, u *proto.Merchant) error
	UpdateFunc func(ctx context.Context, u *proto.Merchant) error
	DeleteFunc func(ctx context.Context, id string) error
	GetFunc    func(ctx context.Context, r *filters.Merchant) (proto.Merchant, error)
	GetAllFunc func(ctx context.Context, r *filters.Merchant) (proto.Merchant, error)
}

func (m Merchant) Create(ctx context.Context, u *proto.Merchant) error {
	return m.CreateFunc(ctx, u)
}

func (m Merchant) Update(ctx context.Context, u *proto.Merchant) error {
	return m.UpdateFunc(ctx, u)
}

func (m Merchant) Delete(ctx context.Context, id string) error {
	return m.DeleteFunc(ctx, id)
}

func (m Merchant) Get(ctx context.Context, r *filters.Merchant) (proto.Merchant, error) {
	return m.GetFunc(ctx, r)
}

func (m Merchant) GetAll(ctx context.Context, r *filters.Merchant) (proto.Merchant, error) {
	return m.GetAllFunc(ctx, r)
}
