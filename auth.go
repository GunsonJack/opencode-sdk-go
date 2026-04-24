// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AuthService contains methods for interacting with the auth resource.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request.
func NewAuthService(opts ...option.RequestOption) (r *AuthService) {
	r = &AuthService{}
	r.Options = opts
	return
}

// Set credentials for a provider.
func (r *AuthService) Set(ctx context.Context, providerID string, params AuthSetParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	providerID, err = requestconfig.EncodePathSegment(providerID, "providerID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("auth/%s", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Remove credentials for a provider.
func (r *AuthService) Remove(ctx context.Context, providerID string, params AuthRemoveParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	providerID, err = requestconfig.EncodePathSegment(providerID, "providerID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("auth/%s", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Auth is the credential union stored for a provider.
//
// Union satisfied by [AuthOAuth], [AuthAPI], or [AuthWellKnown].
type Auth struct {
	Type string `json:"type,required"`
	// OAuth fields
	Refresh       string  `json:"refresh"`
	Access        string  `json:"access"`
	Expires       float64 `json:"expires"`
	AccountID     string  `json:"accountId"`
	EnterpriseURL string  `json:"enterpriseUrl"`
	// API fields
	Key      string            `json:"key"`
	Metadata map[string]string `json:"metadata"`
	// WellKnown fields
	Token string   `json:"token"`
	JSON  authJSON `json:"-"`
	union AuthUnion
}

// authJSON contains the JSON metadata for the struct [Auth]
type authJSON struct {
	Type          apijson.Field
	Refresh       apijson.Field
	Access        apijson.Field
	Expires       apijson.Field
	AccountID     apijson.Field
	EnterpriseURL apijson.Field
	Key           apijson.Field
	Metadata      apijson.Field
	Token         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r authJSON) RawJSON() string {
	return r.raw
}

func (r *Auth) UnmarshalJSON(data []byte) (err error) {
	*r = Auth{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns the underlying [AuthUnion] variant of this union type.
func (r Auth) AsUnion() AuthUnion {
	return r.union
}

// AuthUnion is the interface satisfied by [AuthOAuth], [AuthAPI], and
// [AuthWellKnown].
type AuthUnion interface {
	implementsAuth()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuthUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AuthOAuth{}),
			DiscriminatorValue: "oauth",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AuthAPI{}),
			DiscriminatorValue: "api",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AuthWellKnown{}),
			DiscriminatorValue: "wellknown",
		},
	)
}

type AuthOAuth struct {
	Type          string        `json:"type,required"`
	Refresh       string        `json:"refresh,required"`
	Access        string        `json:"access,required"`
	Expires       float64       `json:"expires,required"`
	AccountID     string        `json:"accountId"`
	EnterpriseURL string        `json:"enterpriseUrl"`
	JSON          authOAuthJSON `json:"-"`
}

// authOAuthJSON contains the JSON metadata for the struct [AuthOAuth]
type authOAuthJSON struct {
	Type          apijson.Field
	Refresh       apijson.Field
	Access        apijson.Field
	Expires       apijson.Field
	AccountID     apijson.Field
	EnterpriseURL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AuthOAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authOAuthJSON) RawJSON() string {
	return r.raw
}

func (r AuthOAuth) implementsAuth() {}

type AuthAPI struct {
	Type     string            `json:"type,required"`
	Key      string            `json:"key,required"`
	Metadata map[string]string `json:"metadata"`
	JSON     authAPIJSON       `json:"-"`
}

// authAPIJSON contains the JSON metadata for the struct [AuthAPI]
type authAPIJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthAPI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authAPIJSON) RawJSON() string {
	return r.raw
}

func (r AuthAPI) implementsAuth() {}

type AuthWellKnown struct {
	Type  string            `json:"type,required"`
	Key   string            `json:"key,required"`
	Token string            `json:"token,required"`
	JSON  authWellKnownJSON `json:"-"`
}

// authWellKnownJSON contains the JSON metadata for the struct [AuthWellKnown]
type authWellKnownJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthWellKnown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authWellKnownJSON) RawJSON() string {
	return r.raw
}

func (r AuthWellKnown) implementsAuth() {}

// Params

type AuthSetParams struct {
	// The credential type: "oauth", "api", or "wellknown".
	Type param.Field[string] `json:"type,required"`
	// OAuth fields
	Refresh       param.Field[string]  `json:"refresh"`
	Access        param.Field[string]  `json:"access"`
	Expires       param.Field[float64] `json:"expires"`
	AccountID     param.Field[string]  `json:"accountId"`
	EnterpriseURL param.Field[string]  `json:"enterpriseUrl"`
	// API fields
	Key      param.Field[string]            `json:"key"`
	Metadata param.Field[map[string]string] `json:"metadata"`
	// WellKnown fields
	Token param.Field[string] `json:"token"`
}

func (r AuthSetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AuthRemoveParams struct{}
