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

// ProviderService contains methods for interacting with the provider resource.
type ProviderService struct {
	Options []option.RequestOption
}

// NewProviderService generates a new service that applies the given options to
// each request.
func NewProviderService(opts ...option.RequestOption) (r *ProviderService) {
	r = &ProviderService{}
	r.Options = opts
	return
}

// List all providers with their models.
func (r *ProviderService) List(ctx context.Context, query ProviderListParams, opts ...option.RequestOption) (res *ProviderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "provider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get available auth methods for all providers.
func (r *ProviderService) Auth(ctx context.Context, query ProviderAuthParams, opts ...option.RequestOption) (res *map[string][]ProviderAuthMethod, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "provider/auth"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Initiate OAuth authorization for a provider.
func (r *ProviderService) OAuthAuthorize(ctx context.Context, providerID string, params ProviderOAuthAuthorizeParams, opts ...option.RequestOption) (res *ProviderAuthAuthorization, err error) {
	opts = slices.Concat(r.Options, opts)
	providerID, err = requestconfig.EncodePathSegment(providerID, "providerID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("provider/%s/oauth/authorize", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Complete OAuth callback for a provider.
func (r *ProviderService) OAuthCallback(ctx context.Context, providerID string, params ProviderOAuthCallbackParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	providerID, err = requestconfig.EncodePathSegment(providerID, "providerID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("provider/%s/oauth/callback", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ProviderListResponse is the response from GET /provider.
type ProviderListResponse struct {
	All       []Provider               `json:"all,required"`
	Default   map[string]string        `json:"default,required"`
	Connected []string                 `json:"connected,required"`
	JSON      providerListResponseJSON `json:"-"`
}

// providerListResponseJSON contains the JSON metadata for the struct
// [ProviderListResponse]
type providerListResponseJSON struct {
	All         apijson.Field
	Default     apijson.Field
	Connected   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseJSON) RawJSON() string {
	return r.raw
}

// ProviderAuthMethod describes an authentication method for a provider.
type ProviderAuthMethod struct {
	Type    ProviderAuthMethodType     `json:"type,required"`
	Label   string                     `json:"label,required"`
	Prompts []ProviderAuthMethodPrompt `json:"prompts"`
	JSON    providerAuthMethodJSON     `json:"-"`
}

// providerAuthMethodJSON contains the JSON metadata for the struct
// [ProviderAuthMethod]
type providerAuthMethodJSON struct {
	Type        apijson.Field
	Label       apijson.Field
	Prompts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthMethodType string

const (
	ProviderAuthMethodTypeOAuth ProviderAuthMethodType = "oauth"
	ProviderAuthMethodTypeAPI   ProviderAuthMethodType = "api"
)

func (r ProviderAuthMethodType) IsKnown() bool {
	switch r {
	case ProviderAuthMethodTypeOAuth, ProviderAuthMethodTypeAPI:
		return true
	}
	return false
}

// ProviderAuthMethodPrompt is a prompt displayed during auth setup.
type ProviderAuthMethodPrompt struct {
	Type        string                           `json:"type,required"`
	Key         string                           `json:"key,required"`
	Message     string                           `json:"message,required"`
	Placeholder string                           `json:"placeholder"`
	Options     []ProviderAuthMethodPromptOption `json:"options"`
	When        *ProviderAuthMethodPromptWhen    `json:"when"`
	JSON        providerAuthMethodPromptJSON     `json:"-"`
}

// providerAuthMethodPromptJSON contains the JSON metadata for the struct
// [ProviderAuthMethodPrompt]
type providerAuthMethodPromptJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Placeholder apijson.Field
	Options     apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthMethodPromptOption struct {
	Label string                             `json:"label,required"`
	Value string                             `json:"value,required"`
	Hint  string                             `json:"hint"`
	JSON  providerAuthMethodPromptOptionJSON `json:"-"`
}

// providerAuthMethodPromptOptionJSON contains the JSON metadata for the struct
// [ProviderAuthMethodPromptOption]
type providerAuthMethodPromptOptionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	Hint        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptOptionJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthMethodPromptWhen struct {
	Key   string                           `json:"key,required"`
	Op    ProviderAuthMethodPromptWhenOp   `json:"op,required"`
	Value string                           `json:"value,required"`
	JSON  providerAuthMethodPromptWhenJSON `json:"-"`
}

// providerAuthMethodPromptWhenJSON contains the JSON metadata for the struct
// [ProviderAuthMethodPromptWhen]
type providerAuthMethodPromptWhenJSON struct {
	Key         apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptWhen) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptWhenJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthMethodPromptWhenOp string

const (
	ProviderAuthMethodPromptWhenOpEq  ProviderAuthMethodPromptWhenOp = "eq"
	ProviderAuthMethodPromptWhenOpNeq ProviderAuthMethodPromptWhenOp = "neq"
)

func (r ProviderAuthMethodPromptWhenOp) IsKnown() bool {
	switch r {
	case ProviderAuthMethodPromptWhenOpEq, ProviderAuthMethodPromptWhenOpNeq:
		return true
	}
	return false
}

// ProviderAuthAuthorization is the response from OAuth authorize.
type ProviderAuthAuthorization struct {
	URL          string                          `json:"url,required"`
	Method       ProviderAuthAuthorizationMethod `json:"method,required"`
	Instructions string                          `json:"instructions,required"`
	JSON         providerAuthAuthorizationJSON   `json:"-"`
}

// providerAuthAuthorizationJSON contains the JSON metadata for the struct
// [ProviderAuthAuthorization]
type providerAuthAuthorizationJSON struct {
	URL          apijson.Field
	Method       apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProviderAuthAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthAuthorizationJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthAuthorizationMethod string

const (
	ProviderAuthAuthorizationMethodAuto ProviderAuthAuthorizationMethod = "auto"
	ProviderAuthAuthorizationMethodCode ProviderAuthAuthorizationMethod = "code"
)

func (r ProviderAuthAuthorizationMethod) IsKnown() bool {
	switch r {
	case ProviderAuthAuthorizationMethodAuto, ProviderAuthAuthorizationMethodCode:
		return true
	}
	return false
}

// Params

type ProviderListParams struct {
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

// URLQuery serializes [ProviderListParams]'s query parameters as `url.Values`.
func (r ProviderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderAuthParams struct {
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

// URLQuery serializes [ProviderAuthParams]'s query parameters as `url.Values`.
func (r ProviderAuthParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderOAuthAuthorizeParams struct {
	// Index of the chosen auth method from the ProviderService.Auth response.
	Method    param.Field[int64]             `json:"method,required"`
	Inputs    param.Field[map[string]string] `json:"inputs"`
	Workspace param.Field[string]            `query:"workspace"`
	Directory param.Field[string]            `query:"directory"`
}

func (r ProviderOAuthAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ProviderOAuthAuthorizeParams]'s query parameters as
// `url.Values`.
func (r ProviderOAuthAuthorizeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProviderOAuthCallbackParams struct {
	// Index of the chosen auth method from the ProviderService.Auth response.
	Method    param.Field[int64]  `json:"method,required"`
	Code      param.Field[string] `json:"code"`
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

func (r ProviderOAuthCallbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ProviderOAuthCallbackParams]'s query parameters as
// `url.Values`.
func (r ProviderOAuthCallbackParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
