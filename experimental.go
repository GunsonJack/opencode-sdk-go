// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import "github.com/GunsonJack/opencode-sdk-go/option"

// ExperimentalService groups experimental API namespaces.
type ExperimentalService struct {
	Options   []option.RequestOption
	Resource  *ResourceService
	Session   *ExperimentalSessionService
	Console   *ConsoleService
	Workspace *WorkspaceService
	Worktree  *WorktreeService
	Tool      *ToolService
}

// NewExperimentalService generates a new service that applies the given options to each request.
func NewExperimentalService(opts ...option.RequestOption) (r *ExperimentalService) {
	r = &ExperimentalService{}
	r.Options = opts
	r.Resource = NewResourceService(opts...)
	r.Session = NewExperimentalSessionService(opts...)
	r.Console = NewConsoleService(opts...)
	r.Workspace = NewWorkspaceService(opts...)
	r.Worktree = NewWorktreeService(opts...)
	r.Tool = NewToolService(opts...)
	return
}
