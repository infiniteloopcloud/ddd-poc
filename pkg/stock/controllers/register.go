package controllers

import "github.com/go-chi/chi/v5"

func Register(r chi.Router) {
	r.Post("/products", CreateProduct)
	r.Put("/products/{id}", UpdateProduct)
	r.Delete("/products/{id}", DeleteProduct)
	r.Get("/products/{id}", GetProduct)
	r.Get("/products", GetAllProducts)

	r.Post("/carts", CreateCart)
	r.Put("/carts/{id}", UpdateCart)
	r.Delete("/carts/{id}", DeleteCart)
	r.Get("/carts/{id}", GetCart)
	r.Get("/carts", GetAllCarts)
}
