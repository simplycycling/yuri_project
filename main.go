package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", endpoint)
	http.ListenAndServe(":3000", nil)
}

func endpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}

// First assignment by Yuri, a simple app that will return "Hello World"
// when called by cURL.
// Revised to use native http router
