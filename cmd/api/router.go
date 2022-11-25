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
		r.Get("/merchants/{id}", accountControllers.GetMerchant)

		r.Post("/transactions", paymentControllers.CreateTransaction)
		r.Get("/transactions/{id}", paymentControllers.GetTransaction)
		r.Get("/transactions", paymentControllers.GetAllTransactions)

		r.Post("/gateways", paymentControllers.SetupGateway)
		r.Put("/gateways/{id}", paymentControllers.UpdateGateway)
		r.Delete("/gateways/{id}", paymentControllers.DeleteGateway)
		r.Get("/gateways/{id}", paymentControllers.GetGateway)
		r.Get("/gateways", paymentControllers.GetAllGateway)

		r.Post("/products", stockControllers.CreateProduct)
		r.Put("/products/{id}", stockControllers.UpdateProduct)
		r.Delete("/products/{id}", stockControllers.DeleteProduct)
		r.Get("/products/{id}", stockControllers.GetProduct)
		r.Get("/products", stockControllers.GetAllProducts)

		r.Post("/carts", stockControllers.CreateCart)
		r.Put("/carts/{id}", stockControllers.UpdateCart)
		r.Delete("/carts/{id}", stockControllers.DeleteCart)
		r.Get("/carts/{id}", stockControllers.GetCart)
		r.Get("/carts", stockControllers.GetAllCarts)
	})

	return http.ListenAndServe(settings.Get().HttpAddress, r)
}
