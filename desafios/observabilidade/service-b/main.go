package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type WeatherResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

var tracer trace.Tracer

func main() {
	initTracer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor padrão para o Serviço B
	}

	r := setupRouter(getLocationFromZip, getWeather)

	log.Printf("Starting Service B on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initTracer() {
	zipkinURL := os.Getenv("OTEL_EXPORTER_ZIPKIN_ENDPOINT")
	if zipkinURL == "" {
		zipkinURL = "http://zipkin:9411/api/v2/spans" // Fallback para o Docker Compose
	}

	// Configurar o exportador do Zipkin
	exporter, err := zipkin.New(zipkinURL)
	if err != nil {
		log.Fatalf("Failed to create Zipkin exporter: %v", err)
	}

	// Configurar o provedor de rastreamento
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewSchemaless(attribute.String("service.name", "service-b"))),
	)

	otel.SetTracerProvider(tp)
	tracer = otel.Tracer("service-b")
}

func setupRouter(locationFunc func(context.Context, string) (string, error), weatherFunc func(context.Context, string) (float64, error)) *gin.Engine {
	r := gin.Default()

	r.GET("/weather/:zipcode", func(c *gin.Context) {
		ctx := c.Request.Context()
		_, span := tracer.Start(ctx, "WeatherHandler")
		defer span.End()

		zipcode := c.Param("zipcode")
		if len(zipcode) != 8 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
			return
		}

		location, err := locationFunc(ctx, zipcode)
		if err != nil || location == "" {
			if err != nil {
				span.SetAttributes(attribute.String("error", err.Error()))
			} else {
				span.SetAttributes(attribute.String("error", "invalid zipcode"))
			}
			c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
			return
		}

		tempC, err := weatherFunc(ctx, location)
		if err != nil {
			span.SetAttributes(attribute.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch weather data"})
			return
		}

		tempF := tempC*1.8 + 32
		tempK := tempC + 273.15

		c.JSON(http.StatusOK, WeatherResponse{
			City:  location,
			TempC: tempC,
			TempF: tempF,
			TempK: tempK,
		})
	})

	return r
}

func getLocationFromZip(ctx context.Context, zipcode string) (string, error) {
	client := resty.New().SetTransport(otelhttp.NewTransport(http.DefaultTransport))
	_, span := tracer.Start(ctx, "getLocationFromZip")
	defer span.End()

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("https://viacep.com.br/ws/" + zipcode + "/json/")

	if err != nil || resp.StatusCode() != 200 {
		span.SetAttributes(attribute.String("error", err.Error()))
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		return "", err
	}

	if _, ok := result["erro"]; ok {
		span.SetAttributes(attribute.String("error", "invalid zip code"))
		return "", err
	}

	city, ok := result["localidade"].(string)
	if !ok {
		span.SetAttributes(attribute.String("error", "city not found"))
		return "", err
	}

	span.SetAttributes(attribute.String("city", city))
	return city, nil
}

func getWeather(ctx context.Context, city string) (float64, error) {
	client := resty.New().SetTransport(otelhttp.NewTransport(http.DefaultTransport))
	_, span := tracer.Start(ctx, "getWeather")
	defer span.End()

	weatherToken := os.Getenv("WEATHER_API_TOKEN") // Substitua pela chave da WeatherAPI
	if weatherToken == "" {
		weatherToken = "CHAVE_VALIDA_TESTE"
	}

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParam("q", city).
		SetQueryParam("key", weatherToken).
		Get("https://api.weatherapi.com/v1/current.json")

	if err != nil || resp.StatusCode() != 200 {
		span.SetAttributes(attribute.String("error", err.Error()))
		return 0, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		return 0, err
	}

	tempC, ok := result["current"].(map[string]interface{})["temp_c"].(float64)
	if !ok {
		span.SetAttributes(attribute.String("error", "temperature not found"))
		return 0, err
	}

	span.SetAttributes(attribute.Float64("temperature_c", tempC))
	return tempC, nil
}
