# Remaining OpenAPI Spec Alignment — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace all remaining `interface{}` fields with properly typed unions and enums to fully align the Go SDK with the OpenAPI spec.

**Architecture:** Each task adds a typed union (response or param) following the SDK's existing patterns: response unions use `apijson.RegisterUnion` + `AsUnion()` method; param unions use marker-method interfaces. Each task is independent and can be committed separately.

**Tech Stack:** Go 1.22, internal `apijson`/`param` packages, `gjson` for union type filtering.

---

## File Map

| File | Changes |
|------|---------|
| `session.go` | Add `OutputFormat` response union, `OutputFormatParam` param union; update `UserMessage.Format`, `SessionPromptParams.Format`, `SessionPromptAsyncParams.Format` |
| `tui.go` | Replace `TuiPublishParams.Body interface{}` with typed union of 4 event param structs; add TUI command enum |
| `provider.go` | Add discriminated union for `ProviderAuthMethodPrompt` (text vs select variants) |
| `mcp.go` | Add `McpAddConfigOAuthUnionParam` that supports both `McpOAuthConfig` object and `false` boolean |

---

### Task 1: Add `OutputFormat` response union types

**Files:**
- Modify: `session.go`

This task adds proper response types for the `OutputFormat` anyOf union (discriminated by `type`: `"text"` | `"json_schema"`) and updates the 3 response fields that currently use `interface{}`.

- [ ] **Step 1: Add the `OutputFormat` response union types**

Add these types after the existing `OutputFormatJsonSchema`/`OutputFormatText` types (if they exist) or near the `UserMessage` type in `session.go`. Follow the SDK's flattened-union-with-AsUnion pattern used by `Message`, `Part`, etc.

```go
// OutputFormat represents the output format configuration.
// Use [OutputFormat.AsUnion] to access the underlying variant.
//
// Union satisfied by [OutputFormatText] or [OutputFormatJsonSchema].
type OutputFormat struct {
	// The format type: "text" or "json_schema".
	Type string          `json:"type,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Schema interface{}   `json:"schema"`
	RetryCount int64     `json:"retryCount"`
	JSON   outputFormatJSON `json:"-"`
	union  OutputFormatUnion
}

// outputFormatJSON contains the JSON metadata for the struct [OutputFormat]
type outputFormatJSON struct {
	Type        apijson.Field
	Schema      apijson.Field
	RetryCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputFormat) UnmarshalJSON(data []byte) (err error) {
	*r = OutputFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r outputFormatJSON) RawJSON() string {
	return r.raw
}

// AsUnion returns the underlying union variant of this OutputFormat.
func (r OutputFormat) AsUnion() OutputFormatUnion {
	return r.union
}

// Union satisfied by [OutputFormatText] or [OutputFormatJsonSchema].
type OutputFormatUnion interface {
	implementsOutputFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OutputFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "text",
			Type:               reflect.TypeOf(OutputFormatText{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "json_schema",
			Type:               reflect.TypeOf(OutputFormatJsonSchema{}),
		},
	)
}

type OutputFormatText struct {
	Type OutputFormatTextType `json:"type,required"`
	JSON outputFormatTextJSON `json:"-"`
}

type outputFormatTextJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputFormatText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputFormatTextJSON) RawJSON() string {
	return r.raw
}

func (r OutputFormatText) implementsOutputFormat() {}

type OutputFormatTextType string

const (
	OutputFormatTextTypeText OutputFormatTextType = "text"
)

func (r OutputFormatTextType) IsKnown() bool {
	switch r {
	case OutputFormatTextTypeText:
		return true
	}
	return false
}

type OutputFormatJsonSchema struct {
	Type       OutputFormatJsonSchemaType `json:"type,required"`
	Schema     map[string]interface{}     `json:"schema,required"`
	RetryCount int64                      `json:"retryCount"`
	JSON       outputFormatJsonSchemaJSON `json:"-"`
}

type outputFormatJsonSchemaJSON struct {
	Type        apijson.Field
	Schema      apijson.Field
	RetryCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputFormatJsonSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputFormatJsonSchemaJSON) RawJSON() string {
	return r.raw
}

func (r OutputFormatJsonSchema) implementsOutputFormat() {}

