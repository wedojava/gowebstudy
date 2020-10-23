package main

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	r.Handle("/", &HelloHandler{})
	r.HandleFunc("/welcome", WelcomeHandler)
	rIndex := r.PathPrefix("/index").Subrouter()
	rIndex.Handle("", &HelloHandler{})
	rUser := r.PathPrefix("/user").Subrouter()
	rUser.HandleFunc("/name/{name}/country/{country}", ShowVisitorInfo)

	// rh := r.Handle("/", &HelloHandler{})
	// rw := r.HandleFunc("/welcome", WelcomeHandler)
	// rh.Methods("GET").Host("localhost").Schemes("http")
	// rw.Methods("GET").Host("localhost").Schemes("http")
	// rIndex.Methods("GET").Host("localhost").Schemes("http")
	// rUser.Methods("GET").Host("localhost").Schemes("http")
}
