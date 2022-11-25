package controllers

import (
	"fmt"
	"net/http"

	"github.com/infiniteloopcloud/webshop-poc-ddd/pkg/account"
	"github.com/infiniteloopcloud/webshop-poc-ddd/types"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}
	if err := account.NewUser().Create(r.Context(), user); err != nil {
		// error response
	}
	// created response
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &types.User{}
	if err := account.NewUser().Update(r.Context(), user); err != nil {
		// error response
	}
	// created response
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := ""
	if err := account.NewUser().Delete(r.Context(), id); err != nil {
		// error response
	}
	// success response
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	resp, err := account.NewUser().Get(r.Context(), nil)
	if err != nil {

	}
	fmt.Println(resp)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := account.NewUser().Get(r.Context(), nil)
	if err != nil {

	}
	fmt.Println(resp)
}
