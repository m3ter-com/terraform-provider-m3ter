// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type CustomFieldDataSourceModel struct {
	OrgID               types.String      `tfsdk:"org_id" path:"orgId,required"`
	CreatedBy           types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated           timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	LastModifiedBy      types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version             types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	Account             types.Dynamic     `tfsdk:"account" json:"account,computed"`
	AccountPlan         types.Dynamic     `tfsdk:"account_plan" json:"accountPlan,computed"`
	Aggregation         types.Dynamic     `tfsdk:"aggregation" json:"aggregation,computed"`
	CompoundAggregation types.Dynamic     `tfsdk:"compound_aggregation" json:"compoundAggregation,computed"`
	Contract            types.Dynamic     `tfsdk:"contract" json:"contract,computed"`
	Meter               types.Dynamic     `tfsdk:"meter" json:"meter,computed"`
	Organization        types.Dynamic     `tfsdk:"organization" json:"organization,computed"`
	Plan                types.Dynamic     `tfsdk:"plan" json:"plan,computed"`
	PlanTemplate        types.Dynamic     `tfsdk:"plan_template" json:"planTemplate,computed"`
	Product             types.Dynamic     `tfsdk:"product" json:"product,computed"`
}

func (m *CustomFieldDataSourceModel) toReadParams(_ context.Context) (params m3ter.CustomFieldGetParams, diags diag.Diagnostics) {
	params = m3ter.CustomFieldGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
