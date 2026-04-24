// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
	"github.com/GunsonJack/opencode-sdk-go/internal/apiquery"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
	"github.com/GunsonJack/opencode-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// ConfigService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// Get config info
func (r *ConfigService) Get(ctx context.Context, query ConfigGetParams, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update config
func (r *ConfigService) Update(ctx context.Context, params ConfigUpdateParams, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all providers
func (r *ConfigService) Providers(ctx context.Context, query ConfigProvidersParams, opts ...option.RequestOption) (res *ConfigProvidersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ConfigProvidersResponse struct {
	Providers []Provider                  `json:"providers,required"`
	Default   map[string]string           `json:"default,required"`
	JSON      configProvidersResponseJSON `json:"-"`
}

// configProvidersResponseJSON contains the JSON metadata for the struct
// [ConfigProvidersResponse]
type configProvidersResponseJSON struct {
	Providers   apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProvidersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProvidersResponseJSON) RawJSON() string {
	return r.raw
}

type Config struct {
	// JSON schema reference for configuration validation
	Schema string `json:"$schema"`
	// Agent configuration, see https://opencode.ai/docs/agents
	Agent ConfigAgent `json:"agent"`
	// @deprecated Use 'share' field instead. Share newly created sessions
	// automatically
	Autoshare bool `json:"autoshare"`
	// Automatically update to the latest version
	Autoupdate ConfigAutoupdateUnion `json:"autoupdate"`
	// Command configuration, see https://opencode.ai/docs/commands
	Command    map[string]ConfigCommand `json:"command"`
	Compaction ConfigCompaction         `json:"compaction"`
	// Default agent to use
	DefaultAgent string `json:"default_agent"`
	// Disable providers that are loaded automatically
	DisabledProviders []string `json:"disabled_providers"`
	// Enable specific providers
	EnabledProviders []string             `json:"enabled_providers"`
	Enterprise       ConfigEnterprise     `json:"enterprise"`
	Experimental     ConfigExperimental   `json:"experimental"`
	Formatter        ConfigFormatterUnion `json:"formatter"`
	// Additional instruction files or patterns to include
	Instructions []string `json:"instructions"`
	// @deprecated Use automatic layout behavior instead.
	Layout   LayoutConfig         `json:"layout"`
	LogLevel ConfigLogLevel       `json:"logLevel"`
	Lsp      ConfigLspConfigUnion `json:"lsp"`
	// MCP (Model Context Protocol) server configurations
	Mcp map[string]ConfigMcp `json:"mcp"`
	// Model to use in the format of provider/model, eg anthropic/claude-2
	Model string `json:"model"`
	// @deprecated Use 'agent' field instead.
	Mode       ConfigMode         `json:"mode"`
	Permission PermissionConfig   `json:"permission"`
	Plugin     []ConfigPluginItem `json:"plugin"`
	// Custom provider configurations and model overrides
	Provider map[string]ConfigProvider `json:"provider"`
	Server   ServerConfig              `json:"server"`
	// Control sharing behavior:'manual' allows manual sharing via commands, 'auto'
	// enables automatic sharing, 'disabled' disables all sharing
	Share  ConfigShare  `json:"share"`
	Skills ConfigSkills `json:"skills"`
	// Small model to use for tasks like title generation in the format of
	// provider/model
	SmallModel string          `json:"small_model"`
	Snapshot   bool            `json:"snapshot"`
	Tools      map[string]bool `json:"tools"`
	// Custom username to display in conversations instead of system username
	Username string        `json:"username"`
	Watcher  ConfigWatcher `json:"watcher"`
	JSON     configJSON    `json:"-"`
}

// configJSON contains the JSON metadata for the struct [Config]
type configJSON struct {
	Schema            apijson.Field
	Agent             apijson.Field
	Autoshare         apijson.Field
	Autoupdate        apijson.Field
	Command           apijson.Field
	Compaction        apijson.Field
	DefaultAgent      apijson.Field
	DisabledProviders apijson.Field
	EnabledProviders  apijson.Field
	Enterprise        apijson.Field
	Experimental      apijson.Field
	Formatter         apijson.Field
	Instructions      apijson.Field
	Layout            apijson.Field
	LogLevel          apijson.Field
	Lsp               apijson.Field
	Mcp               apijson.Field
	Model             apijson.Field
	Mode              apijson.Field
	Permission        apijson.Field
	Plugin            apijson.Field
	Provider          apijson.Field
	Server            apijson.Field
	Share             apijson.Field
	Skills            apijson.Field
	SmallModel        apijson.Field
	Snapshot          apijson.Field
	Tools             apijson.Field
	Username          apijson.Field
	Watcher           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Config) UnmarshalJSON(data []byte) (err error) {
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	parsed := gjson.ParseBytes(data)
	providerTimeouts := parsed.Get("provider")
	if providerTimeouts.Exists() {
		for _, provider := range providerTimeouts.Map() {
			timeout := provider.Get("options.timeout")
			if timeout.Type == gjson.True {
				return fmt.Errorf("invalid provider timeout: expected integer or false")
			}
			if timeout.Exists() && timeout.Type == gjson.Number && (timeout.Num != float64(int64(timeout.Num)) || timeout.Num <= 0) {
				return fmt.Errorf("invalid provider timeout: expected integer or false")
			}
			models := provider.Get("models")
			if models.Exists() {
				for _, model := range models.Map() {
					field := model.Get("interleaved.field")
					if field.Exists() && field.String() != "reasoning_content" && field.String() != "reasoning_details" {
						return fmt.Errorf("invalid interleaved field: %s", field.String())
					}
				}
			}
		}
	}
	mcpConfigs := parsed.Get("mcp")
	if mcpConfigs.Exists() {
		for _, mcp := range mcpConfigs.Map() {
			if !mcp.Get("type").Exists() {
				if !mcp.Get("enabled").Exists() {
					return fmt.Errorf("missing required field: enabled")
				}
				if (mcp.Get("enabled").Type != gjson.False && mcp.Get("enabled").Type != gjson.True) || len(mcp.Map()) > 1 {
					return fmt.Errorf("invalid disabled MCP config: unexpected extra fields")
				}
			}
		}
	}
	return nil
}

func (r configJSON) RawJSON() string {
	return r.raw
}

type ConfigMode struct {
	Build       ConfigAgentEntry            `json:"build"`
	Plan        ConfigAgentEntry            `json:"plan"`
	ExtraFields map[string]ConfigAgentEntry `json:"-,extras"`
	JSON        configModeJSON              `json:"-"`
}

type configModeJSON struct {
	Build       apijson.Field
	Plan        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configModeJSON) RawJSON() string {
	return r.raw
}

// ServerConfig represents the server configuration.
type ServerConfig struct {
	Port       int64            `json:"port"`
	Hostname   string           `json:"hostname"`
	Mdns       bool             `json:"mdns"`
	MdnsDomain string           `json:"mdnsDomain"`
	Cors       []string         `json:"cors"`
	JSON       serverConfigJSON `json:"-"`
}

type serverConfigJSON struct {
	Port        apijson.Field
	Hostname    apijson.Field
	Mdns        apijson.Field
	MdnsDomain  apijson.Field
	Cors        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serverConfigJSON) RawJSON() string {
	return r.raw
}

// ConfigSkills represents skills configuration.
type ConfigSkills struct {
	Paths []string         `json:"paths"`
	URLs  []string         `json:"urls"`
	JSON  configSkillsJSON `json:"-"`
}

type configSkillsJSON struct {
	Paths       apijson.Field
	URLs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigSkills) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configSkillsJSON) RawJSON() string {
	return r.raw
}

// ConfigCompaction represents compaction configuration.
type ConfigCompaction struct {
	Auto                 bool                 `json:"auto"`
	Prune                bool                 `json:"prune"`
	TailTurns            int64                `json:"tail_turns"`
	PreserveRecentTokens int64                `json:"preserve_recent_tokens"`
	Reserved             int64                `json:"reserved"`
	JSON                 configCompactionJSON `json:"-"`
}

type configCompactionJSON struct {
	Auto                 apijson.Field
	Prune                apijson.Field
	TailTurns            apijson.Field
	PreserveRecentTokens apijson.Field
	Reserved             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configCompactionJSON) RawJSON() string {
	return r.raw
}

// ConfigEnterprise represents enterprise configuration.
type ConfigEnterprise struct {
	URL  string               `json:"url"`
	JSON configEnterpriseJSON `json:"-"`
}

type configEnterpriseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigEnterprise) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configEnterpriseJSON) RawJSON() string {
	return r.raw
}

