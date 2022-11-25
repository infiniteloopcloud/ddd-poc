package main

import (
	"fmt"
	"os"

	"github.com/infiniteloopcloud/webshop-poc-ddd/cmd/api"
)

func main() {
	if err := api.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
