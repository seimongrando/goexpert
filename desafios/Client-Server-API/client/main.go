package main

import (
	"client/quote"
	"context"
	"log"
)

func main() {
	log.Printf("starting quote client.")
	defer log.Printf("finishing quote client.")

	quoteService := quote.NewService()
	quoteResponse, err := quoteService.FindLastQuote(context.Background())
	if err != nil {
		log.Printf("failed to get quote information")
		log.Fatal(err)
	}

	if err := quoteService.PersistBID(quoteResponse.Bid); err != nil {
		log.Printf("failed to persist BID")
		log.Fatal(err)
	}

	log.Printf("quote information acquired. Quote Response: %v", quoteResponse)
}
