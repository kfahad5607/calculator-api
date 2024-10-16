package middlewares

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
	"golang.org/x/time/rate"

	"github.com/kfahad5607/calculator-api/api/handlers"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

type Middleware func(http.Handler) http.Handler

func (w *wrappedWriter)WriteHeader(statusCode int){
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

var limiter = rate.NewLimiter(3, 5)

func CreateMiddlewareStack(middlewares ...Middleware) Middleware{
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			m := middlewares[i]
			next = m(next)
		}

		return next
	}
}

func RateLimiter(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		}else{
			w.Header().Set("Content-Type", "apllication/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(handlers.ResponseError{Error: "Rate limit exceeded, please retry after some time."})
		}
	})
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