// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type PlanGroupLinkModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
PlanGroupID types.String `tfsdk:"plan_group_id" json:"planGroupId,required"`
PlanID types.String `tfsdk:"plan_id" json:"planId,required"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m PlanGroupLinkModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m PlanGroupLinkModel) MarshalJSONForUpdate(state PlanGroupLinkModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
