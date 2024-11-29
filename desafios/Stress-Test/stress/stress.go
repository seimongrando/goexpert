package stress

import (
	"net/http"
	"sync"
	"time"
)

// RunLoadTest executa o teste de carga baseado nos parâmetros fornecidos.
func RunLoadTest(url string, totalRequests, concurrency int) Report {
	var wg sync.WaitGroup
	var mu sync.Mutex

	startTime := time.Now()

	// Armazenar resultados
	report := Report{
		StatusOthers: make(map[int]int),
	}

	// Canal para gerenciar as requests
	requestsChan := make(chan int, totalRequests)

	// Preencher o canal com o número de requests
	for i := 0; i < totalRequests; i++ {
		requestsChan <- i
	}
	close(requestsChan)

	// Gerenciar concorrência
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for range requestsChan {
				status := sendRequest(url)
				mu.Lock()
				report.TotalRequests++
				if status == 200 {
					report.Status200++
				} else {
					report.StatusOthers[status]++
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	report.TotalTime = time.Since(startTime)
	return report
}

// sendRequest realiza uma única requisição HTTP e retorna o código de status.
func sendRequest(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		// Retorna um código fictício para erros de rede
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
