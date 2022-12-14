package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Merchant](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := account.NewMerchant().Create(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Merchant](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := account.NewMerchant().Update(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := account.NewMerchant().Delete(r.Context(), id); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, nil)
}

func GetMerchant(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Merchant](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := account.NewMerchant().Get(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}

func GetMerchants(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Merchant](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := account.NewMerchant().GetAll(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}
