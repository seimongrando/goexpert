package service

import (
	"fmt"
)

type CEP interface {
	FindAddress(cep string) (*Address, error)
}

type Address struct {
	From         string
	Cep          string
	State        string
	City         string
	Neighborhood string
	Street       string
}

func (a *Address) String() string {
	return fmt.Sprintf("Information acquired from service: %s | Address: %s, %s, %s - %s, %s",
		a.From, a.Street, a.Neighborhood, a.Cep, a.City, a.State)
}
