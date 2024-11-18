package api

import (
	"log"
	"net/http"
)

type wrappedHandler struct {
	http.ResponseWriter
	statusCode int
}

type Middleware func(http.Handler) http.Handler

func Chain(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, middleware := range xs {
			next = middleware(next)
		}
		return next
	}
}

func (w *wrappedHandler) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &wrappedHandler{w, http.StatusOK}
		next.ServeHTTP(wrapped, r)
		log.Println(wrapped.statusCode, r.Method, r.URL)
	})
}

func Guard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
