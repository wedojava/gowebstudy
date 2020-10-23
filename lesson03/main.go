// https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484172&idx=1&sn=6dc988c86c3572a8092bdc79feb8d4e8&chksm=fa80d29bcdf75b8d06fc56366352671131c06e1c299a4929a56d7f5ab7137d1e1aec213c5e40&scene=178&cur_album_id=1323498303014780929#rd
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
	// r := mux.NewRouter()
	// r.Handle("/", &HelloHandler{})
	// r.HandleFunc("/welcome", WelcomeHandler)
	// r.HandleFunc("/names/{name}/countries/{country}", func(w http.ResponseWriter, r *http.Request) {
	//         vs := mux.Vars(r)
	//         n := vs["name"]
	//         c := vs["country"]
	//         fmt.Fprintf(w, "This guy named %s, was coming from %s.", n, c)
	// }).Methods("GET").Host("localhost").Schemes("http")

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
