package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/infiniteloopcloud/webshop-poc-ddd/dep"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.User](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := dep.User.Create(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.User](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := dep.User.Update(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := dep.User.Delete(r.Context(), id); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, nil)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.User](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := dep.User.Get(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.User](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	resp, err := dep.User.Get(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, resp)
}
