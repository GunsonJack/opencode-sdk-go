// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
	"github.com/GunsonJack/opencode-sdk-go/internal/apiquery"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
	"github.com/tidwall/gjson"
)

// McpService contains methods for interacting with the MCP resource.
type McpService struct {
	Options []option.RequestOption
}

// NewMcpService generates a new service that applies the given options to each
// request.
func NewMcpService(opts ...option.RequestOption) (r *McpService) {
	r = &McpService{}
	r.Options = opts
	return
}

// Status returns the status of all MCP servers.
func (r *McpService) Status(ctx context.Context, query McpStatusParams, opts ...option.RequestOption) (res *map[string]McpStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mcp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Add registers a new MCP server.
func (r *McpService) Add(ctx context.Context, params McpAddParams, opts ...option.RequestOption) (res *map[string]McpStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mcp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Connect connects to an MCP server.
func (r *McpService) Connect(ctx context.Context, name string, params McpConnectParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	name, err = requestconfig.EncodePathSegment(name, "name")
	if err != nil {
		return
	}
	path := fmt.Sprintf("mcp/%s/connect", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Disconnect disconnects from an MCP server.
func (r *McpService) Disconnect(ctx context.Context, name string, params McpDisconnectParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	name, err = requestconfig.EncodePathSegment(name, "name")
	if err != nil {
		return
	}
	path := fmt.Sprintf("mcp/%s/disconnect", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// AuthStart initiates authentication for an MCP server.
func (r *McpService) AuthStart(ctx context.Context, name string, params McpAuthStartParams, opts ...option.RequestOption) (res *McpAuthStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	name, err = requestconfig.EncodePathSegment(name, "name")
	if err != nil {
		return
	}
	path := fmt.Sprintf("mcp/%s/auth", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// AuthRemove removes authentication for an MCP server.
func (r *McpService) AuthRemove(ctx context.Context, name string, params McpAuthRemoveParams, opts ...option.RequestOption) (res *McpAuthRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	name, err = requestconfig.EncodePathSegment(name, "name")
	if err != nil {
		return
	}
	path := fmt.Sprintf("mcp/%s/auth", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// AuthAuthenticate authenticates with an MCP server.
func (r *McpService) AuthAuthenticate(ctx context.Context, name string, params McpAuthAuthenticateParams, opts ...option.RequestOption) (res *McpStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	name, err = requestconfig.EncodePathSegment(name, "name")
	if err != nil {
		return
	}
	path := fmt.Sprintf("mcp/%s/auth/authenticate", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// AuthCallback completes the OAuth callback for an MCP server.
func (r *McpService) AuthCallback(ctx context.Context, name string, params McpAuthCallbackParams, opts ...option.RequestOption) (res *McpStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	name, err = requestconfig.EncodePathSegment(name, "name")
	if err != nil {
		return
	}
	path := fmt.Sprintf("mcp/%s/auth/callback", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// McpStatus is the status of an MCP server.
//
// Union satisfied by [McpStatusConnected], [McpStatusDisabled],
// [McpStatusFailed], [McpStatusNeedsAuth], or
// [McpStatusNeedsClientRegistration].
type McpStatus struct {
	Status string        `json:"status,required"`
	Error  string        `json:"error"`
	JSON   mcpStatusJSON `json:"-"`
	union  McpStatusUnion
}

type mcpStatusJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r mcpStatusJSON) RawJSON() string {
	return r.raw
}

func (r *McpStatus) UnmarshalJSON(data []byte) (err error) {
	*r = McpStatus{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r McpStatus) AsUnion() McpStatusUnion {
	return r.union
}

type McpStatusUnion interface {
	implementsMcpStatus()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*McpStatusUnion)(nil)).Elem(),
		"status",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(McpStatusConnected{}),
			DiscriminatorValue: "connected",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(McpStatusDisabled{}),
			DiscriminatorValue: "disabled",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(McpStatusFailed{}),
			DiscriminatorValue: "failed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(McpStatusNeedsAuth{}),
			DiscriminatorValue: "needs_auth",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(McpStatusNeedsClientRegistration{}),
			DiscriminatorValue: "needs_client_registration",
		},
	)
}

type McpStatusConnected struct {
	Status string                 `json:"status,required"`
	JSON   mcpStatusConnectedJSON `json:"-"`
}

type mcpStatusConnectedJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusConnected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
func (r mcpStatusConnectedJSON) RawJSON() string  { return r.raw }
func (r McpStatusConnected) implementsMcpStatus() {}

type McpStatusDisabled struct {
	Status string                `json:"status,required"`
	JSON   mcpStatusDisabledJSON `json:"-"`
}

type mcpStatusDisabledJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusDisabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
func (r mcpStatusDisabledJSON) RawJSON() string  { return r.raw }
func (r McpStatusDisabled) implementsMcpStatus() {}

type McpStatusFailed struct {
	Status string              `json:"status,required"`
	Error  string              `json:"error,required"`
	JSON   mcpStatusFailedJSON `json:"-"`
}

type mcpStatusFailedJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
func (r mcpStatusFailedJSON) RawJSON() string  { return r.raw }
func (r McpStatusFailed) implementsMcpStatus() {}

type McpStatusNeedsAuth struct {
	Status string                 `json:"status,required"`
	JSON   mcpStatusNeedsAuthJSON `json:"-"`
}

type mcpStatusNeedsAuthJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusNeedsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
func (r mcpStatusNeedsAuthJSON) RawJSON() string  { return r.raw }
func (r McpStatusNeedsAuth) implementsMcpStatus() {}

type McpStatusNeedsClientRegistration struct {
	Status string                               `json:"status,required"`
	Error  string                               `json:"error,required"`
	JSON   mcpStatusNeedsClientRegistrationJSON `json:"-"`
}

type mcpStatusNeedsClientRegistrationJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusNeedsClientRegistration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
func (r mcpStatusNeedsClientRegistrationJSON) RawJSON() string  { return r.raw }
func (r McpStatusNeedsClientRegistration) implementsMcpStatus() {}

// McpAuthStartResponse is the response from POST /mcp/{name}/auth.
type McpAuthStartResponse struct {
	AuthorizationURL string                   `json:"authorizationUrl,required"`
	JSON             mcpAuthStartResponseJSON `json:"-"`
}

type mcpAuthStartResponseJSON struct {
	AuthorizationURL apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *McpAuthStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpAuthStartResponseJSON) RawJSON() string {
	return r.raw
}

// McpAuthRemoveResponse is the response from DELETE /mcp/{name}/auth.
type McpAuthRemoveResponse struct {
	Success bool                      `json:"success,required"`
	JSON    mcpAuthRemoveResponseJSON `json:"-"`
}

type mcpAuthRemoveResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpAuthRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpAuthRemoveResponseJSON) RawJSON() string {
	return r.raw
}

// McpResource describes an MCP-provided resource.
type McpResource struct {
	Name        string          `json:"name,required"`
	URI         string          `json:"uri,required"`
	Client      string          `json:"client,required"`
	Description string          `json:"description"`
	MimeType    string          `json:"mimeType"`
	JSON        mcpResourceJSON `json:"-"`
}

type mcpResourceJSON struct {
	Name        apijson.Field
	URI         apijson.Field
	Client      apijson.Field
	Description apijson.Field
	MimeType    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpResourceJSON) RawJSON() string {
	return r.raw
}

// Params

type McpStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r McpStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type McpAddParams struct {
	Name      param.Field[string]            `json:"name,required"`
	Config    param.Field[McpAddConfigParam] `json:"config,required"`
	Directory param.Field[string]            `query:"directory"`
	Workspace param.Field[string]            `query:"workspace"`
}

func (r McpAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r McpAddParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// McpAddConfigParam is the config body for POST /mcp.
//
// This is a flattened union of McpLocalConfig and McpRemoteConfig
// (anyOf discriminated by Type).
//
// For local servers, set Type to "local" and provide Command (required for local):
//
//	opencode.McpAddConfigParam{
//	    Type:    opencode.F("local"),
//	    Command: opencode.F([]string{"npx", "my-mcp-server"}),
//	}
//
// For remote servers, set Type to "remote" and provide URL (required for remote):
//
//	opencode.McpAddConfigParam{
//	    Type: opencode.F("remote"),
//	    URL:  opencode.F("https://mcp.example.com"),
//	}
//
// Fields from the non-matching variant are ignored during serialization when not
// set (i.e., URL is not serialized for local configs unless explicitly provided).
type McpAddConfigParam struct {
	// The config type: "local" or "remote". (required)
	Type param.Field[string] `json:"type,required"`
	// Command to run (array of strings). Required when Type is "local".
	// Ignored for remote configs unless explicitly set.
	Command     param.Field[[]string]          `json:"command"`
	Environment param.Field[map[string]string] `json:"environment"`
	// URL of the remote server. Required when Type is "remote".
	// Ignored for local configs unless explicitly set.
	URL     param.Field[string]                      `json:"url"`
	Headers param.Field[map[string]string]           `json:"headers"`
	OAuth   param.Field[McpAddConfigOAuthUnionParam] `json:"oauth"`
	// Shared fields (apply to both local and remote configs).
	Enabled param.Field[bool]    `json:"enabled"`
	Timeout param.Field[float64] `json:"timeout"`
}

func (r McpAddConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// McpAddConfigOAuthUnionParam is a param union for the MCP OAuth config.
// Satisfied by [McpAddConfigOAuthParam] (object config) or [McpAddConfigOAuthDisabledParam] (false).
type McpAddConfigOAuthUnionParam interface {
	implementsMcpAddConfigOAuthUnionParam()
}

// McpAddConfigOAuthDisabledParam represents oauth: false to disable OAuth auto-detection.
type McpAddConfigOAuthDisabledParam struct{}

func (r McpAddConfigOAuthDisabledParam) MarshalJSON() (data []byte, err error) {
	return []byte("false"), nil
}

func (r McpAddConfigOAuthDisabledParam) implementsMcpAddConfigOAuthUnionParam() {}

// McpAddConfigOAuthParam represents OAuth configuration for a remote MCP server.
type McpAddConfigOAuthParam struct {
	ClientID     param.Field[string] `json:"clientId"`
	ClientSecret param.Field[string] `json:"clientSecret"`
	Scope        param.Field[string] `json:"scope"`
	RedirectURI  param.Field[string] `json:"redirectUri"`
}

func (r McpAddConfigOAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r McpAddConfigOAuthParam) implementsMcpAddConfigOAuthUnionParam() {}

type McpConnectParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r McpConnectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type McpDisconnectParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r McpDisconnectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type McpAuthStartParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r McpAuthStartParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type McpAuthRemoveParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r McpAuthRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type McpAuthAuthenticateParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r McpAuthAuthenticateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type McpAuthCallbackParams struct {
	Code      param.Field[string] `json:"code,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r McpAuthCallbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r McpAuthCallbackParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
