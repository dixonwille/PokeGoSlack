package adapter

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	status int
	http.ResponseWriter
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

//Logging adds the logging adapter for the handler
func Logging(l *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			myW := &loggingResponseWriter{-1, w}
			before := time.Now().UTC()
			h.ServeHTTP(myW, r)
			totalTime := time.Now().UTC().Sub(before).Nanoseconds()
			total := float64(totalTime) / float64(1000000)
			l.Println(fmt.Sprintf("[%d]\t[%.1fms]\t%s", myW.status, total, r.URL.String()))
		})
	}
}
