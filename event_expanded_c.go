// File auto-generated

package opencode

import (
	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
)

// 1. worktree.ready

type EventListResponseEventWorktreeReady struct {
	Properties EventListResponseEventWorktreeReadyProperties `json:"properties,required"`
	Type       EventListResponseEventWorktreeReadyType       `json:"type,required"`
	JSON       eventListResponseEventWorktreeReadyJSON       `json:"-"`
}

// eventListResponseEventWorktreeReadyJSON contains the JSON metadata for the
// struct [EventListResponseEventWorktreeReady]
type eventListResponseEventWorktreeReadyJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeReady) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeReadyJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorktreeReady) implementsEventListResponse() {}

type EventListResponseEventWorktreeReadyProperties struct {
	Name   string                                             `json:"name,required"`
	Branch string                                             `json:"branch,required"`
	JSON   eventListResponseEventWorktreeReadyPropertiesJSON  `json:"-"`
}

// eventListResponseEventWorktreeReadyPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventWorktreeReadyProperties]
type eventListResponseEventWorktreeReadyPropertiesJSON struct {
	Name        apijson.Field
	Branch      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeReadyProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeReadyPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorktreeReadyType string

const (
	EventListResponseEventWorktreeReadyTypeWorktreeReady EventListResponseEventWorktreeReadyType = "worktree.ready"
)

func (r EventListResponseEventWorktreeReadyType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorktreeReadyTypeWorktreeReady:
		return true
	}
	return false
}

// 2. worktree.failed

type EventListResponseEventWorktreeFailed struct {
	Properties EventListResponseEventWorktreeFailedProperties `json:"properties,required"`
	Type       EventListResponseEventWorktreeFailedType       `json:"type,required"`
	JSON       eventListResponseEventWorktreeFailedJSON       `json:"-"`
}

// eventListResponseEventWorktreeFailedJSON contains the JSON metadata for the
// struct [EventListResponseEventWorktreeFailed]
type eventListResponseEventWorktreeFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorktreeFailed) implementsEventListResponse() {}

type EventListResponseEventWorktreeFailedProperties struct {
	Message string                                              `json:"message,required"`
	JSON    eventListResponseEventWorktreeFailedPropertiesJSON  `json:"-"`
}

// eventListResponseEventWorktreeFailedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventWorktreeFailedProperties]
type eventListResponseEventWorktreeFailedPropertiesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorktreeFailedType string

const (
	EventListResponseEventWorktreeFailedTypeWorktreeFailed EventListResponseEventWorktreeFailedType = "worktree.failed"
)

func (r EventListResponseEventWorktreeFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorktreeFailedTypeWorktreeFailed:
		return true
	}
	return false
}

// 3. pty.created

type EventListResponseEventPtyCreated struct {
	Properties EventListResponseEventPtyCreatedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyCreatedType       `json:"type,required"`
	JSON       eventListResponseEventPtyCreatedJSON       `json:"-"`
}

// eventListResponseEventPtyCreatedJSON contains the JSON metadata for the
// struct [EventListResponseEventPtyCreated]
type eventListResponseEventPtyCreatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyCreatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyCreated) implementsEventListResponse() {}

type EventListResponseEventPtyCreatedProperties struct {
	Info Pty                                             `json:"info,required"`
	JSON eventListResponseEventPtyCreatedPropertiesJSON  `json:"-"`
}

// eventListResponseEventPtyCreatedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventPtyCreatedProperties]
type eventListResponseEventPtyCreatedPropertiesJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyCreatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyCreatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyCreatedType string

const (
	EventListResponseEventPtyCreatedTypePtyCreated EventListResponseEventPtyCreatedType = "pty.created"
)

func (r EventListResponseEventPtyCreatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyCreatedTypePtyCreated:
		return true
	}
	return false
}

// 4. pty.updated

type EventListResponseEventPtyUpdated struct {
	Properties EventListResponseEventPtyUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventPtyUpdatedJSON       `json:"-"`
}

// eventListResponseEventPtyUpdatedJSON contains the JSON metadata for the
// struct [EventListResponseEventPtyUpdated]
type eventListResponseEventPtyUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyUpdated) implementsEventListResponse() {}

