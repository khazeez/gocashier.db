package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type APILog struct {
	Timestamp    string `json:"timestamp"`
	Method       string `json:"method"`
	Endpoint     string `json:"endpoint"`
	StatusCode   int    `json:"status_code"`
	LatencyMs    int64  `json:"latency_ms"`
	ClientIP     string `json:"client_ip"`
	RequestBody  string `json:"request_body,omitempty"`
	ResponseBody string `json:"response_body,omitempty"`
}

// responseWriter wrapper untuk capture response body
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	esURL := os.Getenv("ELASTICSEARCH_URL")
	if esURL == "" {
		esURL = "http://localhost:9200"
	}

	return func(c *gin.Context) {
		start := time.Now()

		// Capture request body
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				// Kembalikan body agar bisa dibaca handler
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// Wrap response writer untuk capture response body
		rw := &responseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer(nil),
		}
		c.Writer = rw

		// Proses request
		c.Next()

		latency := time.Since(start).Milliseconds()

		log := APILog{
			Timestamp:    time.Now().UTC().Format("2006-01-02T15:04:05Z"),
			Method:       c.Request.Method,
			Endpoint:     c.FullPath(),
			StatusCode:   c.Writer.Status(),
			LatencyMs:    latency,
			ClientIP:     c.ClientIP(),
			RequestBody:  requestBody,
			ResponseBody: rw.body.String(),
		}

		// Kirim ke Elasticsearch secara async
		go sendToElasticsearch(esURL, log)
	}
}

func sendToElasticsearch(esURL string, log APILog) {
	data, err := json.Marshal(log)
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s/gocashier-logs/_doc", esURL)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}