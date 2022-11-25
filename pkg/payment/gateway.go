package payment

import (
	"context"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type gatewayService struct {
	storage repository.PaymentStorage
}

func NewGateway() comm.GatewayDescriptor {
	return gatewayService{
		storage: repository.New(),
	}
}

func (g gatewayService) Setup(ctx context.Context, r *proto.Gateway) error {
	if err := r.Validate(); err != nil {
		return err
	}

	// TODO implement some kind of a setup logic
	return g.storage.Gateway.Create(ctx, r)
}

func (g gatewayService) Update(ctx context.Context, r *proto.Gateway) error {
	if err := r.Validate(); err != nil {
		return err
	}

	return g.storage.Gateway.Update(ctx, r)
}

func (g gatewayService) Delete(ctx context.Context, id string) error {
	return g.storage.Gateway.Delete(ctx, id)
}

func (g gatewayService) Get(ctx context.Context, f *filters.Gateway) (*proto.Gateway, error) {
	return g.storage.Gateway.Get(ctx, f)
}

func (g gatewayService) GetAll(ctx context.Context, f *filters.Gateway) ([]proto.Gateway, error) {
	return g.storage.Gateway.GetAll(ctx, f)
}

func (g gatewayService) SendTransaction(ctx context.Context, gateway *proto.Gateway, r *proto.Transaction) error {
	// map types.Transaction into gateway specific request, send the transaction, map gateway response into
	// types.Transaction
	return nil
}
