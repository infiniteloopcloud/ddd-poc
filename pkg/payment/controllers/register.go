package controllers

import "github.com/go-chi/chi/v5"

func Register(r chi.Router) {
	r.Post("/transactions", CreateTransaction)
	r.Get("/transactions/{id}", GetTransaction)
	r.Get("/transactions", GetAllTransactions)

	r.Post("/gateways", SetupGateway)
	r.Put("/gateways/{id}", UpdateGateway)
	r.Delete("/gateways/{id}", DeleteGateway)
	r.Get("/gateways/{id}", GetGateway)
	r.Get("/gateways", GetAllGateway)
}
