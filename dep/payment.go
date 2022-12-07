package dep

import "github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/comm"

var (
	Transaction comm.TransactionDescriptor
	Gateway     comm.GatewayDescriptor
)

func RegisterTransaction(descriptor comm.TransactionDescriptor) {
	Transaction = descriptor
}

func RegisterGateway(descriptor comm.GatewayDescriptor) {
	Gateway = descriptor
}
