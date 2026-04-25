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

// FormatterService contains methods for interacting with the formatter resource.
type FormatterService struct {
	Options []option.RequestOption
}

// NewFormatterService generates a new service that applies the given options to
// each request.
func NewFormatterService(opts ...option.RequestOption) (r *FormatterService) {
	r = &FormatterService{}
	r.Options = opts
	return
}

// Status returns the status of all formatters.
func (r *FormatterService) Status(ctx context.Context, query FormatterStatusParams, opts ...option.RequestOption) (res *[]FormatterStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "formatter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// FormatterStatus describes the status of a single formatter.
type FormatterStatus struct {
	Name       string              `json:"name,required"`
	Extensions []string            `json:"extensions,required"`
	Enabled    bool                `json:"enabled,required"`
	JSON       formatterStatusJSON `json:"-"`
}

// formatterStatusJSON contains the JSON metadata for the struct [FormatterStatus]
type formatterStatusJSON struct {
	Name        apijson.Field
	Extensions  apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FormatterStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r formatterStatusJSON) RawJSON() string {
	return r.raw
}

type FormatterStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [FormatterStatusParams]'s query parameters as `url.Values`.
func (r FormatterStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
