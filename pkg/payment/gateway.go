package payment

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type gatewayService struct {
	gatewayStorage repository.GatewayStorage
}

func NewGateway() comm.GatewayDescriptor {
	return gatewayService{
		gatewayStorage: repository.NewGateway(),
	}
}

func (g gatewayService) Setup(ctx context.Context, r *types.Gateway) error {
	if err := r.Validate(); err != nil {
		return err
	}

	// TODO implement some kind of a setup logic
	return g.gatewayStorage.Create(ctx, r)
}

func (g gatewayService) Update(ctx context.Context, r *types.Gateway) error {
	//TODO implement me
	panic("implement me")
}

func (g gatewayService) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (g gatewayService) Get(ctx context.Context, f *filters.Gateway) (*types.Gateway, error) {
	//TODO implement me
	panic("implement me")
}

func (g gatewayService) GetAll(ctx context.Context, f *filters.Gateway) ([]types.Gateway, error) {
	//TODO implement me
	panic("implement me")
}
