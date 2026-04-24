// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
	"github.com/GunsonJack/opencode-sdk-go/internal/apiquery"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
	"github.com/tidwall/gjson"
)

type SyncService struct {
	Options []option.RequestOption
}

func NewSyncService(opts ...option.RequestOption) (r *SyncService) {
	r = &SyncService{}
	r.Options = opts
	return
}

func (r *SyncService) Start(ctx context.Context, params SyncStartParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *SyncService) Replay(ctx context.Context, params SyncReplayParams, opts ...option.RequestOption) (res *SyncReplayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/replay"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *SyncService) History(ctx context.Context, params SyncHistoryParams, opts ...option.RequestOption) (res *[]SyncHistoryEvent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SyncReplayResponse struct {
	SessionID string                 `json:"sessionID,required"`
	JSON      syncReplayResponseJSON `json:"-"`
}

type syncReplayResponseJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncReplayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncReplayResponseJSON) RawJSON() string { return r.raw }

type SyncHistoryEvent struct {
	ID          string                 `json:"id,required"`
	AggregateID string                 `json:"aggregate_id,required"`
	Seq         int64                  `json:"seq,required"`
	Type        string                 `json:"type,required"`
	Data        map[string]interface{} `json:"data,required"`
	JSON        syncHistoryEventJSON   `json:"-"`
}

type syncHistoryEventJSON struct {
	ID          apijson.Field
	AggregateID apijson.Field
	Seq         apijson.Field
	Type        apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncHistoryEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncHistoryEventJSON) RawJSON() string { return r.raw }

type SyncEvent struct {
	Type        SyncEventType `json:"type,required"`
	Name        string        `json:"name,required"`
	ID          string        `json:"id,required"`
	Seq         float64       `json:"seq,required"`
	AggregateID string        `json:"aggregateID,required"`
	Data        interface{}   `json:"data,required"`
	JSON        syncEventJSON `json:"-"`
	union       SyncEventUnion
}

type syncEventJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventJSON) RawJSON() string { return r.raw }

func (r *SyncEvent) UnmarshalJSON(data []byte) (err error) {
	*r = SyncEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r SyncEvent) AsUnion() SyncEventUnion {
	return r.union
}

