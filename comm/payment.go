package comm

type PaymentDescriptor interface {
	Create() error
}