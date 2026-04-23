// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
	"github.com/GunsonJack/opencode-sdk-go/internal/apiquery"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
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
	AggregateID string                `json:"aggregate_id,required"`
	Seq         float64               `json:"seq,required"`
	Type        string                `json:"type,required"`
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
	// The optional query-level directory is not modeled separately because it
	// conflicts with the required body field name in the SDK framework.
	Directory param.Field[string]              `json:"directory,required"`
	Events    param.Field[[]SyncReplayEvent]   `json:"events,required"`
	Workspace param.Field[string]              `query:"workspace"`
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
	AggregateID param.Field[string]                `json:"aggregateID,required"`
	Seq         param.Field[int64]                 `json:"seq,required"`
	Type        param.Field[string]                `json:"type,required"`
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
