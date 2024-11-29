package stress

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunLoadTest_Success(t *testing.T) {
	// Cria um servidor mock que sempre retorna status 200
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Executa o teste de carga
	report := RunLoadTest(server.URL, 10, 2)

	// Verifica os resultados
	if report.TotalRequests != 10 {
		t.Errorf("expected 10 total requests, got %d", report.TotalRequests)
	}
	if report.Status200 != 10 {
		t.Errorf("expected 10 status 200 responses, got %d", report.Status200)
	}
	if len(report.StatusOthers) != 0 {
		t.Errorf("expected 0 other statuses, got %d", len(report.StatusOthers))
	}
}

func TestRunLoadTest_WithErrors(t *testing.T) {
	// Cria um servidor mock que alterna entre status 200 e 500
	count := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if count%2 == 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		count++
	}))
	defer server.Close()

	// Executa o teste de carga
	report := RunLoadTest(server.URL, 10, 2)

	// Verifica os resultados
	if report.TotalRequests != 10 {
		t.Errorf("expected 10 total requests, got %d", report.TotalRequests)
	}
	if report.Status200 != 5 {
		t.Errorf("expected 5 status 200 responses, got %d", report.Status200)
	}
	if report.StatusOthers[500] != 5 {
		t.Errorf("expected 5 status 500 responses, got %d", report.StatusOthers[500])
	}
}

func TestRunLoadTest_InvalidURL(t *testing.T) {
	// Executa o teste de carga com uma URL inválida
	report := RunLoadTest("http://invalid.url", 5, 2)

	// Verifica os resultados
	if report.TotalRequests != 5 {
		t.Errorf("expected 5 total requests, got %d", report.TotalRequests)
	}
	if report.Status200 != 0 {
		t.Errorf("expected 0 status 200 responses, got %d", report.Status200)
	}
	if report.StatusOthers[0] != 5 {
		t.Errorf("expected 5 error responses (status 0), got %d", report.StatusOthers[0])
	}
}
