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

// SessionPermissionService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionPermissionService] method instead.
type SessionPermissionService struct {
	Options []option.RequestOption
}

// NewSessionPermissionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSessionPermissionService(opts ...option.RequestOption) (r *SessionPermissionService) {
	r = &SessionPermissionService{}
	r.Options = opts
	return
}

// Deprecated: Respond is a legacy compatibility route. Use [PermissionService.Reply] instead.
//
// Respond to a permission request
func (r *SessionPermissionService) Respond(ctx context.Context, id string, permissionID string, params SessionPermissionRespondParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	id, err = requestconfig.EncodePathSegment(id, "id")
	if err != nil {
		return
	}
	permissionID, err = requestconfig.EncodePathSegment(permissionID, "permissionID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("session/%s/permissions/%s", id, permissionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SessionPermissionRespondParams struct {
	Response  param.Field[SessionPermissionRespondParamsResponse] `json:"response,required"`
	Directory param.Field[string]                                 `query:"directory"`
	Workspace param.Field[string]                                 `query:"workspace"`
}

func (r SessionPermissionRespondParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionPermissionRespondParams]'s query parameters as
// `url.Values`.
func (r SessionPermissionRespondParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionPermissionRespondParamsResponse string

const (
	SessionPermissionRespondParamsResponseOnce   SessionPermissionRespondParamsResponse = "once"
	SessionPermissionRespondParamsResponseAlways SessionPermissionRespondParamsResponse = "always"
	SessionPermissionRespondParamsResponseReject SessionPermissionRespondParamsResponse = "reject"
)

func (r SessionPermissionRespondParamsResponse) IsKnown() bool {
	switch r {
	case SessionPermissionRespondParamsResponseOnce, SessionPermissionRespondParamsResponseAlways, SessionPermissionRespondParamsResponseReject:
		return true
	}
	return false
}
