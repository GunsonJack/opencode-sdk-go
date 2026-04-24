package opencode_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
	"github.com/GunsonJack/opencode-sdk-go/option"
)

func TestExperimentalSessionListUsesTypedClientPathAndFullQuerySurface(t *testing.T) {
	var gotPath string
	var gotQuery map[string]string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					gotPath = req.URL.EscapedPath()
					gotQuery = map[string]string{}
					for key, values := range req.URL.Query() {
						if len(values) > 0 {
							gotQuery[key] = values[0]
						}
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`[]`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.ExperimentalSession.List(context.Background(), opencode.ExperimentalSessionListParams{
		Directory: opencode.F("/tmp/project"),
		Workspace: opencode.F("ws_123"),
		Roots:     opencode.F(true),
		Start:     opencode.F(123.5),
		Cursor:    opencode.F(456.5),
		Search:    opencode.F("agent"),
		Limit:     opencode.F(float64(25)),
		Archived:  opencode.F(true),
	})
	if err != nil {
		t.Fatal(err)
	}

	if gotPath != "/experimental/session" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
	for key, want := range map[string]string{
		"directory": "/tmp/project",
		"workspace": "ws_123",
		"roots":     "true",
		"start":     "123.5",
		"cursor":    "456.5",
		"search":    "agent",
		"limit":     "25",
		"archived":  "true",
	} {
		if gotQuery[key] != want {
			t.Fatalf("unexpected %s query: got %q want %q", key, gotQuery[key], want)
		}
	}
}
