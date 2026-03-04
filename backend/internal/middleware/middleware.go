package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	httpdelivery "github.com/abifauzan/property-listing-mini-app/backend/internal/delivery/http"
)

// CORS adds CORS headers. Allowed origins can be configured via ALLOWED_ORIGINS env var.
// Defaults to "*" for development. Use comma-separated list for multiple origins.
func CORS(next http.Handler) http.Handler {
	allowedOrigins := getAllowedOrigins()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if isOriginAllowed(origin, allowedOrigins) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else if len(allowedOrigins) == 1 && allowedOrigins[0] == "*" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Logger logs each request's method, path, status, and duration using slog.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		requestID := getRequestID(r)

		next.ServeHTTP(rw, r)

		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.statusCode,
			"duration", time.Since(start).String(),
			"requestId", requestID,
		)
	})
}

// responseWriter wraps http.ResponseWriter to capture the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// RequestID adds a unique request ID to each request context.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := generateRequestID()
		ctx := context.WithValue(r.Context(), httpdelivery.RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Timeout adds a timeout to each request.
func Timeout(duration time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), duration)
			defer cancel()

			done := make(chan struct{})
			go func() {
				next.ServeHTTP(w, r.WithContext(ctx))
				close(done)
			}()

			select {
			case <-done:
				return
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					slog.Warn("request timeout", "path", r.URL.Path)
					w.WriteHeader(http.StatusGatewayTimeout)
				}
			}
		})
	}
}

// getAllowedOrigins returns the list of allowed origins from env var.
func getAllowedOrigins() []string {
	origins := os.Getenv("ALLOWED_ORIGINS")
	if origins == "" {
		return []string{"*"}
	}
	return strings.Split(origins, ",")
}

// isOriginAllowed checks if the given origin is in the allowed list.
func isOriginAllowed(origin string, allowed []string) bool {
	for _, o := range allowed {
		if strings.TrimSpace(o) == origin {
			return true
		}
	}
	return false
}

// getRequestID extracts request ID from context.
func getRequestID(r *http.Request) string {
	if id := r.Context().Value(httpdelivery.RequestIDKey); id != nil {
		if idStr, ok := id.(string); ok {
			return idStr
		}
	}
	return ""
}

// generateRequestID generates a simple request ID.
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString generates a random alphanumeric string.
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
	}
	return string(b)
}
