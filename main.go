package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", endpoint)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func endpoint(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello world\n")
}