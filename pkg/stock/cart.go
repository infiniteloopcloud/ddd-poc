package stock

import "github.com/infiniteloopcloud/webshop-poc-ddd/comm"

type cartService struct {
	Payment comm.PaymentDescriptor
}

func NewCart(p comm.PaymentDescriptor) comm.CartDescriptor {
	return cartService{
		Payment: p,
	}
}

func (p cartService) Pay() {
	p.Payment.Create()
}

func (p cartService) Create() {
	// TODO implement me
	panic("implement me")
}

func (p cartService) Update() {
	// TODO implement me
	panic("implement me")
}

func (p cartService) Delete() {
	// TODO implement me
	panic("implement me")
}
