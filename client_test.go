// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/GunsonJack/opencode-sdk-go"
	"github.com/GunsonJack/opencode-sdk-go/internal"
	"github.com/GunsonJack/opencode-sdk-go/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.Session.List(context.Background(), opencode.SessionListParams{})
	if userAgent != fmt.Sprintf("Opencode/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Session.List(context.Background(), opencode.SessionListParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	attempts := len(retryCountHeaders)
	if attempts != 3 {
		t.Errorf("Expected %d attempts, got %d", 3, attempts)
	}

	expectedRetryCountHeaders := []string{"0", "1", "2"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestDeleteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeaderDel("X-Stainless-Retry-Count"),
	)
	_, err := client.Session.List(context.Background(), opencode.SessionListParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"", "", ""}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestOverwriteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeader("X-Stainless-Retry-Count", "42"),
	)
	_, err := client.Session.List(context.Background(), opencode.SessionListParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"42", "42", "42"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Session.List(context.Background(), opencode.SessionListParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.Session.List(cancelCtx, opencode.SessionListParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	_, err := client.Session.List(cancelCtx, opencode.SessionListParams{})
	if err == nil {
		t.Error("expected there to be a cancel error")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := opencode.NewClient(
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		_, err := client.Session.List(deadlineCtx, opencode.SessionListParams{})
		if err == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}

func TestContextDeadlineStreaming(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := opencode.NewClient(
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Status:     "200 OK",
							Body: io.NopCloser(
								io.Reader(readerFunc(func([]byte) (int, error) {
									<-req.Context().Done()
									return 0, req.Context().Err()
								})),
							),
						}, nil
					},
				},
			}),
		)
		stream := client.Event.ListStreaming(deadlineCtx, opencode.EventListParams{})
		for stream.Next() {
			_ = stream.Current()
		}
		if stream.Err() == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}

func TestContextDeadlineStreamingWithRequestTimeout(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})
	deadline := time.Now().Add(100 * time.Millisecond)

	go func() {
		client := opencode.NewClient(
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Status:     "200 OK",
							Body: io.NopCloser(
								io.Reader(readerFunc(func([]byte) (int, error) {
									<-req.Context().Done()
									return 0, req.Context().Err()
								})),
							),
						}, nil
					},
				},
			}),
		)
		stream := client.Event.ListStreaming(
			context.Background(),
			opencode.EventListParams{},
			option.WithRequestTimeout((100 * time.Millisecond)),
		)
		for stream.Next() {
			_ = stream.Current()
		}
		if stream.Err() == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}

func TestTuiPublishSendsRawBody(t *testing.T) {
	var body []byte
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					var err error
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

	_, err := client.Tui.Publish(context.Background(), opencode.TuiPublishParams{
		Body:      opencode.F[interface{}](map[string]interface{}{"type": "toast.show"}),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := string(body); got != `{"type":"toast.show"}` {
		t.Fatalf("unexpected body: %s", got)
	}
}

func TestTuiControlResponseSendsRawBody(t *testing.T) {
	var body []byte
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					var err error
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

	_, err := client.Tui.Control.Response(context.Background(), opencode.TuiControlResponseParams{
		Body:      opencode.F[interface{}](map[string]interface{}{"ok": true}),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := string(body); got != `{"ok":true}` {
		t.Fatalf("unexpected body: %s", got)
	}
}

func TestSessionPromptAsyncReturnsMeaningfulSuccess(t *testing.T) {
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusNoContent,
						Body:       io.NopCloser(strings.NewReader("")),
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
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEventQuestionRepliedAnswersDecode(t *testing.T) {
	var evt opencode.EventListResponse
	err := json.Unmarshal([]byte(`{"type":"question.replied","properties":{"sessionID":"ses_1","requestID":"que_1","answers":[["a"],["b","c"]]}}`), &evt)
	if err != nil {
		t.Fatal(err)
	}

	q, ok := evt.AsUnion().(opencode.EventListResponseEventQuestionReplied)
	if !ok {
		t.Fatalf("unexpected event type: %#v", evt.AsUnion())
	}
	want := []opencode.QuestionAnswer{{"a"}, {"b", "c"}}
	if !reflect.DeepEqual(q.Properties.Answers, want) {
		t.Fatalf("unexpected answers: %#v", q.Properties.Answers)
	}
}

func TestEventPermissionRepliedFieldsDecode(t *testing.T) {
	var evt opencode.EventListResponse
	err := json.Unmarshal([]byte(`{"type":"permission.replied","properties":{"sessionID":"ses_1","requestID":"per_1","reply":"once"}}`), &evt)
	if err != nil {
		t.Fatal(err)
	}

	p, ok := evt.AsUnion().(opencode.EventListResponseEventPermissionReplied)
	if !ok {
		t.Fatalf("unexpected event type: %#v", evt.AsUnion())
	}
	if p.Properties.RequestID != "per_1" {
		t.Fatalf("unexpected requestID: %q", p.Properties.RequestID)
	}
	if p.Properties.Reply != opencode.PermissionReplyParamsReplyOnce {
		t.Fatalf("unexpected reply: %q", p.Properties.Reply)
	}
}

func TestAppProvidersAcceptsDirectoryForCompatibility(t *testing.T) {
	var rawQuery string
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					rawQuery = req.URL.RawQuery
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"default":{},"providers":[]}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.App.Providers(context.Background(), opencode.AppProvidersParams{Directory: opencode.F("dir")})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(rawQuery, "directory=dir") {
		t.Fatalf("missing directory query: %s", rawQuery)
	}
}

func TestConfigModeBuildAndPlanCompatibility(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"mode":{"build":{"model":"x"},"plan":{"model":"y"}}}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Mode.Build.Model != "x" || cfg.Mode.Plan.Model != "y" {
		t.Fatalf("bad mode decode: %#v", cfg.Mode)
	}
}

func TestSessionCommandSerializesFileParts(t *testing.T) {
	var body []byte
	client := opencode.NewClient(
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					var err error
					body, err = io.ReadAll(req.Body)
					if err != nil {
						return nil, err
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{"info":{"id":"msg_1","agent":"build","cost":0,"mode":"primary","modelID":"model","parentID":"msg_0","path":{"cwd":"/tmp","root":"/tmp"},"providerID":"provider","role":"assistant","sessionID":"ses_123","time":{"created":0},"tokens":{"cache":{"read":0,"write":0},"input":0,"output":0,"reasoning":0}},"parts":[]}`)),
						Header:     http.Header{"Content-Type": []string{"application/json"}},
					}, nil
				},
			},
		}),
	)

	_, err := client.Session.Command(context.Background(), "ses_123", opencode.SessionCommandParams{
		Arguments: opencode.F("args"),
		Command:   opencode.F("cmd"),
		Parts: opencode.F([]opencode.SessionCommandParamsPart{opencode.FilePartInputParam{
			Mime: opencode.F("text/plain"),
			Type: opencode.F(opencode.FilePartInputTypeFile),
			URL:  opencode.F("file:///tmp/example.txt"),
		}}),
	})
	if err != nil {
		t.Fatal(err)
	}
	if got := string(body); got != `{"arguments":"args","command":"cmd","parts":[{"mime":"text/plain","type":"file","url":"file:///tmp/example.txt"}]}` {
		t.Fatalf("unexpected body: %s", got)
	}
}

type readerFunc func([]byte) (int, error)

func (f readerFunc) Read(p []byte) (int, error) { return f(p) }
func (f readerFunc) Close() error               { return nil }