type OutputFormatJsonSchemaType string

const (
	OutputFormatJsonSchemaTypeJsonSchema OutputFormatJsonSchemaType = "json_schema"
)

func (r OutputFormatJsonSchemaType) IsKnown() bool {
	switch r {
	case OutputFormatJsonSchemaTypeJsonSchema:
		return true
	}
	return false
}
```

- [ ] **Step 2: Update `UserMessage.Format` field**

In `session.go`, change:
```go
Format    interface{}        `json:"format"`
```
to:
```go
Format    OutputFormat       `json:"format"`
```

Also update the corresponding `userMessageJSON` metadata struct if `Format` is listed there (it should already have `Format apijson.Field`).

Remove the comment `// This field can have the runtime type of [interface{}].` above the field if present.

- [ ] **Step 3: Add `OutputFormatParam` union interface for params**

Add param types for the `format` field in request params:

```go
// OutputFormatParam is a param union for OutputFormat.
// Satisfied by [OutputFormatTextParam] or [OutputFormatJsonSchemaParam].
type OutputFormatParam interface {
	implementsOutputFormatParam()
}

type OutputFormatTextParam struct {
	Type param.Field[OutputFormatTextType] `json:"type,required"`
}

func (r OutputFormatTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OutputFormatTextParam) implementsOutputFormatParam() {}

type OutputFormatJsonSchemaParam struct {
	Type       param.Field[OutputFormatJsonSchemaType] `json:"type,required"`
	Schema     param.Field[map[string]interface{}]     `json:"schema,required"`
	RetryCount param.Field[int64]                      `json:"retryCount"`
}

func (r OutputFormatJsonSchemaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OutputFormatJsonSchemaParam) implementsOutputFormatParam() {}
```

- [ ] **Step 4: Update `SessionPromptParams.Format` and `SessionPromptAsyncParams.Format`**

Change both from:
```go
Format    param.Field[interface{}]                    `json:"format"`
```
to:
```go
Format    param.Field[OutputFormatParam]              `json:"format"`
```

- [ ] **Step 5: Also update the flattened `Message` union struct's Format field**

The `Message` response struct (around line 1207) has:
```go
Format interface{} `json:"format"`
```
Change to:
```go
Format OutputFormat `json:"format"`
```

And remove the comment `// This field can have the runtime type of [interface{}].` if present.

- [ ] **Step 6: Verify build**

```bash
cd /home/jack/SemiControlsInnovation/opencode-sdk-go && go build ./...
```

- [ ] **Step 7: Commit**

```bash
git add session.go && git commit -m "feat: add typed OutputFormat union replacing interface{} for format fields"
```

---

### Task 2: Add typed union for `TuiPublishParams.Body`

**Files:**
- Modify: `tui.go`

Replace the untyped `interface{}` body with a proper 4-variant param union discriminated by `type`.

- [ ] **Step 1: Add TUI command enum type**

Add a typed enum for the 16 known TUI commands:

```go
type TuiCommand string

const (
	TuiCommandSessionList        TuiCommand = "session.list"
	TuiCommandSessionNew         TuiCommand = "session.new"
	TuiCommandSessionShare       TuiCommand = "session.share"
	TuiCommandSessionInterrupt   TuiCommand = "session.interrupt"
	TuiCommandSessionCompact     TuiCommand = "session.compact"
	TuiCommandSessionPageUp      TuiCommand = "session.page.up"
	TuiCommandSessionPageDown    TuiCommand = "session.page.down"
	TuiCommandSessionLineUp      TuiCommand = "session.line.up"
	TuiCommandSessionLineDown    TuiCommand = "session.line.down"
	TuiCommandSessionHalfPageUp  TuiCommand = "session.half.page.up"
	TuiCommandSessionHalfPageDown TuiCommand = "session.half.page.down"
	TuiCommandSessionFirst       TuiCommand = "session.first"
	TuiCommandSessionLast        TuiCommand = "session.last"
	TuiCommandPromptClear        TuiCommand = "prompt.clear"
	TuiCommandPromptSubmit       TuiCommand = "prompt.submit"
	TuiCommandAgentCycle         TuiCommand = "agent.cycle"
)

func (r TuiCommand) IsKnown() bool {
	switch r {
	case TuiCommandSessionList, TuiCommandSessionNew, TuiCommandSessionShare,
		TuiCommandSessionInterrupt, TuiCommandSessionCompact,
		TuiCommandSessionPageUp, TuiCommandSessionPageDown,
		TuiCommandSessionLineUp, TuiCommandSessionLineDown,
		TuiCommandSessionHalfPageUp, TuiCommandSessionHalfPageDown,
		TuiCommandSessionFirst, TuiCommandSessionLast,
		TuiCommandPromptClear, TuiCommandPromptSubmit,
		TuiCommandAgentCycle:
		return true
	}
	return false
}
```

