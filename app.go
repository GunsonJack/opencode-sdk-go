// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
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

// AppService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

// Write a log entry to the server logs
func (r *AppService) Log(ctx context.Context, params AppLogParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "log"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Providers lists all providers.
//
// Deprecated: Use [ConfigService.Providers] instead.
func (r *AppService) Providers(ctx context.Context, query AppProvidersParams, opts ...option.RequestOption) (res *AppProvidersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Model struct {
	ID           string                            `json:"id,required"`
	ProviderID   string                            `json:"providerID,required"`
	API          ModelAPI                          `json:"api,required"`
	Name         string                            `json:"name,required"`
	Capabilities ModelCapabilities                 `json:"capabilities,required"`
	Cost         ModelCost                         `json:"cost,required"`
	Limit        ModelLimit                        `json:"limit,required"`
	Status       ModelStatus                       `json:"status,required"`
	Options      map[string]interface{}            `json:"options,required"`
	Headers      map[string]string                 `json:"headers,required"`
	ReleaseDate  string                            `json:"release_date,required"`
	Family       string                            `json:"family"`
	Variants     map[string]map[string]interface{} `json:"variants"`
	JSON         modelJSON                         `json:"-"`
}

// modelJSON contains the JSON metadata for the struct [Model]
type modelJSON struct {
	ID           apijson.Field
	ProviderID   apijson.Field
	API          apijson.Field
	Name         apijson.Field
	Capabilities apijson.Field
	Cost         apijson.Field
	Limit        apijson.Field
	Status       apijson.Field
	Options      apijson.Field
	Headers      apijson.Field
	ReleaseDate  apijson.Field
	Family       apijson.Field
	Variants     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Model) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelJSON) RawJSON() string {
	return r.raw
}

type ModelCost struct {
	Input                float64                       `json:"input,required"`
	Output               float64                       `json:"output,required"`
	Cache                ModelCostCache                `json:"cache,required"`
	ExperimentalOver200K ModelCostExperimentalOver200K `json:"experimentalOver200K"`
	JSON                 modelCostJSON                 `json:"-"`
}

// modelCostJSON contains the JSON metadata for the struct [ModelCost]
type modelCostJSON struct {
	Input                apijson.Field
	Output               apijson.Field
	Cache                apijson.Field
	ExperimentalOver200K apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ModelCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCostJSON) RawJSON() string {
	return r.raw
}

type ModelCostCache struct {
	Read  float64            `json:"read,required"`
	Write float64            `json:"write,required"`
	JSON  modelCostCacheJSON `json:"-"`
}

type modelCostCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelCostCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCostCacheJSON) RawJSON() string {
	return r.raw
}

type ModelCostExperimentalOver200K struct {
	Input  float64                           `json:"input,required"`
	Output float64                           `json:"output,required"`
	Cache  ModelCostCache                    `json:"cache,required"`
	JSON   modelCostExperimentalOver200KJSON `json:"-"`
}

type modelCostExperimentalOver200KJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelCostExperimentalOver200K) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCostExperimentalOver200KJSON) RawJSON() string {
	return r.raw
}

type ModelLimit struct {
	Context float64        `json:"context,required"`
	Output  float64        `json:"output,required"`
	Input   float64        `json:"input"`
	JSON    modelLimitJSON `json:"-"`
}

// modelLimitJSON contains the JSON metadata for the struct [ModelLimit]
type modelLimitJSON struct {
	Context     apijson.Field
	Output      apijson.Field
	Input       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelLimitJSON) RawJSON() string {
	return r.raw
}

type ModelCapabilities struct {
	Temperature bool                         `json:"temperature,required"`
	Reasoning   bool                         `json:"reasoning,required"`
	Attachment  bool                         `json:"attachment,required"`
	Toolcall    bool                         `json:"toolcall,required"`
	Input       ModelCapabilitiesModality    `json:"input,required"`
	Output      ModelCapabilitiesModality    `json:"output,required"`
	Interleaved ModelCapabilitiesInterleaved `json:"interleaved,required"`
	JSON        modelCapabilitiesJSON        `json:"-"`
}

type modelCapabilitiesJSON struct {
	Temperature apijson.Field
	Reasoning   apijson.Field
	Attachment  apijson.Field
	Toolcall    apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	Interleaved apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type ModelCapabilitiesModality struct {
	Text  bool                          `json:"text,required"`
	Audio bool                          `json:"audio,required"`
	Image bool                          `json:"image,required"`
	Video bool                          `json:"video,required"`
	Pdf   bool                          `json:"pdf,required"`
	JSON  modelCapabilitiesModalityJSON `json:"-"`
}

type modelCapabilitiesModalityJSON struct {
	Text        apijson.Field
	Audio       apijson.Field
	Image       apijson.Field
	Video       apijson.Field
	Pdf         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelCapabilitiesModality) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCapabilitiesModalityJSON) RawJSON() string {
	return r.raw
}

type ModelCapabilitiesInterleaved struct {
	Field ModelCapabilitiesInterleavedField `json:"field"`
	JSON  modelCapabilitiesInterleavedJSON  `json:"-"`
	union ModelCapabilitiesInterleavedUnion
}

