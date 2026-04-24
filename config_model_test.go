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
		if got, ok := disabled.Formatter.(opencode.ConfigFormatterBool); !ok || bool(got) != false {
			t.Fatalf("expected formatter=false to preserve a boolean union, got %#v", disabled.Formatter)
		}
	})

	t.Run("object", func(t *testing.T) {
		var cfg opencode.Config
		if err := json.Unmarshal([]byte(`{"formatter":{"prettier":{"disabled":true}}}`), &cfg); err != nil {
			t.Fatal(err)
		}
		formatter, ok := cfg.Formatter.(opencode.ConfigFormatterObject)
		if !ok {
			t.Fatalf("expected formatter object union, got %#v", cfg.Formatter)
		}
		if len(formatter) == 0 {
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
		if got, ok := enabled.Lsp.(opencode.ConfigLspConfigBool); !ok || bool(got) != true {
			t.Fatalf("expected lsp=true to preserve a boolean union, got %#v", enabled.Lsp)
		}
	})

	t.Run("object", func(t *testing.T) {
		var cfg opencode.Config
		if err := json.Unmarshal([]byte(`{"lsp":{"gopls":{"command":["gopls"]}}}`), &cfg); err != nil {
			t.Fatal(err)
		}
		lsp, ok := cfg.Lsp.(opencode.ConfigLspConfigObject)
		if !ok {
			t.Fatalf("expected lsp object union, got %#v", cfg.Lsp)
		}
		if len(lsp) == 0 {
			t.Fatal("lsp object shape did not decode into the public config model")
		}
	})
}

func TestConfigMcpDisabledOnlyObjectUsesDisabledVariant(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"mcp":{"demo":{"enabled":false}}}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	got := cfg.Mcp["demo"].AsUnion()
	if _, ok := got.(opencode.ConfigMcpDisabled); !ok {
		t.Fatalf("expected disabled-only MCP object to decode as disabled variant, got %T", got)
	}
}

func TestConfigMcpDisabledOnlyObjectRejectsExtraFields(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"mcp":{"demo":{"enabled":false,"url":"https://example.com"}}}`), &cfg)
	if err == nil {
		t.Fatal("expected disabled-only MCP object with extra fields to be rejected")
	}
}

func TestConfigMcpDisabledOnlyObjectRequiresEnabled(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"mcp":{"demo":{}}}`), &cfg)
	if err == nil {
		t.Fatal("expected disabled-only MCP object without enabled to be rejected")
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

func TestConfigMcpRejectsLocalPayloadWithoutType(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"command":["mcp-server"]}`), &cfg)
	if err == nil {
		t.Fatal("expected local-like MCP payload without type to be rejected")
	}
}

func TestConfigMcpRejectsRemotePayloadWithoutType(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"url":"https://example.com/mcp"}`), &cfg)
	if err == nil {
		t.Fatal("expected remote-like MCP payload without type to be rejected")
	}
}

func TestConfigMcpRejectsUnknownType(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"type":"bogus","enabled":false}`), &cfg)
	if err == nil {
		t.Fatal("expected unknown MCP type to be rejected")
	}
}

func TestConfigMcpRejectsNonBooleanEnabledWithoutType(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"enabled":"true"}`), &cfg)
	if err == nil {
		t.Fatal("expected non-boolean enabled without type to be rejected")
	}
}

func TestConfigMcpLocalTimeoutDecodes(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"type":"local","command":["mcp-server"],"timeout":1500}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	local, ok := cfg.AsUnion().(opencode.McpLocalConfig)
	if !ok {
		t.Fatalf("expected local MCP payload to decode as local variant, got %T", cfg.AsUnion())
	}
	if local.Timeout != 1500 {
		t.Fatalf("expected timeout 1500, got %v", local.Timeout)
	}
}

func TestConfigMcpRemoteOAuthObjectDecodes(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"type":"remote","url":"https://example.com/mcp","oauth":{"clientId":"abc","redirectUri":"http://127.0.0.1/callback"}}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	remote, ok := cfg.AsUnion().(opencode.McpRemoteConfig)
	if !ok {
		t.Fatalf("expected remote MCP payload to decode as remote variant, got %T", cfg.AsUnion())
	}
	oauth, ok := remote.OAuth.(opencode.McpOAuthConfig)
	if !ok {
		t.Fatalf("expected oauth object union, got %#v", remote.OAuth)
	}
	if oauth.ClientID != "abc" {
		t.Fatalf("expected oauth clientId abc, got %q", oauth.ClientID)
	}
}

func TestConfigMcpRemoteOAuthFalseDecodes(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"type":"remote","url":"https://example.com/mcp","oauth":false}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	remote, ok := cfg.AsUnion().(opencode.McpRemoteConfig)
	if !ok {
		t.Fatalf("expected remote MCP payload to decode as remote variant, got %T", cfg.AsUnion())
	}
	if got, ok := remote.OAuth.(opencode.McpRemoteConfigOAuthFalse); !ok || bool(got) != false {
		t.Fatalf("expected oauth=false union, got %#v", remote.OAuth)
	}
}

