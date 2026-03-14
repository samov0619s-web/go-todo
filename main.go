package main;

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "API run 🚀")
	})

	fmt.Println("server run on :443")
	http.ListenAndServe(":443", nil)
}
