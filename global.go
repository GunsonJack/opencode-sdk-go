// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
	"github.com/tidwall/gjson"
)

// GlobalService contains methods for interacting with the global resource.
type GlobalService struct {
	Options []option.RequestOption
}

// NewGlobalService generates a new service that applies the given options to each
// request.
func NewGlobalService(opts ...option.RequestOption) (r *GlobalService) {
	r = &GlobalService{}
	r.Options = opts
	return
}

// Health checks the health of the server.
func (r *GlobalService) Health(ctx context.Context, opts ...option.RequestOption) (res *GlobalHealthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// ConfigGet retrieves the global configuration.
func (r *GlobalService) ConfigGet(ctx context.Context, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// ConfigUpdate updates the global configuration.
func (r *GlobalService) ConfigUpdate(ctx context.Context, params GlobalConfigUpdateParams, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Dispose shuts down the server.
func (r *GlobalService) Dispose(ctx context.Context, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/dispose"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Upgrade upgrades the server to a target version.
func (r *GlobalService) Upgrade(ctx context.Context, params GlobalUpgradeParams, opts ...option.RequestOption) (res *GlobalUpgradeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/upgrade"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// GlobalHealthResponse is the response from GET /global/health.
type GlobalHealthResponse struct {
	Healthy bool                     `json:"healthy,required"`
	Version string                   `json:"version,required"`
	JSON    globalHealthResponseJSON `json:"-"`
}

type globalHealthResponseJSON struct {
	Healthy     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalHealthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalHealthResponseJSON) RawJSON() string {
	return r.raw
}

// GlobalUpgradeResponse is the response from POST /global/upgrade.
//
// Union satisfied by [GlobalUpgradeResponseSuccess] or
// [GlobalUpgradeResponseFailure].
type GlobalUpgradeResponse struct {
	Success bool                      `json:"success,required"`
	Version string                    `json:"version"`
	Error   string                    `json:"error"`
	JSON    globalUpgradeResponseJSON `json:"-"`
	union   GlobalUpgradeResponseUnion
}

type globalUpgradeResponseJSON struct {
	Success     apijson.Field
	Version     apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r globalUpgradeResponseJSON) RawJSON() string {
	return r.raw
}

func (r *GlobalUpgradeResponse) UnmarshalJSON(data []byte) (err error) {
	*r = GlobalUpgradeResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r GlobalUpgradeResponse) AsUnion() GlobalUpgradeResponseUnion {
	return r.union
}

type GlobalUpgradeResponseUnion interface {
	implementsGlobalUpgradeResponse()
}

type GlobalUpgradeResponseSuccess struct {
	Success bool                             `json:"success,required"`
	Version string                           `json:"version,required"`
	JSON    globalUpgradeResponseSuccessJSON `json:"-"`
}

type globalUpgradeResponseSuccessJSON struct {
	Success     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalUpgradeResponseSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalUpgradeResponseSuccessJSON) RawJSON() string {
	return r.raw
}

func (r GlobalUpgradeResponseSuccess) implementsGlobalUpgradeResponse() {}

type GlobalUpgradeResponseFailure struct {
	Success bool                             `json:"success,required"`
	Error   string                           `json:"error,required"`
	JSON    globalUpgradeResponseFailureJSON `json:"-"`
}

type globalUpgradeResponseFailureJSON struct {
	Success     apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalUpgradeResponseFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalUpgradeResponseFailureJSON) RawJSON() string {
	return r.raw
}

func (r GlobalUpgradeResponseFailure) implementsGlobalUpgradeResponse() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalUpgradeResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GlobalUpgradeResponseSuccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GlobalUpgradeResponseFailure{}),
		},
	)
}

// GlobalSession extends Session with project info for cross-workspace listing.
type GlobalSession struct {
	ID          string            `json:"id,required"`
	Slug        string            `json:"slug,required"`
	ProjectID   string            `json:"projectID,required"`
	Directory   string            `json:"directory,required"`
	Title       string            `json:"title,required"`
	Version     string            `json:"version,required"`
	Time        SessionTime       `json:"time,required"`
	Project     *ProjectSummary   `json:"project,required"`
	WorkspaceID string            `json:"workspaceID"`
	ParentID    string            `json:"parentID"`
	Summary     *SessionSummary   `json:"summary"`
	Share       *SessionShare     `json:"share"`
	Permission  interface{}       `json:"permission"`
	Revert      *SessionRevert    `json:"revert"`
	JSON        globalSessionJSON `json:"-"`
}

type globalSessionJSON struct {
	ID          apijson.Field
	Slug        apijson.Field
	ProjectID   apijson.Field
	Directory   apijson.Field
	Title       apijson.Field
	Version     apijson.Field
	Time        apijson.Field
	Project     apijson.Field
	WorkspaceID apijson.Field
	ParentID    apijson.Field
	Summary     apijson.Field
	Share       apijson.Field
	Permission  apijson.Field
	Revert      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionJSON) RawJSON() string {
	return r.raw
}

// Params

type GlobalConfigUpdateParams struct {
	// Accepts the same shape as ConfigUpdateParams body.
	Model    param.Field[string]      `json:"model"`
	Agent    param.Field[interface{}] `json:"agent"`
	Plugin   param.Field[interface{}] `json:"plugin"`
	Provider param.Field[interface{}] `json:"provider"`
}

func (r GlobalConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GlobalUpgradeParams struct {
	Target param.Field[string] `json:"target"`
}

func (r GlobalUpgradeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
