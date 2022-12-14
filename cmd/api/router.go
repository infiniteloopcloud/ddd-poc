package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	accountControllers "github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account/controllers"
	paymentControllers "github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/controllers"
	stockControllers "github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock/controllers"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/settings"
)

func Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		accountControllers.Register(r)
		paymentControllers.Register(r)
		stockControllers.Register(r)
	})

	return http.ListenAndServe(settings.Get().HttpAddress, r)
}