// LayoutConfig represents the layout mode.
type LayoutConfig string

const (
	LayoutConfigAuto    LayoutConfig = "auto"
	LayoutConfigStretch LayoutConfig = "stretch"
)

func (r LayoutConfig) IsKnown() bool {
	switch r {
	case LayoutConfigAuto, LayoutConfigStretch:
		return true
	}
	return false
}

// ConfigLogLevel represents the log level.
type ConfigLogLevel string

const (
	ConfigLogLevelDebug ConfigLogLevel = "DEBUG"
	ConfigLogLevelInfo  ConfigLogLevel = "INFO"
	ConfigLogLevelWarn  ConfigLogLevel = "WARN"
	ConfigLogLevelError ConfigLogLevel = "ERROR"
)

func (r ConfigLogLevel) IsKnown() bool {
	switch r {
	case ConfigLogLevelDebug, ConfigLogLevelInfo, ConfigLogLevelWarn, ConfigLogLevelError:
		return true
	}
	return false
}

// ConfigAutoupdateUnion represents the autoupdate configuration which can be a
// bool or a string.
type ConfigAutoupdateUnion interface {
	ImplementsConfigAutoupdateUnion()
}

type ConfigAutoupdateString string

const (
	ConfigAutoupdateStringNotify ConfigAutoupdateString = "notify"
)

func (r ConfigAutoupdateString) IsKnown() bool {
	switch r {
	case ConfigAutoupdateStringNotify:
		return true
	}
	return false
}

func (r ConfigAutoupdateString) ImplementsConfigAutoupdateUnion() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigAutoupdateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigAutoupdateString("")),
		},
	)
}

// ConfigPluginItem represents a plugin configuration item which can be a string
// or a tuple.
type ConfigPluginItem struct {
	JSON  configPluginItemJSON `json:"-"`
	union ConfigPluginItemUnion
}

type configPluginItemJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configPluginItemJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigPluginItem) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigPluginItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ConfigPluginItem) AsUnion() ConfigPluginItemUnion {
	return r.union
}

type ConfigPluginItemUnion interface {
	implementsConfigPluginItem()
}

type ConfigPluginItemString string

func (r ConfigPluginItemString) implementsConfigPluginItem() {}

type ConfigPluginItemTuple []interface{}

func (r ConfigPluginItemTuple) implementsConfigPluginItem() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigPluginItemUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigPluginItemString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigPluginItemTuple{}),
		},
	)
}

// Agent configuration, see https://opencode.ai/docs/agent
type ConfigAgent struct {
	Build       ConfigAgentEntry            `json:"build"`
	General     ConfigAgentEntry            `json:"general"`
	Plan        ConfigAgentEntry            `json:"plan"`
	Explore     ConfigAgentEntry            `json:"explore"`
	Title       ConfigAgentEntry            `json:"title"`
	Summary     ConfigAgentEntry            `json:"summary"`
	Compaction  ConfigAgentEntry            `json:"compaction"`
	ExtraFields map[string]ConfigAgentEntry `json:"-,extras"`
	JSON        configAgentJSON             `json:"-"`
}

