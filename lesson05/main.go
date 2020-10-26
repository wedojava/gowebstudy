// https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484180&idx=1&sn=b66497c5428c25577068f18132b2d59d&chksm=fa80d283cdf75b95cc49d08c56d0fa9b00c47d0457c894be3ca3ea02bd3c404cfb5312fa1d93&scene=178&cur_album_id=1323498303014780929#rd
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

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

// func RegisterRoutes(r *mux.Router) {
//         r.Use(middleware.Logging())
//         rIndex := r.PathPrefix("/index").Subrouter()
//         rIndex.Handle("", &handler.HelloHandler{})
//
//         rUser := r.PathPrefix("/user").Subrouter()
//         rUser.HandleFunc("/name/{name}/country/{country}", handler.ShowVisitorInfo)
//         rUser.Use(middleware.Method("GET"))
// }
