// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type PlanGroupLinksDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[PlanGroupLinksItemsDataSourceModel] `json:"data,computed"`
}

type PlanGroupLinksDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
Plan types.String `tfsdk:"plan" query:"plan,optional"`
PlanGroup types.String `tfsdk:"plan_group" query:"planGroup,optional"`
IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[PlanGroupLinksItemsDataSourceModel] `tfsdk:"items"`
}

func (m *PlanGroupLinksDataSourceModel) toListParams(_ context.Context) (params m3ter.PlanGroupLinkListParams, diags diag.Diagnostics) {
  mIDs := []string{}
  for _, item := range *m.IDs {
    mIDs = append(mIDs, item.ValueString())
  }

  params = m3ter.PlanGroupLinkListParams{
    IDs: m3ter.F(mIDs),
  }

  if !m.Plan.IsNull() {
    params.Plan = m3ter.F(m.Plan.ValueString())
  }
  if !m.PlanGroup.IsNull() {
    params.PlanGroup = m3ter.F(m.PlanGroup.ValueString())
  }
  if !m.OrgID.IsNull() {
    params.OrgID = m3ter.F(m.OrgID.ValueString())
  }

  return
}

type PlanGroupLinksItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
PlanGroupID types.String `tfsdk:"plan_group_id" json:"planGroupId,computed"`
PlanID types.String `tfsdk:"plan_id" json:"planId,computed"`
}
