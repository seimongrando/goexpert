package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type CepInput struct {
	Cep string `json:"cep"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Valor padrão para o Serviço A
	}

	r := gin.Default()

	r.POST("/cep", func(c *gin.Context) {
		serviceBURL := os.Getenv("SERVICE_B_URL")
		if serviceBURL == "" {
			serviceBURL = "http://service-b:8080" // Fallback para Docker Compose
		}

		var input CepInput
		if err := c.ShouldBindJSON(&input); err != nil || len(input.Cep) != 8 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
			return
		}

		// Chamar o Serviço B
		responseBody, responseStatus, err := callServiceB(serviceBURL, input.Cep)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not process request"})
			return
		}

		c.Data(responseStatus, "application/json", responseBody)
	})

	log.Printf("Starting Service A on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func callServiceB(serviceBURL, cep string) ([]byte, int, error) {
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	url := serviceBURL + "/weather/" + cep

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
