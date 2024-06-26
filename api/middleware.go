package api

import (
	"bytes"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Set logger to use pretty print format for easier human reading
	logger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	logger.Out = os.Stdout

	// Set logrus as the output for the standard log package
	log.SetOutput(logger.Writer())
}

// RequestLogger logs all requests with their method, path, and body, and all responses
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Save the request body
		var requestBody bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &requestBody)
		body, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&requestBody)

		// Log the incoming request
		logger.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"body":   string(body),
		}).Info("Incoming request")

		// Capture the response by replacing the writer
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// Start timer
		start := time.Now()

		// Process the request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log the outgoing response
		logger.WithFields(logrus.Fields{
			"status":  c.Writer.Status(),
			"latency": latency,
			"body":    blw.body.String(),
		}).Info("Response logged")
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}
