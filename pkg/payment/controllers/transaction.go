package controllers

import (
	"net/http"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types/filters"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	tx := &types.Transaction{}
	if err := payment.NewTransaction().Create(r.Context(), tx); err != nil {
		// error response
	}
	// created response
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	f := &filters.Transaction{}
	tx, err := payment.NewTransaction().Get(r.Context(), f)
	if err != nil {
		// error response
	}
	tx = tx
	// created response
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	f := &filters.Transaction{}
	txs, err := payment.NewTransaction().GetAll(r.Context(), f)
	if err != nil {
		// error response
	}
	txs = txs
	// created response
}
