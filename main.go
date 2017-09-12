package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/hello", endpoint)
	http.ListenAndServe(":3000", nil)
}

func endpoint(w http.ResponseWriter, r *http.Request) {
	urlString := r.URL.String()

	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	m, _ := url.ParseQuery(u.RawQuery)
	nameString := "Hello" + " " + m["name"][0] + "!"
	fmt.Fprintln(w, nameString)
}

// First assignment by Yuri, a simple app that will return "Hello World"
// when called by cURL.
// Second assignment: rewrite so that a url query hello?name=Roger returns
// Hello Roger.
// 226269