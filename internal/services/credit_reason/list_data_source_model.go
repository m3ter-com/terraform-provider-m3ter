// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package credit_reason

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CreditReasonsDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[CreditReasonsItemsDataSourceModel] `json:"data,computed"`
}

type CreditReasonsDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
Archived types.Bool `tfsdk:"archived" query:"archived,optional"`
Codes *[]types.String `tfsdk:"codes" query:"codes,optional"`
IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[CreditReasonsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CreditReasonsDataSourceModel) toListParams(_ context.Context) (params m3ter.CreditReasonListParams, diags diag.Diagnostics) {
  mCodes := []string{}
  for _, item := range *m.Codes {
    mCodes = append(mCodes, item.ValueString())
  }
  mIDs := []string{}
  for _, item := range *m.IDs {
    mIDs = append(mIDs, item.ValueString())
  }

  params = m3ter.CreditReasonListParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
    Codes: m3ter.F(mCodes),
    IDs: m3ter.F(mIDs),
  }

  if !m.Archived.IsNull() {
    params.Archived = m3ter.F(m.Archived.ValueBool())
  }

  return
}

type CreditReasonsItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
Archived types.Bool `tfsdk:"archived" json:"archived,computed"`
Code types.String `tfsdk:"code" json:"code,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}
