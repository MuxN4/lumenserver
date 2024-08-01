package server

import (
	"net/http"
)

type Middleware func(Handler) Handler

func Chain(h Handler, middleware ...Middleware) Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

func (s *Server) ApplyMiddleware(h Handler) Handler {
	return Chain(h, s.middlewares...)
}

func (s *Server) Use(middleware ...Middleware) {
	s.middlewares = append(s.middlewares, middleware...)
}

// Some example middleware functions

// LoggingMiddleware logs information about each request
func LoggingMiddleware(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		println("Request:", r.Method, r.URL.Path)
		next(w, r)
		println("Response sent")
	}
}

// AuthMiddleware is a simple authentication middleware
func AuthMiddleware(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for auth token (this is a very simplistic example)
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// RecoveryMiddleware recovers from panics and returns a 500 error
func RecoveryMiddleware(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}
