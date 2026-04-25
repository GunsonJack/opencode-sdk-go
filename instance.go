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

// InstanceService contains methods for interacting with the instance resource.
type InstanceService struct {
	Options []option.RequestOption
}

// NewInstanceService generates a new service that applies the given options to
// each request.
func NewInstanceService(opts ...option.RequestOption) (r *InstanceService) {
	r = &InstanceService{}
	r.Options = opts
	return
}

// Dispose shuts down the current instance.
func (r *InstanceService) Dispose(ctx context.Context, params InstanceDisposeParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "instance/dispose"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type InstanceDisposeParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [InstanceDisposeParams]'s query parameters as `url.Values`.
func (r InstanceDisposeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
