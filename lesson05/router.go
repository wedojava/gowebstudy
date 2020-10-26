package main

import (
	"github.com/gorilla/mux"
	"github.com/wedojava/gowebstudy/lesson05/handler"
	"github.com/wedojava/gowebstudy/lesson05/middleware"
)

func RegisterRoutes(r *mux.Router) {
	r.Use(middleware.Logging())
	rIndex := r.PathPrefix("/index").Subrouter()
	rIndex.Handle("", &handler.HelloHandler{})

	rUser := r.PathPrefix("/user").Subrouter()
	rUser.HandleFunc("/name/{name}/country/{country}", handler.ShowVisitorInfo)
	rUser.Use(middleware.Method("GET"))
}
