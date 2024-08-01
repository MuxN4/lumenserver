package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	port         int
	Router       *Router
	middlewares  []Middleware
	errorHandler ErrorHandler
	server       *http.Server
}

type ServerOption func(*Server)

func NewServer(port int) *Server {
	s := &Server{
		port:        port,
		middlewares: []Middleware{},
	}
	s.Router = NewRouter(s)
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s,
	}

	fmt.Printf("Server starting on port %d\n", s.port)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) GetRouter() *Router {
	return s.Router
}

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

func WithErrorHandler(handler ErrorHandler) ServerOption {
	return func(s *Server) {
		s.errorHandler = handler
	}
}

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
