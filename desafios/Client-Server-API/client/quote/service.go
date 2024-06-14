package quote

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	httpRequestDefaultTimeout = 300 * time.Millisecond
)

// Service expose quote operations
type Service interface {
	// FindLastQuote returns a Quote or an error
	FindLastQuote(context.Context) (*Response, error)
	// PersistBID BID
	PersistBID(bid string) error
}

// internal struct to handle service operations
type service struct {
	persistence *persistence
	httpClient  *http.Client
}

// NewService creates a new service reference
func NewService() *service {
	return &service{
		persistence: newPersistence(),
		httpClient:  http.DefaultClient,
	}
}

func (s *service) FindLastQuote(ctx context.Context) (*Response, error) {
	log.Printf("starting quote information call.")
	defer log.Printf("finishing quote information call.")

	ctx, cancel := context.WithTimeout(ctx, httpRequestDefaultTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, err
	}

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data Response
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *service) PersistBID(bid string) error {
	return s.persistence.create(bid)
}
