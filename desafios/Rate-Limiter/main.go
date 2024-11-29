package main

import (
	"fmt"
	"log"
	"net/http"
	"rate-limiter/config"
	"rate-limiter/limiter"
	"rate-limiter/middleware"
	"rate-limiter/strategy"
)

func main() {
	// Inicializar configurações
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Criar armazenamento em memória e rate limiter
	store := strategy.NewMemoryStore()
	//store := strategy.NewRedisStore(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB) // descomente essa linha para utilizar Redis
	rateLimiter := limiter.NewRateLimiter(store, cfg.IPLimit, cfg.TokenLimit, cfg.BlockTime)

	// Configurar servidor HTTP com middleware
	mux := http.NewServeMux()
	mux.Handle("/", middleware.RateLimiterMiddleware(rateLimiter)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Request successful"))
	})))

	// Iniciar o servidor na porta configurada
	log.Printf("Server running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerPort), mux))
}
