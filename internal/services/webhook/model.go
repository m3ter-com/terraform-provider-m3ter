// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type WebhookModel struct {
	ID             types.String             `tfsdk:"id" path:"id,optional"`
	OrgID          types.String             `tfsdk:"org_id" path:"orgId,optional"`
	Description    types.String             `tfsdk:"description" json:"description,required"`
	Name           types.String             `tfsdk:"name" json:"name,required"`
	URL            types.String             `tfsdk:"url" json:"url,required"`
	Credentials    *WebhookCredentialsModel `tfsdk:"credentials" json:"credentials,required"`
	Active         types.Bool               `tfsdk:"active" json:"active,optional"`
	Code           types.String             `tfsdk:"code" json:"code,optional"`
	Version        types.Int64              `tfsdk:"version" json:"version,optional"`
	CreatedBy      types.String             `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339        `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339        `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String             `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m WebhookModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WebhookModel) MarshalJSONForUpdate(state WebhookModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type WebhookCredentialsModel struct {
	APIKey  types.String `tfsdk:"api_key" json:"apiKey,required"`
	Secret  types.String `tfsdk:"secret" json:"secret,required"`
	Type    types.String `tfsdk:"type" json:"type,required"`
	Empty   types.Bool   `tfsdk:"empty" json:"empty,optional"`
	Version types.Int64  `tfsdk:"version" json:"version,optional"`
}
