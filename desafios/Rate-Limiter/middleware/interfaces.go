package middleware

// limiter define a interface para o componente de Rate Limiter
type limiter interface {
	IsRateLimited(key string, isToken bool) bool
}
