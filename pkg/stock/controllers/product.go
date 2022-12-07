package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/infiniteloopcloud/webshop-poc-ddd/dep"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Product](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	product, err := dep.Product.Create(r.Context(), &req)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Product](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	product, err := dep.Product.Update(r.Context(), &req)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := dep.Product.Delete(r.Context(), id); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, nil)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Product](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := dep.Product.Get(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Product](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := dep.Product.GetAll(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}
