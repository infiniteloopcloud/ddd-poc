package dep

import (
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/comm"
)

var (
	User     comm.UserDescriptor
	Merchant comm.MerchantDescriptor
)

func RegisterUser(descriptor comm.UserDescriptor) {
	User = descriptor
}

func RegisterMerchant(descriptor comm.MerchantDescriptor) {
	Merchant = descriptor
}
