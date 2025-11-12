package ctx

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

type HandlerFunc func(c *Context)

func Handler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(&Context{c})
	}
}
