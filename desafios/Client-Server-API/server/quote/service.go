package quote

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	httpRequestDefaultTimeout = 200 * time.Millisecond
	persistenceDefaultTimeout = 10 * time.Millisecond
)

// Service expose quote operations
type Service interface {
	// FindLastQuote returns a Quote or an error
	FindLastQuote(context.Context) (*Response, error)
	// Finish the service
	Finish(context.Context) error
}

// internal struct to handle service operations
type service struct {
	persistence *persistence
}

// NewService creates a new service reference
func NewService() (*service, error) {
	persistence, err := newPersistence()
	if err != nil {
		return nil, err
	}
	return &service{persistence: persistence}, nil
}

func (s *service) FindLastQuote(ctx context.Context) (*Response, error) {
	httpClient := http.Client{
		Timeout: httpRequestDefaultTimeout,
	}
	req, err := httpClient.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data Quote
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, persistenceDefaultTimeout)
	defer cancel()
	if err := s.persistence.create(ctx, &CoinEntity{Coin: data.Coin}); err != nil {
		return nil, err
	}

	return &Response{data.Coin.Bid}, nil
}

func (s *service) Finish(context.Context) error {
	if err := s.persistence.close(); err != nil {
		return err
	}
	return nil
}