- [ ] **Step 2: Add `TuiPublishBody` param union interface and 4 variant structs**

```go
// TuiPublishBody is the union type for TUI publish event bodies.
// Satisfied by [TuiPublishBodyPromptAppend], [TuiPublishBodyCommandExecute],
// [TuiPublishBodyToastShow], or [TuiPublishBodySessionSelect].
type TuiPublishBody interface {
	implementsTuiPublishBody()
}

type TuiPublishBodyPromptAppend struct {
	Type       param.Field[string]                                `json:"type,required"` // always "tui.prompt.append"
	Properties param.Field[TuiPublishBodyPromptAppendProperties]  `json:"properties,required"`
}

func (r TuiPublishBodyPromptAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyPromptAppend) implementsTuiPublishBody() {}

type TuiPublishBodyPromptAppendProperties struct {
	Text param.Field[string] `json:"text,required"`
}

func (r TuiPublishBodyPromptAppendProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodyCommandExecute struct {
	Type       param.Field[string]                                  `json:"type,required"` // always "tui.command.execute"
	Properties param.Field[TuiPublishBodyCommandExecuteProperties]  `json:"properties,required"`
}

func (r TuiPublishBodyCommandExecute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyCommandExecute) implementsTuiPublishBody() {}

type TuiPublishBodyCommandExecuteProperties struct {
	Command param.Field[TuiCommand] `json:"command,required"`
}

func (r TuiPublishBodyCommandExecuteProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodyToastShow struct {
	Type       param.Field[string]                             `json:"type,required"` // always "tui.toast.show"
	Properties param.Field[TuiPublishBodyToastShowProperties]  `json:"properties,required"`
}

func (r TuiPublishBodyToastShow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyToastShow) implementsTuiPublishBody() {}

type TuiPublishBodyToastShowProperties struct {
	Message  param.Field[string]                    `json:"message,required"`
	Variant  param.Field[TuiShowToastParamsVariant] `json:"variant,required"`
	Title    param.Field[string]                    `json:"title"`
	Duration param.Field[float64]                   `json:"duration"`
}

func (r TuiPublishBodyToastShowProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodySessionSelect struct {
	Type       param.Field[string]                                  `json:"type,required"` // always "tui.session.select"
	Properties param.Field[TuiPublishBodySessionSelectProperties]   `json:"properties,required"`
}

func (r TuiPublishBodySessionSelect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodySessionSelect) implementsTuiPublishBody() {}

type TuiPublishBodySessionSelectProperties struct {
	SessionID param.Field[string] `json:"sessionID,required"`
}

func (r TuiPublishBodySessionSelectProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
```

- [ ] **Step 3: Update `TuiPublishParams` to use the union**

Change:
```go
type TuiPublishParams struct {
	Body      param.Field[interface{}] `json:"body"`
	Directory param.Field[string]      `query:"directory"`
	Workspace param.Field[string]      `query:"workspace"`
}
```
to:
```go
type TuiPublishParams struct {
	Body      param.Field[TuiPublishBody] `json:"body"`
	Directory param.Field[string]          `query:"directory"`
	Workspace param.Field[string]          `query:"workspace"`
}
```

**Important:** The existing `MarshalJSON` method on `TuiPublishParams` calls `apijson.MarshalRoot(r.Body)` to serialize the body value as the top-level JSON payload (not wrapped in a `"body"` key). Verify this still works correctly with the union interface type — `apijson.MarshalRoot` must handle `param.Field[TuiPublishBody]` by unwrapping the `param.Field` and then calling `MarshalJSON()` on the concrete union variant. If `go build` succeeds but serialization is wrong, add a manual test to confirm the JSON output matches expected shape.

