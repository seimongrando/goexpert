package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor padrão
	}

	r := setupRouter(getLocationFromZip, getWeather)

	log.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// setupRouter cria e retorna o roteador
func setupRouter(locationFunc func(string) (string, error), weatherFunc func(string) (float64, error)) *gin.Engine {
	r := gin.Default()

	r.GET("/weather/:zipcode", func(c *gin.Context) {
		zipcode := c.Param("zipcode")
		if len(zipcode) != 8 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
			return
		}

		location, err := locationFunc(zipcode)
		if err != nil || location == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
			return
		}

		tempC, err := weatherFunc(location)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch weather data"})
			return
		}

		tempF := tempC*1.8 + 32
		tempK := tempC + 273.15

		c.JSON(http.StatusOK, WeatherResponse{
			TempC: tempC,
			TempF: tempF,
			TempK: tempK,
		})
	})

	return r
}

func getLocationFromZip(zipcode string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("https://viacep.com.br/ws/" + zipcode + "/json/")

	if err != nil || resp.StatusCode() != 200 {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", err
	}

	if _, ok := result["erro"]; ok {
		return "", err
	}

	city, ok := result["localidade"].(string)
	if !ok {
		return "", err
	}

	return city, nil
}

func getWeather(city string) (float64, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParam("q", city).
		SetQueryParam("key", "API+KEY"). // Substitua pela chave da WeatherAPI
		Get("https://api.weatherapi.com/v1/current.json")

	if err != nil || resp.StatusCode() != 200 {
		return 0, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return 0, err
	}

	tempC, ok := result["current"].(map[string]interface{})["temp_c"].(float64)
	if !ok {
		return 0, err
	}

	return tempC, nil
}
