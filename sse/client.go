package sse

import "github.com/gin-gonic/gin"

// SSEClient SSE客户端连接
type SSEClient struct {
	ClientID string
	Writer   gin.ResponseWriter
	Notify   <-chan bool
}
