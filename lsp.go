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

// LspService contains methods for interacting with the LSP resource.
type LspService struct {
	Options []option.RequestOption
}

// NewLspService generates a new service that applies the given options to each
// request.
func NewLspService(opts ...option.RequestOption) (r *LspService) {
	r = &LspService{}
	r.Options = opts
	return
}

// Status returns the status of all LSP servers.
func (r *LspService) Status(ctx context.Context, query LspStatusParams, opts ...option.RequestOption) (res *[]LspStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lsp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// LspStatus describes the status of a single LSP server.
type LspStatus struct {
	ID     string          `json:"id,required"`
	Name   string          `json:"name,required"`
	Root   string          `json:"root,required"`
	Status LspStatusStatus `json:"status,required"`
	JSON   lspStatusJSON   `json:"-"`
}

// lspStatusJSON contains the JSON metadata for the struct [LspStatus]
type lspStatusJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Root        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LspStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lspStatusJSON) RawJSON() string {
	return r.raw
}

type LspStatusStatus string

const (
	LspStatusStatusConnected LspStatusStatus = "connected"
	LspStatusStatusError     LspStatusStatus = "error"
)

func (r LspStatusStatus) IsKnown() bool {
	switch r {
	case LspStatusStatusConnected, LspStatusStatusError:
		return true
	}
	return false
}

type LspStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [LspStatusParams]'s query parameters as `url.Values`.
func (r LspStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
