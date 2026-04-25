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

type ToolService struct {
	Options []option.RequestOption
}

func NewToolService(opts ...option.RequestOption) (r *ToolService) {
	r = &ToolService{}
	r.Options = opts
	return
}

func (r *ToolService) List(ctx context.Context, query ToolListParams, opts ...option.RequestOption) (res *[]ToolListItem, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/tool"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *ToolService) IDs(ctx context.Context, query ToolIDsParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/tool/ids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ToolListItem struct {
	ID          string           `json:"id,required"`
	Description string           `json:"description,required"`
	Parameters  interface{}      `json:"parameters,required"`
	JSON        toolListItemJSON `json:"-"`
}

type toolListItemJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Parameters  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListItemJSON) RawJSON() string { return r.raw }

type ToolListParams struct {
	Provider  param.Field[string] `query:"provider,required"`
	Model     param.Field[string] `query:"model,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ToolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolIDsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ToolIDsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
