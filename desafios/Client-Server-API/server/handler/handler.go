package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"server/quote"
)

type Handler interface {
	QuoteHandler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	quoteService quote.Service
}

func NewHandler(quoteService quote.Service) *handler {
	return &handler{
		quoteService: quoteService,
	}
}

func (h *handler) QuoteHandler(w http.ResponseWriter, r *http.Request) {
	lastQuote, err := h.quoteService.FindLastQuote(context.Background())
	if err != nil {
		log.Printf("error finding last quote: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(lastQuote)
	if err != nil {
		log.Printf("error marshalling quote: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
}
