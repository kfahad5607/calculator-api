package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter)WriteHeader(statusCode int){
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		wrappedWriter := wrappedWriter{ResponseWriter: w, statusCode:http.StatusOK}

		next.ServeHTTP(&wrappedWriter, r)
		logger.Info(fmt.Sprintf("%v %v %v %v", r.Method, r.URL.Path, wrappedWriter.statusCode, time.Since(startTime)))
	})
}