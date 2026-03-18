package main

import (
	"fmt"
	"net/http"
	"go-todo/internal/handler"
	"go-todo/internal/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	handler := handler.NewTodoHandler(store)

	handler.RegisterRoutes()

	fmt.Println("server run on :443")
	err := http.ListenAndServeTLS(":443",
		"/etc/xray/cert.pem",
		"/etc/xray/key.pem",
		nil)
	if err != nil {
		panic(err)
	}
}
