package opencode_test

import (
	"encoding/json"
	"fmt"
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

func TestTodoIncludesPriority(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.Todo{}).FieldByName("Priority"); !ok {
		t.Fatal("Todo is missing priority")
	}
}

func TestTodoDoesNotIncludeLegacyFields(t *testing.T) {
	todoType := reflect.TypeOf(opencode.Todo{})
	if _, ok := todoType.FieldByName("ID"); ok {
		t.Fatal("Todo should not include id")
	}
	if _, ok := todoType.FieldByName("SessionID"); ok {
		t.Fatal("Todo should not include sessionID")
	}
}

func TestFilePartSourceRecognizesResourceType(t *testing.T) {
	if !opencode.FilePartSourceType("resource").IsKnown() {
		t.Fatal("resource is not a known FilePartSource type")
	}
}

func TestFilePartSourceUnionIncludesResourceSource(t *testing.T) {
	types := packageTypeSpecsSessionTest(t)
	for _, typeName := range []string{"ResourceSource", "ResourceSourceParam"} {
		if _, ok := types[typeName]; !ok {
			t.Fatalf("missing type declaration: %s", typeName)
		}
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

func TestAssistantMessageErrorRecognizesStructuredAndContextOverflowErrors(t *testing.T) {
	if !opencode.AssistantMessageErrorName("StructuredOutputError").IsKnown() {
		t.Fatal("assistant message error is missing StructuredOutputError")
	}
	if !opencode.AssistantMessageErrorName("ContextOverflowError").IsKnown() {
		t.Fatal("assistant message error is missing ContextOverflowError")
	}
}

func TestAssistantMessageErrorDecodesStructuredOutputAndContextOverflowVariants(t *testing.T) {
	for _, tc := range []struct {
		name string
		body string
		want string
	}{
		{name: "structured", body: `{"name":"StructuredOutputError","data":{"message":"bad schema","retries":2}}`, want: ".AssistantMessageErrorStructuredOutputError"},
		{name: "context overflow", body: `{"name":"ContextOverflowError","data":{"message":"too long","responseBody":"overflow"}}`, want: ".AssistantMessageErrorContextOverflowError"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var msgErr opencode.AssistantMessageError
			err := json.Unmarshal([]byte(tc.body), &msgErr)
			if err != nil {
				t.Fatal(err)
			}
			if got := fmt.Sprintf("%T", msgErr.AsUnion()); !strings.HasSuffix(got, tc.want) {
				t.Fatalf("unexpected assistant message error type: %s", got)
			}
		})
	}
}

func packageTypeSpecsSessionTest(t *testing.T) map[string]*ast.TypeSpec {
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
