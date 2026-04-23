// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apiquery"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
)

type ResourceService struct {
	Options []option.RequestOption
}

func NewResourceService(opts ...option.RequestOption) (r *ResourceService) {
	r = &ResourceService{}
	r.Options = opts
	return
}

// List returns all MCP resources.
// Note: McpResource type is defined in mcp.go.
func (r *ResourceService) List(ctx context.Context, query ResourceListParams, opts ...option.RequestOption) (res *map[string]McpResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/resource"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ResourceListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
