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

type SkillService struct {
	Options []option.RequestOption
}

func NewSkillService(opts ...option.RequestOption) (r *SkillService) {
	r = &SkillService{}
	r.Options = opts
	return
}

func (r *SkillService) List(ctx context.Context, query SkillListParams, opts ...option.RequestOption) (res *[]SkillItem, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skill"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SkillItem struct {
	Name        string        `json:"name,required"`
	Description string        `json:"description,required"`
	Location    string        `json:"location,required"`
	Content     string        `json:"content,required"`
	JSON        skillItemJSON `json:"-"`
}

type skillItemJSON struct {
	Name        apijson.Field
	Description apijson.Field
	Location    apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillItemJSON) RawJSON() string { return r.raw }

type SkillListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r SkillListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
