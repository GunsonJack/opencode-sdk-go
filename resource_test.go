// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
	"github.com/GunsonJack/opencode-sdk-go/internal/testutil"
	"github.com/GunsonJack/opencode-sdk-go/option"
)

func TestResourceList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Resource.List(context.TODO(), opencode.ResourceListParams{
		Directory: opencode.F("/tmp/test"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourceListUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Resource.List(context.Background(), opencode.ResourceListParams{
		Directory: opencode.F("/tmp/test"),
		Workspace: opencode.F("ws_123"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodGet {
		t.Fatalf("expected GET, got: %s", method)
	}
	if gotPath != "/experimental/resource" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}
