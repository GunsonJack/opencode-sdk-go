package opencode_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
)

func TestTodoIncludesPriority(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.Todo{}).FieldByName("Priority"); !ok {
		t.Fatal("Todo is missing priority")
	}
}

func TestFilePartSourceRecognizesResourceType(t *testing.T) {
	if !opencode.FilePartSourceType("resource").IsKnown() {
		t.Fatal("resource is not a known FilePartSource type")
	}
}

func TestFilePartSourceDecodesResourceVariant(t *testing.T) {
	var source opencode.FilePartSource
	err := json.Unmarshal([]byte(`{"type":"resource","clientName":"docs","uri":"resource://doc","text":{"start":0,"end":4,"value":"body"}}`), &source)
	if err != nil {
		t.Fatal(err)
	}
	if got := fmt.Sprintf("%T", source.AsUnion()); !strings.HasSuffix(got, ".ResourceSource") {
		t.Fatalf("unexpected source type: %s", got)
	}
}
