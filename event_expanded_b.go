// File auto-generated

package opencode

import (
	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
)

// 1. question.replied

type EventListResponseEventQuestionReplied struct {
	Properties QuestionReplied                           `json:"properties,required"`
	Type       EventListResponseEventQuestionRepliedType `json:"type,required"`
	JSON       eventListResponseEventQuestionRepliedJSON `json:"-"`
}

// eventListResponseEventQuestionRepliedJSON contains the JSON metadata for the
// struct [EventListResponseEventQuestionReplied]
type eventListResponseEventQuestionRepliedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionReplied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRepliedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionReplied) implementsEventListResponse() {}

type EventListResponseEventQuestionRepliedType string

const (
	EventListResponseEventQuestionRepliedTypeQuestionReplied EventListResponseEventQuestionRepliedType = "question.replied"
)

func (r EventListResponseEventQuestionRepliedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionRepliedTypeQuestionReplied:
		return true
	}
	return false
}

// 2. question.rejected

type EventListResponseEventQuestionRejected struct {
	Properties QuestionRejected                           `json:"properties,required"`
	Type       EventListResponseEventQuestionRejectedType `json:"type,required"`
	JSON       eventListResponseEventQuestionRejectedJSON `json:"-"`
}

// eventListResponseEventQuestionRejectedJSON contains the JSON metadata for the
// struct [EventListResponseEventQuestionRejected]
type eventListResponseEventQuestionRejectedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionRejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRejectedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionRejected) implementsEventListResponse() {}

type EventListResponseEventQuestionRejectedType string

const (
	EventListResponseEventQuestionRejectedTypeQuestionRejected EventListResponseEventQuestionRejectedType = "question.rejected"
)

func (r EventListResponseEventQuestionRejectedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionRejectedTypeQuestionRejected:
		return true
	}
	return false
}

// 3. tui.prompt.append

type EventListResponseEventTuiPromptAppend struct {
	Properties EventListResponseEventTuiPromptAppendProperties `json:"properties,required"`
	Type       EventListResponseEventTuiPromptAppendType       `json:"type,required"`
	JSON       eventListResponseEventTuiPromptAppendJSON       `json:"-"`
}

// eventListResponseEventTuiPromptAppendJSON contains the JSON metadata for the
// struct [EventListResponseEventTuiPromptAppend]
type eventListResponseEventTuiPromptAppendJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiPromptAppend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiPromptAppendJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiPromptAppend) implementsEventListResponse() {}

type EventListResponseEventTuiPromptAppendProperties struct {
	Text string                                              `json:"text,required"`
	JSON eventListResponseEventTuiPromptAppendPropertiesJSON `json:"-"`
}

// eventListResponseEventTuiPromptAppendPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventTuiPromptAppendProperties]
type eventListResponseEventTuiPromptAppendPropertiesJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiPromptAppendProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiPromptAppendPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiPromptAppendType string

const (
	EventListResponseEventTuiPromptAppendTypeTuiPromptAppend EventListResponseEventTuiPromptAppendType = "tui.prompt.append"
)

func (r EventListResponseEventTuiPromptAppendType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiPromptAppendTypeTuiPromptAppend:
		return true
	}
	return false
}

// 4. tui.command.execute

type EventListResponseEventTuiCommandExecute struct {
	Properties EventListResponseEventTuiCommandExecuteProperties `json:"properties,required"`
	Type       EventListResponseEventTuiCommandExecuteType       `json:"type,required"`
	JSON       eventListResponseEventTuiCommandExecuteJSON       `json:"-"`
}

// eventListResponseEventTuiCommandExecuteJSON contains the JSON metadata for the
// struct [EventListResponseEventTuiCommandExecute]
type eventListResponseEventTuiCommandExecuteJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiCommandExecute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiCommandExecuteJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiCommandExecute) implementsEventListResponse() {}

type EventListResponseEventTuiCommandExecuteProperties struct {
	Command string                                                `json:"command,required"`
	JSON    eventListResponseEventTuiCommandExecutePropertiesJSON `json:"-"`
}

// eventListResponseEventTuiCommandExecutePropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventTuiCommandExecuteProperties]
type eventListResponseEventTuiCommandExecutePropertiesJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiCommandExecuteProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiCommandExecutePropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiCommandExecuteType string

const (
	EventListResponseEventTuiCommandExecuteTypeTuiCommandExecute EventListResponseEventTuiCommandExecuteType = "tui.command.execute"
)

func (r EventListResponseEventTuiCommandExecuteType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiCommandExecuteTypeTuiCommandExecute:
		return true
	}
	return false
}

// 5. tui.toast.show

type EventListResponseEventTuiToastShow struct {
	Properties EventListResponseEventTuiToastShowProperties `json:"properties,required"`
	Type       EventListResponseEventTuiToastShowType       `json:"type,required"`
	JSON       eventListResponseEventTuiToastShowJSON       `json:"-"`
}

