package controllers

import (
	"net/http"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/utils/httpio"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	req, err := httpio.Bind[types.Transaction](r.Body)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	if err := payment.NewTransaction(account.NewMerchant(), account.NewUser()).Create(r.Context(), &req); err != nil {
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
	tx, err := payment.NewTransaction(account.NewMerchant(), account.NewUser()).Get(r.Context(), &f)
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
	tx, err := payment.NewTransaction(account.NewMerchant(), account.NewUser()).GetAll(r.Context(), &f)
	if err != nil {
		httpio.ResponseBadRequest(w)
		return
	}
	httpio.ResponseSuccess(w, tx)
}
