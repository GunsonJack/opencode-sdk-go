// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
	"github.com/GunsonJack/opencode-sdk-go/internal/testutil"
	"github.com/GunsonJack/opencode-sdk-go/option"
)

func TestGlobalHealth(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Global.Health(context.TODO())
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalDispose(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Global.Dispose(context.TODO())
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalUpgrade(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Global.Upgrade(context.TODO(), opencode.GlobalUpgradeParams{
		Target: opencode.F("1.0.0"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalEventStreaming(t *testing.T) {
	var gotPath string
	var gotAccept string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					gotPath = req.URL.EscapedPath()
					gotAccept = req.Header.Get("Accept")
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(
							"event: message\ndata: {\"directory\":\"/tmp/project\",\"payload\":{\"type\":\"session.idle\",\"properties\":{\"sessionID\":\"ses_123\"}}}\n\n",
						)),
						Header: http.Header{"Content-Type": []string{"text/event-stream"}},
					}, nil
				},
			},
		}),
	)

	stream := client.Global.Event(context.Background())
	if !stream.Next() {
		t.Fatalf("expected stream event, err=%v", stream.Err())
	}
	if gotPath != "/global/event" {
		t.Fatalf("unexpected path: %s", gotPath)
	}
	if gotAccept != "text/event-stream" {
		t.Fatalf("unexpected Accept header: %q", gotAccept)
	}

	evt := stream.Current()
	idle, ok := evt.Payload.(opencode.EventListResponseEventSessionIdle)
	if !ok {
		t.Fatalf("unexpected payload type: %T", evt.Payload)
	}
	if idle.Properties.SessionID != "ses_123" {
		t.Fatalf("unexpected sessionID: %q", idle.Properties.SessionID)
	}
	if evt.Directory != "/tmp/project" {
		t.Fatalf("unexpected directory: %q", evt.Directory)
	}
	if stream.Err() != nil {
		t.Fatalf("unexpected stream error: %v", stream.Err())
	}
}

func TestGlobalEventDecodesFullEnvelopeWithSharedEventPayload(t *testing.T) {
	var evt opencode.GlobalEvent
	err := json.Unmarshal([]byte(`{"directory":"/tmp/project","project":"proj_123","workspace":"ws_123","payload":{"type":"session.idle","properties":{"sessionID":"ses_123"}}}`), &evt)
	if err != nil {
		t.Fatal(err)
	}

	if evt.Directory != "/tmp/project" {
		t.Fatalf("unexpected directory: %q", evt.Directory)
	}
	if evt.Project != "proj_123" {
		t.Fatalf("unexpected project: %q", evt.Project)
	}
	if evt.Workspace != "ws_123" {
		t.Fatalf("unexpected workspace: %q", evt.Workspace)
	}

	idle, ok := evt.Payload.(opencode.EventListResponseEventSessionIdle)
	if !ok {
		t.Fatalf("expected shared event payload type, got %T", evt.Payload)
	}
	if idle.Properties.SessionID != "ses_123" {
		t.Fatalf("unexpected sessionID: %q", idle.Properties.SessionID)
	}
	if strings.Contains(fmt.Sprintf("%T", evt.Payload), "GlobalEvent") {
		t.Fatal("global event payload should reuse shared event payload types")
	}
}

func TestGlobalEventDecodesSyncEventPayload(t *testing.T) {
	var evt opencode.GlobalEvent
	err := json.Unmarshal([]byte(`{"directory":"/tmp/project","payload":{"type":"sync","name":"session.updated.1","id":"evt_123","seq":1,"aggregateID":"sessionID","data":{"sessionID":"ses_123","info":{"id":null,"slug":null,"projectID":null,"workspaceID":null,"directory":null,"parentID":null,"summary":null,"share":{"url":null},"title":null,"version":null,"time":{"created":null,"updated":null,"compacting":null,"archived":null},"permission":null,"revert":null}}}}`), &evt)
	if err != nil {
		t.Fatal(err)
	}

	updated, ok := evt.Payload.(opencode.SyncEventSessionUpdated)
	if !ok {
		t.Fatalf("expected shared sync event payload type, got %T", evt.Payload)
	}
	if updated.Data.SessionID != "ses_123" {
		t.Fatalf("unexpected sessionID: %q", updated.Data.SessionID)
	}
}
