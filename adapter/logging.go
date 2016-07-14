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
			before := time.Now().UTC().Unix()
			h.ServeHTTP(myW, r)
			totalTime := time.Now().UTC().Unix() - before
			l.Println(fmt.Sprintf("[%s]\t[%d]\t[%dms]\t%s", r.Method, myW.status, totalTime, r.Host+r.URL.String()))
		})
	}
}