// configAgentJSON contains the JSON metadata for the struct [ConfigAgent]
type configAgentJSON struct {
	Build       apijson.Field
	General     apijson.Field
	Plan        apijson.Field
	Explore     apijson.Field
	Title       apijson.Field
	Summary     apijson.Field
	Compaction  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentEntry struct {
	Description string                 `json:"description"`
	Disable     bool                   `json:"disable"`
	Mode        ConfigAgentEntryMode   `json:"mode"`
	Model       string                 `json:"model"`
	Permission  PermissionConfig       `json:"permission"`
	Prompt      string                 `json:"prompt"`
	Temperature float64                `json:"temperature"`
	Tools       map[string]bool        `json:"tools"`
	TopP        float64                `json:"top_p"`
	Hidden      bool                   `json:"hidden"`
	Color       string                 `json:"color"`
	Variant     string                 `json:"variant"`
	Steps       int64                  `json:"steps"`
	MaxSteps    int64                  `json:"maxSteps"`
	Options     map[string]interface{} `json:"options"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        configAgentEntryJSON   `json:"-"`
}

type configAgentEntryJSON struct {
	Description apijson.Field
	Disable     apijson.Field
	Mode        apijson.Field
	Model       apijson.Field
	Permission  apijson.Field
	Prompt      apijson.Field
	Temperature apijson.Field
	Tools       apijson.Field
	TopP        apijson.Field
	Hidden      apijson.Field
	Color       apijson.Field
	Variant     apijson.Field
	Steps       apijson.Field
	MaxSteps    apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigAgentEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentEntryJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentEntryMode string

const (
	ConfigAgentEntryModeSubagent ConfigAgentEntryMode = "subagent"
	ConfigAgentEntryModePrimary  ConfigAgentEntryMode = "primary"
	ConfigAgentEntryModeAll      ConfigAgentEntryMode = "all"
)

func (r ConfigAgentEntryMode) IsKnown() bool {
	switch r {
	case ConfigAgentEntryModeSubagent, ConfigAgentEntryModePrimary, ConfigAgentEntryModeAll:
		return true
	}
	return false
}

type ConfigCommand struct {
	Template    string            `json:"template,required"`
	Agent       string            `json:"agent"`
	Description string            `json:"description"`
	Model       string            `json:"model"`
	Subtask     bool              `json:"subtask"`
	JSON        configCommandJSON `json:"-"`
}

// configCommandJSON contains the JSON metadata for the struct [ConfigCommand]
type configCommandJSON struct {
	Template    apijson.Field
	Agent       apijson.Field
	Description apijson.Field
	Model       apijson.Field
	Subtask     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigCommand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configCommandJSON) RawJSON() string {
	return r.raw
}

type ConfigExperimental struct {
	DisablePasteSummary bool                   `json:"disable_paste_summary"`
	BatchTool           bool                   `json:"batch_tool"`
	OpenTelemetry       bool                   `json:"openTelemetry"`
	PrimaryTools        []string               `json:"primary_tools"`
	ContinueLoopOnDeny  bool                   `json:"continue_loop_on_deny"`
	McpTimeout          int64                  `json:"mcp_timeout"`
	JSON                configExperimentalJSON `json:"-"`
}

// configExperimentalJSON contains the JSON metadata for the struct
// [ConfigExperimental]
type configExperimentalJSON struct {
	DisablePasteSummary apijson.Field
	BatchTool           apijson.Field
	OpenTelemetry       apijson.Field
	PrimaryTools        apijson.Field
	ContinueLoopOnDeny  apijson.Field
	McpTimeout          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConfigExperimental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configExperimentalJSON) RawJSON() string {
	return r.raw
}

type ConfigFormatter struct {
	Command     []string            `json:"command"`
	Disabled    bool                `json:"disabled"`
	Environment map[string]string   `json:"environment"`
	Extensions  []string            `json:"extensions"`
	JSON        configFormatterJSON `json:"-"`
}

type ConfigFormatterObject map[string]ConfigFormatter

func (ConfigFormatterObject) implementsConfigFormatterUnion() {}

type ConfigFormatterBool bool

func (ConfigFormatterBool) implementsConfigFormatterUnion() {}

type ConfigFormatterUnion interface {
	implementsConfigFormatterUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigFormatterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(ConfigFormatterBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(ConfigFormatterBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigFormatterObject{}),
		},
	)
}

// configFormatterJSON contains the JSON metadata for the struct [ConfigFormatter]
type configFormatterJSON struct {
	Command     apijson.Field
	Disabled    apijson.Field
	Environment apijson.Field
	Extensions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigFormatter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFormatterJSON) RawJSON() string {
	return r.raw
}

type ConfigLspConfigObject map[string]ConfigLsp

func (ConfigLspConfigObject) implementsConfigLspConfigUnion() {}

type ConfigLspConfigBool bool

func (ConfigLspConfigBool) implementsConfigLspConfigUnion() {}

type ConfigLspConfigUnion interface {
	implementsConfigLspConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigLspConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(ConfigLspConfigBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(ConfigLspConfigBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigLspConfigObject{}),
		},
	)
}

type ConfigLsp struct {
	// This field can have the runtime type of [[]string].
	Command  interface{} `json:"command"`
	Disabled bool        `json:"disabled"`
	// This field can have the runtime type of [map[string]string].
	Env interface{} `json:"env"`
	// This field can have the runtime type of [[]string].
	Extensions interface{} `json:"extensions"`
	// This field can have the runtime type of [map[string]interface{}].
	Initialization interface{}   `json:"initialization"`
	JSON           configLspJSON `json:"-"`
	union          ConfigLspUnion
}

// configLspJSON contains the JSON metadata for the struct [ConfigLsp]
type configLspJSON struct {
	Command        apijson.Field
	Disabled       apijson.Field
	Env            apijson.Field
	Extensions     apijson.Field
	Initialization apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r configLspJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigLsp) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigLsp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigLspUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ConfigLspDisabled], [ConfigLspObject].
func (r ConfigLsp) AsUnion() ConfigLspUnion {
	return r.union
}

// Union satisfied by [ConfigLspDisabled] or [ConfigLspObject].
type ConfigLspUnion interface {
	implementsConfigLsp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigLspUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigLspDisabled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigLspObject{}),
		},
	)
}

type ConfigLspDisabled struct {
	Disabled ConfigLspDisabledDisabled `json:"disabled,required"`
	JSON     configLspDisabledJSON     `json:"-"`
}

// configLspDisabledJSON contains the JSON metadata for the struct
// [ConfigLspDisabled]
type configLspDisabledJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigLspDisabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configLspDisabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigLspDisabled) implementsConfigLsp() {}

type ConfigLspDisabledDisabled bool

const (
	ConfigLspDisabledDisabledTrue ConfigLspDisabledDisabled = true
)

func (r ConfigLspDisabledDisabled) IsKnown() bool {
	switch r {
	case ConfigLspDisabledDisabledTrue:
		return true
	}
	return false
}

type ConfigLspObject struct {
	Command        []string               `json:"command,required"`
	Disabled       bool                   `json:"disabled"`
	Env            map[string]string      `json:"env"`
	Extensions     []string               `json:"extensions"`
	Initialization map[string]interface{} `json:"initialization"`
	JSON           configLspObjectJSON    `json:"-"`
}

// configLspObjectJSON contains the JSON metadata for the struct [ConfigLspObject]
type configLspObjectJSON struct {
	Command        apijson.Field
	Disabled       apijson.Field
	Env            apijson.Field
	Extensions     apijson.Field
	Initialization apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConfigLspObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configLspObjectJSON) RawJSON() string {
	return r.raw
}

func (r ConfigLspObject) implementsConfigLsp() {}

type ConfigMcp struct {
	// Type of MCP server connection
	Type ConfigMcpType `json:"type,required"`
	// This field can have the runtime type of [[]string].
	Command interface{} `json:"command"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [map[string]string].
	Environment interface{} `json:"environment"`
	// This field can have the runtime type of [map[string]string].
	Headers interface{} `json:"headers"`
	// URL of the remote MCP server
	URL     string               `json:"url"`
	OAuth   McpRemoteConfigOAuth `json:"oauth"`
	Timeout float64              `json:"timeout"`
	JSON    configMcpJSON        `json:"-"`
	union   ConfigMcpUnion
}

// configMcpJSON contains the JSON metadata for the struct [ConfigMcp]
type configMcpJSON struct {
	Type        apijson.Field
	Command     apijson.Field
	Enabled     apijson.Field
	Environment apijson.Field
	Headers     apijson.Field
	URL         apijson.Field
	OAuth       apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configMcpJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigMcp) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigMcp{}
	parsed := gjson.ParseBytes(data)
	if !parsed.Get("type").Exists() {
		if !parsed.Get("enabled").Exists() {
			return fmt.Errorf("missing required field: enabled")
		}
		if parsed.Get("enabled").Type != gjson.True && parsed.Get("enabled").Type != gjson.False {
			return fmt.Errorf("invalid disabled MCP config: enabled must be boolean")
		}
		if len(parsed.Map()) > 1 {
			return fmt.Errorf("invalid disabled MCP config: unexpected extra fields")
		}
	} else if parsed.Get("type").String() != string(ConfigMcpTypeLocal) && parsed.Get("type").String() != string(ConfigMcpTypeRemote) {
		return fmt.Errorf("invalid MCP type: %s", parsed.Get("type").String())
	}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	if err = apijson.Port(r.union, &r); err != nil {
		return err
	}
	switch r.Type {
	case ConfigMcpTypeLocal:
		if !parsed.Get("command").Exists() {
			return fmt.Errorf("missing required field: command")
		}
	case ConfigMcpTypeRemote:
		if !parsed.Get("url").Exists() {
			return fmt.Errorf("missing required field: url")
		}
	}
	return nil
}

