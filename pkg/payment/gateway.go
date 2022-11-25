package stock

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

type gatewayService struct {
	gatewayStorage repository.GatewayStorage
}

func NewGateway(gatewayStorage repository.GatewayStorage) comm.GatewayDescriptor {
	return gatewayService{
		gatewayStorage: gatewayStorage,
	}
}

func (g gatewayService) Setup(ctx context.Context, r *types.Gateway) (*types.Gateway, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	// TODO implement some kind of a setup logic

	return g.gatewayStorage.Create(ctx, r)
}
