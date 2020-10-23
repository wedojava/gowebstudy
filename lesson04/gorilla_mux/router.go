package main

import (
	"github.com/gorilla/mux"
	"github.com/wedojava/gowebstudy/lesson04/gorilla_mux/handler"
	"github.com/wedojava/gowebstudy/lesson04/gorilla_mux/middleware"
)

func RegisterRoutes(r *mux.Router) {
	r.Use(middleware.Logging())
	rIndex := r.PathPrefix("/index").Subrouter()
	rIndex.Handle("", &handler.HelloHandler{})

	rUser := r.PathPrefix("/user").Subrouter()
	rUser.HandleFunc("/name/{name}/country/{country}", handler.ShowVisitorInfo)
	rUser.Use(middleware.Method("GET"))
}
