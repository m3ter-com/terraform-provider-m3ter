// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type CustomFieldModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	OrgID               types.String      `tfsdk:"org_id" path:"orgId,optional"`
	Account             types.Dynamic     `tfsdk:"account" json:"account,optional"`
	AccountPlan         types.Dynamic     `tfsdk:"account_plan" json:"accountPlan,optional"`
	Aggregation         types.Dynamic     `tfsdk:"aggregation" json:"aggregation,optional"`
	CompoundAggregation types.Dynamic     `tfsdk:"compound_aggregation" json:"compoundAggregation,optional"`
	Contract            types.Dynamic     `tfsdk:"contract" json:"contract,optional"`
	Meter               types.Dynamic     `tfsdk:"meter" json:"meter,optional"`
	Organization        types.Dynamic     `tfsdk:"organization" json:"organization,optional"`
	Plan                types.Dynamic     `tfsdk:"plan" json:"plan,optional"`
	PlanTemplate        types.Dynamic     `tfsdk:"plan_template" json:"planTemplate,optional"`
	Product             types.Dynamic     `tfsdk:"product" json:"product,optional"`
	CreatedBy           types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated           timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy      types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version             types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m CustomFieldModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomFieldModel) MarshalJSONForUpdate(state CustomFieldModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
