// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
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

func TestAuthSetUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`true`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Auth.Set(context.Background(), "openai", opencode.AuthSetParamsAPI{
		Type: opencode.F("api"),
		Key:  opencode.F("sk-test-key"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPut {
		t.Fatalf("expected PUT, got: %s", method)
	}
	if gotPath != "/auth/openai" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestAuthDeleteUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`true`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Auth.Remove(context.Background(), "openai")
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodDelete {
		t.Fatalf("expected DELETE, got: %s", method)
	}
	if gotPath != "/auth/openai" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestAuthDeleteSendsNoRequestBody(t *testing.T) {
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					contentType = req.Header.Get("Content-Type")
					if req.Body != nil {
						var err error
						body, err = io.ReadAll(req.Body)
						if err != nil {
							return nil, err
						}
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`true`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Auth.Remove(context.Background(), "openai")
	if err != nil {
		t.Fatal(err)
	}
	if got := string(body); got != "" {
		t.Fatalf("expected empty body, got: %q", got)
	}
	if contentType != "" {
		t.Fatalf("expected no content type, got: %q", contentType)
	}
}

func TestAuthSetParamsMarshalVariantShapes(t *testing.T) {
	tests := []struct {
		name   string
		params opencode.AuthSetParamsUnion
		want   string
	}{
		{
			name: "api",
			params: opencode.AuthSetParamsAPI{
				Type: opencode.F("api"),
				Key:  opencode.F("sk-test-key"),
			},
			want: `{"type":"api","key":"sk-test-key"}`,
		},
		{
			name: "oauth",
			params: opencode.AuthSetParamsOAuth{
				Type:    opencode.F("oauth"),
				Refresh: opencode.F("refresh-token"),
				Access:  opencode.F("access-token"),
				Expires: opencode.F(123.0),
			},
			want: `{"type":"oauth","refresh":"refresh-token","access":"access-token","expires":123}`,
		},
		{
			name: "wellknown",
			params: opencode.AuthSetParamsWellKnown{
				Type:  opencode.F("wellknown"),
				Key:   opencode.F("wk-key"),
				Token: opencode.F("wk-token"),
			},
			want: `{"type":"wellknown","key":"wk-key","token":"wk-token"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.params)
			if err != nil {
				t.Fatal(err)
			}

			var gotJSON map[string]any
			if err := json.Unmarshal(got, &gotJSON); err != nil {
				t.Fatal(err)
			}

			var wantJSON map[string]any
			if err := json.Unmarshal([]byte(tt.want), &wantJSON); err != nil {
				t.Fatal(err)
			}

			if !jsonEqual(gotJSON, wantJSON) {
				t.Fatalf("unexpected json: %s", got)
			}
		})
	}
}

func TestAuthSetParamsDeprecatedFlattenedShapeStillMarshals(t *testing.T) {
	got, err := json.Marshal(opencode.AuthSetParams{
		Type: opencode.F("api"),
		Key:  opencode.F("sk-test-key"),
	})
	if err != nil {
		t.Fatal(err)
	}

	var gotJSON map[string]any
	if err := json.Unmarshal(got, &gotJSON); err != nil {
		t.Fatal(err)
	}

	var wantJSON map[string]any
	if err := json.Unmarshal([]byte(`{"type":"api","key":"sk-test-key"}`), &wantJSON); err != nil {
		t.Fatal(err)
	}

	if !jsonEqual(gotJSON, wantJSON) {
		t.Fatalf("unexpected json: %s", got)
	}
}

func TestAuthSetAcceptsDeprecatedFlattenedParamsOnTransport(t *testing.T) {
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					contentType = req.Header.Get("Content-Type")
					if req.Body != nil {
						var err error
						body, err = io.ReadAll(req.Body)
						if err != nil {
							return nil, err
						}
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`true`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Auth.Set(context.Background(), "openai", opencode.AuthSetParams{
		Type: opencode.F("api"),
		Key:  opencode.F("sk-test-key"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if contentType != "application/json" {
		t.Fatalf("expected application/json content type, got: %q", contentType)
	}

	var gotJSON map[string]any
	if err := json.Unmarshal(body, &gotJSON); err != nil {
		t.Fatal(err)
	}

	var wantJSON map[string]any
	if err := json.Unmarshal([]byte(`{"type":"api","key":"sk-test-key"}`), &wantJSON); err != nil {
		t.Fatal(err)
	}

	if !jsonEqual(gotJSON, wantJSON) {
		t.Fatalf("unexpected json: %s", body)
	}
}

func TestAuthDeleteAcceptsDeprecatedRemoveParamsCompatibility(t *testing.T) {
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					contentType = req.Header.Get("Content-Type")
					if req.Body != nil {
						var err error
						body, err = io.ReadAll(req.Body)
						if err != nil {
							return nil, err
						}
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`true`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Auth.Remove(context.Background(), "openai", opencode.AuthRemoveParams{})
	if err != nil {
		t.Fatal(err)
	}
	if got := string(body); got != "" {
		t.Fatalf("expected empty body, got: %q", got)
	}
	if contentType != "" {
		t.Fatalf("expected no content type, got: %q", contentType)
	}
}

func TestAuthUnmarshalOAuthVariant(t *testing.T) {
	var auth opencode.Auth
	err := json.Unmarshal([]byte(`{"type":"oauth","refresh":"refresh-token","access":"access-token","expires":123,"accountId":"acct_123"}`), &auth)
	if err != nil {
		t.Fatal(err)
	}

	oauth, ok := auth.AsUnion().(opencode.AuthOAuth)
	if !ok {
		t.Fatalf("expected oauth auth variant, got %T", auth.AsUnion())
	}
	if oauth.Refresh != "refresh-token" {
		t.Fatalf("expected refresh token to round-trip, got %q", oauth.Refresh)
	}
	if oauth.Access != "access-token" {
		t.Fatalf("expected access token to round-trip, got %q", oauth.Access)
	}
	if oauth.Expires != 123 {
		t.Fatalf("expected expires to round-trip, got %v", oauth.Expires)
	}
}

func TestAuthUnmarshalAPIVariant(t *testing.T) {
	var auth opencode.Auth
	err := json.Unmarshal([]byte(`{"type":"api","key":"sk-test-key","metadata":{"env":"test"}}`), &auth)
	if err != nil {
		t.Fatal(err)
	}

	api, ok := auth.AsUnion().(opencode.AuthAPI)
	if !ok {
		t.Fatalf("expected api auth variant, got %T", auth.AsUnion())
	}
	if api.Key != "sk-test-key" {
		t.Fatalf("expected key to round-trip, got %q", api.Key)
	}
	if api.Metadata["env"] != "test" {
		t.Fatalf("expected metadata to round-trip, got %#v", api.Metadata)
	}
}

func TestAuthUnmarshalWellKnownVariant(t *testing.T) {
	var auth opencode.Auth
	err := json.Unmarshal([]byte(`{"type":"wellknown","key":"wk-key","token":"wk-token"}`), &auth)
	if err != nil {
		t.Fatal(err)
	}

	wellKnown, ok := auth.AsUnion().(opencode.AuthWellKnown)
	if !ok {
		t.Fatalf("expected wellknown auth variant, got %T", auth.AsUnion())
	}
	if wellKnown.Key != "wk-key" {
		t.Fatalf("expected key to round-trip, got %q", wellKnown.Key)
	}
	if wellKnown.Token != "wk-token" {
		t.Fatalf("expected token to round-trip, got %q", wellKnown.Token)
	}
}

func jsonEqual(got, want map[string]any) bool {
	gotJSON, err := json.Marshal(got)
	if err != nil {
		return false
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		return false
	}
	return string(gotJSON) == string(wantJSON)
}

func TestAuthSet(t *testing.T) {
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
	_, err := client.Auth.Set(context.TODO(), "openai", opencode.AuthSetParamsAPI{
		Type: opencode.F("api"),
		Key:  opencode.F("sk-test-key"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRemoveWithOptionalParams(t *testing.T) {
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
	_, err := client.Auth.Remove(context.TODO(), "openai", opencode.AuthRemoveParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
