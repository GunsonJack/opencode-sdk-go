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
	requestID, err = requestconfig.EncodePathSegment(requestID, "requestID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("permission/%s/reply", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PermissionRequest struct {
	ID         string                 `json:"id,required"`
	SessionID  string                 `json:"sessionID,required"`
	Permission string                 `json:"permission,required"`
	Patterns   []string               `json:"patterns,required"`
	Metadata   map[string]interface{} `json:"metadata,required"`
	Always     []string               `json:"always,required"`
	Tool       PermissionRequestTool  `json:"tool"`
	JSON       permissionRequestJSON  `json:"-"`
}

// PermissionRequestTool represents the optional tool reference on a permission
// request.
type PermissionRequestTool struct {
	MessageID string                    `json:"messageID"`
	CallID    string                    `json:"callID"`
	JSON      permissionRequestToolJSON `json:"-"`
}

// permissionRequestToolJSON contains the JSON metadata for the struct
// [PermissionRequestTool]
type permissionRequestToolJSON struct {
	MessageID   apijson.Field
	CallID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionRequestTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionRequestToolJSON) RawJSON() string {
	return r.raw
}

// permissionRequestJSON contains the JSON metadata for the struct [PermissionRequest]
type permissionRequestJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Permission  apijson.Field
	Patterns    apijson.Field
	Metadata    apijson.Field
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

// PermissionReplyParamsReply is the enum type for the reply field.
type PermissionReplyParamsReply string

const (
	PermissionReplyParamsReplyOnce   PermissionReplyParamsReply = "once"
	PermissionReplyParamsReplyAlways PermissionReplyParamsReply = "always"
	PermissionReplyParamsReplyReject PermissionReplyParamsReply = "reject"
)

type PermissionReplyParams struct {
	Reply     param.Field[PermissionReplyParamsReply] `json:"reply,required"`
	Message   param.Field[string]                     `json:"message"`
	Workspace param.Field[string]                     `query:"workspace"`
	Directory param.Field[string]                     `query:"directory"`
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
