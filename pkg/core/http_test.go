package core

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/zhihao0924/amapSdk/pkg/common"
	"github.com/zhihao0924/amapSdk/pkg/models"
)

type stubHTTPClient struct {
	do func(req *http.Request) (*http.Response, error)
}

func (s *stubHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return s.do(req)
}

func TestHTTPClientGetReturnsAPIErrorForBusinessFailure(t *testing.T) {
	t.Parallel()

	client := NewHTTPClient(
		&stubHTTPClient{
			do: func(req *http.Request) (*http.Response, error) {
				return jsonResponse(http.StatusOK, `{"status":"0","info":"key不正确或过期","infocode":"10001"}`), nil
			},
		},
		"https://restapi.amap.com/v3",
		"test-key-123456",
		common.NewLogger(false),
		&RetryConfig{MaxRetries: 0, RetryDelay: 0},
	)

	var resp models.GeocodeResponse
	err := client.Get(context.Background(), "/geocode/geo", nil, &resp)
	if err == nil {
		t.Fatal("expected api error, got nil")
	}
	if !common.IsAPIError(err) {
		t.Fatalf("expected api error type, got %v", err)
	}
	if !common.IsAuthError(err) {
		t.Fatalf("expected auth error classification, got %v", err)
	}
}

func TestHTTPClientPostRetriesServerErrorsWithFreshBody(t *testing.T) {
	t.Parallel()

	var bodies []string
	attempts := 0

	client := NewHTTPClient(
		&stubHTTPClient{
			do: func(req *http.Request) (*http.Response, error) {
				attempts++

				body, err := io.ReadAll(req.Body)
				if err != nil {
					t.Fatalf("read body: %v", err)
				}
				bodies = append(bodies, string(body))

				if attempts == 1 {
					return jsonResponse(http.StatusInternalServerError, `{"status":"0","info":"server error","infocode":"50000"}`), nil
				}

				return jsonResponse(http.StatusOK, `{"status":"1","info":"OK","infocode":"10000","geocodes":[]}`), nil
			},
		},
		"https://restapi.amap.com/v3",
		"test-key-123456",
		common.NewLogger(false),
		&RetryConfig{MaxRetries: 1, RetryDelay: 0},
	)

	var resp models.GeocodeResponse
	err := client.Post(context.Background(), "/geocode/geo", nil, map[string]string{"foo": "bar"}, &resp)
	if err != nil {
		t.Fatalf("expected retry to succeed, got %v", err)
	}
	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
	if len(bodies) != 2 {
		t.Fatalf("expected 2 captured bodies, got %d", len(bodies))
	}
	for i, body := range bodies {
		if body != `{"foo":"bar"}` {
			t.Fatalf("attempt %d body mismatch: %s", i+1, body)
		}
	}
}

func TestLoggingRequestInterceptorPreservesBody(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodPost, "https://example.com", bytes.NewReader([]byte(strings.Repeat("x", 2048))))
	if err != nil {
		t.Fatalf("new request: %v", err)
	}

	interceptor := LoggingRequestInterceptor(common.NewLogger(false))
	if err := interceptor(req); err != nil {
		t.Fatalf("interceptor returned error: %v", err)
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("read preserved body: %v", err)
	}
	if len(body) != 2048 {
		t.Fatalf("expected full body to remain readable, got %d bytes", len(body))
	}
}

func jsonResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
}
