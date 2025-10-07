// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CustomFieldModel struct {
	ID                  types.String                       `tfsdk:"id" json:"id,computed"`
	OrgID               types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	Account             customfield.NormalizedDynamicValue `tfsdk:"account" json:"account,optional"`
	AccountPlan         customfield.NormalizedDynamicValue `tfsdk:"account_plan" json:"accountPlan,optional"`
	Aggregation         customfield.NormalizedDynamicValue `tfsdk:"aggregation" json:"aggregation,optional"`
	CompoundAggregation customfield.NormalizedDynamicValue `tfsdk:"compound_aggregation" json:"compoundAggregation,optional"`
	Contract            customfield.NormalizedDynamicValue `tfsdk:"contract" json:"contract,optional"`
	Meter               customfield.NormalizedDynamicValue `tfsdk:"meter" json:"meter,optional"`
	Organization        customfield.NormalizedDynamicValue `tfsdk:"organization" json:"organization,optional"`
	Plan                customfield.NormalizedDynamicValue `tfsdk:"plan" json:"plan,optional"`
	PlanTemplate        customfield.NormalizedDynamicValue `tfsdk:"plan_template" json:"planTemplate,optional"`
	Product             customfield.NormalizedDynamicValue `tfsdk:"product" json:"product,optional"`
	Version             types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m CustomFieldModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomFieldModel) MarshalJSONForUpdate(state CustomFieldModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
