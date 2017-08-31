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
	s := "localhost:3000/hello?name=Roger"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	m, _ := url.ParseQuery(u.RawQuery)
	nameString := "Hello" + " " + m["name"][0] + "!"
	fmt.Fprint(w, nameString)

}


// First assignment by Yuri, a simple app that will return "Hello World"
// when called by cURL.
// Revised to use native http router
