package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type brasilAPI struct{}

type brasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func NewBrasilAPI() *brasilAPI {
	return &brasilAPI{}
}

func (b *brasilAPI) FindAddress(cep string) (*Address, error) {
	httpClient := http.Client{}
	req, err := httpClient.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data brasilAPIResponse
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return &Address{
		From:         "BrasilAPI",
		Cep:          data.Cep,
		State:        data.State,
		City:         data.City,
		Neighborhood: data.Neighborhood,
		Street:       data.Street,
	}, nil
}
