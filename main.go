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
	http.Handle("/", Server("Hello World."))
	http.ListenAndServe("localhost:8080", nil)
}
