package _GinZap

import (
	"fmt"
	"project/modules/_Log"
	"time"

	"github.com/gin-gonic/gin"
)

// Ginzap returns a gin.HandlerFunc (middleware) that logs requests using uber-go/zap.
//
// Requests with errors are logged using zap.Error().
// Requests without errors are logged using zap.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.

func Ginzap() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 開始時間
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		//結束時間
		end := time.Now()
		latency := end.Sub(start)

		// if utc {
		// 	end = end.UTC()
		// }

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				_Log.MainLogger.Error(e)
			}
		} else {
			_Log.WebLogger.Info(fmt.Sprintf("| %3d | %13v | %15s | %s  %s | %s", statusCode, latency, clientIP, method, path, query)) // zap.Int("status", statusCode),
			// zap.String("method", method),
			// zap.String("path", path),
			// zap.Duration("latency", latency),
			// zap.String("ip", clientIP),
			// zap.String("user-agent", c.Request.UserAgent()),
			// zap.String("query", query),

		}
	}
}
