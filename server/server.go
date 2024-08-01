package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	port         int
	router       *Router
	middlewares  []Middleware
	errorHandler ErrorHandler
	server       *http.Server
}

type ServerOption func(*Server)

func NewServer(port int, options ...ServerOption) *Server {
	s := &Server{
		port:   port,
		router: NewRouter(),
	}

	for _, option := range options {
		option(s)
	}

	return s
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	fmt.Printf("Server starting on port %d\n", s.port)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Use(middleware Middleware) {
	s.middlewares = append(s.middlewares, middleware)
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
