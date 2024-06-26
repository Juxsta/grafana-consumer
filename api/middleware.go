package api

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger logs all requests with their method, path, and body, and all responses
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Save the request body
		var requestBody bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &requestBody)
		body, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&requestBody)

		// Log the incoming request
		log.Printf("Incoming request - Method: %s, Path: %s, Body: %s", c.Request.Method, c.Request.URL.Path, string(body))

		// Capture the response
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// Start timer
		start := time.Now()

		// Process the request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log the outgoing response
		log.Printf("Response - Status: %d, Latency: %s, Body: %s", c.Writer.Status(), latency, blw.body.String())

		// Check for errors and log them
		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				log.Printf("Error: %s", e)
			}
		}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	return w.body.WriteString(s)
}
