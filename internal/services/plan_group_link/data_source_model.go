// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PlanGroupLinkDataSourceModel struct {
	ID             types.String      `tfsdk:"id" path:"id,required"`
	OrgID          types.String      `tfsdk:"org_id" path:"orgId,required"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	PlanGroupID    types.String      `tfsdk:"plan_group_id" json:"planGroupId,computed"`
	PlanID         types.String      `tfsdk:"plan_id" json:"planId,computed"`
	Version        types.Int64       `tfsdk:"version" json:"version,computed"`
}
