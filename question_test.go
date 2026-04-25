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

func TestQuestionListUsesCorrectMethodAndPath(t *testing.T) {
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
	_, err := client.Question.List(context.Background(), opencode.QuestionListParams{})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodGet {
		t.Fatalf("expected GET, got: %s", method)
	}
	if gotPath != "/question" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestQuestionReplyUsesCorrectMethodAndPath(t *testing.T) {
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
	_, err := client.Question.Reply(context.Background(), "que_123", opencode.QuestionReplyParams{
		Answers: opencode.F([]opencode.QuestionAnswer{{"answer1"}}),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if gotPath != "/question/que_123/reply" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestQuestionRejectUsesCorrectMethodAndPath(t *testing.T) {
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
	_, err := client.Question.Reject(context.Background(), "que_123", opencode.QuestionRejectParams{})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if gotPath != "/question/que_123/reject" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
}

func TestQuestionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Question.List(
		context.TODO(),
		opencode.QuestionListParams{
			Workspace: opencode.F("workspace"),
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuestionReplyWithOptionalParams(t *testing.T) {
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
	_, err := client.Question.Reply(
		context.TODO(),
		"que_test123",
		opencode.QuestionReplyParams{
			Answers:   opencode.F([]opencode.QuestionAnswer{{"option1", "option2"}}),
			Workspace: opencode.F("workspace"),
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuestionRejectWithOptionalParams(t *testing.T) {
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
	_, err := client.Question.Reject(
		context.TODO(),
		"que_test123",
		opencode.QuestionRejectParams{
			Workspace: opencode.F("workspace"),
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
