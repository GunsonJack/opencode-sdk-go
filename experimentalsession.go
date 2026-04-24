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

// ExperimentalSessionService contains methods for interacting with the
// experimental session resource.
type ExperimentalSessionService struct {
	Options []option.RequestOption
}

// NewExperimentalSessionService generates a new service that applies the given
// options to each request.
func NewExperimentalSessionService(opts ...option.RequestOption) (r *ExperimentalSessionService) {
	r = &ExperimentalSessionService{}
	r.Options = opts
	return
}

// List all sessions across projects.
func (r *ExperimentalSessionService) List(ctx context.Context, query ExperimentalSessionListParams, opts ...option.RequestOption) (res *[]GlobalSession, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ExperimentalSessionListParams struct {
	Directory param.Field[string]  `query:"directory"`
	Workspace param.Field[string]  `query:"workspace"`
	Roots     param.Field[bool]    `query:"roots"`
	Start     param.Field[float64] `query:"start"`
	Cursor    param.Field[float64] `query:"cursor"`
	Search    param.Field[string]  `query:"search"`
	Limit     param.Field[float64] `query:"limit"`
	Archived  param.Field[bool]    `query:"archived"`
}

func (r ExperimentalSessionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
