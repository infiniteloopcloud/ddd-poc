package stock

import "github.com/infiniteloopcloud/webshop-poc-ddd/comm"

type productService struct{}

func NewProduct() comm.ProductDescriptor {
	return productService{}
}

func (p productService) Create() {
	// TODO implement me
	panic("implement me")
}

func (p productService) Update() {
	// TODO implement me
	panic("implement me")
}

func (p productService) Delete() {
	// TODO implement me
	panic("implement me")
}
