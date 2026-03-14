package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "API run 🚀")
	})

	fmt.Println("server run on :443")
	err := http.ListenAndServeTLS(":443",
		"/etc/xray/cert.pem",
		"/etc/xray/key.pem",
		nil)
	if err != nil {
		panic(err)
	}
}
