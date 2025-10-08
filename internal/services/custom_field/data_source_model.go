// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CustomFieldDataSourceModel struct {
	ID                  types.String                       `tfsdk:"id" path:"orgId,computed"`
	OrgID               types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	Version             types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	Account             customfield.NormalizedDynamicValue `tfsdk:"account" json:"account,computed"`
	AccountPlan         customfield.NormalizedDynamicValue `tfsdk:"account_plan" json:"accountPlan,computed"`
	Aggregation         customfield.NormalizedDynamicValue `tfsdk:"aggregation" json:"aggregation,computed"`
	CompoundAggregation customfield.NormalizedDynamicValue `tfsdk:"compound_aggregation" json:"compoundAggregation,computed"`
	Contract            customfield.NormalizedDynamicValue `tfsdk:"contract" json:"contract,computed"`
	Meter               customfield.NormalizedDynamicValue `tfsdk:"meter" json:"meter,computed"`
	Organization        customfield.NormalizedDynamicValue `tfsdk:"organization" json:"organization,computed"`
	Plan                customfield.NormalizedDynamicValue `tfsdk:"plan" json:"plan,computed"`
	PlanTemplate        customfield.NormalizedDynamicValue `tfsdk:"plan_template" json:"planTemplate,computed"`
	Product             customfield.NormalizedDynamicValue `tfsdk:"product" json:"product,computed"`
}

func (m *CustomFieldDataSourceModel) toReadParams(_ context.Context) (params m3ter.CustomFieldGetParams, diags diag.Diagnostics) {
	params = m3ter.CustomFieldGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
