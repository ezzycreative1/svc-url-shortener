package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ezzycreative1/svc-url-shortener/app/middleware"
)

// Middleware to set content type
func restMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set header
		w.Header().Set("Content-Type", "application/json")

		// Continue flow
		next.ServeHTTP(w, r)
	})
}

func errHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		err := f(w, r)
		var errMessage string

		if err != nil {
			errMessage = err.Error()
			switch err.Error() {
			case model.ErrNotFound.Error():
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusBadRequest)
			}
		}

		if err := json.NewEncoder(log.Writer()).Encode(logRequest{
			Method:   r.Method,
			Path:     r.URL.Path,
			Duration: time.Since(start).Microseconds(),
			Err:      errMessage,
		}); err != nil {
			log.Fatal(err)
		}
	}
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" || r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/create")

	return *middleware.LoadConfig(mux)
}
