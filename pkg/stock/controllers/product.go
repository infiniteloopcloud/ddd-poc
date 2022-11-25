package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[types.Product](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	product, err := stock.NewProduct().Create(r.Context(), &req)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[types.Product](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	product, err := stock.NewProduct().Update(r.Context(), &req)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := stock.NewProduct().Delete(r.Context(), id); err != nil {
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
	resp, err := stock.NewProduct().Get(r.Context(), &f)
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
	resp, err := stock.NewProduct().GetAll(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}
