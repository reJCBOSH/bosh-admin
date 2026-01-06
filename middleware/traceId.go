package middleware

import (
	"context"

	"bosh-admin/core/ctx"

	"github.com/duke-git/lancet/v2/random"
	"github.com/gin-gonic/gin"
)

func TraceId() gin.HandlerFunc {
	return ctx.Handler(func(c *ctx.Context) {
		traceId := c.GetHeader("X-Trace-Id")
		if traceId == "" {
			traceId, _ = random.UUIdV4()
			c.Header("X-Trace-Id", traceId)
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "traceId", traceId))
	})
}
