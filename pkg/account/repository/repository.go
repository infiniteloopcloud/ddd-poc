package repository

import (
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/repository/merchant"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/repository/user"
)

type AccountStorage struct {
	User     user.Storage
	Merchant merchant.Storage
}

func New() AccountStorage {
	return AccountStorage{
		Merchant: merchant.NewMerchant(),
		User:     user.NewUser(),
	}
}
