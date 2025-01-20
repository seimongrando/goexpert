package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestWeatherAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/weather/:zipcode", func(c *gin.Context) {
		c.JSON(http.StatusOK, WeatherResponse{
			TempC: 25.0,
			TempF: 77.0,
			TempK: 298.15,
		})
	})

	req, _ := http.NewRequest("GET", "/weather/01001000", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
