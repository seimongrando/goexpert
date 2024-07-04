package main

import (
	"log"
	"multithreading/service"
	"os"
	"time"
)

func main() {
	log.Println("Cep service starting")
	defer log.Println("Cep service done")

	// Validate the arguments
	if len(os.Args) < 2 {
		log.Println("Usage: go run main.go <cep>")
		return
	}

	// Get the arguments
	cep := os.Args[1]

	// Response channel
	response := make(chan *service.Address)

	// find the address inside the GoRoutine
	go findAddress(cep, service.NewBrasilAPI(), response)
	go findAddress(cep, service.NewViaCEP(), response)

	// wait for the 1st response
	select {
	case res := <-response:
		{
			log.Println(res)
		}
	case <-time.After(time.Second):
		log.Println("Request has timed out - over 1 second")
	}

}

func findAddress(cep string, cepService service.CEP, cepResponse chan<- *service.Address) {
	address, err := cepService.FindAddress(cep)
	if err != nil {
		log.Printf("API FindAddress error: %s", err.Error())
		return
	}
	cepResponse <- address
}
