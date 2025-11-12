package middleware

import (
	"sync"

	"bosh-admin/core/ctx"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// IPRateLimiter 为每个IP地址维护一个限流器
type IPRateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
	r        rate.Limit // 每秒生成的令牌数
	b        int        // 桶容量
}

// NewIPRateLimiter 创建一个新的IP限流器
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		limiters: make(map[string]*rate.Limiter),
		r:        r,
		b:        b,
	}
}

// AddIP 添加一个IP
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.limiters[ip] = limiter
	return limiter
}

// GetLimiter 获取或创建指定IP的限流器
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.limiters[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}

// RateLimiter 限流中间件
func RateLimiter(rps int, burst int) gin.HandlerFunc {
	rateLimiter := NewIPRateLimiter(rate.Limit(rps), burst)
	return ctx.Handler(func(c *ctx.Context) {
		limiter := rateLimiter.GetLimiter(c.ClientIP())
		if !limiter.Allow() {
			c.TooManyRequests()
			c.Abort()
			return
		}
		c.Next()
	})
}
