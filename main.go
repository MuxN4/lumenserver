package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lumenserver/handlers"
	"lumenserver/server"
)

func main() {
	s := server.NewServer(8080)

	s.Use(server.LoggingMiddleware)
	s.Use(server.RecoveryMiddleware)

	// Register routes
	s.Router.GET("/", handlers.HomeHandler)
	s.Router.GET("/time", handlers.TimeHandler)
	s.Router.GET("/users", handlers.UsersHandler)
	s.Router.POST("/users", handlers.CreateUserHandler)
	s.Router.NotFound(handlers.NotFoundHandler)

	// Set up graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.Run(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	fmt.Println("Server is running on http://localhost:8080")

	//Interrupt signal
	<-stop

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server gracefully stopped")
}
