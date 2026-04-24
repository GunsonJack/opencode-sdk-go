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

func TestSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.New(context.TODO(), opencode.SessionNewParams{
		Directory: opencode.F("directory"),
		ParentID:  opencode.F("sesJ!"),
		Title:     opencode.F("title"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionAbortUsesQueryWithoutJSONBody(t *testing.T) {
	var rawQuery string
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					rawQuery = req.URL.RawQuery
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
						Body:       io.NopCloser(strings.NewReader("true")),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Session.Abort(context.Background(), "ses_123", opencode.SessionAbortParams{
		Directory: opencode.F("/tmp/project"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Fproject") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if !strings.Contains(rawQuery, "workspace=workspace") {
		t.Fatalf("missing workspace query: %s", rawQuery)
	}
	if got := string(body); got != "" {
		t.Fatalf("expected empty body, got: %s", got)
	}
	if contentType != "" {
		t.Fatalf("expected no content-type for body-less request, got: %s", contentType)
	}
}

func TestSessionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Update(
		context.TODO(),
		"id",
		opencode.SessionUpdateParams{
			Directory: opencode.F("directory"),
			Title:     opencode.F("title"),
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

func TestSessionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.List(context.TODO(), opencode.SessionListParams{
		Directory: opencode.F("directory"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Delete(
		context.TODO(),
		"sesJ!",
		opencode.SessionDeleteParams{
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

func TestSessionDeleteMessageEscapesPathParameters(t *testing.T) {
	var path string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					path = req.URL.EscapedPath()
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("true")),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Session.DeleteMessage(context.Background(), "ses/123", "msg with space", opencode.SessionDeleteMessageParams{})
	if err != nil {
		t.Fatal(err)
	}
	if path != "/session/ses%2F123/message/msg%20with%20space" {
		t.Fatalf("unexpected path: %s", path)
	}
}

func TestSessionDeleteMessageRejectsEmptyPathIDs(t *testing.T) {
	tests := []struct {
		name      string
		sessionID string
		messageID string
	}{
		{name: "empty session id", sessionID: "", messageID: "msg_123"},
		{name: "empty message id", sessionID: "ses_123", messageID: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			client := opencode.NewClient(
				option.WithHTTPClient(&http.Client{
					Transport: &closureTransport{
						fn: func(req *http.Request) (*http.Response, error) {
							called = true
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader("true")),
								Header:     http.Header{"Content-Type": []string{"application/json"}},
							}, nil
						},
					},
				}),
			)

			_, err := client.Session.DeleteMessage(context.Background(), tt.sessionID, tt.messageID, opencode.SessionDeleteMessageParams{})
			if err == nil {
				t.Fatal("expected missing path id error")
			}
			if called {
				t.Fatal("request should not be sent when a required path id is empty")
			}
		})
	}
}

func TestSessionAbortWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Abort(
		context.TODO(),
		"id",
		opencode.SessionAbortParams{
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

func TestSessionChildrenWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Children(
		context.TODO(),
		"sesJ!",
		opencode.SessionChildrenParams{
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

func TestSessionCommandWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Command(
		context.TODO(),
		"id",
		opencode.SessionCommandParams{
			Arguments: opencode.F("arguments"),
			Command:   opencode.F("command"),
			Directory: opencode.F("directory"),
			Agent:     opencode.F("agent"),
			MessageID: opencode.F("msgJ!"),
			Model:     opencode.F("model"),
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

func TestSessionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Get(
		context.TODO(),
		"sesJ!",
		opencode.SessionGetParams{
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

func TestSessionInitWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Init(
		context.TODO(),
		"id",
		opencode.SessionInitParams{
			MessageID:  opencode.F("msgJ!"),
			ModelID:    opencode.F("modelID"),
			ProviderID: opencode.F("providerID"),
			Directory:  opencode.F("directory"),
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

func TestSessionMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Message(
		context.TODO(),
		"id",
		"messageID",
		opencode.SessionMessageParams{
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

func TestSessionMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Messages(
		context.TODO(),
		"id",
		opencode.SessionMessagesParams{
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

func TestSessionPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Prompt(
		context.TODO(),
		"id",
		opencode.SessionPromptParams{
			Parts: opencode.F([]opencode.SessionPromptParamsPartUnion{opencode.TextPartInputParam{
				Text: opencode.F("text"),
				Type: opencode.F(opencode.TextPartInputTypeText),
				ID:   opencode.F("id"),
				Metadata: opencode.F(map[string]interface{}{
					"foo": "bar",
				}),
				Synthetic: opencode.F(true),
				Time: opencode.F(opencode.TextPartInputTimeParam{
					Start: opencode.F(0.000000),
					End:   opencode.F(0.000000),
				}),
			}}),
			Directory: opencode.F("directory"),
			Agent:     opencode.F("agent"),
			MessageID: opencode.F("msgJ!"),
			Model: opencode.F(opencode.SessionPromptParamsModel{
				ModelID:    opencode.F("modelID"),
				ProviderID: opencode.F("providerID"),
			}),
			NoReply: opencode.F(true),
			System:  opencode.F("system"),
			Tools: opencode.F(map[string]bool{
				"foo": true,
			}),
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

func TestSessionRevertWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Revert(
		context.TODO(),
		"id",
		opencode.SessionRevertParams{
			MessageID: opencode.F("msgJ!"),
			Directory: opencode.F("directory"),
			PartID:    opencode.F("prtJ!"),
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

func TestSessionShareWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Share(
		context.TODO(),
		"id",
		opencode.SessionShareParams{
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

func TestSessionShareUsesQueryWithoutJSONBody(t *testing.T) {
	var rawQuery string
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					rawQuery = req.URL.RawQuery
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
						Body:       io.NopCloser(strings.NewReader(`{"id":"ses_123","title":"shared","versionID":"ver_123","time":{"created":0,"updated":0},"projectID":"proj_123","parentSessionID":"","share":{"url":"https://example.com/share","id":"shr_123"},"messages":[]}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Session.Share(context.Background(), "ses_123", opencode.SessionShareParams{
		Directory: opencode.F("/tmp/project"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Fproject") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if !strings.Contains(rawQuery, "workspace=workspace") {
		t.Fatalf("missing workspace query: %s", rawQuery)
	}
	if got := string(body); got != "" {
		t.Fatalf("expected empty body, got: %s", got)
	}
	if contentType != "" {
		t.Fatalf("expected no content-type for body-less request, got: %s", contentType)
	}
}

func TestSessionShellWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Shell(
		context.TODO(),
		"id",
		opencode.SessionShellParams{
			Agent:     opencode.F("agent"),
			Command:   opencode.F("command"),
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

func TestSessionSummarizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Summarize(
		context.TODO(),
		"id",
		opencode.SessionSummarizeParams{
			ModelID:    opencode.F("modelID"),
			ProviderID: opencode.F("providerID"),
			Directory:  opencode.F("directory"),
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

func TestSessionUnrevertWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Unrevert(
		context.TODO(),
		"id",
		opencode.SessionUnrevertParams{
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

func TestSessionUnrevertUsesQueryWithoutJSONBody(t *testing.T) {
	var rawQuery string
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					rawQuery = req.URL.RawQuery
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
						Body:       io.NopCloser(strings.NewReader(`{"id":"ses_123","title":"session","versionID":"ver_123","time":{"created":0,"updated":0},"projectID":"proj_123","parentSessionID":"","share":{"url":"https://example.com/share","id":"shr_123"},"messages":[]}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Session.Unrevert(context.Background(), "ses_123", opencode.SessionUnrevertParams{
		Directory: opencode.F("/tmp/project"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Fproject") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if !strings.Contains(rawQuery, "workspace=workspace") {
		t.Fatalf("missing workspace query: %s", rawQuery)
	}
	if got := string(body); got != "" {
		t.Fatalf("expected empty body, got: %s", got)
	}
	if contentType != "" {
		t.Fatalf("expected no content-type for body-less request, got: %s", contentType)
	}
}

func TestSessionUnshareWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Unshare(
		context.TODO(),
		"sesJ!",
		opencode.SessionUnshareParams{
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

func TestSessionUnshareUsesQueryWithoutJSONBody(t *testing.T) {
	var rawQuery string
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					rawQuery = req.URL.RawQuery
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
						Body:       io.NopCloser(strings.NewReader(`{"id":"ses_123","title":"unshared","versionID":"ver_123","time":{"created":0,"updated":0},"projectID":"proj_123","parentSessionID":"","share":{"url":"","id":""},"messages":[]}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Session.Unshare(context.Background(), "ses_123", opencode.SessionUnshareParams{
		Directory: opencode.F("/tmp/project"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Fproject") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if !strings.Contains(rawQuery, "workspace=workspace") {
		t.Fatalf("missing workspace query: %s", rawQuery)
	}
	if got := string(body); got != "" {
		t.Fatalf("expected empty body, got: %s", got)
	}
	if contentType != "" {
		t.Fatalf("expected no content-type for body-less request, got: %s", contentType)
	}
}

func TestSessionDeletePartUsesQueryWithoutJSONBody(t *testing.T) {
	var rawQuery string
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					rawQuery = req.URL.RawQuery
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
						Body:       io.NopCloser(strings.NewReader("true")),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Session.DeletePart(context.Background(), "ses_123", "msg_123", "part_123", opencode.SessionDeletePartParams{
		Directory: opencode.F("/tmp/project"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Fproject") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if !strings.Contains(rawQuery, "workspace=workspace") {
		t.Fatalf("missing workspace query: %s", rawQuery)
	}
	if got := string(body); got != "" {
		t.Fatalf("expected empty body, got: %s", got)
	}
	if contentType != "" {
		t.Fatalf("expected no content-type for body-less request, got: %s", contentType)
	}
}

func TestSessionPromptAsyncUsesQueryAndJSONBody(t *testing.T) {
	var rawQuery string
	var body []byte
	var contentType string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					var err error
					rawQuery = req.URL.RawQuery
					contentType = req.Header.Get("Content-Type")
					body, err = io.ReadAll(req.Body)
					if err != nil {
						return nil, err
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader("")),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	err := client.Session.PromptAsync(context.Background(), "ses_123", opencode.SessionPromptAsyncParams{
		Parts: opencode.F([]opencode.SessionPromptParamsPartUnion{opencode.TextPartInputParam{
			Text: opencode.F("hello"),
			Type: opencode.F(opencode.TextPartInputTypeText),
		}}),
		Directory: opencode.F("/tmp/project"),
		Workspace: opencode.F("workspace"),
		Agent:     opencode.F("agent"),
		MessageID: opencode.F("msg_123"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Fproject") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if !strings.Contains(rawQuery, "workspace=workspace") {
		t.Fatalf("missing workspace query: %s", rawQuery)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("invalid json body: %v", err)
	}
	if got["agent"] != "agent" {
		t.Fatalf("unexpected agent field: %#v", got["agent"])
	}
	if got["messageID"] != "msg_123" {
		t.Fatalf("unexpected messageID field: %#v", got["messageID"])
	}
	parts, ok := got["parts"].([]any)
	if !ok || len(parts) != 1 {
		t.Fatalf("unexpected parts payload: %#v", got["parts"])
	}
	part, ok := parts[0].(map[string]any)
	if !ok {
		t.Fatalf("unexpected part payload: %#v", parts[0])
	}
	if part["text"] != "hello" || part["type"] != "text" {
		t.Fatalf("unexpected part payload: %#v", part)
	}
	if contentType != "application/json" {
		t.Fatalf("expected JSON content-type, got: %s", contentType)
	}
}
