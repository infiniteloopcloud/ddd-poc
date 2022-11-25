package controllers

import (
	"github.com/go-chi/chi/v5"
)

func Register(r chi.Router) {
	r.Post("/merchants", CreateMerchant)
	r.Put("/merchants/{id}", UpdateMerchant)
	r.Delete("/merchants/{id}", DeleteMerchant)
	r.Get("/merchants/{id}", GetMerchant)
	r.Get("/merchants", GetMerchants)

	r.Post("/users", CreateUser)
	r.Put("/users/{id}", UpdateUser)
	r.Delete("/users/{id}", DeleteUser)
	r.Get("/users/{id}", GetUser)
	r.Get("/users", GetUsers)
}
