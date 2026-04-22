package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogsMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		placeholder := make(map[string]any)
		var body any = nil

		rawBody, err := io.ReadAll(c.Request.Body)
		if err == nil {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))
		}

		if err := c.ShouldBindJSON(&placeholder); len(rawBody) > 0 && err == nil {
			body = placeholder
		}

		c.Next()

		logger.Info("request",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Any("body", body),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
		)

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.Error("request error", zap.String("error", err.Error()))
			}
		}
	}
}
