package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"server/handler"
	"server/quote"
	"syscall"
)

// to execute the server CGO_ENABLED must be 1 | e.g CGO_ENABLED=1
func main() {
	// create service
	quoteService, err := quote.NewService()
	if err != nil {
		log.Printf("failed to start quote service")
		log.Fatal(err)
	}

	// server quoteHandler
	quoteHandler := handler.NewHandler(quoteService)

	// server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", quoteHandler.QuoteHandler)

	// starting the server
	go func() {
		log.Printf("server listening on 8080")
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	// handle signals to graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	log.Printf("server is shutting down")
	// finishing the service
	if err := quoteService.Finish(context.Background()); err != nil {
		log.Printf("error finishing quote service. Err: %v\n", err)
	}
	log.Printf("server stopped")
}
