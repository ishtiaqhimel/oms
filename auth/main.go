package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ishtiaqhimel/oms/auth/handlers"
	"github.com/ishtiaqhimel/oms/auth/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/signup", handlers.SignUp)
	r.Post("/login", handlers.Login)

	bindAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	l := log.New(os.Stdout, "auth-api: ", log.LstdFlags)
	server := &http.Server{
		Addr:         bindAddr,
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on", bindAddr)
		if err := server.ListenAndServe(); err != nil {
			l.Printf("Error listening to server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(tc)
}
