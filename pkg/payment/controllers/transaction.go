package controllers

import (
	"net/http"

	"github.com/infiniteloopcloud/webshop-poc-ddd/dep"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto"
	"github.com/infiniteloopcloud/webshop-poc-ddd/proto/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[proto.Transaction](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := dep.Transaction.Create(r.Context(), &req); err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, req)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Transaction](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	tx, err := dep.Transaction.Get(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, tx)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	f, err := httpio.Bind[filters.Transaction](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	tx, err := dep.Transaction.GetAll(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, tx)
}
