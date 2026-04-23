// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
	"github.com/GunsonJack/opencode-sdk-go/internal/apiquery"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
)

type WorkspaceService struct {
	Options []option.RequestOption
}

func NewWorkspaceService(opts ...option.RequestOption) (r *WorkspaceService) {
	r = &WorkspaceService{}
	r.Options = opts
	return
}

func (r *WorkspaceService) Create(ctx context.Context, params WorkspaceCreateParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *WorkspaceService) List(ctx context.Context, query WorkspaceListParams, opts ...option.RequestOption) (res *[]Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *WorkspaceService) Status(ctx context.Context, query WorkspaceStatusParams, opts ...option.RequestOption) (res *[]WorkspaceStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *WorkspaceService) Adaptors(ctx context.Context, query WorkspaceAdaptorsParams, opts ...option.RequestOption) (res *[]WorkspaceAdaptor, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/adaptor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *WorkspaceService) Remove(ctx context.Context, id string, params WorkspaceRemoveParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("experimental/workspace/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

func (r *WorkspaceService) SessionRestore(ctx context.Context, id string, params WorkspaceSessionRestoreParams, opts ...option.RequestOption) (res *WorkspaceSessionRestoreResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("experimental/workspace/%s/session-restore", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Workspace struct {
	ID        string        `json:"id,required"`
	Type      string        `json:"type,required"`
	Name      string        `json:"name,required"`
	Branch    *string       `json:"branch,required"`
	Directory *string       `json:"directory,required"`
	Extra     interface{}   `json:"extra,required"`
	ProjectID string        `json:"projectID,required"`
	JSON      workspaceJSON `json:"-"`
}

type workspaceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Branch      apijson.Field
	Directory   apijson.Field
	Extra       apijson.Field
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Workspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceJSON) RawJSON() string { return r.raw }

type WorkspaceStatusResponse struct {
	WorkspaceID string                        `json:"workspaceID,required"`
	Status      WorkspaceStatusResponseStatus `json:"status,required"`
	JSON        workspaceStatusResponseJSON   `json:"-"`
}

type workspaceStatusResponseJSON struct {
	WorkspaceID apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceStatusResponseJSON) RawJSON() string { return r.raw }

type WorkspaceStatusResponseStatus string

const (
	WorkspaceStatusResponseStatusConnected    WorkspaceStatusResponseStatus = "connected"
	WorkspaceStatusResponseStatusConnecting   WorkspaceStatusResponseStatus = "connecting"
	WorkspaceStatusResponseStatusDisconnected WorkspaceStatusResponseStatus = "disconnected"
	WorkspaceStatusResponseStatusError        WorkspaceStatusResponseStatus = "error"
)

func (r WorkspaceStatusResponseStatus) IsKnown() bool {
	switch r {
	case WorkspaceStatusResponseStatusConnected, WorkspaceStatusResponseStatusConnecting, WorkspaceStatusResponseStatusDisconnected, WorkspaceStatusResponseStatusError:
		return true
	}
	return false
}

type WorkspaceAdaptor struct {
	Type        string               `json:"type,required"`
	Name        string               `json:"name,required"`
	Description string               `json:"description,required"`
	JSON        workspaceAdaptorJSON `json:"-"`
}

type workspaceAdaptorJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceAdaptor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceAdaptorJSON) RawJSON() string { return r.raw }

type WorkspaceSessionRestoreResponse struct {
	Total int64                               `json:"total,required"`
	JSON  workspaceSessionRestoreResponseJSON `json:"-"`
}

type workspaceSessionRestoreResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceSessionRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceSessionRestoreResponseJSON) RawJSON() string { return r.raw }

// Params

type WorkspaceCreateParams struct {
	Type      param.Field[string]      `json:"type,required"`
	Branch    param.Field[string]      `json:"branch,required"`
	Extra     param.Field[interface{}] `json:"extra,required"`
	ID        param.Field[string]      `json:"id"`
	Directory param.Field[string]      `query:"directory"`
	Workspace param.Field[string]      `query:"workspace"`
}

func (r WorkspaceCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkspaceCreateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorkspaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorkspaceStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceAdaptorsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorkspaceAdaptorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceRemoveParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorkspaceRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceSessionRestoreParams struct {
	SessionID param.Field[string] `json:"sessionID,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorkspaceSessionRestoreParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkspaceSessionRestoreParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
