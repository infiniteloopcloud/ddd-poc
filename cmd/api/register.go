package api

import (
	"github.com/infiniteloopcloud/webshop-poc-ddd/dep"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock"
)

func init() {
	dep.RegisterCart(stock.NewCart())
	dep.RegisterProduct(stock.NewProduct())
	dep.RegisterTransaction(payment.NewTransaction())
	dep.RegisterGateway(payment.NewGateway())
	dep.RegisterMerchant(account.NewMerchant())
	dep.RegisterUser(account.NewUser())
}
