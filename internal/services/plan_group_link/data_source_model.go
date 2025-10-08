// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type PlanGroupLinkDataSourceModel struct {
	ID          types.String                           `tfsdk:"id" path:"id,computed_optional"`
	OrgID       types.String                           `tfsdk:"org_id" path:"orgId,optional"`
	PlanGroupID types.String                           `tfsdk:"plan_group_id" json:"planGroupId,computed"`
	PlanID      types.String                           `tfsdk:"plan_id" json:"planId,computed"`
	Version     types.Int64                            `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	FindOneBy   *PlanGroupLinkFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *PlanGroupLinkDataSourceModel) toReadParams(_ context.Context) (params m3ter.PlanGroupLinkGetParams, diags diag.Diagnostics) {
	params = m3ter.PlanGroupLinkGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *PlanGroupLinkDataSourceModel) toListParams(_ context.Context) (params m3ter.PlanGroupLinkListParams, diags diag.Diagnostics) {
	mFindOneByIDs := []string{}
	if m.FindOneBy.IDs != nil {
		for _, item := range *m.FindOneBy.IDs {
			mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
		}
	}

	params = m3ter.PlanGroupLinkListParams{
		IDs: m3ter.F(mFindOneByIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.FindOneBy.Plan.IsNull() {
		params.Plan = m3ter.F(m.FindOneBy.Plan.ValueString())
	}
	if !m.FindOneBy.PlanGroup.IsNull() {
		params.PlanGroup = m3ter.F(m.FindOneBy.PlanGroup.ValueString())
	}

	return
}

type PlanGroupLinkFindOneByDataSourceModel struct {
	IDs       *[]types.String `tfsdk:"ids" query:"ids,optional"`
	Plan      types.String    `tfsdk:"plan" query:"plan,optional"`
	PlanGroup types.String    `tfsdk:"plan_group" query:"planGroup,optional"`
}
