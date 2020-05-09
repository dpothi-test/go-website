package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Main Page </h1>")
	} else if r.URL.Path == "/contact" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Contact Info: <a href=\"test@gmail.com\">link</a>")
	}
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8000", nil)
}
