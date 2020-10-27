package main

// func RegisterRoutes(r *mux.Router) {
//         r.Use(middleware.Logging())
//         rIndex := r.PathPrefix("/index").Subrouter()
//         rIndex.Handle("", &handler.HelloHandler{})
//         // curl http://localhost:8080/index/display_headers
//         rIndex.HandleFunc("/display_headers", handler.DisplayHeadersHandler)
//         // curl http://localhost:8080/index/display_url_params\?a\=b\&c\=d\&a\=c
//         rIndex.HandleFunc("/display_url_params", handler.DisplayUrlParamsHandler)
//         // curl -X POST -d 'username=James&password=123' http://localhost:8080/index/display_form_data
//         rIndex.HandleFunc("/display_form_data", handler.DisplayFormDataHandler)
//         // curl --cookie "USER_TOKEN=Yes" http://localhost:8080/index/read_cookie
//         rIndex.HandleFunc("/read_cookie", handler.ReadCookieHandler)
//         // curl -X POST -d '{"name": "James", "age": 18}' -H "Content-Type: application/json" http://localhost:8080/index/parse_json_request
//         rIndex.HandleFunc("/parse_json_request", handler.DisplayPersonHandler)
//
//         rUser := r.PathPrefix("/user").Subrouter()
//         rUser.HandleFunc("/name/{name}/country/{country}", handler.ShowVisitorInfo)
//         rUser.Use(middleware.Method("GET"))
//
//         rView := r.PathPrefix("/view").Subrouter()
//         rView.HandleFunc("/index", handler.ShowIndexView)
// }
