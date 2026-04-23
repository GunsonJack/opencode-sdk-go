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
func (r *QuestionService) Reply(ctx context.Context, requestID string, params QuestionReplyParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	requestID, err = requestconfig.EncodePathSegment(requestID, "requestID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("question/%s/reply", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Reject a question
func (r *QuestionService) Reject(ctx context.Context, requestID string, params QuestionRejectParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	requestID, err = requestconfig.EncodePathSegment(requestID, "requestID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("question/%s/reject", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type QuestionRequest struct {
	ID        string              `json:"id,required"`
	SessionID string              `json:"sessionID,required"`
	Questions []QuestionInfo      `json:"questions,required"`
	Tool      QuestionTool        `json:"tool"`
	JSON      questionRequestJSON `json:"-"`
}

type questionRequestJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Questions   apijson.Field
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
	Question string           `json:"question,required"`
	Header   string           `json:"header,required"`
	Options  []QuestionOption `json:"options,required"`
	Multiple bool             `json:"multiple"`
	Custom   bool             `json:"custom"`
	JSON     questionInfoJSON `json:"-"`
}

type questionInfoJSON struct {
	Question    apijson.Field
	Header      apijson.Field
	Options     apijson.Field
	Multiple    apijson.Field
	Custom      apijson.Field
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
	Label       string             `json:"label,required"`
	Description string             `json:"description,required"`
	JSON        questionOptionJSON `json:"-"`
}

type questionOptionJSON struct {
	Label       apijson.Field
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
	MessageID string           `json:"messageID,required"`
	CallID    string           `json:"callID,required"`
	JSON      questionToolJSON `json:"-"`
}

type questionToolJSON struct {
	MessageID   apijson.Field
	CallID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionToolJSON) RawJSON() string {
	return r.raw
}

type QuestionAnswer = []string

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
	Answers   param.Field[[]QuestionAnswer] `json:"answers,required"`
	Workspace param.Field[string]           `query:"workspace"`
	Directory param.Field[string]           `query:"directory"`
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
	Workspace param.Field[string] `query:"workspace"`
	Directory param.Field[string] `query:"directory"`
}

// URLQuery serializes [QuestionRejectParams]'s query parameters as `url.Values`.
func (r QuestionRejectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
