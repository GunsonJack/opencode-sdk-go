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

func TestWorkspaceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Experimental.Workspace.List(context.TODO(), opencode.WorkspaceListParams{
		Directory: opencode.F("directory"),
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

func TestWorkspaceCreate(t *testing.T) {
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
	_, err := client.Experimental.Workspace.Create(context.TODO(), opencode.WorkspaceCreateParams{
		Type:      opencode.F("type"),
		Branch:    opencode.F("branch"),
		Extra:     opencode.F[interface{}](map[string]interface{}{}),
		ID:        opencode.F("id"),
		Directory: opencode.F("directory"),
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

func TestWorkspaceCreateUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"id":"wrk_1","type":"git","name":"name","branch":"main","directory":"/tmp/project","extra":{},"projectID":"proj_1"}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Experimental.Workspace.Create(context.Background(), opencode.WorkspaceCreateParams{
		Type:   opencode.F("git"),
		Branch: opencode.F("main"),
		Extra:  opencode.F[interface{}](map[string]interface{}{}),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if gotPath != "/experimental/workspace" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestWorkspaceListUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`[]`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Experimental.Workspace.List(context.Background(), opencode.WorkspaceListParams{})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodGet {
		t.Fatalf("expected GET, got: %s", method)
	}
	if gotPath != "/experimental/workspace" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestWorkspaceStatusWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Experimental.Workspace.Status(context.TODO(), opencode.WorkspaceStatusParams{
		Directory: opencode.F("directory"),
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

func TestWorkspaceStatusUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`[]`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Experimental.Workspace.Status(context.Background(), opencode.WorkspaceStatusParams{})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodGet {
		t.Fatalf("expected GET, got: %s", method)
	}
	if gotPath != "/experimental/workspace/status" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestWorkspaceAdaptorsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Experimental.Workspace.Adaptors(context.TODO(), opencode.WorkspaceAdaptorsParams{
		Directory: opencode.F("directory"),
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

func TestWorkspaceAdaptorsUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`[]`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Experimental.Workspace.Adaptors(context.Background(), opencode.WorkspaceAdaptorsParams{})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodGet {
		t.Fatalf("expected GET, got: %s", method)
	}
	if gotPath != "/experimental/workspace/adaptor" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestWorkspaceRemoveWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Experimental.Workspace.Remove(context.TODO(), "wrk_test123", opencode.WorkspaceRemoveParams{
		Directory: opencode.F("directory"),
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

func TestWorkspaceRemoveUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"id":"wrk_1","type":"git","name":"name","branch":"main","directory":"/tmp/project","extra":{},"projectID":"proj_1"}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Experimental.Workspace.Remove(context.Background(), "wrk_test123", opencode.WorkspaceRemoveParams{})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodDelete {
		t.Fatalf("expected DELETE, got: %s", method)
	}
	if gotPath != "/experimental/workspace/wrk_test123" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestWorkspaceSessionRestoreWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Experimental.Workspace.SessionRestore(context.TODO(), "wrk_test123", opencode.WorkspaceSessionRestoreParams{
		SessionID: opencode.F("ses_abc123"),
		Directory: opencode.F("directory"),
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

func TestWorkspaceSessionRestoreUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"total":1}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Experimental.Workspace.SessionRestore(context.Background(), "wrk_test123", opencode.WorkspaceSessionRestoreParams{
		SessionID: opencode.F("ses_abc123"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if gotPath != "/experimental/workspace/wrk_test123/session-restore" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}
