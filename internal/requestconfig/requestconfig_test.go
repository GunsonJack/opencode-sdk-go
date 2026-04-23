package requestconfig

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type closureTransport struct {
	fn func(*http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestExecuteNewRequestEscapesPathSegments(t *testing.T) {
	var path string
	err := ExecuteNewRequest(
		context.Background(),
		http.MethodGet,
		"auth/provider/with space",
		nil,
		new(bool),
		WithDefaultBaseURL("https://example.com/"),
		RequestOptionFunc(func(cfg *RequestConfig) error {
			cfg.HTTPClient = &http.Client{Transport: &closureTransport{fn: func(req *http.Request) (*http.Response, error) {
				path = req.URL.EscapedPath()
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("true")),
					Header:     http.Header{"Content-Type": []string{"application/json"}},
				}, nil
			}}}
			return nil
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	if path != "/auth/provider%2Fwith%20space" {
		t.Fatalf("unexpected path: %s", path)
	}
}
