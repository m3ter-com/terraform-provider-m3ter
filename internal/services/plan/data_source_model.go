// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type PlanDataSourceModel struct {
	ID                                types.String  `tfsdk:"id" path:"id,required"`
	OrgID                             types.String  `tfsdk:"org_id" path:"orgId,required"`
	AccountID                         types.String  `tfsdk:"account_id" json:"accountId,computed"`
	Bespoke                           types.Bool    `tfsdk:"bespoke" json:"bespoke,computed"`
	Code                              types.String  `tfsdk:"code" json:"code,computed"`
	MinimumSpend                      types.Float64 `tfsdk:"minimum_spend" json:"minimumSpend,computed"`
	MinimumSpendAccountingProductID   types.String  `tfsdk:"minimum_spend_accounting_product_id" json:"minimumSpendAccountingProductId,computed"`
	MinimumSpendBillInAdvance         types.Bool    `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,computed"`
	MinimumSpendDescription           types.String  `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,computed"`
	Name                              types.String  `tfsdk:"name" json:"name,computed"`
	Ordinal                           types.Int64   `tfsdk:"ordinal" json:"ordinal,computed"`
	PlanTemplateID                    types.String  `tfsdk:"plan_template_id" json:"planTemplateId,computed"`
	ProductID                         types.String  `tfsdk:"product_id" json:"productId,computed"`
	StandingCharge                    types.Float64 `tfsdk:"standing_charge" json:"standingCharge,computed"`
	StandingChargeAccountingProductID types.String  `tfsdk:"standing_charge_accounting_product_id" json:"standingChargeAccountingProductId,computed"`
	StandingChargeBillInAdvance       types.Bool    `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,computed"`
	StandingChargeDescription         types.String  `tfsdk:"standing_charge_description" json:"standingChargeDescription,computed"`
	Version                           types.Int64   `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	CustomFields                      types.Dynamic `tfsdk:"custom_fields" json:"customFields,computed"`
}

func (m *PlanDataSourceModel) toReadParams(_ context.Context) (params m3ter.PlanGetParams, diags diag.Diagnostics) {
	params = m3ter.PlanGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