// eventListResponseEventTuiToastShowJSON contains the JSON metadata for the struct
// [EventListResponseEventTuiToastShow]
type eventListResponseEventTuiToastShowJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiToastShow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiToastShowJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiToastShow) implementsEventListResponse() {}

type EventListResponseEventTuiToastShowProperties struct {
	Message  string                                              `json:"message,required"`
	Variant  EventListResponseEventTuiToastShowPropertiesVariant `json:"variant,required"`
	Title    string                                              `json:"title"`
	Duration float64                                             `json:"duration"`
	JSON     eventListResponseEventTuiToastShowPropertiesJSON    `json:"-"`
}

// eventListResponseEventTuiToastShowPropertiesJSON contains the JSON metadata for
// the struct [EventListResponseEventTuiToastShowProperties]
type eventListResponseEventTuiToastShowPropertiesJSON struct {
	Message     apijson.Field
	Variant     apijson.Field
	Title       apijson.Field
	Duration    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiToastShowProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiToastShowPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiToastShowPropertiesVariant string

const (
	EventListResponseEventTuiToastShowPropertiesVariantInfo    EventListResponseEventTuiToastShowPropertiesVariant = "info"
	EventListResponseEventTuiToastShowPropertiesVariantSuccess EventListResponseEventTuiToastShowPropertiesVariant = "success"
	EventListResponseEventTuiToastShowPropertiesVariantWarning EventListResponseEventTuiToastShowPropertiesVariant = "warning"
	EventListResponseEventTuiToastShowPropertiesVariantError   EventListResponseEventTuiToastShowPropertiesVariant = "error"
)

func (r EventListResponseEventTuiToastShowPropertiesVariant) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiToastShowPropertiesVariantInfo, EventListResponseEventTuiToastShowPropertiesVariantSuccess, EventListResponseEventTuiToastShowPropertiesVariantWarning, EventListResponseEventTuiToastShowPropertiesVariantError:
		return true
	}
	return false
}

type EventListResponseEventTuiToastShowType string

const (
	EventListResponseEventTuiToastShowTypeTuiToastShow EventListResponseEventTuiToastShowType = "tui.toast.show"
)

func (r EventListResponseEventTuiToastShowType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiToastShowTypeTuiToastShow:
		return true
	}
	return false
}

// 6. tui.session.select

type EventListResponseEventTuiSessionSelect struct {
	Properties EventListResponseEventTuiSessionSelectProperties `json:"properties,required"`
	Type       EventListResponseEventTuiSessionSelectType       `json:"type,required"`
	JSON       eventListResponseEventTuiSessionSelectJSON       `json:"-"`
}

// eventListResponseEventTuiSessionSelectJSON contains the JSON metadata for the
// struct [EventListResponseEventTuiSessionSelect]
type eventListResponseEventTuiSessionSelectJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiSessionSelect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiSessionSelectJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiSessionSelect) implementsEventListResponse() {}

type EventListResponseEventTuiSessionSelectProperties struct {
	SessionID string                                               `json:"sessionID,required"`
	JSON      eventListResponseEventTuiSessionSelectPropertiesJSON `json:"-"`
}

// eventListResponseEventTuiSessionSelectPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventTuiSessionSelectProperties]
type eventListResponseEventTuiSessionSelectPropertiesJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiSessionSelectProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiSessionSelectPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiSessionSelectType string

const (
	EventListResponseEventTuiSessionSelectTypeTuiSessionSelect EventListResponseEventTuiSessionSelectType = "tui.session.select"
)

func (r EventListResponseEventTuiSessionSelectType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiSessionSelectTypeTuiSessionSelect:
		return true
	}
	return false
}

// 7. mcp.tools.changed

type EventListResponseEventMcpToolsChanged struct {
	Properties EventListResponseEventMcpToolsChangedProperties `json:"properties,required"`
	Type       EventListResponseEventMcpToolsChangedType       `json:"type,required"`
	JSON       eventListResponseEventMcpToolsChangedJSON       `json:"-"`
}

// eventListResponseEventMcpToolsChangedJSON contains the JSON metadata for the
// struct [EventListResponseEventMcpToolsChanged]
type eventListResponseEventMcpToolsChangedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpToolsChanged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpToolsChangedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMcpToolsChanged) implementsEventListResponse() {}

type EventListResponseEventMcpToolsChangedProperties struct {
	Server string                                              `json:"server,required"`
	JSON   eventListResponseEventMcpToolsChangedPropertiesJSON `json:"-"`
}

// eventListResponseEventMcpToolsChangedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventMcpToolsChangedProperties]
type eventListResponseEventMcpToolsChangedPropertiesJSON struct {
	Server      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpToolsChangedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpToolsChangedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMcpToolsChangedType string

const (
	EventListResponseEventMcpToolsChangedTypeMcpToolsChanged EventListResponseEventMcpToolsChangedType = "mcp.tools.changed"
)

func (r EventListResponseEventMcpToolsChangedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMcpToolsChangedTypeMcpToolsChanged:
		return true
	}
	return false
}

