package opencode_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
)

func TestConfigFormatterSupportsBooleanAndObjectShapes(t *testing.T) {
	t.Run("boolean", func(t *testing.T) {
		var absent opencode.Config
		if err := json.Unmarshal([]byte(`{}`), &absent); err != nil {
			t.Fatal(err)
		}

		var disabled opencode.Config
		if err := json.Unmarshal([]byte(`{"formatter":false}`), &disabled); err != nil {
			t.Fatal(err)
		}

		if reflect.DeepEqual(disabled.Formatter, absent.Formatter) {
			t.Fatal("formatter=false decodes identically to an absent formatter field")
		}
	})

	t.Run("object", func(t *testing.T) {
		var cfg opencode.Config
		if err := json.Unmarshal([]byte(`{"formatter":{"prettier":{"disabled":true}}}`), &cfg); err != nil {
			t.Fatal(err)
		}
		if len(cfg.Formatter) == 0 {
			t.Fatal("formatter object shape did not decode into the public config model")
		}
	})
}

func TestConfigLSPSupportsBooleanAndObjectShapes(t *testing.T) {
	t.Run("boolean", func(t *testing.T) {
		var absent opencode.Config
		if err := json.Unmarshal([]byte(`{}`), &absent); err != nil {
			t.Fatal(err)
		}

		var enabled opencode.Config
		if err := json.Unmarshal([]byte(`{"lsp":true}`), &enabled); err != nil {
			t.Fatal(err)
		}

		if reflect.DeepEqual(enabled.Lsp, absent.Lsp) {
			t.Fatal("lsp=true decodes identically to an absent lsp field")
		}
	})

	t.Run("object", func(t *testing.T) {
		var cfg opencode.Config
		if err := json.Unmarshal([]byte(`{"lsp":{"gopls":{"command":["gopls"]}}}`), &cfg); err != nil {
			t.Fatal(err)
		}
		if len(cfg.Lsp) == 0 {
			t.Fatal("lsp object shape did not decode into the public config model")
		}
	})
}

func TestMcpLocalConfigIncludesTimeoutField(t *testing.T) {
	if _, ok := reflect.TypeOf(opencode.McpLocalConfig{}).FieldByName("Timeout"); !ok {
		t.Fatal("McpLocalConfig is missing timeout")
	}
}

func TestConfigMcpRemotePayloadUsesRemoteVariant(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"type":"remote","url":"https://example.com/mcp"}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	if got := fmt.Sprintf("%T", cfg.AsUnion()); got != "opencode.McpRemoteConfig" {
		t.Fatalf("expected remote MCP payload to decode as remote variant, got %s", got)
	}
}

func TestConfigMcpRejectsMalformedRemoteVariant(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"type":"remote","command":["mcp-server"]}`), &cfg)
	if err == nil {
		t.Fatal("expected malformed remote MCP variant to be rejected")
	}
}

func TestMcpRemoteConfigIncludesOAuthAndTimeoutFields(t *testing.T) {
	rt := reflect.TypeOf(opencode.McpRemoteConfig{})
	if _, ok := rt.FieldByName("Timeout"); !ok {
		t.Fatal("McpRemoteConfig is missing timeout")
	}
	if _, ok := rt.FieldByName("OAuth"); !ok {
		if _, ok := rt.FieldByName("Oauth"); !ok {
			t.Fatal("McpRemoteConfig is missing oauth")
		}
	}
}
