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

func TestWorktreeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Experimental.Worktree.List(context.TODO(), opencode.WorktreeListParams{
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

func TestWorktreeCreateWithOptionalParams(t *testing.T) {
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
	_, err := client.Experimental.Worktree.Create(context.TODO(), opencode.WorktreeCreateParams{
		Name:         opencode.F("name"),
		StartCommand: opencode.F("startCommand"),
		Directory:    opencode.F("directory"),
		Workspace:    opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorktreeCreateUsesCorrectMethodAndPath(t *testing.T) {
	var method, gotPath string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					gotPath = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"name":"wt","branch":"main","directory":"/tmp/wt"}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)
	_, err := client.Experimental.Worktree.Create(context.Background(), opencode.WorktreeCreateParams{
		Name:      opencode.F("wt"),
		Directory: opencode.F("/tmp/project"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if gotPath != "/experimental/worktree" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestWorktreeListUsesCorrectMethodAndPath(t *testing.T) {
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
	_, err := client.Experimental.Worktree.List(context.Background(), opencode.WorktreeListParams{})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodGet {
		t.Fatalf("expected GET, got: %s", method)
	}
	if gotPath != "/experimental/worktree" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestWorktreeRemoveSendsDirectoryInQueryAndBody(t *testing.T) {
	var method, gotPath, rawQuery string
	var body []byte
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					var err error
					method = req.Method
					gotPath = req.URL.EscapedPath()
					rawQuery = req.URL.RawQuery
					body, err = io.ReadAll(req.Body)
					if err != nil {
						return nil, err
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("true")),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Experimental.Worktree.Remove(context.Background(), opencode.WorktreeRemoveParams{
		Directory:      opencode.F("/tmp/worktree"),
		QueryDirectory: opencode.F("/tmp/worktree"),
		Workspace:      opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodDelete {
		t.Fatalf("expected DELETE, got: %s", method)
	}
	if gotPath != "/experimental/worktree" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Fworktree") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if !strings.Contains(rawQuery, "workspace=workspace") {
		t.Fatalf("missing workspace query: %s", rawQuery)
	}
	if got := string(body); got != `{"directory":"/tmp/worktree"}` {
		t.Fatalf("unexpected body: %s", got)
	}
}

func TestWorktreeResetSendsDirectoryInQueryAndBody(t *testing.T) {
	var method, gotPath, rawQuery string
	var body []byte
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					var err error
					method = req.Method
					gotPath = req.URL.EscapedPath()
					rawQuery = req.URL.RawQuery
					body, err = io.ReadAll(req.Body)
					if err != nil {
						return nil, err
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("true")),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Experimental.Worktree.Reset(context.Background(), opencode.WorktreeResetParams{
		Directory:      opencode.F("/tmp/worktree"),
		QueryDirectory: opencode.F("/tmp/worktree"),
		Workspace:      opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if gotPath != "/experimental/worktree/reset" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Fworktree") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if !strings.Contains(rawQuery, "workspace=workspace") {
		t.Fatalf("missing workspace query: %s", rawQuery)
	}
	if got := string(body); got != `{"directory":"/tmp/worktree"}` {
		t.Fatalf("unexpected body: %s", got)
	}
}
