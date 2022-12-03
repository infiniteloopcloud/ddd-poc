package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/stock"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Cart](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	product, err := stock.NewCart(payment.NewTransaction(account.NewMerchant(), account.NewUser())).Create(r.Context(), &req)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, product)
}

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Cart](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	product, err := stock.NewCart(payment.NewTransaction(account.NewMerchant(), account.NewUser())).Update(r.Context(), &req)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, product)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := stock.NewCart(payment.NewTransaction(account.NewMerchant(), account.NewUser())).Delete(r.Context(), id); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, nil)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Cart](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := stock.NewCart(payment.NewTransaction(account.NewMerchant(), account.NewUser())).Get(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}

func GetAllCarts(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Cart](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := stock.NewCart(payment.NewTransaction(account.NewMerchant(), account.NewUser())).GetAll(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}
