package middleware

import (
	"net"
	"net/http"
)

func RateLimiterMiddleware(limiter limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extrai o IP de RemoteAddr
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				http.Error(w, "invalid IP address", http.StatusBadRequest)
				return
			}

			token := r.Header.Get("API_KEY")

			// Verifica o token primeiro
			if token != "" {
				if limiter.IsRateLimited(token, true) {
					http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
					return
				}
			} else {
				// Se não houver token, verifica o IP
				if limiter.IsRateLimited(ip, false) {
					http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