- [ ] **Step 4: Verify build**

```bash
cd /home/jack/SemiControlsInnovation/opencode-sdk-go && go build ./...
```

- [ ] **Step 5: Commit**

```bash
git add tui.go && git commit -m "feat: add typed TuiPublishBody union with 4 event variants and TuiCommand enum"
```

---

### Task 3: Add discriminated union for `ProviderAuthMethodPrompt`

**Files:**
- Modify: `provider.go`

Replace the flattened `ProviderAuthMethodPrompt` struct with a proper response union discriminated by `type` (`"text"` vs `"select"`).

- [ ] **Step 1: Replace `ProviderAuthMethodPrompt` with a discriminated union**

Replace the existing `ProviderAuthMethodPrompt` struct (lines 140-148) with a union pattern:

```go
// ProviderAuthMethodPrompt is a union type for auth prompts.
// Use [ProviderAuthMethodPrompt.AsUnion] to access the underlying variant.
//
// Union satisfied by [ProviderAuthMethodPromptText] or [ProviderAuthMethodPromptSelect].
type ProviderAuthMethodPrompt struct {
	Type    string                           `json:"type,required"`
	Key     string                           `json:"key,required"`
	Message string                           `json:"message,required"`
	// This field can have the runtime type of [[]ProviderAuthMethodPromptSelectOption].
	Options interface{}                      `json:"options"`
	Placeholder string                       `json:"placeholder"`
	When    *ProviderAuthMethodPromptWhen    `json:"when"`
	JSON    providerAuthMethodPromptJSON     `json:"-"`
	union   ProviderAuthMethodPromptUnion
}

type providerAuthMethodPromptJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Options     apijson.Field
	Placeholder apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPrompt) UnmarshalJSON(data []byte) (err error) {
	*r = ProviderAuthMethodPrompt{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r providerAuthMethodPromptJSON) RawJSON() string {
	return r.raw
}

// AsUnion returns the underlying union variant.
func (r ProviderAuthMethodPrompt) AsUnion() ProviderAuthMethodPromptUnion {
	return r.union
}

// Union satisfied by [ProviderAuthMethodPromptText] or [ProviderAuthMethodPromptSelect].
type ProviderAuthMethodPromptUnion interface {
	implementsProviderAuthMethodPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProviderAuthMethodPromptUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "text",
			Type:               reflect.TypeOf(ProviderAuthMethodPromptText{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "select",
			Type:               reflect.TypeOf(ProviderAuthMethodPromptSelect{}),
		},
	)
}
```

- [ ] **Step 2: Add the text and select variant types**

```go
type ProviderAuthMethodPromptText struct {
	Type        ProviderAuthMethodPromptTextType `json:"type,required"`
	Key         string                           `json:"key,required"`
	Message     string                           `json:"message,required"`
	Placeholder string                           `json:"placeholder"`
	When        *ProviderAuthMethodPromptWhen    `json:"when"`
	JSON        providerAuthMethodPromptTextJSON `json:"-"`
}

type providerAuthMethodPromptTextJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Placeholder apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptTextJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthMethodPromptText) implementsProviderAuthMethodPrompt() {}

type ProviderAuthMethodPromptTextType string

const (
	ProviderAuthMethodPromptTextTypeText ProviderAuthMethodPromptTextType = "text"
)

func (r ProviderAuthMethodPromptTextType) IsKnown() bool {
	switch r {
	case ProviderAuthMethodPromptTextTypeText:
		return true
	}
	return false
}

type ProviderAuthMethodPromptSelect struct {
	Type    ProviderAuthMethodPromptSelectType     `json:"type,required"`
	Key     string                                 `json:"key,required"`
	Message string                                 `json:"message,required"`
	Options []ProviderAuthMethodPromptSelectOption  `json:"options,required"`
	When    *ProviderAuthMethodPromptWhen           `json:"when"`
	JSON    providerAuthMethodPromptSelectJSON      `json:"-"`
}

type providerAuthMethodPromptSelectJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Options     apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptSelect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptSelectJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthMethodPromptSelect) implementsProviderAuthMethodPrompt() {}

type ProviderAuthMethodPromptSelectType string

const (
	ProviderAuthMethodPromptSelectTypeSelect ProviderAuthMethodPromptSelectType = "select"
)

func (r ProviderAuthMethodPromptSelectType) IsKnown() bool {
	switch r {
	case ProviderAuthMethodPromptSelectTypeSelect:
		return true
	}
	return false
}

type ProviderAuthMethodPromptSelectOption struct {
	Label string                                   `json:"label,required"`
	Value string                                   `json:"value,required"`
	Hint  string                                   `json:"hint"`
	JSON  providerAuthMethodPromptSelectOptionJSON `json:"-"`
}

type providerAuthMethodPromptSelectOptionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	Hint        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptSelectOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptSelectOptionJSON) RawJSON() string {
	return r.raw
}
```

