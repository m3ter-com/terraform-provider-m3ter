// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CustomFieldDataSourceModel struct {
	OrgID               types.String                   `tfsdk:"org_id" path:"orgId,required"`
	CreatedBy           types.String                   `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated           timetypes.RFC3339              `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339              `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	ID                  types.String                   `tfsdk:"id" json:"id,computed"`
	LastModifiedBy      types.String                   `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version             types.Int64                    `tfsdk:"version" json:"version,computed"`
	Account             customfield.Map[types.Dynamic] `tfsdk:"account" json:"account,computed"`
	AccountPlan         customfield.Map[types.Dynamic] `tfsdk:"account_plan" json:"accountPlan,computed"`
	Aggregation         customfield.Map[types.Dynamic] `tfsdk:"aggregation" json:"aggregation,computed"`
	CompoundAggregation customfield.Map[types.Dynamic] `tfsdk:"compound_aggregation" json:"compoundAggregation,computed"`
	Meter               customfield.Map[types.Dynamic] `tfsdk:"meter" json:"meter,computed"`
	Organization        customfield.Map[types.Dynamic] `tfsdk:"organization" json:"organization,computed"`
	Plan                customfield.Map[types.Dynamic] `tfsdk:"plan" json:"plan,computed"`
	PlanTemplate        customfield.Map[types.Dynamic] `tfsdk:"plan_template" json:"planTemplate,computed"`
	Product             customfield.Map[types.Dynamic] `tfsdk:"product" json:"product,computed"`
}

func (m *CustomFieldDataSourceModel) toReadParams(_ context.Context) (params m3ter.CustomFieldGetParams, diags diag.Diagnostics) {
	params = m3ter.CustomFieldGetParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
	}

	return
}
