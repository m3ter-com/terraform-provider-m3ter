// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type WebhooksDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[WebhooksItemsDataSourceModel] `json:"data,computed"`
}

type WebhooksDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[WebhooksItemsDataSourceModel] `tfsdk:"items"`
}

func (m *WebhooksDataSourceModel) toListParams(_ context.Context) (params m3ter.WebhookListParams, diags diag.Diagnostics) {
  mIDs := []string{}
  for _, item := range *m.IDs {
    mIDs = append(mIDs, item.ValueString())
  }

  params = m3ter.WebhookListParams{
    IDs: m3ter.F(mIDs),
  }

  if !m.OrgID.IsNull() {
    params.OrgID = m3ter.F(m.OrgID.ValueString())
  }

  return
}

type WebhooksItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
Active types.Bool `tfsdk:"active" json:"active,computed"`
Code types.String `tfsdk:"code" json:"code,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
Credentials customfield.NestedObject[WebhooksCredentialsDataSourceModel] `tfsdk:"credentials" json:"credentials,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
URL types.String `tfsdk:"url" json:"url,computed"`
}

type WebhooksCredentialsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Destination types.String `tfsdk:"destination" json:"destination,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
APIKey types.String `tfsdk:"api_key" json:"apiKey,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DestinationID types.String `tfsdk:"destination_id" json:"destinationId,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Secret types.String `tfsdk:"secret" json:"secret,computed"`
}
