package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rms-diego/book-my-session/internal/utils/token"
	"github.com/rms-diego/book-my-session/pkg/logger"
	"go.uber.org/zap"
)

var skipPaths = map[string]bool{
	"/health": true,
}

func LogsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		if skipPaths[path] {
			c.Next()
			return
		}

		var body map[string]any
		if raw, err := io.ReadAll(c.Request.Body); err == nil {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(raw))
			if len(raw) > 0 {
				_ = json.Unmarshal(raw, &body)
			}
		}

		c.Next()

		claims, _ := token.FromContext(c.Request.Context())
		var userInfo any
		if claims != nil {
			userInfo = map[string]any{
				"id":    claims.ID,
				"name":  claims.Name,
				"email": claims.Email,
				"role":  claims.Role,
			}
		}

		status := c.Writer.Status()
		latency := time.Since(start)
		method := c.Request.Method
		query := c.Request.URL.RawQuery

		fields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.Any("user", userInfo),
		}

		if query != "" {
			fields = append(fields, zap.String("query", query))
		}

		if status >= 400 && body != nil {
			fields = append(fields, zap.Any("body", body))
		}

		for _, err := range c.Errors {
			fields = append(fields, zap.String("error", err.Error()))
		}

		switch {
		case status >= 500:
			logger.Log.Error("request", fields...)
		case status >= 400:
			logger.Log.Warn("request", fields...)
		default:
			logger.Log.Info("request", fields...)
		}
	}
}
