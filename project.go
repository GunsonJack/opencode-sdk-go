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

// ProjectService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r *ProjectService) {
	r = &ProjectService{}
	r.Options = opts
	return
}

// List all projects
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *[]Project, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "project"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the current project
func (r *ProjectService) Current(ctx context.Context, query ProjectCurrentParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "project/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a project
func (r *ProjectService) Update(ctx context.Context, projectID string, params ProjectUpdateParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.Options, opts)
	projectID, err = requestconfig.EncodePathSegment(projectID, "projectID")
	if err != nil {
		return
	}
	path := fmt.Sprintf("project/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Initialize git for a project
func (r *ProjectService) InitGit(ctx context.Context, params ProjectInitGitParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "project/git/init"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ProjectIcon represents a project's icon configuration.
type ProjectIcon struct {
	URL      string          `json:"url"`
	Override string          `json:"override"`
	Color    string          `json:"color"`
	JSON     projectIconJSON `json:"-"`
}

type projectIconJSON struct {
	URL         apijson.Field
	Override    apijson.Field
	Color       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectIcon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectIconJSON) RawJSON() string {
	return r.raw
}

// ProjectCommands holds the project's command configuration.
type ProjectCommands struct {
	Start string              `json:"start"`
	JSON  projectCommandsJSON `json:"-"`
}

type projectCommandsJSON struct {
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectCommands) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCommandsJSON) RawJSON() string {
	return r.raw
}

// ProjectSummary is a lightweight project reference.
type ProjectSummary struct {
	ID       string             `json:"id,required"`
	Worktree string             `json:"worktree,required"`
	Name     string             `json:"name"`
	JSON     projectSummaryJSON `json:"-"`
}

type projectSummaryJSON struct {
	ID          apijson.Field
	Worktree    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectSummaryJSON) RawJSON() string {
	return r.raw
}

type Project struct {
	ID        string          `json:"id,required"`
	Time      ProjectTime     `json:"time,required"`
	Worktree  string          `json:"worktree,required"`
	Sandboxes []string        `json:"sandboxes,required"`
	Name      string          `json:"name"`
	Icon      ProjectIcon     `json:"icon"`
	Commands  ProjectCommands `json:"commands"`
	Vcs       ProjectVcs      `json:"vcs"`
	JSON      projectJSON     `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Worktree    apijson.Field
	Sandboxes   apijson.Field
	Name        apijson.Field
	Icon        apijson.Field
	Commands    apijson.Field
	Vcs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectJSON) RawJSON() string {
	return r.raw
}

type ProjectTime struct {
	Created     float64         `json:"created,required"`
	Updated     float64         `json:"updated,required"`
	Initialized float64         `json:"initialized"`
	JSON        projectTimeJSON `json:"-"`
}

// projectTimeJSON contains the JSON metadata for the struct [ProjectTime]
type projectTimeJSON struct {
	Created     apijson.Field
	Updated     apijson.Field
	Initialized apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectTimeJSON) RawJSON() string {
	return r.raw
}

type ProjectVcs string

const (
	ProjectVcsGit ProjectVcs = "git"
)

func (r ProjectVcs) IsKnown() bool {
	switch r {
	case ProjectVcsGit:
		return true
	}
	return false
}

type ProjectListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectCurrentParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ProjectCurrentParams]'s query parameters as `url.Values`.
func (r ProjectCurrentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectUpdateParams struct {
	Name      param.Field[string]                      `json:"name"`
	Icon      param.Field[ProjectUpdateParamsIcon]     `json:"icon"`
	Commands  param.Field[ProjectUpdateParamsCommands] `json:"commands"`
	Directory param.Field[string]                      `query:"directory"`
	Workspace param.Field[string]                      `query:"workspace"`
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectUpdateParamsIcon struct {
	URL      param.Field[string] `json:"url"`
	Override param.Field[string] `json:"override"`
	Color    param.Field[string] `json:"color"`
}

type ProjectUpdateParamsCommands struct {
	Start param.Field[string] `json:"start"`
}

type ProjectInitGitParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ProjectInitGitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
