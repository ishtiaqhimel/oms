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

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/ishtiaqhimel/oms/oms-server/handlers"
	"github.com/ishtiaqhimel/oms/oms-server/initializers"
	mw "github.com/ishtiaqhimel/oms/oms-server/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.AuthMiddleware)

	r.Post("/api/order", handlers.CreateOrder)
	r.Get("/api/order/{id}", handlers.GetOrder)
	r.Put("/api/order/{id}", handlers.UpdateOrder)
	r.Delete("/api/order/{id}", handlers.DeleteOrder)

	// handler for documentation
	// opts := openapi.RedocOpts{BasePath: "/auth", SpecURL: "/swagger.yaml"}
	// sh := openapi.Redoc(opts, nil)

	// r.Handle("/auth/docs", sh)
	// r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	bindAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := &http.Server{
		Addr:         bindAddr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Starting server on", bindAddr)
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Error listening to server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(tc)
}

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}
