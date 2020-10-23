package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HelloHandler struct{}

func (*HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func ShowVisitorInfo(w http.ResponseWriter, r *http.Request) {
	vs := mux.Vars(r)
	n := vs["name"]
	c := vs["country"]
	fmt.Fprintf(w, "This guy named %s was coming from %s.", n, c)
}
