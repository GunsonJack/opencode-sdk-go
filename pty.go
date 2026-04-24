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

// PtyService contains methods for interacting with the PTY resource.
type PtyService struct {
	Options []option.RequestOption
}

// NewPtyService generates a new service that applies the given options to each
// request.
func NewPtyService(opts ...option.RequestOption) (r *PtyService) {
	r = &PtyService{}
	r.Options = opts
	return
}

// List returns all PTY sessions.
func (r *PtyService) List(ctx context.Context, query PtyListParams, opts ...option.RequestOption) (res *[]Pty, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pty"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create creates a new PTY session.
func (r *PtyService) Create(ctx context.Context, params PtyCreateParams, opts ...option.RequestOption) (res *Pty, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pty"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get retrieves a PTY session by ID.
func (r *PtyService) Get(ctx context.Context, ptyID string, query PtyGetParams, opts ...option.RequestOption) (res *Pty, err error) {
	opts = slices.Concat(r.Options, opts)
	ptyID, err = requestconfig.EncodePathSegment(ptyID, "ptyID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update updates a PTY session.
func (r *PtyService) Update(ctx context.Context, ptyID string, params PtyUpdateParams, opts ...option.RequestOption) (res *Pty, err error) {
	opts = slices.Concat(r.Options, opts)
	ptyID, err = requestconfig.EncodePathSegment(ptyID, "ptyID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Remove deletes a PTY session.
func (r *PtyService) Remove(ctx context.Context, ptyID string, params PtyRemoveParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	ptyID, err = requestconfig.EncodePathSegment(ptyID, "ptyID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Pty represents a pseudo-terminal session.
type Pty struct {
	ID      string    `json:"id,required"`
	Title   string    `json:"title,required"`
	Command string    `json:"command,required"`
	Args    []string  `json:"args,required"`
	Cwd     string    `json:"cwd,required"`
	Status  PtyStatus `json:"status,required"`
	Pid     float64   `json:"pid,required"`
	JSON    ptyJSON   `json:"-"`
}

type ptyJSON struct {
	ID          apijson.Field
	Title       apijson.Field
	Command     apijson.Field
	Args        apijson.Field
	Cwd         apijson.Field
	Status      apijson.Field
	Pid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyJSON) RawJSON() string {
	return r.raw
}

type PtyStatus string

const (
	PtyStatusRunning PtyStatus = "running"
	PtyStatusExited  PtyStatus = "exited"
)

func (r PtyStatus) IsKnown() bool {
	switch r {
	case PtyStatusRunning, PtyStatusExited:
		return true
	}
	return false
}

// Params

type PtyListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r PtyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyCreateParams struct {
	Command   param.Field[string]            `json:"command"`
	Args      param.Field[[]string]          `json:"args"`
	Cwd       param.Field[string]            `json:"cwd"`
	Title     param.Field[string]            `json:"title"`
	Env       param.Field[map[string]string] `json:"env"`
	Directory param.Field[string]            `query:"directory"`
	Workspace param.Field[string]            `query:"workspace"`
}

func (r PtyCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PtyCreateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r PtyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyUpdateParams struct {
	Title     param.Field[string]        `json:"title"`
	Size      param.Field[PtyUpdateSize] `json:"size"`
	Directory param.Field[string]        `query:"directory"`
	Workspace param.Field[string]        `query:"workspace"`
}

func (r PtyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PtyUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyUpdateSize struct {
	Rows param.Field[float64] `json:"rows,required"`
	Cols param.Field[float64] `json:"cols,required"`
}

type PtyRemoveParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r PtyRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