// AsUnion returns a [ConfigMcpUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [McpLocalConfig], [McpRemoteConfig],
// [ConfigMcpDisabled].
func (r ConfigMcp) AsUnion() ConfigMcpUnion {
	return r.union
}

// Union satisfied by [McpLocalConfig], [McpRemoteConfig], or [ConfigMcpDisabled].
type ConfigMcpUnion interface {
	implementsConfigMcp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigMcpUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "local",
			Type:               reflect.TypeOf(McpLocalConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "remote",
			Type:               reflect.TypeOf(McpRemoteConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigMcpDisabled{}),
		},
	)
}

// Type of MCP server connection
type ConfigMcpType string

const (
	ConfigMcpTypeLocal  ConfigMcpType = "local"
	ConfigMcpTypeRemote ConfigMcpType = "remote"
)

func (r ConfigMcpType) IsKnown() bool {
	switch r {
	case ConfigMcpTypeLocal, ConfigMcpTypeRemote:
		return true
	}
	return false
}

// PermissionActionConfig is the simple action enum for config permissions.
// Spec: PermissionActionConfig (anyOf variant 1 of PermissionConfig)
type PermissionActionConfig string

const (
	PermissionActionConfigAsk   PermissionActionConfig = "ask"
	PermissionActionConfigAllow PermissionActionConfig = "allow"
	PermissionActionConfigDeny  PermissionActionConfig = "deny"
)

func (r PermissionActionConfig) IsKnown() bool {
	switch r {
	case PermissionActionConfigAsk, PermissionActionConfigAllow, PermissionActionConfigDeny:
		return true
	}
	return false
}

func (r PermissionActionConfig) implementsPermissionRuleConfig() {}
func (r PermissionActionConfig) implementsPermissionConfig()     {}

// PermissionObjectConfig is a map of tool names to permission actions.
// Spec: openapi.json:10962
type PermissionObjectConfig map[string]PermissionActionConfig

func (r PermissionObjectConfig) implementsPermissionRuleConfig() {}

// PermissionRuleConfig is a union: PermissionActionConfig | PermissionObjectConfig.
// Spec: openapi.json:10971
type PermissionRuleConfig interface {
	implementsPermissionRuleConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PermissionRuleConfig)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PermissionActionConfig("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PermissionObjectConfig{}),
		},
	)
}

