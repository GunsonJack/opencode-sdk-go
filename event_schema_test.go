package opencode_test

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
)

func TestMessageUpdatedEventPropertiesIncludeSessionID(t *testing.T) {
	propertiesType := reflect.TypeOf(opencode.EventListResponseEventMessageUpdatedProperties{})
	if _, ok := propertiesType.FieldByName("SessionID"); !ok {
		t.Fatal("message.updated properties are missing sessionID")
	}
	if _, ok := propertiesType.FieldByName("Info"); !ok {
		t.Fatal("message.updated properties are missing info")
	}
}

func TestMessagePartUpdatedEventPropertiesIncludeSessionIDPartAndTime(t *testing.T) {
	propertiesType := reflect.TypeOf(opencode.EventListResponseEventMessagePartUpdatedProperties{})
	if _, ok := propertiesType.FieldByName("SessionID"); !ok {
		t.Fatal("message.part.updated properties are missing sessionID")
	}
	if _, ok := propertiesType.FieldByName("Part"); !ok {
		t.Fatal("message.part.updated properties are missing part")
	}
	if _, ok := propertiesType.FieldByName("Time"); !ok {
		t.Fatal("message.part.updated properties are missing time")
	}
}

func TestSessionCreatedEventPropertiesIncludeSessionID(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionCreatedProperties{}).FieldByName("SessionID"); !ok {
		t.Fatal("session.created properties are missing sessionID")
	}
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionCreatedProperties{}).FieldByName("Info"); !ok {
		t.Fatal("session.created properties are missing info")
	}
}

func TestSessionUpdatedEventPropertiesIncludeSessionID(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionUpdatedProperties{}).FieldByName("SessionID"); !ok {
		t.Fatal("session.updated properties are missing sessionID")
	}
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionUpdatedProperties{}).FieldByName("Info"); !ok {
		t.Fatal("session.updated properties are missing info")
	}
}