type EventListResponseEventPtyUpdatedProperties struct {
	Info Pty                                             `json:"info,required"`
	JSON eventListResponseEventPtyUpdatedPropertiesJSON  `json:"-"`
}

// eventListResponseEventPtyUpdatedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventPtyUpdatedProperties]
type eventListResponseEventPtyUpdatedPropertiesJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyUpdatedType string

const (
	EventListResponseEventPtyUpdatedTypePtyUpdated EventListResponseEventPtyUpdatedType = "pty.updated"
)

func (r EventListResponseEventPtyUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyUpdatedTypePtyUpdated:
		return true
	}
	return false
}

// 5. pty.exited

type EventListResponseEventPtyExited struct {
	Properties EventListResponseEventPtyExitedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyExitedType       `json:"type,required"`
	JSON       eventListResponseEventPtyExitedJSON       `json:"-"`
}

// eventListResponseEventPtyExitedJSON contains the JSON metadata for the
// struct [EventListResponseEventPtyExited]
type eventListResponseEventPtyExitedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyExited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyExitedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyExited) implementsEventListResponse() {}

type EventListResponseEventPtyExitedProperties struct {
	ID       string                                             `json:"id,required"`
	ExitCode float64                                            `json:"exitCode,required"`
	JSON     eventListResponseEventPtyExitedPropertiesJSON      `json:"-"`
}

// eventListResponseEventPtyExitedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventPtyExitedProperties]
type eventListResponseEventPtyExitedPropertiesJSON struct {
	ID          apijson.Field
	ExitCode    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyExitedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyExitedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyExitedType string

const (
	EventListResponseEventPtyExitedTypePtyExited EventListResponseEventPtyExitedType = "pty.exited"
)

func (r EventListResponseEventPtyExitedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyExitedTypePtyExited:
		return true
	}
	return false
}

// 6. pty.deleted

type EventListResponseEventPtyDeleted struct {
	Properties EventListResponseEventPtyDeletedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyDeletedType       `json:"type,required"`
	JSON       eventListResponseEventPtyDeletedJSON       `json:"-"`
}

// eventListResponseEventPtyDeletedJSON contains the JSON metadata for the
// struct [EventListResponseEventPtyDeleted]
type eventListResponseEventPtyDeletedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyDeletedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyDeleted) implementsEventListResponse() {}

type EventListResponseEventPtyDeletedProperties struct {
	ID   string                                          `json:"id,required"`
	JSON eventListResponseEventPtyDeletedPropertiesJSON  `json:"-"`
}

// eventListResponseEventPtyDeletedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventPtyDeletedProperties]
type eventListResponseEventPtyDeletedPropertiesJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyDeletedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyDeletedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyDeletedType string

const (
	EventListResponseEventPtyDeletedTypePtyDeleted EventListResponseEventPtyDeletedType = "pty.deleted"
)

func (r EventListResponseEventPtyDeletedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyDeletedTypePtyDeleted:
		return true
	}
	return false
}

// 7. workspace.ready

type EventListResponseEventWorkspaceReady struct {
	Properties EventListResponseEventWorkspaceReadyProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceReadyType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceReadyJSON       `json:"-"`
}

// eventListResponseEventWorkspaceReadyJSON contains the JSON metadata for the
// struct [EventListResponseEventWorkspaceReady]
type eventListResponseEventWorkspaceReadyJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceReady) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceReadyJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceReady) implementsEventListResponse() {}

type EventListResponseEventWorkspaceReadyProperties struct {
	Name string                                              `json:"name,required"`
	JSON eventListResponseEventWorkspaceReadyPropertiesJSON  `json:"-"`
}

// eventListResponseEventWorkspaceReadyPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventWorkspaceReadyProperties]
type eventListResponseEventWorkspaceReadyPropertiesJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceReadyProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceReadyPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceReadyType string

const (
	EventListResponseEventWorkspaceReadyTypeWorkspaceReady EventListResponseEventWorkspaceReadyType = "workspace.ready"
)