type SyncEventUnion interface {
	implementsSyncEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SyncEventUnion)(nil)).Elem(),
		"name",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "message.updated.1",
			Type:               reflect.TypeOf(SyncEventMessageUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "message.removed.1",
			Type:               reflect.TypeOf(SyncEventMessageRemoved{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "message.part.updated.1",
			Type:               reflect.TypeOf(SyncEventMessagePartUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "message.part.removed.1",
			Type:               reflect.TypeOf(SyncEventMessagePartRemoved{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.created.1",
			Type:               reflect.TypeOf(SyncEventSessionCreated{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.updated.1",
			Type:               reflect.TypeOf(SyncEventSessionUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.deleted.1",
			Type:               reflect.TypeOf(SyncEventSessionDeleted{}),
		},
	)
}

type SyncEventType string

const (
	SyncEventTypeSync SyncEventType = "sync"
)

func (r SyncEventType) IsKnown() bool {
	switch r {
	case SyncEventTypeSync:
		return true
	}
	return false
}

type SyncEventMessageUpdated struct {
	Type        SyncEventType                                  `json:"type,required"`
	Name        string                                         `json:"name,required"`
	ID          string                                         `json:"id,required"`
	Seq         float64                                        `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventMessageUpdatedProperties `json:"data,required"`
	JSON        syncEventMessageUpdatedJSON                    `json:"-"`
}

type syncEventMessageUpdatedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessageUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessageUpdatedJSON) RawJSON() string { return r.raw }

func (r SyncEventMessageUpdated) implementsSyncEvent() {}

type SyncEventMessageRemoved struct {
	Type        SyncEventType                                  `json:"type,required"`
	Name        string                                         `json:"name,required"`
	ID          string                                         `json:"id,required"`
	Seq         float64                                        `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventMessageRemovedProperties `json:"data,required"`
	JSON        syncEventMessageRemovedJSON                    `json:"-"`
}

type syncEventMessageRemovedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessageRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessageRemovedJSON) RawJSON() string { return r.raw }

func (r SyncEventMessageRemoved) implementsSyncEvent() {}

type SyncEventMessagePartUpdated struct {
	Type        SyncEventType                                      `json:"type,required"`
	Name        string                                             `json:"name,required"`
	ID          string                                             `json:"id,required"`
	Seq         float64                                            `json:"seq,required"`
	AggregateID string                                             `json:"aggregateID,required"`
	Data        EventListResponseEventMessagePartUpdatedProperties `json:"data,required"`
	JSON        syncEventMessagePartUpdatedJSON                    `json:"-"`
}

type syncEventMessagePartUpdatedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessagePartUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessagePartUpdatedJSON) RawJSON() string { return r.raw }

func (r SyncEventMessagePartUpdated) implementsSyncEvent() {}

type SyncEventMessagePartRemoved struct {
	Type        SyncEventType                                      `json:"type,required"`
	Name        string                                             `json:"name,required"`
	ID          string                                             `json:"id,required"`
	Seq         float64                                            `json:"seq,required"`
	AggregateID string                                             `json:"aggregateID,required"`
	Data        EventListResponseEventMessagePartRemovedProperties `json:"data,required"`
	JSON        syncEventMessagePartRemovedJSON                    `json:"-"`
}

type syncEventMessagePartRemovedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessagePartRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessagePartRemovedJSON) RawJSON() string { return r.raw }

func (r SyncEventMessagePartRemoved) implementsSyncEvent() {}

type SyncEventSessionCreated struct {
	Type        SyncEventType                                  `json:"type,required"`
	Name        string                                         `json:"name,required"`
	ID          string                                         `json:"id,required"`
	Seq         float64                                        `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventSessionCreatedProperties `json:"data,required"`
	JSON        syncEventSessionCreatedJSON                    `json:"-"`
}

type syncEventSessionCreatedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionCreatedJSON) RawJSON() string { return r.raw }

func (r SyncEventSessionCreated) implementsSyncEvent() {}

type SyncEventSessionUpdated struct {
	Type        SyncEventType               `json:"type,required"`
	Name        string                      `json:"name,required"`
	ID          string                      `json:"id,required"`
	Seq         float64                     `json:"seq,required"`
	AggregateID string                      `json:"aggregateID,required"`
	Data        SyncEventSessionUpdatedData `json:"data,required"`
	JSON        syncEventSessionUpdatedJSON `json:"-"`
}

type syncEventSessionUpdatedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedJSON) RawJSON() string { return r.raw }

func (r SyncEventSessionUpdated) implementsSyncEvent() {}

type SyncEventSessionUpdatedData struct {
	SessionID string                          `json:"sessionID,required"`
	Info      SyncEventSessionUpdatedDataInfo `json:"info,required"`
	JSON      syncEventSessionUpdatedDataJSON `json:"-"`
}

type syncEventSessionUpdatedDataJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataJSON) RawJSON() string { return r.raw }

type SyncEventSessionUpdatedDataInfo struct {
	ID          *string                                 `json:"id,required"`
	Slug        *string                                 `json:"slug,required"`
	ProjectID   *string                                 `json:"projectID,required"`
	WorkspaceID *string                                 `json:"workspaceID,required"`
	Directory   *string                                 `json:"directory,required"`
	ParentID    *string                                 `json:"parentID,required"`
	Summary     *SyncEventSessionUpdatedDataInfoSummary `json:"summary,required"`
	Share       SyncEventSessionUpdatedDataInfoShare    `json:"share"`
	Title       *string                                 `json:"title,required"`
	Version     *string                                 `json:"version,required"`
	Time        SyncEventSessionUpdatedDataInfoTime     `json:"time"`
	Permission  *[]PermissionRule                       `json:"permission,required"`
	Revert      *SyncEventSessionUpdatedDataInfoRevert  `json:"revert,required"`
	JSON        syncEventSessionUpdatedDataInfoJSON     `json:"-"`
}

type syncEventSessionUpdatedDataInfoJSON struct {
	ID          apijson.Field
	Slug        apijson.Field
	ProjectID   apijson.Field
	WorkspaceID apijson.Field
	Directory   apijson.Field
	ParentID    apijson.Field
	Summary     apijson.Field
	Share       apijson.Field
	Title       apijson.Field
	Version     apijson.Field
	Time        apijson.Field
	Permission  apijson.Field
	Revert      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoJSON) RawJSON() string { return r.raw }

type SyncEventSessionUpdatedDataInfoSummary struct {
	Additions *float64                                   `json:"additions,required"`
	Deletions *float64                                   `json:"deletions,required"`
	Files     *float64                                   `json:"files,required"`
	Diffs     []SnapshotFileDiff                         `json:"diffs"`
	JSON      syncEventSessionUpdatedDataInfoSummaryJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoSummaryJSON struct {
	Additions   apijson.Field
	Deletions   apijson.Field
	Files       apijson.Field
	Diffs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoSummaryJSON) RawJSON() string { return r.raw }

type SyncEventSessionUpdatedDataInfoShare struct {
	URL  *string                                  `json:"url,required"`
	JSON syncEventSessionUpdatedDataInfoShareJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoShareJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoShareJSON) RawJSON() string { return r.raw }

type SyncEventSessionUpdatedDataInfoTime struct {
	Created    *float64                                `json:"created,required"`
	Updated    *float64                                `json:"updated,required"`
	Compacting *float64                                `json:"compacting,required"`
	Archived   *float64                                `json:"archived,required"`
	JSON       syncEventSessionUpdatedDataInfoTimeJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoTimeJSON struct {
	Created     apijson.Field
	Updated     apijson.Field
	Compacting  apijson.Field
	Archived    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoTimeJSON) RawJSON() string { return r.raw }

type SyncEventSessionUpdatedDataInfoRevert struct {
	MessageID *string                                   `json:"messageID,required"`
	PartID    *string                                   `json:"partID"`
	Snapshot  *string                                   `json:"snapshot"`
	Diff      *string                                   `json:"diff"`
	JSON      syncEventSessionUpdatedDataInfoRevertJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoRevertJSON struct {
	MessageID   apijson.Field
	PartID      apijson.Field
	Snapshot    apijson.Field
	Diff        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoRevert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoRevertJSON) RawJSON() string { return r.raw }

type SyncEventSessionDeleted struct {
	Type        SyncEventType                                  `json:"type,required"`
	Name        string                                         `json:"name,required"`
	ID          string                                         `json:"id,required"`
	Seq         float64                                        `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventSessionDeletedProperties `json:"data,required"`
	JSON        syncEventSessionDeletedJSON                    `json:"-"`
}

type syncEventSessionDeletedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionDeletedJSON) RawJSON() string { return r.raw }

func (r SyncEventSessionDeleted) implementsSyncEvent() {}

// Params

type SyncStartParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r SyncStartParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SyncReplayParams struct {
	// Directory is the target directory for the replay (body field).
	Directory param.Field[string]            `json:"directory,required"`
	Events    param.Field[[]SyncReplayEvent] `json:"events,required"`
	// QueryDirectory is the optional project directory context (query parameter).
	QueryDirectory param.Field[string] `query:"directory"`
	Workspace      param.Field[string] `query:"workspace"`
}

func (r SyncReplayParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SyncReplayParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SyncReplayEvent struct {
	ID          param.Field[string]                 `json:"id,required"`
	AggregateID param.Field[string]                 `json:"aggregateID,required"`
	Seq         param.Field[int64]                  `json:"seq,required"`
	Type        param.Field[string]                 `json:"type,required"`
	Data        param.Field[map[string]interface{}] `json:"data,required"`
}

type SyncHistoryParams struct {
	Body      param.Field[map[string]int64] `json:"-"`
	Directory param.Field[string]           `query:"directory"`
	Workspace param.Field[string]           `query:"workspace"`
}

func (r SyncHistoryParams) MarshalJSON() (data []byte, err error) {
	if r.Body.Present {
		return apijson.MarshalRoot(r.Body.Value)
	}
	return []byte("{}"), nil
}

func (r SyncHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
