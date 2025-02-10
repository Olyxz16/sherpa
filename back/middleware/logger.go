package middleware

import (
	"net/http"
	"time"
	"bytes"

	"github.com/Olyxz16/sherpa/config"
)

var (
	logger = config.DefaultLogger()
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
	buffer      *bytes.Buffer
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		buffer:         bytes.NewBuffer([]byte{}),
	}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}
	rw.buffer.Write(b)
	return rw.ResponseWriter.Write(b)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer logger.Sync()
		sugar := logger.Sugar()

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		
		next.ServeHTTP(wrapped, r)
		
		delta := time.Now().Sub(start)
		sugar.Infof("%s \"%s %s %s\" from %s - %d %dB in %s",
			time.Now().Format("2006/01/02 03:04:05"),
			r.Method,
			r.URL.String(),
			r.Proto,
			r.RemoteAddr,
			wrapped.Status(),
			wrapped.buffer.Len(),
			delta.String(),
			)
	})
}
