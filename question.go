// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
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

// QuestionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuestionService] method instead.
type QuestionService struct {
	Options []option.RequestOption
}

// NewQuestionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQuestionService(opts ...option.RequestOption) (r *QuestionService) {
	r = &QuestionService{}
	r.Options = opts
	return
}

// List pending questions
func (r *QuestionService) List(ctx context.Context, query QuestionListParams, opts ...option.RequestOption) (res *[]QuestionRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Reply to a question
func (r *QuestionService) Reply(ctx context.Context, requestID string, params QuestionReplyParams, opts ...option.RequestOption) (res *QuestionReplied, err error) {
	opts = slices.Concat(r.Options, opts)
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return
	}
	path := fmt.Sprintf("question/%s/reply", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Reject a question
func (r *QuestionService) Reject(ctx context.Context, requestID string, params QuestionRejectParams, opts ...option.RequestOption) (res *QuestionRejected, err error) {
	opts = slices.Concat(r.Options, opts)
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return
	}
	path := fmt.Sprintf("question/%s/reject", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type QuestionRequest struct {
	ID        string           `json:"id,required"`
	SessionID string           `json:"sessionID,required"`
	Info      QuestionInfo     `json:"info"`
	Options   []QuestionOption `json:"options"`
	Tool      QuestionTool     `json:"tool"`
	JSON      questionRequestJSON `json:"-"`
}

type questionRequestJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Info        apijson.Field
	Options     apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionRequestJSON) RawJSON() string {
	return r.raw
}

type QuestionInfo struct {
	Label       string           `json:"label"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	JSON        questionInfoJSON `json:"-"`
}

type questionInfoJSON struct {
	Label       apijson.Field
	Title       apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionInfoJSON) RawJSON() string {
	return r.raw
}

type QuestionOption struct {
	Label       string             `json:"label"`
	Value       string             `json:"value"`
	Description string             `json:"description"`
	JSON        questionOptionJSON `json:"-"`
}

type questionOptionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionOptionJSON) RawJSON() string {
	return r.raw
}

type QuestionTool struct {
	Name   string           `json:"name"`
	CallID string           `json:"callID"`
	State  string           `json:"state"`
	JSON   questionToolJSON `json:"-"`
}

type questionToolJSON struct {
	Name        apijson.Field
	CallID      apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionToolJSON) RawJSON() string {
	return r.raw
}

type QuestionAnswer struct {
	Value string             `json:"value"`
	JSON  questionAnswerJSON `json:"-"`
}

type questionAnswerJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionAnswer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionAnswerJSON) RawJSON() string {
	return r.raw
}

type QuestionReplied struct {
	RequestID string              `json:"requestID"`
	Answer    QuestionAnswer      `json:"answer"`
	JSON      questionRepliedJSON `json:"-"`
}

type questionRepliedJSON struct {
	RequestID   apijson.Field
	Answer      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionReplied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionRepliedJSON) RawJSON() string {
	return r.raw
}

type QuestionRejected struct {
	RequestID string               `json:"requestID"`
	Reason    string               `json:"reason"`
	JSON      questionRejectedJSON `json:"-"`
}

type questionRejectedJSON struct {
	RequestID   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionRejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionRejectedJSON) RawJSON() string {
	return r.raw
}

type QuestionListParams struct {
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

// URLQuery serializes [QuestionListParams]'s query parameters as `url.Values`.
func (r QuestionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuestionReplyParams struct {
	Value     param.Field[string] `json:"value,required"`
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

func (r QuestionReplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [QuestionReplyParams]'s query parameters as `url.Values`.
func (r QuestionReplyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuestionRejectParams struct {
	Reason    param.Field[string] `json:"reason"`
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

func (r QuestionRejectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [QuestionRejectParams]'s query parameters as `url.Values`.
func (r QuestionRejectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
