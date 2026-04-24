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

func TestSyncStart(t *testing.T) {
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
	_, err := client.Sync.Start(context.TODO(), opencode.SyncStartParams{
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

func TestSyncReplayWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Sync.Replay(context.TODO(), opencode.SyncReplayParams{
		Directory: opencode.F("/tmp/test"),
		Events: opencode.F([]opencode.SyncReplayEvent{{
			ID:          opencode.F("evt_001"),
			AggregateID: opencode.F("ses_abc"),
			Seq:         opencode.F(int64(1)),
			Type:        opencode.F("session.created"),
			Data:        opencode.F(map[string]interface{}{"title": "test"}),
		}}),
		QueryDirectory: opencode.F("/tmp/test"),
		Workspace:      opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSyncHistoryWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Sync.History(context.TODO(), opencode.SyncHistoryParams{
		Body:      opencode.F(map[string]int64{"ses_abc": 0}),
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

func TestSyncStartUsesCorrectMethodAndPath(t *testing.T) {
	var method, path, rawQuery string
	var body []byte
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					path = req.URL.Path
					rawQuery = req.URL.RawQuery
					if req.Body != nil {
						body, _ = io.ReadAll(req.Body)
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

	_, err := client.Sync.Start(context.Background(), opencode.SyncStartParams{
		Directory: opencode.F("/tmp/test"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if path != "/sync/start" {
		t.Fatalf("unexpected path: %s", path)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Ftest") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	if got := string(body); got != "" {
		t.Fatalf("expected empty body, got: %s", got)
	}
}

func TestSyncReplayUsesCorrectMethodAndBody(t *testing.T) {
	var method, path, rawQuery string
	var body []byte
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					path = req.URL.Path
					rawQuery = req.URL.RawQuery
					body, _ = io.ReadAll(req.Body)
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"sessionID":"ses_abc"}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Sync.Replay(context.Background(), opencode.SyncReplayParams{
		Directory: opencode.F("/tmp/test"),
		Events: opencode.F([]opencode.SyncReplayEvent{{
			ID:          opencode.F("evt_001"),
			AggregateID: opencode.F("ses_abc"),
			Seq:         opencode.F(int64(1)),
			Type:        opencode.F("session.created"),
			Data:        opencode.F(map[string]interface{}{"title": "test"}),
		}}),
		QueryDirectory: opencode.F("/tmp/test"),
		Workspace:      opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if path != "/sync/replay" {
		t.Fatalf("unexpected path: %s", path)
	}
	if !strings.Contains(rawQuery, "directory=%2Ftmp%2Ftest") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("invalid json body: %v", err)
	}
	events, ok := got["events"].([]any)
	if !ok || len(events) != 1 {
		t.Fatalf("unexpected events: %#v", got["events"])
	}
}

func TestSyncHistoryUsesCorrectMethodAndBody(t *testing.T) {
	var method, path string
	var body []byte
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					method = req.Method
					path = req.URL.Path
					body, _ = io.ReadAll(req.Body)
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`[{"id":"evt_001","aggregate_id":"ses_abc","seq":1,"type":"session.created","data":{}}]`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	results, err := client.Sync.History(context.Background(), opencode.SyncHistoryParams{
		Body:      opencode.F(map[string]int64{"ses_abc": 0}),
		Directory: opencode.F("/tmp/test"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost {
		t.Fatalf("expected POST, got: %s", method)
	}
	if path != "/sync/history" {
		t.Fatalf("unexpected path: %s", path)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("invalid json body: %v", err)
	}
	if got["ses_abc"] != float64(0) {
		t.Fatalf("unexpected body: %#v", got)
	}
	if len(*results) != 1 || (*results)[0].ID != "evt_001" {
		t.Fatalf("unexpected results: %+v", results)
	}
	if (*results)[0].AggregateID != "ses_abc" {
		t.Fatalf("unexpected aggregate_id: %q", (*results)[0].AggregateID)
	}
}
