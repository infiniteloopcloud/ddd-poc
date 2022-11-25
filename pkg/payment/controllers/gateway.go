package controllers

import (
	"net/http"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

var gatewayService = payment.NewGateway(repository.NewGateway())

func SetupGateway(w http.ResponseWriter, r *http.Request) {
	gw := &types.Gateway{}
	res, err := gatewayService.Setup(r.Context(), gw)
	if err != nil {
		// error response
	}
	res = res
	// created response
}