// PermissionConfigObject is the per-tool permission configuration.
// Spec: PermissionConfig anyOf variant 2 (openapi.json:10981)
type PermissionConfigObject struct {
	Read              PermissionRuleConfig            `json:"read"`
	Edit              PermissionRuleConfig            `json:"edit"`
	Glob              PermissionRuleConfig            `json:"glob"`
	Grep              PermissionRuleConfig            `json:"grep"`
	List              PermissionRuleConfig            `json:"list"`
	Bash              PermissionRuleConfig            `json:"bash"`
	Task              PermissionRuleConfig            `json:"task"`
	ExternalDirectory PermissionRuleConfig            `json:"external_directory"`
	Lsp               PermissionRuleConfig            `json:"lsp"`
	Skill             PermissionRuleConfig            `json:"skill"`
	Todowrite         PermissionActionConfig          `json:"todowrite"`
	Question          PermissionActionConfig          `json:"question"`
	Webfetch          PermissionActionConfig          `json:"webfetch"`
	Websearch         PermissionActionConfig          `json:"websearch"`
	Codesearch        PermissionActionConfig          `json:"codesearch"`
	DoomLoop          PermissionActionConfig          `json:"doom_loop"`
	ExtraFields       map[string]PermissionRuleConfig `json:"-,extras"`
	JSON              permissionConfigObjectJSON      `json:"-"`
}

type permissionConfigObjectJSON struct {
	Read              apijson.Field
	Edit              apijson.Field
	Glob              apijson.Field
	Grep              apijson.Field
	List              apijson.Field
	Bash              apijson.Field
	Task              apijson.Field
	ExternalDirectory apijson.Field
	Lsp               apijson.Field
	Skill             apijson.Field
	Todowrite         apijson.Field
	Question          apijson.Field
	Webfetch          apijson.Field
	Websearch         apijson.Field
	Codesearch        apijson.Field
	DoomLoop          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PermissionConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r PermissionConfigObject) implementsPermissionConfig() {}

// PermissionConfig is a union: PermissionActionConfig | PermissionConfigObject.
// Spec: openapi.json:10981
type PermissionConfig interface {
	implementsPermissionConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PermissionConfig)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(PermissionActionConfig("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PermissionConfigObject{}),
		},
	)
}

type ConfigProvider struct {
	ID        string                         `json:"id"`
	API       string                         `json:"api"`
	Blacklist []string                       `json:"blacklist"`
	Env       []string                       `json:"env"`
	Models    map[string]ConfigProviderModel `json:"models"`
	Name      string                         `json:"name"`
	Npm       string                         `json:"npm"`
	Options   ConfigProviderOptions          `json:"options"`
	Whitelist []string                       `json:"whitelist"`
	JSON      configProviderJSON             `json:"-"`
}

// configProviderJSON contains the JSON metadata for the struct [ConfigProvider]
type configProviderJSON struct {
	ID          apijson.Field
	API         apijson.Field
	Blacklist   apijson.Field
	Env         apijson.Field
	Models      apijson.Field
	Name        apijson.Field
	Npm         apijson.Field
	Options     apijson.Field
	Whitelist   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModel struct {
	ID           string                            `json:"id"`
	Attachment   bool                              `json:"attachment"`
	Cost         ConfigProviderModelsCost          `json:"cost"`
	Experimental bool                              `json:"experimental"`
	Family       string                            `json:"family"`
	Headers      map[string]string                 `json:"headers"`
	Interleaved  ConfigProviderModelInterleaved    `json:"interleaved"`
	Limit        ConfigProviderModelsLimit         `json:"limit"`
	Modalities   ConfigProviderModelsModalities    `json:"modalities"`
	Name         string                            `json:"name"`
	Options      map[string]interface{}            `json:"options"`
	Provider     ConfigProviderModelsProvider      `json:"provider"`
	Reasoning    bool                              `json:"reasoning"`
	ReleaseDate  string                            `json:"release_date"`
	Status       ConfigProviderModelsStatus        `json:"status"`
	Temperature  bool                              `json:"temperature"`
	ToolCall     bool                              `json:"tool_call"`
	Variants     map[string]map[string]interface{} `json:"variants"`
	JSON         configProviderModelJSON           `json:"-"`
}

// configProviderModelJSON contains the JSON metadata for the struct
// [ConfigProviderModel]
type configProviderModelJSON struct {
	ID           apijson.Field
	Attachment   apijson.Field
	Cost         apijson.Field
	Experimental apijson.Field
	Family       apijson.Field
	Headers      apijson.Field
	Interleaved  apijson.Field
	Limit        apijson.Field
	Modalities   apijson.Field
	Name         apijson.Field
	Options      apijson.Field
	Provider     apijson.Field
	Reasoning    apijson.Field
	ReleaseDate  apijson.Field
	Status       apijson.Field
	Temperature  apijson.Field
	ToolCall     apijson.Field
	Variants     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConfigProviderModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelInterleaved struct {
	Field string                             `json:"field"`
	JSON  configProviderModelInterleavedJSON `json:"-"`
	union ConfigProviderModelInterleavedUnion
}

type configProviderModelInterleavedJSON struct {
	Field       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configProviderModelInterleavedJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigProviderModelInterleaved) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigProviderModelInterleaved{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	if object, ok := r.union.(ConfigProviderModelInterleavedObject); ok {
		r.Field = object.Field
		return apijson.Port(object, &r)
	}
	r.JSON.raw = string(data)
	return nil
}

func (r ConfigProviderModelInterleaved) AsUnion() ConfigProviderModelInterleavedUnion {
	return r.union
}

type ConfigProviderModelInterleavedUnion interface {
	implementsConfigProviderModelInterleaved()
}

type ConfigProviderModelInterleavedBool bool

func (ConfigProviderModelInterleavedBool) implementsConfigProviderModelInterleaved() {}

type ConfigProviderModelInterleavedObject struct {
	Field string                                   `json:"field,required"`
	JSON  configProviderModelInterleavedObjectJSON `json:"-"`
}

type configProviderModelInterleavedObjectJSON struct {
	Field       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelInterleavedObject) UnmarshalJSON(data []byte) (err error) {
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	if r.JSON.Field.IsMissing() {
		return fmt.Errorf("missing required field: field")
	}
	if r.Field != "reasoning_content" && r.Field != "reasoning_details" {
		return fmt.Errorf("invalid interleaved field: %s", r.Field)
	}
	return nil
}

func (r configProviderModelInterleavedObjectJSON) RawJSON() string {
	return r.raw
}

func (ConfigProviderModelInterleavedObject) implementsConfigProviderModelInterleaved() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigProviderModelInterleavedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(ConfigProviderModelInterleavedBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigProviderModelInterleavedObject{}),
		},
	)
}

