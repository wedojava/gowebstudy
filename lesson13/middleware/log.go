package middleware

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/wedojava/gowebstudy/lesson13/utils/vlog"
)

func Logging() mux.MiddlewareFunc {
	// Create middleware
	return func(next http.Handler) http.Handler {
		// Create a new handler encapsulate http.HandlerFunc
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()
				defer func() { log.Println(r.URL.Path, time.Since(start)) }()
				next.ServeHTTP(w, r)
			})
	}

}

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rwr *ResponseWithRecorder) WriteHeader(statusCode int) {
	rwr.ResponseWriter.WriteHeader(statusCode)
	rwr.statusCode = statusCode
}

func (rwr *ResponseWithRecorder) Write(b []byte) (n int, err error) {
	n, err = rwr.ResponseWriter.Write(b)
	if err != nil {
		return
	}
	rwr.body.Write(b)
	return
}

func AccessLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			logEntry := vlog.AccessLog.WithFields(logrus.Fields{
				"ip":           r.RemoteAddr,
				"method":       r.Method,
				"path":         r.RequestURI,
				"query":        r.URL.RawQuery,
				"request_body": buf.String(),
			})
			wc := &ResponseWithRecorder{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
				body:           bytes.Buffer{},
			}
			next.ServeHTTP(wc, r)

			defer logEntry.WithFields(logrus.Fields{
				"status":        wc.statusCode,
				"response_body": wc.body.String(),
			}).Info()
		})
}
