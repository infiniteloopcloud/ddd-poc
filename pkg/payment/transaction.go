package payment

import (
	"context"
	"fmt"

	comm2 "github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/comm"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

type transactionService struct {
	paymentStorage repository.PaymentStorage

	merchant comm2.MerchantDescriptor
	user     comm2.UserDescriptor
}

func NewTransaction(merchant comm2.MerchantDescriptor, user comm2.UserDescriptor) comm.TransactionDescriptor {
	return transactionService{
		paymentStorage: repository.New(),
		merchant:       merchant,
		user:           user,
	}
}

func (t transactionService) Create(ctx context.Context, r *types.Transaction) error {
	if err := r.Validate(); err != nil {
		return err
	}
	user, err := t.user.Get(ctx, &filters.User{
		ID: r.UserID,
	})
	if err != nil {
		return err
	}
	merchant, err := t.merchant.Get(ctx, &filters.Merchant{ID: r.MerchantID})
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

func (t transactionService) Get(ctx context.Context, f *filters.Transaction) (*types.Transaction, error) {
	return t.paymentStorage.Transaction.Get(ctx, f)
}

func (t transactionService) GetAll(ctx context.Context, f *filters.Transaction) ([]types.Transaction, error) {
	return t.paymentStorage.Transaction.GetAll(ctx, f)
}
