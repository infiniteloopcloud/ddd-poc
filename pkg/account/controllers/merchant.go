package controllers

import (
	"fmt"
	"net/http"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	user := &types.Merchant{}
	if err := account.NewMerchant().Create(r.Context(), user); err != nil {
		// error response
	}
	// created response
}

func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	user := &types.Merchant{}
	if err := account.NewMerchant().Update(r.Context(), user); err != nil {
		// error response
	}
	// created response
}

func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	id := ""
	if err := account.NewMerchant().Delete(r.Context(), id); err != nil {
		// error response
	}
	// success response
}

func GetMerchant(w http.ResponseWriter, r *http.Request) {
	resp, err := account.NewMerchant().Get(r.Context(), nil)
	if err != nil {

	}
	fmt.Println(resp)
}

func GetMerchants(w http.ResponseWriter, r *http.Request) {
	resp, err := account.NewMerchant().Get(r.Context(), nil)
	if err != nil {

	}
	fmt.Println(resp)
}
