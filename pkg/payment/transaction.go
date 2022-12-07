package payment

import (
	"context"
	"fmt"

	"github.com/infiniteloopcloud/webshop-poc-ddd/dep"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
)

type transactionService struct {
	paymentStorage repository.PaymentStorage
}

func NewTransaction() comm.TransactionDescriptor {
	return transactionService{
		paymentStorage: repository.New(),
	}
}

func (t transactionService) Create(ctx context.Context, r *proto.Transaction) error {
	if err := r.Validate(); err != nil {
		return err
	}
	user, err := dep.User.Get(ctx, &filters.User{
		ID: r.UserID,
	})
	if err != nil {
		return err
	}
	merchant, err := dep.Merchant.Get(ctx, &filters.Merchant{ID: r.MerchantID})
	if err != nil {
		return err
	}
	// do something with merchant and user
	fmt.Println(user, merchant)
	gateway, err := t.paymentStorage.Gateway.Get(ctx, &filters.Gateway{
		ID: r.GatewayID,
	})
	if err != nil {
		return err
	}
	if err = t.paymentStorage.Transaction.Create(ctx, r); err != nil {
		return err
	}

	// send transaction to gateway
	err = gatewayService{}.SendTransaction(ctx, gateway, r)
	if err != nil {
		return err
	}

	return t.paymentStorage.Transaction.Update(ctx, r)
}

func (t transactionService) Get(ctx context.Context, f *filters.Transaction) (*proto.Transaction, error) {
	return t.paymentStorage.Transaction.Get(ctx, f)
}

func (t transactionService) GetAll(ctx context.Context, f *filters.Transaction) ([]proto.Transaction, error) {
	return t.paymentStorage.Transaction.GetAll(ctx, f)
}
