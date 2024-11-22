package api

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

type wrappedHandler struct {
	http.ResponseWriter
	statusCode int
}

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Special handling for WebSocket endpoint
			if r.URL.Path == "/ws" {
				// Skip other middleware for WebSocket connections
				next.ServeHTTP(w, r)
				return
			}

			// Apply middleware chain for non-WebSocket requests
			handler := next
			for i := len(middlewares) - 1; i >= 0; i-- {
				handler = middlewares[i](handler)
			}
			handler.ServeHTTP(w, r)
		})
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

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		next.ServeHTTP(w, r)
	})
}

type ResponseWriterWrapper struct {
	http.ResponseWriter
}

func Hijack(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		if hijacker, ok := w.(http.Hijacker); ok {
			hijacker.Hijack()
		} else {
			http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		}
	})
}

func (w *ResponseWriterWrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, http.ErrNotSupported
}