// 8. mcp.browser.open.failed

type EventListResponseEventMcpBrowserOpenFailed struct {
	Properties EventListResponseEventMcpBrowserOpenFailedProperties `json:"properties,required"`
	Type       EventListResponseEventMcpBrowserOpenFailedType       `json:"type,required"`
	JSON       eventListResponseEventMcpBrowserOpenFailedJSON       `json:"-"`
}

// eventListResponseEventMcpBrowserOpenFailedJSON contains the JSON metadata for
// the struct [EventListResponseEventMcpBrowserOpenFailed]
type eventListResponseEventMcpBrowserOpenFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpBrowserOpenFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpBrowserOpenFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMcpBrowserOpenFailed) implementsEventListResponse() {}

type EventListResponseEventMcpBrowserOpenFailedProperties struct {
	McpName string                                                   `json:"mcpName,required"`
	URL     string                                                   `json:"url,required"`
	JSON    eventListResponseEventMcpBrowserOpenFailedPropertiesJSON `json:"-"`
}

// eventListResponseEventMcpBrowserOpenFailedPropertiesJSON contains the JSON
// metadata for the struct
// [EventListResponseEventMcpBrowserOpenFailedProperties]
type eventListResponseEventMcpBrowserOpenFailedPropertiesJSON struct {
	McpName     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpBrowserOpenFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpBrowserOpenFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMcpBrowserOpenFailedType string

const (
	EventListResponseEventMcpBrowserOpenFailedTypeMcpBrowserOpenFailed EventListResponseEventMcpBrowserOpenFailedType = "mcp.browser.open.failed"
)

func (r EventListResponseEventMcpBrowserOpenFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMcpBrowserOpenFailedTypeMcpBrowserOpenFailed:
		return true
	}
	return false
}

// 9. command.executed

type EventListResponseEventCommandExecuted struct {
	Properties EventListResponseEventCommandExecutedProperties `json:"properties,required"`
	Type       EventListResponseEventCommandExecutedType       `json:"type,required"`
	JSON       eventListResponseEventCommandExecutedJSON       `json:"-"`
}

// eventListResponseEventCommandExecutedJSON contains the JSON metadata for the
// struct [EventListResponseEventCommandExecuted]
type eventListResponseEventCommandExecutedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCommandExecuted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCommandExecutedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventCommandExecuted) implementsEventListResponse() {}

type EventListResponseEventCommandExecutedProperties struct {
	Name      string                                              `json:"name,required"`
	SessionID string                                              `json:"sessionID,required"`
	Arguments string                                              `json:"arguments,required"`
	MessageID string                                              `json:"messageID,required"`
	JSON      eventListResponseEventCommandExecutedPropertiesJSON `json:"-"`
}

// eventListResponseEventCommandExecutedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventCommandExecutedProperties]
type eventListResponseEventCommandExecutedPropertiesJSON struct {
	Name        apijson.Field
	SessionID   apijson.Field
	Arguments   apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCommandExecutedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCommandExecutedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventCommandExecutedType string

const (
	EventListResponseEventCommandExecutedTypeCommandExecuted EventListResponseEventCommandExecutedType = "command.executed"
)

func (r EventListResponseEventCommandExecutedType) IsKnown() bool {
	switch r {
	case EventListResponseEventCommandExecutedTypeCommandExecuted:
		return true
	}
	return false
}

// 10. vcs.branch.updated

type EventListResponseEventVcsBranchUpdated struct {
	Properties EventListResponseEventVcsBranchUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventVcsBranchUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventVcsBranchUpdatedJSON       `json:"-"`
}

// eventListResponseEventVcsBranchUpdatedJSON contains the JSON metadata for the
// struct [EventListResponseEventVcsBranchUpdated]
type eventListResponseEventVcsBranchUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventVcsBranchUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventVcsBranchUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventVcsBranchUpdated) implementsEventListResponse() {}

type EventListResponseEventVcsBranchUpdatedProperties struct {
	Branch string                                               `json:"branch"`
	JSON   eventListResponseEventVcsBranchUpdatedPropertiesJSON `json:"-"`
}

// eventListResponseEventVcsBranchUpdatedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventVcsBranchUpdatedProperties]
type eventListResponseEventVcsBranchUpdatedPropertiesJSON struct {
	Branch      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventVcsBranchUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventVcsBranchUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventVcsBranchUpdatedType string

const (
	EventListResponseEventVcsBranchUpdatedTypeVcsBranchUpdated EventListResponseEventVcsBranchUpdatedType = "vcs.branch.updated"
)

func (r EventListResponseEventVcsBranchUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventVcsBranchUpdatedTypeVcsBranchUpdated:
		return true
	}
	return false
}
