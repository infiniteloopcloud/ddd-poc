package repository

import (
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository/cart"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/repository/product"
)

type StockStorage struct {
	Cart    cart.Storage
	Product product.Storage
}

func New() StockStorage {
	return StockStorage{
		Cart:    cart.New(),
		Product: product.New(),
	}
}