type ConfigProviderModelsCost struct {
	Input           float64                                 `json:"input,required"`
	Output          float64                                 `json:"output,required"`
	CacheRead       float64                                 `json:"cache_read"`
	CacheWrite      float64                                 `json:"cache_write"`
	ContextOver200K ConfigProviderModelsCostContextOver200K `json:"context_over_200k"`
	JSON            configProviderModelsCostJSON            `json:"-"`
}

// configProviderModelsCostJSON contains the JSON metadata for the struct
// [ConfigProviderModelsCost]
type configProviderModelsCostJSON struct {
	Input           apijson.Field
	Output          apijson.Field
	CacheRead       apijson.Field
	CacheWrite      apijson.Field
	ContextOver200K apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigProviderModelsCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsCostJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsCostContextOver200K struct {
	Input      float64                                     `json:"input,required"`
	Output     float64                                     `json:"output,required"`
	CacheRead  float64                                     `json:"cache_read"`
	CacheWrite float64                                     `json:"cache_write"`
	JSON       configProviderModelsCostContextOver200KJSON `json:"-"`
}

type configProviderModelsCostContextOver200KJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	CacheRead   apijson.Field
	CacheWrite  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelsCostContextOver200K) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsCostContextOver200KJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsLimit struct {
	Context float64                       `json:"context,required"`
	Input   float64                       `json:"input"`
	Output  float64                       `json:"output,required"`
	JSON    configProviderModelsLimitJSON `json:"-"`
}

// configProviderModelsLimitJSON contains the JSON metadata for the struct
// [ConfigProviderModelsLimit]
type configProviderModelsLimitJSON struct {
	Context     apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelsLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsLimitJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsModalities struct {
	Input  []ConfigProviderModelsModalitiesInput  `json:"input,required"`
	Output []ConfigProviderModelsModalitiesOutput `json:"output,required"`
	JSON   configProviderModelsModalitiesJSON     `json:"-"`
}

// configProviderModelsModalitiesJSON contains the JSON metadata for the struct
// [ConfigProviderModelsModalities]
type configProviderModelsModalitiesJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelsModalities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsModalitiesJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsModalitiesInput string

const (
	ConfigProviderModelsModalitiesInputText  ConfigProviderModelsModalitiesInput = "text"
	ConfigProviderModelsModalitiesInputAudio ConfigProviderModelsModalitiesInput = "audio"
	ConfigProviderModelsModalitiesInputImage ConfigProviderModelsModalitiesInput = "image"
	ConfigProviderModelsModalitiesInputVideo ConfigProviderModelsModalitiesInput = "video"
	ConfigProviderModelsModalitiesInputPdf   ConfigProviderModelsModalitiesInput = "pdf"
)

func (r ConfigProviderModelsModalitiesInput) IsKnown() bool {
	switch r {
	case ConfigProviderModelsModalitiesInputText, ConfigProviderModelsModalitiesInputAudio, ConfigProviderModelsModalitiesInputImage, ConfigProviderModelsModalitiesInputVideo, ConfigProviderModelsModalitiesInputPdf:
		return true
	}
	return false
}

type ConfigProviderModelsModalitiesOutput string

const (
	ConfigProviderModelsModalitiesOutputText  ConfigProviderModelsModalitiesOutput = "text"
	ConfigProviderModelsModalitiesOutputAudio ConfigProviderModelsModalitiesOutput = "audio"
	ConfigProviderModelsModalitiesOutputImage ConfigProviderModelsModalitiesOutput = "image"
	ConfigProviderModelsModalitiesOutputVideo ConfigProviderModelsModalitiesOutput = "video"
	ConfigProviderModelsModalitiesOutputPdf   ConfigProviderModelsModalitiesOutput = "pdf"
)

func (r ConfigProviderModelsModalitiesOutput) IsKnown() bool {
	switch r {
	case ConfigProviderModelsModalitiesOutputText, ConfigProviderModelsModalitiesOutputAudio, ConfigProviderModelsModalitiesOutputImage, ConfigProviderModelsModalitiesOutputVideo, ConfigProviderModelsModalitiesOutputPdf:
		return true
	}
	return false
}

type ConfigProviderModelsProvider struct {
	Npm  string                           `json:"npm"`
	API  string                           `json:"api"`
	JSON configProviderModelsProviderJSON `json:"-"`
}

// configProviderModelsProviderJSON contains the JSON metadata for the struct
// [ConfigProviderModelsProvider]
type configProviderModelsProviderJSON struct {
	Npm         apijson.Field
	API         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelsProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsProviderJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsStatus string

const (
	ConfigProviderModelsStatusAlpha      ConfigProviderModelsStatus = "alpha"
	ConfigProviderModelsStatusBeta       ConfigProviderModelsStatus = "beta"
	ConfigProviderModelsStatusDeprecated ConfigProviderModelsStatus = "deprecated"
)

func (r ConfigProviderModelsStatus) IsKnown() bool {
	switch r {
	case ConfigProviderModelsStatusAlpha, ConfigProviderModelsStatusBeta, ConfigProviderModelsStatusDeprecated:
		return true
	}
	return false
}

type ConfigProviderOptions struct {
	APIKey  string `json:"apiKey"`
	BaseURL string `json:"baseURL"`
	// Timeout in milliseconds for requests to this provider. Default is 300000 (5
	// minutes). Set to false to disable timeout.
	Timeout       ConfigProviderOptionsTimeoutUnion `json:"timeout"`
	EnterpriseURL string                            `json:"enterpriseUrl"`
	SetCacheKey   bool                              `json:"setCacheKey"`
	ChunkTimeout  int64                             `json:"chunkTimeout"`
	ExtraFields   map[string]interface{}            `json:"-,extras"`
	JSON          configProviderOptionsJSON         `json:"-"`
}

// configProviderOptionsJSON contains the JSON metadata for the struct
// [ConfigProviderOptions]
type configProviderOptionsJSON struct {
	APIKey        apijson.Field
	BaseURL       apijson.Field
	Timeout       apijson.Field
	EnterpriseURL apijson.Field
	SetCacheKey   apijson.Field
	ChunkTimeout  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConfigProviderOptions) UnmarshalJSON(data []byte) (err error) {
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	if !r.JSON.Timeout.IsMissing() && r.Timeout == nil {
		return fmt.Errorf("invalid provider timeout: expected integer or false")
	}
	if !r.JSON.Timeout.IsMissing() {
		timeout := gjson.Parse(r.JSON.Timeout.Raw())
		if timeout.Type == gjson.Number && (timeout.Num != float64(int64(timeout.Num)) || timeout.Num <= 0) {
			return fmt.Errorf("invalid provider timeout: expected integer or false")
		}
	}
	return nil
}

func (r configProviderOptionsJSON) RawJSON() string {
	return r.raw
}

// Timeout in milliseconds for requests to this provider. Default is 300000 (5
// minutes). Set to false to disable timeout.
//
// Union satisfied by [shared.UnionInt] or [ConfigProviderOptionsTimeoutFalse].
type ConfigProviderOptionsTimeoutUnion interface {
	ImplementsConfigProviderOptionsTimeoutUnion()
}

type ConfigProviderOptionsTimeoutFalse bool

func (ConfigProviderOptionsTimeoutFalse) ImplementsConfigProviderOptionsTimeoutUnion() {}

func (r *ConfigProviderOptionsTimeoutFalse) UnmarshalJSON(data []byte) error {
	var value bool
	if err := apijson.Unmarshal(data, &value); err != nil {
		return err
	}
	if value {
		return fmt.Errorf("expected false")
	}
	*r = ConfigProviderOptionsTimeoutFalse(false)
	return nil
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigProviderOptionsTimeoutUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(ConfigProviderOptionsTimeoutFalse(false)),
		},
	)
}

