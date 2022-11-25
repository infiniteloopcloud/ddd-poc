package controllers

import (
	"net/http"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

func SetupGateway(w http.ResponseWriter, r *http.Request) {
	gw := &types.Gateway{}
	if err := payment.NewGateway().Setup(r.Context(), gw); err != nil {
		// error response
	}
	// created response
}
