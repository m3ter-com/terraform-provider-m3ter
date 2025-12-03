// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type WebhookDataSourceModel struct {
	ID          types.String                                                `tfsdk:"id" path:"id,computed_optional"`
	OrgID       types.String                                                `tfsdk:"org_id" path:"orgId,optional"`
	Active      types.Bool                                                  `tfsdk:"active" json:"active,computed"`
	Code        types.String                                                `tfsdk:"code" json:"code,computed"`
	Description types.String                                                `tfsdk:"description" json:"description,computed"`
	Name        types.String                                                `tfsdk:"name" json:"name,computed"`
	URL         types.String                                                `tfsdk:"url" json:"url,computed"`
	Version     types.Int64                                                 `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	Credentials customfield.NestedObject[WebhookCredentialsDataSourceModel] `tfsdk:"credentials" json:"credentials,computed"`
	FindOneBy   *WebhookFindOneByDataSourceModel                            `tfsdk:"find_one_by"`
}

func (m *WebhookDataSourceModel) toReadParams(_ context.Context) (params m3ter.WebhookGetParams, diags diag.Diagnostics) {
	params = m3ter.WebhookGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *WebhookDataSourceModel) toListParams(_ context.Context) (params m3ter.WebhookListParams, diags diag.Diagnostics) {
	mFindOneByIDs := []string{}
	if m.FindOneBy.IDs != nil {
		for _, item := range *m.FindOneBy.IDs {
			mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
		}
	}

	params = m3ter.WebhookListParams{
		IDs: m3ter.F(mFindOneByIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type WebhookCredentialsDataSourceModel struct {
	ID            types.String `tfsdk:"id" json:"id,computed"`
	Destination   types.String `tfsdk:"destination" json:"destination,computed"`
	Type          types.String `tfsdk:"type" json:"type,computed"`
	APIKey        types.String `tfsdk:"api_key" json:"apiKey,computed"`
	DestinationID types.String `tfsdk:"destination_id" json:"destinationId,computed"`
	Name          types.String `tfsdk:"name" json:"name,computed"`
	Secret        types.String `tfsdk:"secret" json:"secret,computed"`
	Version       types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

type WebhookFindOneByDataSourceModel struct {
	IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
}
