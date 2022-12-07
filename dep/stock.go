package dep

import (
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/comm"
)

var (
	Product comm.ProductDescriptor
	Cart    comm.CartDescriptor
)

func RegisterProduct(descriptor comm.ProductDescriptor) {
	Product = descriptor
}

func RegisterCart(descriptor comm.CartDescriptor) {
	Cart = descriptor
}
