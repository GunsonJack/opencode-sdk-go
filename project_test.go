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

func TestProjectUpdateUsesCorrectMethodAndPath(t *testing.T) {
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
	_, err := client.Project.Update(context.Background(), "proj/special id", opencode.ProjectUpdateParams{
		Name: opencode.F("my-project"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPatch {
		t.Fatalf("expected PATCH, got: %s", method)
	}
	if gotPath != "/project/proj%2Fspecial%20id" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestProjectInitGitUsesCorrectMethodAndPath(t *testing.T) {
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
	_, err := client.Project.InitGit(context.Background(), opencode.ProjectInitGitParams{})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if gotPath != "/project/git/init" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestProjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Project.List(context.TODO(), opencode.ProjectListParams{
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

func TestProjectCurrentWithOptionalParams(t *testing.T) {
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
	_, err := client.Project.Current(context.TODO(), opencode.ProjectCurrentParams{
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

func TestProjectUpdate(t *testing.T) {
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
	_, err := client.Project.Update(context.TODO(), "proj_123", opencode.ProjectUpdateParams{
		Name: opencode.F("my-project"),
		Icon: opencode.F(opencode.ProjectUpdateParamsIcon{
			URL:      opencode.F("https://example.com/icon.png"),
			Override: opencode.F("custom-icon"),
			Color:    opencode.F("#FF0000"),
		}),
		Commands: opencode.F(opencode.ProjectUpdateParamsCommands{
			Start: opencode.F("npm run dev"),
		}),
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

func TestProjectInitGit(t *testing.T) {
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
	_, err := client.Project.InitGit(context.TODO(), opencode.ProjectInitGitParams{
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
