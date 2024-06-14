package quote

import (
	"fmt"
	"log"
	"os"
)

type persistence struct{}

func newPersistence() *persistence {
	return &persistence{}
}

func (p *persistence) create(bid string) error {
	log.Printf("starting bid file creation.")
	defer log.Printf("finishing bid file creation.")

	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	value := fmt.Sprintf("Dólar: %s\n", bid)
	if _, err := file.Write([]byte(value)); err != nil {
		return err
	}

	return nil
}