// Control sharing behavior:'manual' allows manual sharing via commands, 'auto'
// enables automatic sharing, 'disabled' disables all sharing
type ConfigShare string

const (
	ConfigShareManual   ConfigShare = "manual"
	ConfigShareAuto     ConfigShare = "auto"
	ConfigShareDisabled ConfigShare = "disabled"
)

func (r ConfigShare) IsKnown() bool {
	switch r {
	case ConfigShareManual, ConfigShareAuto, ConfigShareDisabled:
		return true
	}
	return false
}

// TUI specific settings
type ConfigWatcher struct {
	Ignore []string          `json:"ignore"`
	JSON   configWatcherJSON `json:"-"`
}

// configWatcherJSON contains the JSON metadata for the struct [ConfigWatcher]
type configWatcherJSON struct {
	Ignore      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigWatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configWatcherJSON) RawJSON() string {
	return r.raw
}

type McpLocalConfig struct {
	// Command and arguments to run the MCP server
	Command []string `json:"command,required"`
	// Type of MCP server connection
	Type McpLocalConfigType `json:"type,required"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Environment variables to set when running the MCP server
	Environment map[string]string  `json:"environment"`
	Timeout     float64            `json:"timeout"`
	JSON        mcpLocalConfigJSON `json:"-"`
}

// mcpLocalConfigJSON contains the JSON metadata for the struct [McpLocalConfig]
type mcpLocalConfigJSON struct {
	Command     apijson.Field
	Type        apijson.Field
	Enabled     apijson.Field
	Environment apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpLocalConfig) UnmarshalJSON(data []byte) (err error) {
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	if r.JSON.Command.IsMissing() {
		return fmt.Errorf("missing required field: command")
	}
	return nil
}

func (r mcpLocalConfigJSON) RawJSON() string {
	return r.raw
}

func (r McpLocalConfig) implementsConfigMcp() {}

// Type of MCP server connection
type McpLocalConfigType string

const (
	McpLocalConfigTypeLocal McpLocalConfigType = "local"
)

func (r McpLocalConfigType) IsKnown() bool {
	switch r {
	case McpLocalConfigTypeLocal:
		return true
	}
	return false
}

type McpRemoteConfig struct {
	// Type of MCP server connection
	Type McpRemoteConfigType `json:"type,required"`
	// URL of the remote MCP server
	URL string `json:"url,required"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Headers to send with the request
	Headers map[string]string    `json:"headers"`
	OAuth   McpRemoteConfigOAuth `json:"oauth"`
	Timeout float64              `json:"timeout"`
	JSON    mcpRemoteConfigJSON  `json:"-"`
}

// mcpRemoteConfigJSON contains the JSON metadata for the struct [McpRemoteConfig]
type mcpRemoteConfigJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	Enabled     apijson.Field
	Headers     apijson.Field
	OAuth       apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpRemoteConfig) UnmarshalJSON(data []byte) (err error) {
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	if r.JSON.URL.IsMissing() {
		return fmt.Errorf("missing required field: url")
	}
	return nil
}

func (r mcpRemoteConfigJSON) RawJSON() string {
	return r.raw
}

func (r McpRemoteConfig) implementsConfigMcp() {}

type ConfigMcpDisabled struct {
	Enabled bool                  `json:"enabled,required"`
	JSON    configMcpDisabledJSON `json:"-"`
}

type configMcpDisabledJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigMcpDisabled) UnmarshalJSON(data []byte) (err error) {
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	if r.JSON.Enabled.IsMissing() {
		return fmt.Errorf("missing required field: enabled")
	}
	if len(gjson.ParseBytes(data).Map()) > 1 {
		return fmt.Errorf("invalid disabled MCP config: unexpected extra fields")
	}
	return nil
}

func (r configMcpDisabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigMcpDisabled) implementsConfigMcp() {}

type McpOAuthConfig struct {
	ClientID     string             `json:"clientId"`
	ClientSecret string             `json:"clientSecret"`
	Scope        string             `json:"scope"`
	RedirectURI  string             `json:"redirectUri"`
	JSON         mcpOAuthConfigJSON `json:"-"`
}

type mcpOAuthConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	Scope        apijson.Field
	RedirectURI  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *McpOAuthConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpOAuthConfigJSON) RawJSON() string {
	return r.raw
}

func (McpOAuthConfig) implementsMcpRemoteConfigOAuthUnion() {}

type McpRemoteConfigOAuth interface {
	implementsMcpRemoteConfigOAuthUnion()
}

type McpRemoteConfigOAuthFalse bool

func (McpRemoteConfigOAuthFalse) implementsMcpRemoteConfigOAuthUnion() {}

func (r *McpRemoteConfigOAuthFalse) UnmarshalJSON(data []byte) error {
	var value bool
	if err := apijson.Unmarshal(data, &value); err != nil {
		return err
	}
	if value {
		return fmt.Errorf("expected false")
	}
	*r = McpRemoteConfigOAuthFalse(false)
	return nil
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*McpRemoteConfigOAuth)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McpOAuthConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(McpRemoteConfigOAuthFalse(false)),
		},
	)
}

// Type of MCP server connection
type McpRemoteConfigType string

const (
	McpRemoteConfigTypeRemote McpRemoteConfigType = "remote"
)

func (r McpRemoteConfigType) IsKnown() bool {
	switch r {
	case McpRemoteConfigTypeRemote:
		return true
	}
	return false
}

type ConfigGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ConfigGetParams]'s query parameters as `url.Values`.
func (r ConfigGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigUpdateParams struct {
	Schema            param.Field[string]                       `json:"$schema"`
	Agent             param.Field[map[string]interface{}]       `json:"agent"`
	Autoshare         param.Field[bool]                         `json:"autoshare"`
	Autoupdate        param.Field[interface{}]                  `json:"autoupdate"`
	Command           param.Field[map[string]interface{}]       `json:"command"`
	Compaction        param.Field[ConfigUpdateParamsCompaction] `json:"compaction"`
	DefaultAgent      param.Field[string]                       `json:"default_agent"`
	DisabledProviders param.Field[[]string]                     `json:"disabled_providers"`
	EnabledProviders  param.Field[[]string]                     `json:"enabled_providers"`
	Enterprise        param.Field[ConfigUpdateParamsEnterprise] `json:"enterprise"`
	Experimental      param.Field[interface{}]                  `json:"experimental"`
	Formatter         param.Field[interface{}]                  `json:"formatter"`
	Instructions      param.Field[[]string]                     `json:"instructions"`
	Layout            param.Field[LayoutConfig]                 `json:"layout"`
	LogLevel          param.Field[ConfigLogLevel]               `json:"logLevel"`
	Lsp               param.Field[interface{}]                  `json:"lsp"`
	Mcp               param.Field[interface{}]                  `json:"mcp"`
	Model             param.Field[string]                       `json:"model"`
	Mode              param.Field[map[string]interface{}]       `json:"mode"`
	Permission        param.Field[interface{}]                  `json:"permission"`
	Plugin            param.Field[[]interface{}]                `json:"plugin"`
	Provider          param.Field[map[string]interface{}]       `json:"provider"`
	Server            param.Field[ConfigUpdateParamsServer]     `json:"server"`
	Share             param.Field[ConfigShare]                  `json:"share"`
	Skills            param.Field[ConfigUpdateParamsSkills]     `json:"skills"`
	SmallModel        param.Field[string]                       `json:"small_model"`
	Snapshot          param.Field[bool]                         `json:"snapshot"`
	Tools             param.Field[map[string]bool]              `json:"tools"`
	Username          param.Field[string]                       `json:"username"`
	Watcher           param.Field[ConfigUpdateParamsWatcher]    `json:"watcher"`
	Directory         param.Field[string]                       `query:"directory"`
	Workspace         param.Field[string]                       `query:"workspace"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ConfigUpdateParams]'s query parameters as `url.Values`.
func (r ConfigUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigUpdateParamsCompaction struct {
	Auto                 param.Field[bool]  `json:"auto"`
	Prune                param.Field[bool]  `json:"prune"`
	TailTurns            param.Field[int64] `json:"tail_turns"`
	PreserveRecentTokens param.Field[int64] `json:"preserve_recent_tokens"`
	Reserved             param.Field[int64] `json:"reserved"`
}

func (r ConfigUpdateParamsCompaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsEnterprise struct {
	URL param.Field[string] `json:"url"`
}

func (r ConfigUpdateParamsEnterprise) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsServer struct {
	Port       param.Field[int64]    `json:"port"`
	Hostname   param.Field[string]   `json:"hostname"`
	Mdns       param.Field[bool]     `json:"mdns"`
	MdnsDomain param.Field[string]   `json:"mdnsDomain"`
	Cors       param.Field[[]string] `json:"cors"`
}

func (r ConfigUpdateParamsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsSkills struct {
	Paths param.Field[[]string] `json:"paths"`
	URLs  param.Field[[]string] `json:"urls"`
}

func (r ConfigUpdateParamsSkills) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigUpdateParamsWatcher struct {
	Ignore param.Field[[]string] `json:"ignore"`
}

func (r ConfigUpdateParamsWatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigProvidersParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ConfigProvidersParams]'s query parameters as `url.Values`.
func (r ConfigProvidersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
