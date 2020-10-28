package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type HelloHandler struct{}

func (*HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	// fmt.Fprintf(w, "%v", []int{0, 1, 2}[0:5]) // err log test
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

func DisplayHeadersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s URL: %s Protocol: %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])
}

func DisplayUrlParamsHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.URL.Query() {
		fmt.Fprintf(w, "ParamName %q, Value %q\n", k, v)
		fmt.Fprintf(w, "ParamName %q, Get Value %q\n", k, r.URL.Query().Get(k))
	}
}

func DisplayFormDataHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	for key, values := range r.Form {
		fmt.Fprintf(w, "Formfield %q, Values %q\n", key, values)
		fmt.Fprintf(w, "form field %q, Value %q\n", key, r.FormValue(key))
	}
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		fmt.Fprintf(w, "Cookie field %q, Value %q\n", cookie.Name, cookie.Value)
	}
}

type Person struct {
	Name string
	Age  int
}

func DisplayPersonHandler(w http.ResponseWriter, r *http.Request) {
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Person: %+v", p)
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // 32M
	var buf bytes.Buffer

	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])

	io.Copy(&buf, file)
	contents := buf.String()
	fmt.Println(contents)
	buf.Reset()

	return
}

type Teacher struct {
	Name    string
	Subject string
}

type Student struct {
	Id      int
	Name    string
	Country string
}

type Rooster struct {
	Teacher  Teacher
	Students []Student
}

func ShowIndexView(w http.ResponseWriter, r *http.Request) {
	teacher := Teacher{
		Name:    "Alex",
		Subject: "Physics",
	}
	students := []Student{
		{Id: 1001, Name: "Peter", Country: "China"},
		{Id: 1002, Name: "Jeniffer", Country: "Sweden"},
	}
	rooster := Rooster{
		Teacher:  teacher,
		Students: students,
	}

	tmpl, err := template.ParseFiles("./views/layout.html",
		"./views/nav.html",
		"./views/content.html")
	if err != nil {
		fmt.Println("Error " + err.Error())
	}
	tmpl.Execute(w, rooster)
}
