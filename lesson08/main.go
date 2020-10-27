// https://learnku.com/articles/40948
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/wedojava/gowebstudy/lesson08/handler"
	"github.com/wedojava/gowebstudy/lesson08/middleware"
)

func RegisterRoutes(r *mux.Router) {
	r.Use(middleware.Logging())
	rIndex := r.PathPrefix("/index").Subrouter()
	rIndex.Handle("", &handler.HelloHandler{})
	// curl http://localhost:8080/index/display_headers
	rIndex.HandleFunc("/display_headers", handler.DisplayHeadersHandler)
	// curl http://localhost:8080/index/display_url_params\?a\=b\&c\=d\&a\=c
	rIndex.HandleFunc("/display_url_params", handler.DisplayUrlParamsHandler)
	// curl -X POST -d 'username=James&password=123' http://localhost:8080/index/display_form_data
	rIndex.HandleFunc("/display_form_data", handler.DisplayFormDataHandler)
	// curl --cookie "USER_TOKEN=Yes" http://localhost:8080/index/read_cookie
	rIndex.HandleFunc("/read_cookie", handler.ReadCookieHandler)
	// curl -X POST -d '{"name": "James", "age": 18}' -H "Content-Type: application/json" http://localhost:8080/index/parse_json_request
	rIndex.HandleFunc("/parse_json_request", handler.DisplayPersonHandler)

	rUser := r.PathPrefix("/user").Subrouter()
	rUser.HandleFunc("/name/{name}/country/{country}", handler.ShowVisitorInfo)
	rUser.Use(middleware.Method("GET"))

	rView := r.PathPrefix("/view").Subrouter()
	rView.HandleFunc("/index", handler.ShowIndexView)

	fs := http.FileServer(http.Dir("assets/"))
	serveFileHandler := http.StripPrefix("/static/", fs)
	r.PathPrefix("/static/").Handler(serveFileHandler)
}

func main() {
	r := mux.NewRouter()
	RegisterRoutes(r)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// create signal reciver
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server: ", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}
