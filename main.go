package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", endpoint)
	http.ListenAndServe(":3000", r)
}

func endpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}
