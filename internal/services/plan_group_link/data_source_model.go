// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type PlanGroupLinkDataSourceModel struct {
	ID          types.String `tfsdk:"id" path:"id,required"`
	OrgID       types.String `tfsdk:"org_id" path:"orgId,required"`
	PlanGroupID types.String `tfsdk:"plan_group_id" json:"planGroupId,computed"`
	PlanID      types.String `tfsdk:"plan_id" json:"planId,computed"`
	Version     types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m *PlanGroupLinkDataSourceModel) toReadParams(_ context.Context) (params m3ter.PlanGroupLinkGetParams, diags diag.Diagnostics) {
	params = m3ter.PlanGroupLinkGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
