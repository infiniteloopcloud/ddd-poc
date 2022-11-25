package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func SetupGateway(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[types.Gateway](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := payment.NewGateway().Setup(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func Update(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[types.Gateway](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := payment.NewGateway().Update(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := payment.NewGateway().Delete(r.Context(), id); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, nil)
}

func Get(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Gateway](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := payment.NewGateway().Get(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Gateway](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := payment.NewGateway().GetAll(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}
