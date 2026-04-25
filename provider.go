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

// ProviderAuthMethodPrompt is a union type for auth prompts.
// Use [ProviderAuthMethodPrompt.AsUnion] to access the underlying variant.
//
// Union satisfied by [ProviderAuthMethodPromptText] or [ProviderAuthMethodPromptSelect].
type ProviderAuthMethodPrompt struct {
	Type    string `json:"type,required"`
	Key     string `json:"key,required"`
	Message string `json:"message,required"`
	// This field can have the runtime type of [[]ProviderAuthMethodPromptSelectOption].
	Options     interface{}                   `json:"options"`
	Placeholder string                        `json:"placeholder"`
	When        *ProviderAuthMethodPromptWhen `json:"when"`
	JSON        providerAuthMethodPromptJSON  `json:"-"`
	union       ProviderAuthMethodPromptUnion
}

type providerAuthMethodPromptJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Options     apijson.Field
	Placeholder apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPrompt) UnmarshalJSON(data []byte) (err error) {
	*r = ProviderAuthMethodPrompt{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r providerAuthMethodPromptJSON) RawJSON() string {
	return r.raw
}

// AsUnion returns the underlying union variant.
func (r ProviderAuthMethodPrompt) AsUnion() ProviderAuthMethodPromptUnion {
	return r.union
}

// Union satisfied by [ProviderAuthMethodPromptText] or [ProviderAuthMethodPromptSelect].
type ProviderAuthMethodPromptUnion interface {
	implementsProviderAuthMethodPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProviderAuthMethodPromptUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "text",
			Type:               reflect.TypeOf(ProviderAuthMethodPromptText{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "select",
			Type:               reflect.TypeOf(ProviderAuthMethodPromptSelect{}),
		},
	)
}

type ProviderAuthMethodPromptText struct {
	Type        ProviderAuthMethodPromptTextType `json:"type,required"`
	Key         string                           `json:"key,required"`
	Message     string                           `json:"message,required"`
	Placeholder string                           `json:"placeholder"`
	When        *ProviderAuthMethodPromptWhen    `json:"when"`
	JSON        providerAuthMethodPromptTextJSON `json:"-"`
}

type providerAuthMethodPromptTextJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Placeholder apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptTextJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthMethodPromptText) implementsProviderAuthMethodPrompt() {}

type ProviderAuthMethodPromptTextType string

const (
	ProviderAuthMethodPromptTextTypeText ProviderAuthMethodPromptTextType = "text"
)

func (r ProviderAuthMethodPromptTextType) IsKnown() bool {
	switch r {
	case ProviderAuthMethodPromptTextTypeText:
		return true
	}
	return false
}

type ProviderAuthMethodPromptSelect struct {
	Type    ProviderAuthMethodPromptSelectType     `json:"type,required"`
	Key     string                                 `json:"key,required"`
	Message string                                 `json:"message,required"`
	Options []ProviderAuthMethodPromptSelectOption `json:"options,required"`
	When    *ProviderAuthMethodPromptWhen          `json:"when"`
	JSON    providerAuthMethodPromptSelectJSON     `json:"-"`
}

type providerAuthMethodPromptSelectJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Options     apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptSelect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptSelectJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthMethodPromptSelect) implementsProviderAuthMethodPrompt() {}

type ProviderAuthMethodPromptSelectType string

const (
	ProviderAuthMethodPromptSelectTypeSelect ProviderAuthMethodPromptSelectType = "select"
)

func (r ProviderAuthMethodPromptSelectType) IsKnown() bool {
	switch r {
	case ProviderAuthMethodPromptSelectTypeSelect:
		return true
	}
	return false
}

type ProviderAuthMethodPromptSelectOption struct {
	Label string                                   `json:"label,required"`
	Value string                                   `json:"value,required"`
	Hint  string                                   `json:"hint"`
	JSON  providerAuthMethodPromptSelectOptionJSON `json:"-"`
}

type providerAuthMethodPromptSelectOptionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	Hint        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptSelectOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptSelectOptionJSON) RawJSON() string {
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
	Method    param.Field[float64]           `json:"method,required"`
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
	Method    param.Field[float64] `json:"method,required"`
	Code      param.Field[string]  `json:"code"`
	Workspace param.Field[string]  `query:"workspace"`
	Directory param.Field[string]  `query:"directory"`
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
