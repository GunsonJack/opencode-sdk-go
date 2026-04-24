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

// TuiService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTuiService] method instead.
type TuiService struct {
	Options []option.RequestOption
	Control *TuiControlService
}

// NewTuiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTuiService(opts ...option.RequestOption) (r *TuiService) {
	r = &TuiService{}
	r.Options = opts
	r.Control = NewTuiControlService(opts...)
	return
}

// Append prompt to the TUI
func (r *TuiService) AppendPrompt(ctx context.Context, params TuiAppendPromptParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/append-prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Clear the prompt
func (r *TuiService) ClearPrompt(ctx context.Context, body TuiClearPromptParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/clear-prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute a TUI command (e.g. agent_cycle)
func (r *TuiService) ExecuteCommand(ctx context.Context, params TuiExecuteCommandParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/execute-command"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// OpenHelp opens the help dialog
func (r *TuiService) OpenHelp(ctx context.Context, body TuiOpenHelpParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/open-help"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// OpenModels opens the model dialog
func (r *TuiService) OpenModels(ctx context.Context, body TuiOpenModelsParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/open-models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// OpenSessions opens the session dialog
func (r *TuiService) OpenSessions(ctx context.Context, body TuiOpenSessionsParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/open-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// OpenThemes opens the theme dialog
func (r *TuiService) OpenThemes(ctx context.Context, body TuiOpenThemesParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/open-themes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Publish publishes an event to the TUI.
func (r *TuiService) Publish(ctx context.Context, params TuiPublishParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/publish"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// SelectSession selects a session in the TUI.
func (r *TuiService) SelectSession(ctx context.Context, params TuiSelectSessionParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/select-session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Show a toast notification in the TUI
func (r *TuiService) ShowToast(ctx context.Context, params TuiShowToastParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/show-toast"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Submit the prompt
func (r *TuiService) SubmitPrompt(ctx context.Context, body TuiSubmitPromptParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/submit-prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TuiAppendPromptParams struct {
	Text      param.Field[string] `json:"text,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r TuiAppendPromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [TuiAppendPromptParams]'s query parameters as `url.Values`.
func (r TuiAppendPromptParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiClearPromptParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiClearPromptParams]'s query parameters as `url.Values`.
func (r TuiClearPromptParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiExecuteCommandParams struct {
	// The command to execute. This is a free-form string, not restricted to the
	// TuiCommand enum (which only applies to TUI publish events).
	Command   param.Field[string] `json:"command,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r TuiExecuteCommandParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [TuiExecuteCommandParams]'s query parameters as
// `url.Values`.
func (r TuiExecuteCommandParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiOpenHelpParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiOpenHelpParams]'s query parameters as `url.Values`.
func (r TuiOpenHelpParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiOpenModelsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiOpenModelsParams]'s query parameters as `url.Values`.
func (r TuiOpenModelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiOpenSessionsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiOpenSessionsParams]'s query parameters as `url.Values`.
func (r TuiOpenSessionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiOpenThemesParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiOpenThemesParams]'s query parameters as `url.Values`.
func (r TuiOpenThemesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiPublishParams struct {
	Body      param.Field[TuiPublishBody] `json:"body"`
	Directory param.Field[string]          `query:"directory"`
	Workspace param.Field[string]          `query:"workspace"`
}

func (r TuiPublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

func (r TuiPublishParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// TuiCommand represents one of the 16 known TUI commands.
type TuiCommand string

const (
	TuiCommandSessionList         TuiCommand = "session.list"
	TuiCommandSessionNew          TuiCommand = "session.new"
	TuiCommandSessionShare        TuiCommand = "session.share"
	TuiCommandSessionInterrupt    TuiCommand = "session.interrupt"
	TuiCommandSessionCompact      TuiCommand = "session.compact"
	TuiCommandSessionPageUp       TuiCommand = "session.page.up"
	TuiCommandSessionPageDown     TuiCommand = "session.page.down"
	TuiCommandSessionLineUp       TuiCommand = "session.line.up"
	TuiCommandSessionLineDown     TuiCommand = "session.line.down"
	TuiCommandSessionHalfPageUp   TuiCommand = "session.half.page.up"
	TuiCommandSessionHalfPageDown TuiCommand = "session.half.page.down"
	TuiCommandSessionFirst        TuiCommand = "session.first"
	TuiCommandSessionLast         TuiCommand = "session.last"
	TuiCommandPromptClear         TuiCommand = "prompt.clear"
	TuiCommandPromptSubmit        TuiCommand = "prompt.submit"
	TuiCommandAgentCycle          TuiCommand = "agent.cycle"
)

func (r TuiCommand) IsKnown() bool {
	switch r {
	case TuiCommandSessionList, TuiCommandSessionNew, TuiCommandSessionShare,
		TuiCommandSessionInterrupt, TuiCommandSessionCompact,
		TuiCommandSessionPageUp, TuiCommandSessionPageDown,
		TuiCommandSessionLineUp, TuiCommandSessionLineDown,
		TuiCommandSessionHalfPageUp, TuiCommandSessionHalfPageDown,
		TuiCommandSessionFirst, TuiCommandSessionLast,
		TuiCommandPromptClear, TuiCommandPromptSubmit,
		TuiCommandAgentCycle:
		return true
	}
	return false
}

// TuiPublishBody is the union type for TUI publish event bodies.
// Satisfied by [TuiPublishBodyPromptAppend], [TuiPublishBodyCommandExecute],
// [TuiPublishBodyToastShow], or [TuiPublishBodySessionSelect].
type TuiPublishBody interface {
	implementsTuiPublishBody()
}

type TuiPublishBodyPromptAppend struct {
	Type       param.Field[string]                               `json:"type,required"` // always "tui.prompt.append"
	Properties param.Field[TuiPublishBodyPromptAppendProperties] `json:"properties,required"`
}

func (r TuiPublishBodyPromptAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyPromptAppend) implementsTuiPublishBody() {}

type TuiPublishBodyPromptAppendProperties struct {
	Text param.Field[string] `json:"text,required"`
}

func (r TuiPublishBodyPromptAppendProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodyCommandExecute struct {
	Type       param.Field[string]                                 `json:"type,required"` // always "tui.command.execute"
	Properties param.Field[TuiPublishBodyCommandExecuteProperties] `json:"properties,required"`
}

func (r TuiPublishBodyCommandExecute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyCommandExecute) implementsTuiPublishBody() {}

type TuiPublishBodyCommandExecuteProperties struct {
	Command param.Field[TuiCommand] `json:"command,required"`
}

func (r TuiPublishBodyCommandExecuteProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodyToastShow struct {
	Type       param.Field[string]                            `json:"type,required"` // always "tui.toast.show"
	Properties param.Field[TuiPublishBodyToastShowProperties] `json:"properties,required"`
}

func (r TuiPublishBodyToastShow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyToastShow) implementsTuiPublishBody() {}

type TuiPublishBodyToastShowProperties struct {
	Message  param.Field[string]                    `json:"message,required"`
	Variant  param.Field[TuiShowToastParamsVariant] `json:"variant,required"`
	Title    param.Field[string]                    `json:"title"`
	Duration param.Field[float64]                   `json:"duration"`
}

func (r TuiPublishBodyToastShowProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodySessionSelect struct {
	Type       param.Field[string]                                `json:"type,required"` // always "tui.session.select"
	Properties param.Field[TuiPublishBodySessionSelectProperties] `json:"properties,required"`
}

func (r TuiPublishBodySessionSelect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodySessionSelect) implementsTuiPublishBody() {}

type TuiPublishBodySessionSelectProperties struct {
	SessionID param.Field[string] `json:"sessionID,required"`
}

func (r TuiPublishBodySessionSelectProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiSelectSessionParams struct {
	SessionID param.Field[string] `json:"sessionID,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r TuiSelectSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiSelectSessionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiShowToastParams struct {
	Message   param.Field[string]                    `json:"message,required"`
	Variant   param.Field[TuiShowToastParamsVariant] `json:"variant,required"`
	Directory param.Field[string]                    `query:"directory"`
	Workspace param.Field[string]                    `query:"workspace"`
	Duration  param.Field[float64]                   `json:"duration"`
	Title     param.Field[string]                    `json:"title"`
}

func (r TuiShowToastParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [TuiShowToastParams]'s query parameters as `url.Values`.
func (r TuiShowToastParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiShowToastParamsVariant string

const (
	TuiShowToastParamsVariantInfo    TuiShowToastParamsVariant = "info"
	TuiShowToastParamsVariantSuccess TuiShowToastParamsVariant = "success"
	TuiShowToastParamsVariantWarning TuiShowToastParamsVariant = "warning"
	TuiShowToastParamsVariantError   TuiShowToastParamsVariant = "error"
)

func (r TuiShowToastParamsVariant) IsKnown() bool {
	switch r {
	case TuiShowToastParamsVariantInfo, TuiShowToastParamsVariantSuccess, TuiShowToastParamsVariantWarning, TuiShowToastParamsVariantError:
		return true
	}
	return false
}

type TuiSubmitPromptParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiSubmitPromptParams]'s query parameters as `url.Values`.
func (r TuiSubmitPromptParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// TuiControlService contains methods for interacting with the TUI control resource.
type TuiControlService struct {
	Options []option.RequestOption
}

// NewTuiControlService generates a new service.
func NewTuiControlService(opts ...option.RequestOption) (r *TuiControlService) {
	r = &TuiControlService{}
	r.Options = opts
	return
}

// Next gets the next pending control request from the TUI.
func (r *TuiControlService) Next(ctx context.Context, query TuiControlNextParams, opts ...option.RequestOption) (res *TuiControlNextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/control/next"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response sends a response to a TUI control request.
func (r *TuiControlService) Response(ctx context.Context, params TuiControlResponseParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/control/response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// TuiControlNextResponse is the response from GET /tui/control/next.
type TuiControlNextResponse struct {
	Path string      `json:"path,required"`
	Body interface{} `json:"body,required"`
	JSON tuiControlNextResponseJSON `json:"-"`
}

type tuiControlNextResponseJSON struct {
	Path        apijson.Field
	Body        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TuiControlNextResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tuiControlNextResponseJSON) RawJSON() string {
	return r.raw
}

type TuiControlNextParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r TuiControlNextParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiControlResponseParams struct {
	Body      param.Field[interface{}] `json:"body"`
	Directory param.Field[string]      `query:"directory"`
	Workspace param.Field[string]      `query:"workspace"`
}

func (r TuiControlResponseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

func (r TuiControlResponseParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
