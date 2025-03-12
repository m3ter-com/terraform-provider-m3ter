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

type WebhookDataSourceModel struct {
ID types.String `tfsdk:"id" path:"id,required"`
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
Active types.Bool `tfsdk:"active" json:"active,computed"`
Code types.String `tfsdk:"code" json:"code,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
URL types.String `tfsdk:"url" json:"url,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
Credentials customfield.NestedObject[WebhookCredentialsDataSourceModel] `tfsdk:"credentials" json:"credentials,computed"`
}

func (m *WebhookDataSourceModel) toReadParams(_ context.Context) (params m3ter.WebhookGetParams, diags diag.Diagnostics) {
  params = m3ter.WebhookGetParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
  }

  return
}

type WebhookCredentialsDataSourceModel struct {
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