- [ ] **Step 3: Remove the old `ProviderAuthMethodPromptOption` type and its associated types**

Remove ALL of the following from `provider.go` (they are replaced by `ProviderAuthMethodPromptSelectOption` and the new union types):
- `ProviderAuthMethodPromptOption` struct (lines ~171-176)
- `providerAuthMethodPromptOptionJSON` metadata struct (lines ~178-185)
- `ProviderAuthMethodPromptOption.UnmarshalJSON` method
- `providerAuthMethodPromptOptionJSON.RawJSON` method

Also remove the old `providerAuthMethodPromptJSON` metadata struct and the old `ProviderAuthMethodPrompt.UnmarshalJSON`/`RawJSON` methods since Step 1 replaces them entirely.

Verify no external references to `ProviderAuthMethodPromptOption` remain (search the codebase). If any exist, update them to `ProviderAuthMethodPromptSelectOption`.

- [ ] **Step 4: Verify build**

```bash
cd /home/jack/SemiControlsInnovation/opencode-sdk-go && go build ./...
```

- [ ] **Step 5: Commit**

```bash
git add provider.go && git commit -m "feat: add discriminated union for ProviderAuthMethodPrompt (text vs select)"
```

---

### Task 4: Support `false` boolean variant for `McpAddConfigParam.OAuth`

**Files:**
- Modify: `mcp.go`

The spec defines `McpRemoteConfig.oauth` as `anyOf: [McpOAuthConfig, {type: boolean, const: false}]`. Currently the SDK only supports the object variant. Add a param union that also accepts `false`.

- [ ] **Step 1: Add `McpAddConfigOAuthUnionParam` interface**

```go
// McpAddConfigOAuthUnionParam is a param union for the MCP OAuth config.
// Satisfied by [McpAddConfigOAuthParam] (object config) or [McpAddConfigOAuthDisabledParam] (false).
type McpAddConfigOAuthUnionParam interface {
	implementsMcpAddConfigOAuthUnionParam()
}

// McpAddConfigOAuthDisabledParam represents oauth: false to disable OAuth auto-detection.
type McpAddConfigOAuthDisabledParam struct{}

func (r McpAddConfigOAuthDisabledParam) MarshalJSON() (data []byte, err error) {
	return []byte("false"), nil
}

func (r McpAddConfigOAuthDisabledParam) implementsMcpAddConfigOAuthUnionParam() {}
```

- [ ] **Step 2: Add the interface marker to the existing `McpAddConfigOAuthParam`**

Add this method to the existing `McpAddConfigOAuthParam` struct:

```go
func (r McpAddConfigOAuthParam) implementsMcpAddConfigOAuthUnionParam() {}
```

- [ ] **Step 3: Update `McpAddConfigParam.OAuth` field type**

Change:
```go
OAuth   param.Field[McpAddConfigOAuthParam] `json:"oauth"`
```
to:
```go
OAuth   param.Field[McpAddConfigOAuthUnionParam] `json:"oauth"`
```

- [ ] **Step 4: Verify build**

```bash
cd /home/jack/SemiControlsInnovation/opencode-sdk-go && go build ./...
```

- [ ] **Step 5: Commit**

```bash
git add mcp.go && git commit -m "feat: support oauth: false for McpRemoteConfig to disable OAuth auto-detection"
```
