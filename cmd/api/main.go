package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"Authentication/internaal/app"
	"Authentication/internaal/httpserver"
)

func main() {
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Startup failed: %v", err)
	}
	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Panicf("Shutdown warning: %v", err)
		}
	}()

	router := httpserver.NewRouter(a)
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Api running on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server closed")
			return
		}
		log.Fatalf("Server Error: %v", err)
	}
}
