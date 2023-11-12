package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	l "github.com/prayogatriady/ecommerce-order/utils/logger"
	"github.com/sirupsen/logrus"
)

type ResponseWriterInterceptor struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w *ResponseWriterInterceptor) Write(b []byte) (int, error) {
	// Capture the response body
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware(c *gin.Context) {

	var errorMessage strings.Builder

	start := time.Now()

	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errorMessage.WriteString("Error reading request body")
		l.Slog.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn(errorMessage.String())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errorMessage.String(),
		})
		return
	}
	defer c.Request.Body.Close()

	var requestData interface{}
	json.Unmarshal([]byte(requestBody), &requestData)

	logMiddleware := l.Slog.WithFields(logrus.Fields{
		"http_method": c.Request.Method,
		"path":        c.Request.URL.Path,
		"request":     requestData,
		"query":       c.Request.URL.Query(),
	})

	// Create an interceptor that captures the response
	interceptor := &ResponseWriterInterceptor{
		ResponseWriter: c.Writer,
		Body:           &bytes.Buffer{},
	}
	c.Writer = interceptor // Replace the default writer

	// Run the handlers
	c.Next()

	responseBody := interceptor.Body.String()

	var responseData interface{}
	json.Unmarshal([]byte(responseBody), &responseData)

	duration := time.Since(start)

	logMiddleware.WithFields(logrus.Fields{
		"duration": duration.String(),
		"status":   c.Writer.Status(),
		"response": responseData,
	}).Debug("Log Middleware")
}
