package main

import (
	"fmt"
	"net/http"
	"go-todo/internal/handler"
	"go-todo/internal/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	h := handler.NewTodoHandler(store)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	loggedMux := handler.LoggingMiddleware(mux)

	fmt.Println("server run on :443")
	err := http.ListenAndServeTLS(":443",
		"/etc/xray/cert.pem",
		"/etc/xray/key.pem",
		loggedMux)
	if err != nil {
		panic(err)
	}
}
