package main

import (
	"fmt"
	"net/http"

	"github.com/hive-bootcamp/go-rest-api-homework/internal/route"
)

func main() {
	r := route.New()

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
