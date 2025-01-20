package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Teste com table-driven tests
func TestWeatherAPI(t *testing.T) {
	mockGetLocationFromZip := func(zipcode string) (string, error) {
		if zipcode == "00000000" {
			return "", errors.New("invalid zip")
		}
		return "São Paulo", nil
	}

	mockGetWeather := func(city string) (float64, error) {
		if city == "ErrorCity" {
			return 0, errors.New("weather error")
		}
		return 25.0, nil
	}

	router := setupRouter(mockGetLocationFromZip, mockGetWeather)

	tests := []struct {
		name           string
		zipcode        string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "Valid zipcode",
			zipcode:        "01001000",
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"temp_C": 25.0,
				"temp_F": 77.0,
				"temp_K": 298.15,
			},
		},
		{
			name:           "Invalid zipcode length",
			zipcode:        "123",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: map[string]interface{}{
				"message": "invalid zipcode",
			},
		},
		{
			name:           "Zipcode not found",
			zipcode:        "00000000",
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"message": "can not find zipcode invalid zip",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/weather/"+test.zipcode, nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, test.expectedStatus, resp.Code)

			var actualBody map[string]interface{}
			if err := json.Unmarshal(resp.Body.Bytes(), &actualBody); err == nil {
				assert.Equal(t, test.expectedBody, actualBody)
			}
		})
	}
}