func (r EventListResponseEventWorkspaceReadyType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceReadyTypeWorkspaceReady:
		return true
	}
	return false
}

// 8. workspace.failed

type EventListResponseEventWorkspaceFailed struct {
	Properties EventListResponseEventWorkspaceFailedProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceFailedType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceFailedJSON       `json:"-"`
}

// eventListResponseEventWorkspaceFailedJSON contains the JSON metadata for the
// struct [EventListResponseEventWorkspaceFailed]
type eventListResponseEventWorkspaceFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceFailed) implementsEventListResponse() {}

type EventListResponseEventWorkspaceFailedProperties struct {
	Message string                                               `json:"message,required"`
	JSON    eventListResponseEventWorkspaceFailedPropertiesJSON  `json:"-"`
}

// eventListResponseEventWorkspaceFailedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventWorkspaceFailedProperties]
type eventListResponseEventWorkspaceFailedPropertiesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceFailedType string

const (
	EventListResponseEventWorkspaceFailedTypeWorkspaceFailed EventListResponseEventWorkspaceFailedType = "workspace.failed"
)

func (r EventListResponseEventWorkspaceFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceFailedTypeWorkspaceFailed:
		return true
	}
	return false
}

// 9. workspace.restore

type EventListResponseEventWorkspaceRestore struct {
	Properties EventListResponseEventWorkspaceRestoreProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceRestoreType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceRestoreJSON       `json:"-"`
}

// eventListResponseEventWorkspaceRestoreJSON contains the JSON metadata for the
// struct [EventListResponseEventWorkspaceRestore]
type eventListResponseEventWorkspaceRestoreJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceRestore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceRestoreJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceRestore) implementsEventListResponse() {}

type EventListResponseEventWorkspaceRestoreProperties struct {
	WorkspaceID string                                                `json:"workspaceID,required"`
	SessionID   string                                                `json:"sessionID,required"`
	Total       int64                                                 `json:"total,required"`
	Step        int64                                                 `json:"step,required"`
	JSON        eventListResponseEventWorkspaceRestorePropertiesJSON  `json:"-"`
}

// eventListResponseEventWorkspaceRestorePropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventWorkspaceRestoreProperties]
type eventListResponseEventWorkspaceRestorePropertiesJSON struct {
	WorkspaceID apijson.Field
	SessionID   apijson.Field
	Total       apijson.Field
	Step        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceRestoreProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceRestorePropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceRestoreType string

const (
	EventListResponseEventWorkspaceRestoreTypeWorkspaceRestore EventListResponseEventWorkspaceRestoreType = "workspace.restore"
)

func (r EventListResponseEventWorkspaceRestoreType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceRestoreTypeWorkspaceRestore:
		return true
	}
	return false
}

// 10. workspace.status

type EventListResponseEventWorkspaceStatus struct {
	Properties EventListResponseEventWorkspaceStatusProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceStatusType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceStatusJSON       `json:"-"`
}

// eventListResponseEventWorkspaceStatusJSON contains the JSON metadata for the
// struct [EventListResponseEventWorkspaceStatus]
type eventListResponseEventWorkspaceStatusJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceStatusJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceStatus) implementsEventListResponse() {}

type EventListResponseEventWorkspaceStatusProperties struct {
	WorkspaceID string                                               `json:"workspaceID,required"`
	Status      string                                               `json:"status,required"`
	JSON        eventListResponseEventWorkspaceStatusPropertiesJSON  `json:"-"`
}

// eventListResponseEventWorkspaceStatusPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventWorkspaceStatusProperties]
type eventListResponseEventWorkspaceStatusPropertiesJSON struct {
	WorkspaceID apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceStatusProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceStatusPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceStatusType string

const (
	EventListResponseEventWorkspaceStatusTypeWorkspaceStatus EventListResponseEventWorkspaceStatusType = "workspace.status"
)

func (r EventListResponseEventWorkspaceStatusType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceStatusTypeWorkspaceStatus:
		return true
	}
	return false
}