func TestSessionDeletedEventPropertiesIncludeSessionID(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionDeletedProperties{}).FieldByName("SessionID"); !ok {
		t.Fatal("session.deleted properties are missing sessionID")
	}
	if _, ok := reflect.TypeOf(opencode.EventListResponseEventSessionDeletedProperties{}).FieldByName("Info"); !ok {
		t.Fatal("session.deleted properties are missing info")
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

func TestMessageUpdatedDecodeDoesNotTreatSessionIDAsExtraField(t *testing.T) {
	var evt opencode.EventListResponse
	err := json.Unmarshal([]byte(`{"type":"message.updated","properties":{"sessionID":"ses_123","info":{"id":"msg_123","role":"user","agent":"assistant","model":{"modelID":"gpt-5","providerID":"openai"},"sessionID":"ses_123","time":{"created":0}}}}`), &evt)
	if err != nil {
		t.Fatal(err)
	}

	updated, ok := evt.AsUnion().(opencode.EventListResponseEventMessageUpdated)
	if !ok {
		t.Fatalf("unexpected event type: %#v", evt.AsUnion())
	}
	if hasJSONExtraField(updated.Properties, "sessionID") {
		t.Fatal("message.updated sessionID decoded as an unknown extra field")
	}
	sessionID := reflect.ValueOf(updated.Properties).FieldByName("SessionID")
	if !sessionID.IsValid() {
		t.Fatal("message.updated properties are missing sessionID after decode")
	}
	if sessionID.String() != "ses_123" {
		t.Fatalf("unexpected sessionID: %q", sessionID.String())
	}
}

func TestMessagePartUpdatedDecodeDoesNotTreatFieldsAsExtra(t *testing.T) {
	var evt opencode.EventListResponse
	err := json.Unmarshal([]byte(`{"type":"message.part.updated","properties":{"sessionID":"ses_123","part":{"id":"prt_123","messageID":"msg_123","sessionID":"ses_123","type":"text","text":"hello"},"time":1.5}}`), &evt)
	if err != nil {
		t.Fatal(err)
	}

	updated, ok := evt.AsUnion().(opencode.EventListResponseEventMessagePartUpdated)
	if !ok {
		t.Fatalf("unexpected event type: %#v", evt.AsUnion())
	}
	for _, field := range []string{"sessionID", "time"} {
		if hasJSONExtraField(updated.Properties, field) {
			t.Fatalf("message.part.updated %s decoded as an unknown extra field", field)
		}
	}
	sessionID := reflect.ValueOf(updated.Properties).FieldByName("SessionID")
	if !sessionID.IsValid() {
		t.Fatal("message.part.updated properties are missing sessionID after decode")
	}
	if sessionID.String() != "ses_123" {
		t.Fatalf("unexpected sessionID: %q", sessionID.String())
	}
	timeField := reflect.ValueOf(updated.Properties).FieldByName("Time")
	if !timeField.IsValid() {
		t.Fatal("message.part.updated properties are missing time after decode")
	}
	if timeField.Float() != 1.5 {
		t.Fatalf("unexpected time: %v", timeField.Float())
	}
}

func TestSyncEventTypesExistForGlobalEventReuse(t *testing.T) {
	types := packageTypeSpecsEventTest(t)
	for _, want := range []string{
		"SyncEvent",
		"SyncEventMessageUpdated",
		"SyncEventMessageRemoved",
		"SyncEventMessagePartUpdated",
		"SyncEventMessagePartRemoved",
		"SyncEventSessionCreated",
		"SyncEventSessionUpdated",
		"SyncEventSessionDeleted",
	} {
		if _, ok := types[want]; !ok {
			t.Fatalf("missing type declaration: %s", want)
		}
	}
}

func TestSyncEventStructsMatchSpecDefinedShapes(t *testing.T) {
	types := packageTypeSpecsEventTest(t)
	for _, tc := range []struct {
		name   string
		fields []string
	}{
		{name: "SyncEvent", fields: []string{"Type", "Name", "ID", "Seq", "AggregateID", "Data"}},
		{name: "SyncEventMessageUpdated", fields: []string{"Type", "Name", "ID", "Seq", "AggregateID", "Data"}},
		{name: "SyncEventMessageRemoved", fields: []string{"Type", "Name", "ID", "Seq", "AggregateID", "Data"}},
		{name: "SyncEventMessagePartUpdated", fields: []string{"Type", "Name", "ID", "Seq", "AggregateID", "Data"}},
		{name: "SyncEventMessagePartRemoved", fields: []string{"Type", "Name", "ID", "Seq", "AggregateID", "Data"}},
		{name: "SyncEventSessionCreated", fields: []string{"Type", "Name", "ID", "Seq", "AggregateID", "Data"}},
		{name: "SyncEventSessionUpdated", fields: []string{"Type", "Name", "ID", "Seq", "AggregateID", "Data"}},
		{name: "SyncEventSessionDeleted", fields: []string{"Type", "Name", "ID", "Seq", "AggregateID", "Data"}},
	} {
		spec := types[tc.name]
		st, ok := spec.Type.(*ast.StructType)
		if !ok {
			t.Fatalf("%s is not a struct type", tc.name)
		}
		for _, field := range tc.fields {
			if !astStructHasField(st, field) {
				t.Fatalf("%s is missing field %s", tc.name, field)
			}
		}
	}
}

func TestSyncEventSessionUpdatedUsesSparseInfoPayloadShape(t *testing.T) {
	var evt opencode.SyncEvent
	err := json.Unmarshal([]byte(`{"type":"sync","name":"session.updated.1","id":"evt_123","seq":1,"aggregateID":"sessionID","data":{"sessionID":"ses_123","info":{"id":null,"slug":null,"projectID":null,"workspaceID":null,"directory":null,"parentID":null,"summary":null,"share":{"url":null},"title":null,"version":null,"time":{"created":null,"updated":null,"compacting":null,"archived":null},"permission":null,"revert":null}}}`), &evt)
	if err != nil {
		t.Fatal(err)
	}

	updated, ok := evt.AsUnion().(opencode.SyncEventSessionUpdated)
	if !ok {
		t.Fatalf("unexpected sync event type: %#v", evt.AsUnion())
	}
	if updated.Data.SessionID != "ses_123" {
		t.Fatalf("unexpected sessionID: %q", updated.Data.SessionID)
	}
	if updated.Data.Info.ID != nil {
		t.Fatal("expected sparse info.id to accept null")
	}
	if updated.Data.Info.Time.Created != nil || updated.Data.Info.Time.Updated != nil || updated.Data.Info.Time.Compacting != nil || updated.Data.Info.Time.Archived != nil {
		t.Fatal("expected sparse info.time fields to accept null")
	}
	if updated.Data.Info.Share.URL != nil {
		t.Fatal("expected sparse info.share.url to accept null")
	}
	if updated.Data.Info.Permission != nil {
		t.Fatal("expected sparse info.permission to accept null")
	}
	if updated.Data.Info.Revert != nil {
		t.Fatal("expected sparse info.revert to accept null")
	}
	if hasJSONExtraField(updated.Data, "sessionID") {
		t.Fatal("sync session.updated data.sessionID decoded as an unknown extra field")
	}
	if hasJSONExtraField(updated.Data.Info, "id") || hasJSONExtraField(updated.Data.Info, "workspaceID") || hasJSONExtraField(updated.Data.Info, "permission") {
		t.Fatal("sync session.updated info fields decoded as unknown extra fields")
	}
}

func packageTypeSpecsEventTest(t *testing.T) map[string]*ast.TypeSpec {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	root := wd
	if filepath.Base(root) == "opencode_test" {
		root = filepath.Dir(root)
	}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, root, func(info os.FileInfo) bool {
		return strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "_test.go")
	}, 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg, ok := pkgs["opencode"]
	if !ok {
		t.Fatal("package opencode not found")
	}
	types := map[string]*ast.TypeSpec{}
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				typeSpec := spec.(*ast.TypeSpec)
				types[typeSpec.Name.Name] = typeSpec
			}
		}
	}
	return types
}

func astStructHasField(st *ast.StructType, name string) bool {
	for _, field := range st.Fields.List {
		for _, ident := range field.Names {
			if ident.Name == name {
				return true
			}
		}
	}
	return false
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
