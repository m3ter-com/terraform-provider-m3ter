// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_template

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type PlanTemplateDataSourceModel struct {
	ID                          types.String                          `tfsdk:"id" path:"id,computed_optional"`
	OrgID                       types.String                          `tfsdk:"org_id" path:"orgId,optional"`
	BillFrequency               types.String                          `tfsdk:"bill_frequency" json:"billFrequency,computed"`
	BillFrequencyInterval       types.Int64                           `tfsdk:"bill_frequency_interval" json:"billFrequencyInterval,computed"`
	Code                        types.String                          `tfsdk:"code" json:"code,computed"`
	Currency                    types.String                          `tfsdk:"currency" json:"currency,computed"`
	MinimumSpend                types.Float64                         `tfsdk:"minimum_spend" json:"minimumSpend,computed"`
	MinimumSpendBillInAdvance   types.Bool                            `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,computed"`
	MinimumSpendDescription     types.String                          `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,computed"`
	Name                        types.String                          `tfsdk:"name" json:"name,computed"`
	Ordinal                     types.Int64                           `tfsdk:"ordinal" json:"ordinal,computed"`
	ProductID                   types.String                          `tfsdk:"product_id" json:"productId,computed"`
	StandingCharge              types.Float64                         `tfsdk:"standing_charge" json:"standingCharge,computed"`
	StandingChargeBillInAdvance types.Bool                            `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,computed"`
	StandingChargeDescription   types.String                          `tfsdk:"standing_charge_description" json:"standingChargeDescription,computed"`
	StandingChargeInterval      types.Int64                           `tfsdk:"standing_charge_interval" json:"standingChargeInterval,computed"`
	StandingChargeOffset        types.Int64                           `tfsdk:"standing_charge_offset" json:"standingChargeOffset,computed"`
	Version                     types.Int64                           `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	CustomFields                customfield.NormalizedDynamicValue    `tfsdk:"custom_fields" json:"customFields,computed"`
	FindOneBy                   *PlanTemplateFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *PlanTemplateDataSourceModel) toReadParams(_ context.Context) (params m3ter.PlanTemplateGetParams, diags diag.Diagnostics) {
	params = m3ter.PlanTemplateGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *PlanTemplateDataSourceModel) toListParams(_ context.Context) (params m3ter.PlanTemplateListParams, diags diag.Diagnostics) {
	mFindOneByIDs := []string{}
	if m.FindOneBy.IDs != nil {
		for _, item := range *m.FindOneBy.IDs {
			mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
		}
	}

	params = m3ter.PlanTemplateListParams{
		IDs: m3ter.F(mFindOneByIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.FindOneBy.ProductID.IsNull() {
		params.ProductID = m3ter.F(m.FindOneBy.ProductID.ValueString())
	}

	return
}

type PlanTemplateFindOneByDataSourceModel struct {
	IDs       *[]types.String `tfsdk:"ids" query:"ids,optional"`
	ProductID types.String    `tfsdk:"product_id" query:"productId,optional"`
}
