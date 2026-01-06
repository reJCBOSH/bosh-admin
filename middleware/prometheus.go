package middleware

import (
	"strconv"
	"time"

	"bosh-admin/core/ctx"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// 定义指标
var (
	// 请求计数器：按方法、路径、状态码统计
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// 请求延迟直方图：统计请求耗时分布
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency distribution",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1, 2, 5}, // 延迟区间（秒）
		},
		[]string{"method", "path"},
	)
)

func init() {
	// 注册指标到默认的注册表
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

func Prometheus() gin.HandlerFunc {
	return ctx.Handler(func(c *ctx.Context) {
		start := time.Now()
		// 处理请求
		c.Next()
		// 计算耗时
		latency := time.Since(start).Seconds()
		path := c.FullPath() // 获取路由路径（如 /api/user/:id）
		if path == "" {
			path = c.Request.URL.Path // 回退到原始路径
		}
		// 标签值
		method := c.Request.Method
		status := strconv.Itoa(c.Writer.Status())
		// 收集指标
		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
		httpRequestDuration.WithLabelValues(method, path).Observe(latency)
	})
}
