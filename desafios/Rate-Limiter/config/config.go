package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string
	IPLimit       int
	TokenLimit    int
	BlockTime     time.Duration
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

// NewConfig carrega as variáveis de configuração e as expõe por meio da struct Config.
func NewConfig() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Tenta carregar o arquivo .env, mas não falha se ele não for encontrado
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Garante que todas as variáveis necessárias estão definidas
	cfg := &Config{
		ServerPort:    getStringOrDefault("SERVER_PORT", "8080"),
		IPLimit:       getIntOrDefault("RATE_LIMITER_IP_LIMIT", 10),
		TokenLimit:    getIntOrDefault("RATE_LIMITER_TOKEN_LIMIT", 100),
		BlockTime:     getDurationOrDefault("RATE_LIMITER_BLOCK_TIME", 300) * time.Second,
		RedisAddr:     getStringOrDefault("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getStringOrDefault("REDIS_PASSWORD", ""),
		RedisDB:       getIntOrDefault("REDIS_DB", 0),
	}

	log.Println("Configuration loaded successfully")
	return cfg, nil
}

func getStringOrDefault(key, defaultValue string) string {
	if value := viper.GetString(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntOrDefault(key string, defaultValue int) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}
	return defaultValue
}

func getDurationOrDefault(key string, defaultValue int) time.Duration {
	if viper.IsSet(key) {
		return viper.GetDuration(key)
	}
	return time.Duration(defaultValue)
}
