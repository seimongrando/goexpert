package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type viaCEP struct{}

type viaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCEP() *viaCEP {
	return &viaCEP{}
}

func (v *viaCEP) FindAddress(cep string) (*Address, error) {
	httpClient := http.Client{}
	req, err := httpClient.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data viaCEPResponse
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return &Address{
		From:         "ViaCEP",
		Cep:          data.Cep,
		State:        data.Uf,
		City:         data.Localidade,
		Neighborhood: data.Bairro,
		Street:       data.Logradouro,
	}, nil
}
