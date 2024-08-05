package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		log.Printf("Latency:%v| path:%s | method:%s | status:%d",
			latency,
			c.Request.URL.Path,
			c.Request.Method,
			c.Writer.Status(),
		)
	}
}