type modelCapabilitiesInterleavedJSON struct {
	Field       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r modelCapabilitiesInterleavedJSON) RawJSON() string {
	return r.raw
}

func (r *ModelCapabilitiesInterleaved) UnmarshalJSON(data []byte) (err error) {
	*r = ModelCapabilitiesInterleaved{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ModelCapabilitiesInterleaved) AsUnion() ModelCapabilitiesInterleavedUnion {
	return r.union
}

type ModelCapabilitiesInterleavedUnion interface {
	implementsModelCapabilitiesInterleaved()
}

type ModelCapabilitiesInterleavedBool bool

func (r ModelCapabilitiesInterleavedBool) implementsModelCapabilitiesInterleaved() {}

type ModelCapabilitiesInterleavedObject struct {
	Field ModelCapabilitiesInterleavedField      `json:"field,required"`
	JSON  modelCapabilitiesInterleavedObjectJSON `json:"-"`
}

type modelCapabilitiesInterleavedObjectJSON struct {
	Field       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelCapabilitiesInterleavedObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCapabilitiesInterleavedObjectJSON) RawJSON() string {
	return r.raw
}

func (r ModelCapabilitiesInterleavedObject) implementsModelCapabilitiesInterleaved() {}

type ModelCapabilitiesInterleavedField string

const (
	ModelCapabilitiesInterleavedFieldReasoningContent ModelCapabilitiesInterleavedField = "reasoning_content"
	ModelCapabilitiesInterleavedFieldReasoningDetails ModelCapabilitiesInterleavedField = "reasoning_details"
)

func (r ModelCapabilitiesInterleavedField) IsKnown() bool {
	switch r {
	case ModelCapabilitiesInterleavedFieldReasoningContent, ModelCapabilitiesInterleavedFieldReasoningDetails:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ModelCapabilitiesInterleavedUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(ModelCapabilitiesInterleavedBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(ModelCapabilitiesInterleavedBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ModelCapabilitiesInterleavedObject{}),
		},
	)
}

type ModelAPI struct {
	ID   string       `json:"id,required"`
	URL  string       `json:"url,required"`
	Npm  string       `json:"npm,required"`
	JSON modelAPIJSON `json:"-"`
}

type modelAPIJSON struct {
	ID          apijson.Field
	URL         apijson.Field
	Npm         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelAPI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelAPIJSON) RawJSON() string {
	return r.raw
}

type ModelStatus string

const (
	ModelStatusAlpha      ModelStatus = "alpha"
	ModelStatusBeta       ModelStatus = "beta"
	ModelStatusDeprecated ModelStatus = "deprecated"
	ModelStatusActive     ModelStatus = "active"
)

func (r ModelStatus) IsKnown() bool {
	switch r {
	case ModelStatusAlpha, ModelStatusBeta, ModelStatusDeprecated, ModelStatusActive:
		return true
	}
	return false
}

type ProviderSource string

const (
	ProviderSourceEnv    ProviderSource = "env"
	ProviderSourceConfig ProviderSource = "config"
	ProviderSourceCustom ProviderSource = "custom"
	ProviderSourceAPI    ProviderSource = "api"
)

func (r ProviderSource) IsKnown() bool {
	switch r {
	case ProviderSourceEnv, ProviderSourceConfig, ProviderSourceCustom, ProviderSourceAPI:
		return true
	}
	return false
}

type Provider struct {
	ID      string                 `json:"id,required"`
	Name    string                 `json:"name,required"`
	Source  ProviderSource         `json:"source,required"`
	Env     []string               `json:"env,required"`
	Options map[string]interface{} `json:"options,required"`
	Models  map[string]Model       `json:"models,required"`
	Key     string                 `json:"key"`
	JSON    providerJSON           `json:"-"`
}

// providerJSON contains the JSON metadata for the struct [Provider]
type providerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Env         apijson.Field
	Options     apijson.Field
	Models      apijson.Field
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Provider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerJSON) RawJSON() string {
	return r.raw
}

// Deprecated: use ConfigProvidersResponse.
type AppProvidersResponse = ConfigProvidersResponse

type AppLogParams struct {
	// Log level
	Level param.Field[AppLogParamsLevel] `json:"level,required"`
	// Log message
	Message param.Field[string] `json:"message,required"`
	// Service name for the log entry
	Service   param.Field[string] `json:"service,required"`
	Directory param.Field[string] `query:"directory"`
	// Additional metadata for the log entry
	Extra param.Field[map[string]interface{}] `json:"extra"`
}

func (r AppLogParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AppLogParams]'s query parameters as `url.Values`.
func (r AppLogParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Log level
type AppLogParamsLevel string

const (
	AppLogParamsLevelDebug AppLogParamsLevel = "debug"
	AppLogParamsLevelInfo  AppLogParamsLevel = "info"
	AppLogParamsLevelError AppLogParamsLevel = "error"
	AppLogParamsLevelWarn  AppLogParamsLevel = "warn"
)

func (r AppLogParamsLevel) IsKnown() bool {
	switch r {
	case AppLogParamsLevelDebug, AppLogParamsLevelInfo, AppLogParamsLevelError, AppLogParamsLevelWarn:
		return true
	}
	return false
}

type AppProvidersParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [AppProvidersParams]'s query parameters as `url.Values`.
func (r AppProvidersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
