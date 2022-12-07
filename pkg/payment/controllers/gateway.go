package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/infiniteloopcloud/webshop-poc-ddd/dep"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func SetupGateway(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Gateway](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := dep.Gateway.Setup(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func UpdateGateway(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Gateway](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := dep.Gateway.Update(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func DeleteGateway(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := dep.Gateway.Delete(r.Context(), id); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, nil)
}

func GetGateway(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Gateway](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := dep.Gateway.Get(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}

func GetAllGateway(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Gateway](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := dep.Gateway.GetAll(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}
