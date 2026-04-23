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

type ConsoleService struct {
	Options []option.RequestOption
}

func NewConsoleService(opts ...option.RequestOption) (r *ConsoleService) {
	r = &ConsoleService{}
	r.Options = opts
	return
}

func (r *ConsoleService) Get(ctx context.Context, query ConsoleGetParams, opts ...option.RequestOption) (res *ConsoleState, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *ConsoleService) ListOrgs(ctx context.Context, query ConsoleListOrgsParams, opts ...option.RequestOption) (res *ConsoleListOrgsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *ConsoleService) SwitchOrg(ctx context.Context, params ConsoleSwitchOrgParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console/switch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ConsoleState struct {
	ConsoleManagedProviders []string         `json:"consoleManagedProviders,required"`
	SwitchableOrgCount      float64          `json:"switchableOrgCount,required"`
	ActiveOrgName           string           `json:"activeOrgName"`
	JSON                    consoleStateJSON `json:"-"`
}

type consoleStateJSON struct {
	ConsoleManagedProviders apijson.Field
	SwitchableOrgCount      apijson.Field
	ActiveOrgName           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ConsoleState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleStateJSON) RawJSON() string { return r.raw }

type ConsoleListOrgsResponse struct {
	Orgs []ConsoleOrg                `json:"orgs,required"`
	JSON consoleListOrgsResponseJSON `json:"-"`
}

type consoleListOrgsResponseJSON struct {
	Orgs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsoleListOrgsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleListOrgsResponseJSON) RawJSON() string { return r.raw }

type ConsoleOrg struct {
	AccountID    string         `json:"accountID,required"`
	AccountEmail string         `json:"accountEmail,required"`
	AccountURL   string         `json:"accountUrl,required"`
	OrgID        string         `json:"orgID,required"`
	OrgName      string         `json:"orgName,required"`
	Active       bool           `json:"active,required"`
	JSON         consoleOrgJSON `json:"-"`
}

type consoleOrgJSON struct {
	AccountID    apijson.Field
	AccountEmail apijson.Field
	AccountURL   apijson.Field
	OrgID        apijson.Field
	OrgName      apijson.Field
	Active       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConsoleOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleOrgJSON) RawJSON() string { return r.raw }

// Params

type ConsoleGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ConsoleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConsoleListOrgsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ConsoleListOrgsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConsoleSwitchOrgParams struct {
	AccountID param.Field[string] `json:"accountID,required"`
	OrgID     param.Field[string] `json:"orgID,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ConsoleSwitchOrgParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConsoleSwitchOrgParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
