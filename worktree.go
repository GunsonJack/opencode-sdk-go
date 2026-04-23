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

type WorktreeService struct {
	Options []option.RequestOption
}

func NewWorktreeService(opts ...option.RequestOption) (r *WorktreeService) {
	r = &WorktreeService{}
	r.Options = opts
	return
}

func (r *WorktreeService) Create(ctx context.Context, params WorktreeCreateParams, opts ...option.RequestOption) (res *Worktree, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/worktree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

func (r *WorktreeService) List(ctx context.Context, query WorktreeListParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/worktree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *WorktreeService) Remove(ctx context.Context, params WorktreeRemoveParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/worktree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

func (r *WorktreeService) Reset(ctx context.Context, params WorktreeResetParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/worktree/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Worktree struct {
	Name      string       `json:"name,required"`
	Branch    string       `json:"branch,required"`
	Directory string       `json:"directory,required"`
	JSON      worktreeJSON `json:"-"`
}

type worktreeJSON struct {
	Name        apijson.Field
	Branch      apijson.Field
	Directory   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Worktree) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r worktreeJSON) RawJSON() string { return r.raw }

type WorktreeCreateParams struct {
	Name         param.Field[string] `json:"name"`
	StartCommand param.Field[string] `json:"startCommand"`
	Directory    param.Field[string] `query:"directory"`
	Workspace    param.Field[string] `query:"workspace"`
}

func (r WorktreeCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorktreeCreateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorktreeListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorktreeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorktreeRemoveParams struct {
	// The optional query-level directory is not modeled separately because it
	// conflicts with the required body field name in the SDK framework.
	Directory param.Field[string] `json:"directory,required"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorktreeRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorktreeRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorktreeResetParams struct {
	// The optional query-level directory is not modeled separately because it
	// conflicts with the required body field name in the SDK framework.
	Directory param.Field[string] `json:"directory,required"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorktreeResetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorktreeResetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
