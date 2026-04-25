package ssestream

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestSSEDecoderFlushesFinalEventAtEOF(t *testing.T) {
	decoder := NewDecoder(&http.Response{
		Header: http.Header{"Content-Type": []string{"text/event-stream"}},
		Body:   io.NopCloser(strings.NewReader("event: message\ndata: {\"ok\":true}")),
	})
	if !decoder.Next() {
		t.Fatalf("expected final event before EOF, err=%v", decoder.Err())
	}
	evt := decoder.Event()
	if evt.Type != "message" {
		t.Fatalf("unexpected event type: %q", evt.Type)
	}
	if got := string(evt.Data); !strings.Contains(got, `{"ok":true}`) {
		t.Fatalf("unexpected event data: %q", got)
	}
}
