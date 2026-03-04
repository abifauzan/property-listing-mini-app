package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	httpdelivery "github.com/abifauzan/property-listing-mini-app/backend/internal/delivery/http"
)

func TestCORS(t *testing.T) {
	tests := []struct {
		name           string
		envOrigins     string
		requestOrigin  string
		requestMethod  string
		wantOrigin     string
		wantStatus     int
	}{
		{
			name:          "wildcard allows any origin",
			envOrigins:    "",
			requestOrigin: "https://example.com",
			requestMethod: "GET",
			wantOrigin:    "*",
			wantStatus:    http.StatusOK,
		},
		{
			name:          "specific origin allowed",
			envOrigins:    "https://example.com,https://test.com",
			requestOrigin: "https://example.com",
			requestMethod: "GET",
			wantOrigin:    "https://example.com",
			wantStatus:    http.StatusOK,
		},
		{
			name:          "origin not in allowed list",
			envOrigins:    "https://example.com",
			requestOrigin: "https://evil.com",
			requestMethod: "GET",
			wantOrigin:    "",
			wantStatus:    http.StatusOK,
		},
		{
			name:          "OPTIONS request returns NoContent",
			envOrigins:    "",
			requestOrigin: "https://example.com",
			requestMethod: "OPTIONS",
			wantOrigin:    "*",
			wantStatus:    http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envOrigins != "" {
				os.Setenv("ALLOWED_ORIGINS", tt.envOrigins)
				defer os.Unsetenv("ALLOWED_ORIGINS")
			}

			handler := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			req := httptest.NewRequest(tt.requestMethod, "/test", nil)
			if tt.requestOrigin != "" {
				req.Header.Set("Origin", tt.requestOrigin)
			}
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}

			gotOrigin := rec.Header().Get("Access-Control-Allow-Origin")
			if gotOrigin != tt.wantOrigin {
				t.Errorf("Access-Control-Allow-Origin = %q, want %q", gotOrigin, tt.wantOrigin)
			}

			if tt.requestMethod != "OPTIONS" {
				methods := rec.Header().Get("Access-Control-Allow-Methods")
				if methods == "" {
					t.Error("Access-Control-Allow-Methods header not set")
				}
			}
		})
	}
}

func TestLogger(t *testing.T) {
	handler := Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	ctx := context.WithValue(req.Context(), httpdelivery.RequestIDKey, "test-id-123")
	req = req.WithContext(ctx)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestLoggerWithCustomStatus(t *testing.T) {
	handler := Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	req := httptest.NewRequest("POST", "/api/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestRequestID(t *testing.T) {
	handler := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(httpdelivery.RequestIDKey)
		if id == nil {
			t.Error("request ID not found in context")
			return
		}
		if _, ok := id.(string); !ok {
			t.Error("request ID is not a string")
		}
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestTimeout(t *testing.T) {
	t.Run("request completes before timeout", func(t *testing.T) {
		handler := Timeout(100 * time.Millisecond)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(10 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
		}
	})

	t.Run("request times out", func(t *testing.T) {
		handler := Timeout(50 * time.Millisecond)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusGatewayTimeout {
			t.Errorf("status = %d, want %d", rec.Code, http.StatusGatewayTimeout)
		}
	})
}

func TestGetAllowedOrigins(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		want     []string
	}{
		{
			name:     "default wildcard",
			envValue: "",
			want:     []string{"*"},
		},
		{
			name:     "single origin",
			envValue: "https://example.com",
			want:     []string{"https://example.com"},
		},
		{
			name:     "multiple origins",
			envValue: "https://example.com,https://test.com",
			want:     []string{"https://example.com", "https://test.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("ALLOWED_ORIGINS", tt.envValue)
				defer os.Unsetenv("ALLOWED_ORIGINS")
			} else {
				os.Unsetenv("ALLOWED_ORIGINS")
			}

			got := getAllowedOrigins()
			if len(got) != len(tt.want) {
				t.Fatalf("length = %d, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("origin[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestIsOriginAllowed(t *testing.T) {
	tests := []struct {
		name    string
		origin  string
		allowed []string
		want    bool
	}{
		{
			name:    "origin in list",
			origin:  "https://example.com",
			allowed: []string{"https://example.com", "https://test.com"},
			want:    true,
		},
		{
			name:    "origin not in list",
			origin:  "https://evil.com",
			allowed: []string{"https://example.com", "https://test.com"},
			want:    false,
		},
		{
			name:    "origin with whitespace",
			origin:  "https://example.com",
			allowed: []string{" https://example.com ", "https://test.com"},
			want:    true,
		},
		{
			name:    "empty origin",
			origin:  "",
			allowed: []string{"https://example.com"},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isOriginAllowed(tt.origin, tt.allowed)
			if got != tt.want {
				t.Errorf("isOriginAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRequestID(t *testing.T) {
	t.Run("request ID exists in context", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		ctx := context.WithValue(req.Context(), httpdelivery.RequestIDKey, "test-id-123")
		req = req.WithContext(ctx)

		got := getRequestID(req)
		if got != "test-id-123" {
			t.Errorf("getRequestID() = %q, want %q", got, "test-id-123")
		}
	})

	t.Run("request ID not in context", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)

		got := getRequestID(req)
		if got != "" {
			t.Errorf("getRequestID() = %q, want empty string", got)
		}
	})

	t.Run("request ID is not a string", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		ctx := context.WithValue(req.Context(), httpdelivery.RequestIDKey, 12345)
		req = req.WithContext(ctx)

		got := getRequestID(req)
		if got != "" {
			t.Errorf("getRequestID() = %q, want empty string", got)
		}
	})
}

func TestGenerateRequestID(t *testing.T) {
	id1 := generateRequestID()
	if id1 == "" {
		t.Error("generateRequestID() returned empty string")
	}

	if !strings.Contains(id1, "-") {
		t.Error("generateRequestID() should contain a hyphen")
	}

	parts := strings.Split(id1, "-")
	if len(parts) != 2 {
		t.Errorf("generateRequestID() should have 2 parts, got %d", len(parts))
	}

	if len(parts[1]) != 8 {
		t.Errorf("random part length = %d, want 8", len(parts[1]))
	}
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"length 8", 8},
		{"length 16", 16},
		{"length 1", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomString(tt.length)
			if len(got) != tt.length {
				t.Errorf("randomString(%d) length = %d, want %d", tt.length, len(got), tt.length)
			}

			for _, c := range got {
				if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')) {
					t.Errorf("randomString() contains invalid character: %c", c)
				}
			}
		})
	}
}

func TestResponseWriter(t *testing.T) {
	rec := httptest.NewRecorder()
	rw := &responseWriter{ResponseWriter: rec, statusCode: http.StatusOK}

	rw.WriteHeader(http.StatusNotFound)

	if rw.statusCode != http.StatusNotFound {
		t.Errorf("statusCode = %d, want %d", rw.statusCode, http.StatusNotFound)
	}

	if rec.Code != http.StatusNotFound {
		t.Errorf("underlying ResponseWriter code = %d, want %d", rec.Code, http.StatusNotFound)
	}
}
