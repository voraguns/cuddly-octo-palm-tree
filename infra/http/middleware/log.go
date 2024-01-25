package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

// LogMiddleware is the logging middleware
func LogMiddleware(logger *log.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := GetDurationInMillseconds(start)

		entry := logger.WithFields(log.Fields{
			"type":        "router",
			"duration_ms": duration,
			"method":      c.Request.Method,
			"path":        c.Request.RequestURI,
			"status":      c.Writer.Status(),
			"referrer":    c.Request.Referer(),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

// GetDurationInMillseconds takes a start time and returns a duration in milliseconds
func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
