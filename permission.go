// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
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

// PermissionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPermissionService] method instead.
type PermissionService struct {
	Options []option.RequestOption
}

// NewPermissionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPermissionService(opts ...option.RequestOption) (r *PermissionService) {
	r = &PermissionService{}
	r.Options = opts
	return
}

// List pending permission requests
func (r *PermissionService) List(ctx context.Context, query PermissionListParams, opts ...option.RequestOption) (res *[]PermissionRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "permission"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Reply to a permission request
func (r *PermissionService) Reply(ctx context.Context, requestID string, params PermissionReplyParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return
	}
	path := fmt.Sprintf("permission/%s/reply", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PermissionRequest struct {
	ID         string                `json:"id,required"`
	Permission string                `json:"permission,required"`
	Patterns   []string              `json:"patterns"`
	Always     bool                  `json:"always"`
	Tool       string                `json:"tool"`
	JSON       permissionRequestJSON `json:"-"`
}

// permissionRequestJSON contains the JSON metadata for the struct [PermissionRequest]
type permissionRequestJSON struct {
	ID          apijson.Field
	Permission  apijson.Field
	Patterns    apijson.Field
	Always      apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionRequestJSON) RawJSON() string {
	return r.raw
}

type PermissionListParams struct {
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

// URLQuery serializes [PermissionListParams]'s query parameters as `url.Values`.
func (r PermissionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PermissionReplyParams struct {
	Response  param.Field[string] `json:"response,required"`
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

func (r PermissionReplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PermissionReplyParams]'s query parameters as `url.Values`.
func (r PermissionReplyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
