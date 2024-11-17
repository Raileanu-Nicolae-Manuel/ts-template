package api

import "net/http"

func Get(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("GET "+path, handler)
}

func Post(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("POST "+path, handler)
}

func Put(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("PUT "+path, handler)
}

func Delete(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("DELETE "+path, handler)
}

func Patch(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("PATCH "+path, handler)
}

func Options(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("OPTIONS "+path, handler)
}
