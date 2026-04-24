// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/GunsonJack/opencode-sdk-go/internal/apijson"
	"github.com/GunsonJack/opencode-sdk-go/internal/param"
	"github.com/GunsonJack/opencode-sdk-go/internal/requestconfig"
	"github.com/GunsonJack/opencode-sdk-go/option"
	"github.com/GunsonJack/opencode-sdk-go/packages/ssestream"
	"github.com/tidwall/gjson"
)

// GlobalService contains methods for interacting with the global resource.
type GlobalService struct {
	Options []option.RequestOption
	Config  *GlobalConfigService
}

// NewGlobalService generates a new service that applies the given options to each
// request.
func NewGlobalService(opts ...option.RequestOption) (r *GlobalService) {
	r = &GlobalService{}
	r.Options = opts
	r.Config = NewGlobalConfigService(opts...)
	return
}

// Health checks the health of the server.
func (r *GlobalService) Health(ctx context.Context, opts ...option.RequestOption) (res *GlobalHealthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get global events.
func (r *GlobalService) Event(ctx context.Context, opts ...option.RequestOption) (stream *ssestream.Stream[GlobalEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "global/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[GlobalEvent](ssestream.NewDecoder(raw), err)
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

// GlobalConfigService contains methods for interacting with the global config resource.
type GlobalConfigService struct {
	Options []option.RequestOption
}

// NewGlobalConfigService generates a new service.
func NewGlobalConfigService(opts ...option.RequestOption) (r *GlobalConfigService) {
	r = &GlobalConfigService{}
	r.Options = opts
	return
}

// Get retrieves the global configuration.
func (r *GlobalConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update updates the global configuration.
func (r *GlobalConfigService) Update(ctx context.Context, params GlobalConfigUpdateParams, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
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

// GlobalEvent is the response envelope from GET /global/event.
type GlobalEvent struct {
	Directory string          `json:"directory,required"`
	Project   string          `json:"project"`
	Workspace string          `json:"workspace"`
	Payload   interface{}     `json:"payload,required"`
	JSON      globalEventJSON `json:"-"`
}

type globalEventJSON struct {
	Directory   apijson.Field
	Project     apijson.Field
	Workspace   apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalEvent) UnmarshalJSON(data []byte) (err error) {
	type globalEventEnvelope struct {
		Directory string          `json:"directory,required"`
		Project   string          `json:"project"`
		Workspace string          `json:"workspace"`
		Payload   json.RawMessage `json:"payload,required"`
		JSON      globalEventJSON `json:"-"`
	}

	var envelope globalEventEnvelope
	if err = apijson.UnmarshalRoot(data, &envelope); err != nil {
		return err
	}
	if envelope.JSON.Directory.IsMissing() {
		return fmt.Errorf("missing required field: directory")
	}
	if envelope.JSON.Payload.IsMissing() {
		return fmt.Errorf("missing required field: payload")
	}

	*r = GlobalEvent{
		Directory: envelope.Directory,
		Project:   envelope.Project,
		Workspace: envelope.Workspace,
		JSON:      envelope.JSON,
	}

	parsed := gjson.ParseBytes(envelope.Payload)
	if parsed.Get("type").String() == string(SyncEventTypeSync) {
		var evt SyncEvent
		if err = json.Unmarshal(envelope.Payload, &evt); err != nil {
			return err
		}
		r.Payload = evt.AsUnion()
		return nil
	}

	var evt EventListResponse
	if err = json.Unmarshal(envelope.Payload, &evt); err != nil {
		return err
	}
	r.Payload = evt.AsUnion()
	return nil
}

func (r globalEventJSON) RawJSON() string {
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
	parsed := gjson.ParseBytes(data)
	if !parsed.Get("success").Exists() {
		return fmt.Errorf("missing required field: success")
	}
	if parsed.Get("success").Type != gjson.True && parsed.Get("success").Type != gjson.False {
		return fmt.Errorf("invalid field type: success must be boolean")
	}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	if err = apijson.Port(r.union, &r); err != nil {
		return err
	}
	if r.Success && !parsed.Get("version").Exists() {
		return fmt.Errorf("missing required field: version")
	}
	if !r.Success && !parsed.Get("error").Exists() {
		return fmt.Errorf("missing required field: error")
	}
	return nil
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
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	if r.JSON.Version.IsMissing() {
		return fmt.Errorf("missing required field: version")
	}
	return nil
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
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	if r.JSON.Error.IsMissing() {
		return fmt.Errorf("missing required field: error")
	}
	return nil
}

func (r globalUpgradeResponseFailureJSON) RawJSON() string {
	return r.raw
}

func (r GlobalUpgradeResponseFailure) implementsGlobalUpgradeResponse() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalUpgradeResponseUnion)(nil)).Elem(),
		"success",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: true,
			Type:               reflect.TypeOf(GlobalUpgradeResponseSuccess{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: false,
			Type:               reflect.TypeOf(GlobalUpgradeResponseFailure{}),
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
	Permission  []PermissionRule  `json:"permission"`
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
	Schema            param.Field[string]                       `json:"$schema"`
	Agent             param.Field[map[string]interface{}]       `json:"agent"`
	Autoshare         param.Field[bool]                         `json:"autoshare"`
	Autoupdate        param.Field[interface{}]                  `json:"autoupdate"`
	Command           param.Field[map[string]interface{}]       `json:"command"`
	Compaction        param.Field[ConfigUpdateParamsCompaction] `json:"compaction"`
	DefaultAgent      param.Field[string]                       `json:"default_agent"`
	DisabledProviders param.Field[[]string]                     `json:"disabled_providers"`
	EnabledProviders  param.Field[[]string]                     `json:"enabled_providers"`
	Enterprise        param.Field[ConfigUpdateParamsEnterprise] `json:"enterprise"`
	Experimental      param.Field[interface{}]                  `json:"experimental"`
	Formatter         param.Field[interface{}]                  `json:"formatter"`
	Instructions      param.Field[[]string]                     `json:"instructions"`
	Layout            param.Field[LayoutConfig]                 `json:"layout"`
	LogLevel          param.Field[ConfigLogLevel]               `json:"logLevel"`
	Lsp               param.Field[interface{}]                  `json:"lsp"`
	Mcp               param.Field[interface{}]                  `json:"mcp"`
	Mode              param.Field[map[string]interface{}]       `json:"mode"`
	Model             param.Field[string]                       `json:"model"`
	Permission        param.Field[interface{}]                  `json:"permission"`
	Plugin            param.Field[[]interface{}]                `json:"plugin"`
	Provider          param.Field[map[string]interface{}]       `json:"provider"`
	Server            param.Field[ConfigUpdateParamsServer]     `json:"server"`
	Share             param.Field[ConfigShare]                  `json:"share"`
	Skills            param.Field[ConfigUpdateParamsSkills]     `json:"skills"`
	SmallModel        param.Field[string]                       `json:"small_model"`
	Snapshot          param.Field[bool]                         `json:"snapshot"`
	Tools             param.Field[map[string]bool]              `json:"tools"`
	Username          param.Field[string]                       `json:"username"`
	Watcher           param.Field[ConfigUpdateParamsWatcher]    `json:"watcher"`
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
