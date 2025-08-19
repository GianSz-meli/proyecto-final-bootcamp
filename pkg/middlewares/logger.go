package middlewares

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"time"
)

// LoggerEntry stores details about an HTTP request and response.
type LoggerEntry struct {
	Method        string
	RequestPath   string
	RemoteAddress string
	DurationMs    time.Duration
	Response      string
	StatusCode    int
}

// ResponseWriterCapture wraps the ResponseWriter to capture its output.
type ResponseWriterCapture struct {
	http.ResponseWriter
	Body       *bytes.Buffer
	StatusCode int
}

// Write captures the body of the response.
func (rw *ResponseWriterCapture) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
}

// WriteHeader captures the status code of the response.
func (rw *ResponseWriterCapture) WriteHeader(statusCode int) {
	rw.StatusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Logger creates middleware for logging HTTP requests and responses.
func Logger(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &ResponseWriterCapture{ResponseWriter: w, Body: &bytes.Buffer{}}
			next.ServeHTTP(rw, r)

			duration := time.Since(start)
			logEntry := LoggerEntry{
				Method:        r.Method,
				RequestPath:   r.RequestURI,
				RemoteAddress: r.RemoteAddr,
				DurationMs:    duration,
				Response:      rw.Body.String(),
				StatusCode:    rw.StatusCode,
			}
			_, err := db.Exec(SQL_INSERT_LOGGER, logEntry.Method, logEntry.RequestPath, logEntry.RemoteAddress, logEntry.DurationMs, logEntry.Response, logEntry.StatusCode)
			if err != nil {
				log.Printf("Error inserting log entry: %v", err)
				return
			}
			log.Printf(
				"[%s] [%v] [%s] [%s] [%v]",
				logEntry.Method,
				logEntry.StatusCode,
				logEntry.RequestPath,
				logEntry.RemoteAddress,
				logEntry.DurationMs,
			)
		})
	}
}
