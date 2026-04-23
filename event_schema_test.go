package opencode_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
)

func TestSessionCreatedEventPropertiesIncludeSessionID(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionCreatedProperties{}).FieldByName("SessionID"); !ok {
		t.Fatal("session.created properties are missing sessionID")
	}
}

func TestSessionUpdatedEventPropertiesIncludeSessionID(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionUpdatedProperties{}).FieldByName("SessionID"); !ok {
		t.Fatal("session.updated properties are missing sessionID")
	}
}

func TestSessionDeletedEventPropertiesIncludeSessionID(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionDeletedProperties{}).FieldByName("SessionID"); !ok {
		t.Fatal("session.deleted properties are missing sessionID")
	}
}

func TestSessionErrorEventRecognizesStructuredAndContextOverflowErrors(t *testing.T) {
	if !opencode.EventListResponseEventSessionErrorPropertiesErrorName("StructuredOutputError").IsKnown() {
		t.Fatal("session.error is missing StructuredOutputError")
	}
	if !opencode.EventListResponseEventSessionErrorPropertiesErrorName("ContextOverflowError").IsKnown() {
		t.Fatal("session.error is missing ContextOverflowError")
	}
}

func TestSessionCreatedDecodeDoesNotTreatSessionIDAsExtraField(t *testing.T) {
	var evt opencode.EventListResponse
	err := json.Unmarshal([]byte(`{"type":"session.created","properties":{"sessionID":"ses_123","info":{}}}`), &evt)
	if err != nil {
		t.Fatal(err)
	}

	created, ok := evt.AsUnion().(opencode.EventListResponseEventSessionCreated)
	if !ok {
		t.Fatalf("unexpected event type: %#v", evt.AsUnion())
	}
	if hasJSONExtraField(created.Properties, "sessionID") {
		t.Fatal("session.created sessionID decoded as an unknown extra field")
	}
}

func hasJSONExtraField(v interface{}, field string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	jsonField := rv.FieldByName("JSON")
	if !jsonField.IsValid() {
		return false
	}
	extra := jsonField.FieldByName("ExtraFields")
	if !extra.IsValid() || extra.IsNil() {
		return false
	}
	return extra.MapIndex(reflect.ValueOf(field)).IsValid()
}
