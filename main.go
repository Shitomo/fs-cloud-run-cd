package main

import (
	"fmt"
	"net/http"
)

type Server string

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/", Server("Hello CloudRun with ko."))
	http.ListenAndServe(":8080", nil)
}
