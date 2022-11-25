package controllers

import (
	"net/http"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository"
	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/payment/repository/filters"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

var transactionService = payment.NewTransaction(repository.NewTransaction())

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	tx := &types.Transaction{}
	if err := transactionService.Create(r.Context(), tx); err != nil {
		// error response
	}
	// created response
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	f := &filters.Transaction{}
	tx, err := transactionService.Get(r.Context(), f)
	if err != nil {
		// error response
	}
	tx = tx
	// created response
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	f := &filters.Transaction{}
	txs, err := transactionService.GetAll(r.Context(), f)
	if err != nil {
		// error response
	}
	txs = txs
	// created response
}
