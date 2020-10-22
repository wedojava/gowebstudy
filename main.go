// https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzUzNTY5MzU2MA==&action=getalbum&album_id=1323498303014780929&scene=173&from_msgid=2247484271&from_itemidx=1&count=3#wechat_redirect
package main

import (
	"fmt"
	"net/http"
)

// # Method 1
/*
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
*/

// #Method 2
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
