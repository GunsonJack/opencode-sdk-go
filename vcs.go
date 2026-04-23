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

type VcsService struct {
	Options []option.RequestOption
}

func NewVcsService(opts ...option.RequestOption) (r *VcsService) {
	r = &VcsService{}
	r.Options = opts
	return
}

func (r *VcsService) Get(ctx context.Context, query VcsGetParams, opts ...option.RequestOption) (res *VcsInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *VcsService) Diff(ctx context.Context, query VcsDiffParams, opts ...option.RequestOption) (res *[]VcsFileDiff, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vcs/diff"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VcsInfo struct {
	Branch        string      `json:"branch"`
	DefaultBranch string      `json:"default_branch"`
	JSON          vcsInfoJSON `json:"-"`
}

type vcsInfoJSON struct {
	Branch        apijson.Field
	DefaultBranch apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VcsInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vcsInfoJSON) RawJSON() string { return r.raw }

type VcsFileDiff struct {
	File      string              `json:"file,required"`
	Patch     string              `json:"patch,required"`
	Additions float64             `json:"additions,required"`
	Deletions float64             `json:"deletions,required"`
	Status    VcsFileDiffStatus   `json:"status"`
	JSON      vcsFileDiffJSON     `json:"-"`
}

type vcsFileDiffJSON struct {
	File        apijson.Field
	Patch       apijson.Field
	Additions   apijson.Field
	Deletions   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VcsFileDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vcsFileDiffJSON) RawJSON() string { return r.raw }

type VcsFileDiffStatus string

const (
	VcsFileDiffStatusAdded    VcsFileDiffStatus = "added"
	VcsFileDiffStatusDeleted  VcsFileDiffStatus = "deleted"
	VcsFileDiffStatusModified VcsFileDiffStatus = "modified"
)

func (r VcsFileDiffStatus) IsKnown() bool {
	switch r {
	case VcsFileDiffStatusAdded, VcsFileDiffStatusDeleted, VcsFileDiffStatusModified:
		return true
	}
	return false
}

type VcsGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r VcsGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VcsDiffParams struct {
	Mode      param.Field[VcsDiffParamsMode] `query:"mode,required"`
	Directory param.Field[string]            `query:"directory"`
	Workspace param.Field[string]            `query:"workspace"`
}

func (r VcsDiffParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VcsDiffParamsMode string

const (
	VcsDiffParamsModeGit    VcsDiffParamsMode = "git"
	VcsDiffParamsModeBranch VcsDiffParamsMode = "branch"
)

func (r VcsDiffParamsMode) IsKnown() bool {
	switch r {
	case VcsDiffParamsModeGit, VcsDiffParamsModeBranch:
		return true
	}
	return false
}