func TestConfigMcpRemoteTimeoutDecodes(t *testing.T) {
	var cfg opencode.ConfigMcp
	err := json.Unmarshal([]byte(`{"type":"remote","url":"https://example.com/mcp","timeout":2750}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	remote, ok := cfg.AsUnion().(opencode.McpRemoteConfig)
	if !ok {
		t.Fatalf("expected remote MCP payload to decode as remote variant, got %T", cfg.AsUnion())
	}
	if remote.Timeout != 2750 {
		t.Fatalf("expected timeout 2750, got %v", remote.Timeout)
	}
}

func TestConfigProviderModelMatchesSpecSurface(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{
		"provider": {
			"demo": {
				"models": {
					"gpt-demo": {
						"id": "gpt-demo",
						"name": "Demo",
					"family": "gpt",
					"interleaved": {
						"field": "reasoning_content"
					},
					"release_date": "2025-01-01",
					"cost": {
							"input": 1,
							"output": 2,
							"context_over_200k": {
								"input": 3,
								"output": 4,
								"cache_read": 5,
								"cache_write": 6
							}
						},
						"limit": {
							"context": 128000,
							"input": 64000,
							"output": 8192
						},
						"modalities": {
							"input": ["text"],
							"output": ["text"]
						},
						"provider": {
							"npm": "@demo/provider",
							"api": "responses"
						},
						"status": "deprecated",
						"headers": {
							"x-demo": "1"
						},
						"variants": {
							"fast": {
								"disabled": true,
								"note": "keep"
							}
						}
					}
				}
			}
		}
	}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	model := cfg.Provider["demo"].Models["gpt-demo"]
	if model.Family != "gpt" {
		t.Fatalf("expected family gpt, got %q", model.Family)
	}
	interleaved, ok := model.Interleaved.AsUnion().(opencode.ConfigProviderModelInterleavedObject)
	if !ok {
		t.Fatalf("expected interleaved object union, got %#v", model.Interleaved)
	}
	if interleaved.Field != "reasoning_content" {
		t.Fatalf("expected interleaved field reasoning_content, got %q", interleaved.Field)
	}
	if model.Provider.API != "responses" {
		t.Fatalf("expected provider.api responses, got %q", model.Provider.API)
	}
	if model.Limit.Input != 64000 {
		t.Fatalf("expected limit.input 64000, got %v", model.Limit.Input)
	}
	if model.Status != opencode.ConfigProviderModelsStatusDeprecated {
		t.Fatalf("expected deprecated status, got %q", model.Status)
	}
	if model.Cost.ContextOver200K.Input != 3 {
		t.Fatalf("expected context_over_200k.input 3, got %v", model.Cost.ContextOver200K.Input)
	}
	if model.Headers["x-demo"] != "1" {
		t.Fatalf("expected header x-demo=1, got %#v", model.Headers)
	}
	if _, ok := model.Variants["fast"]; !ok {
		t.Fatalf("expected fast variant, got %#v", model.Variants)
	}
	if got := model.Variants["fast"]["note"]; got != "keep" {
		t.Fatalf("expected variant note keep, got %#v", got)
	}
	if model.ReleaseDate != "2025-01-01" {
		t.Fatalf("expected release date to round-trip, got %q", model.ReleaseDate)
	}
}

func TestConfigProviderModelInterleavedBooleanDecodes(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{
		"provider": {
			"demo": {
				"models": {
					"gpt-demo": {
						"id": "gpt-demo",
						"name": "Demo",
						"interleaved": true,
						"cost": {"input": 1, "output": 2},
						"limit": {"context": 1, "output": 1},
						"modalities": {"input": ["text"], "output": ["text"]},
						"provider": {},
						"status": "alpha"
					}
				}
			}
		}
	}`), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	if got, ok := cfg.Provider["demo"].Models["gpt-demo"].Interleaved.AsUnion().(opencode.ConfigProviderModelInterleavedBool); !ok || bool(got) != true {
		t.Fatalf("expected interleaved=true union, got %#v", cfg.Provider["demo"].Models["gpt-demo"].Interleaved)
	}
}

func TestConfigProviderOptionsTimeoutRejectsTrue(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"provider":{"demo":{"options":{"timeout":true}}}}`), &cfg)
	if err == nil {
		t.Fatal("expected provider options timeout=true to be rejected")
	}
}

func TestConfigProviderOptionsTimeoutAcceptsFalse(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"provider":{"demo":{"options":{"timeout":false}}}}`), &cfg)
	if err != nil {
		t.Fatalf("expected provider options timeout=false to be accepted, got %v", err)
	}
}

func TestConfigProviderOptionsTimeoutRejectsFractionalNumber(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"provider":{"demo":{"options":{"timeout":1.5}}}}`), &cfg)
	if err == nil {
		t.Fatal("expected provider options timeout=1.5 to be rejected")
	}
}

func TestConfigProviderOptionsTimeoutRejectsZero(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"provider":{"demo":{"options":{"timeout":0}}}}`), &cfg)
	if err == nil {
		t.Fatal("expected provider options timeout=0 to be rejected")
	}
}

func TestConfigProviderOptionsTimeoutRejectsNegativeInteger(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{"provider":{"demo":{"options":{"timeout":-1}}}}`), &cfg)
	if err == nil {
		t.Fatal("expected provider options timeout=-1 to be rejected")
	}
}

func TestConfigProviderModelInterleavedRejectsUnknownField(t *testing.T) {
	var cfg opencode.Config
	err := json.Unmarshal([]byte(`{
		"provider": {
			"demo": {
				"models": {
					"gpt-demo": {
						"id": "gpt-demo",
						"name": "Demo",
						"interleaved": {
							"field": "nope"
						},
						"cost": {"input": 1, "output": 2},
						"limit": {"context": 1, "output": 1},
						"modalities": {"input": ["text"], "output": ["text"]},
						"provider": {},
						"status": "alpha"
					}
				}
			}
		}
	}`), &cfg)
	if err == nil {
		t.Fatal("expected invalid interleaved.field to be rejected")
	}
}
