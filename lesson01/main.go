// https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484112&idx=1&sn=79d0d3167d0d962fe41ec00cdafffbb0&chksm=fa80d347cdf75a51183182f14622af766538ca0c5335012e5e1cc50b100e78f2954fa3943770&scene=21#wechat_redirect
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// # Stage 1
/*
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
*/

// # Stage 2
/*
type HelloHandlerStruct struct {
	content string
}

func (handler *HelloHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.content)
}

func main() {
	http.Handle("/", &HelloHandlerStruct{content: "Hello World"})
	http.ListenAndServe(":8080", nil)
}
*/

// # Stage 3
// type WelcomeHandlerStruct struct {
// }
//
// func HelloHandler(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Hello World")
// }
//
// func (*WelcomeHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Welcome")
// }
//
// func main() {
//         mux := http.NewServeMux()
//         mux.HandleFunc("/", HelloHandler)
//         mux.Handle("/welcome", &WelcomeHandlerStruct{})
//         http.ListenAndServe(":8080", mux)
// }

// # Stage 4

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
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
