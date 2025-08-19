package middlewares

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"time"
)

type LoggerEntry struct {
	Method        string
	RequestPath   string
	RemoteAddress string
	DurationMs    time.Duration
	Response      string
	StatusCode    int
}

type ResponseCapture struct {
	http.ResponseWriter
	Body       *bytes.Buffer
	StatusCode int
}

func (rc *ResponseCapture) Write(b []byte) (int, error) {
	rc.Body.Write(b)
	return rc.ResponseWriter.Write(b)
}

func (rc *ResponseCapture) WriteHeader(statusCode int) {
	rc.StatusCode = statusCode
	rc.ResponseWriter.WriteHeader(statusCode)
}

func Logger(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rc := &ResponseCapture{ResponseWriter: w, Body: &bytes.Buffer{}}
			next.ServeHTTP(rc, r)

			duration := time.Since(start)
			logEntry := LoggerEntry{
				Method:        r.Method,
				RequestPath:   r.RequestURI,
				RemoteAddress: r.RemoteAddr,
				DurationMs:    duration,
				Response:      rc.Body.String(),
				StatusCode:    rc.StatusCode,
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
				duration,
			)
		})
	}
}
