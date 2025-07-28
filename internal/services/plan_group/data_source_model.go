// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type PlanGroupDataSourceModel struct {
	ID                                types.String  `tfsdk:"id" path:"id,required"`
	OrgID                             types.String  `tfsdk:"org_id" path:"orgId,required"`
	AccountID                         types.String  `tfsdk:"account_id" json:"accountId,computed"`
	Code                              types.String  `tfsdk:"code" json:"code,computed"`
	Currency                          types.String  `tfsdk:"currency" json:"currency,computed"`
	MinimumSpend                      types.Float64 `tfsdk:"minimum_spend" json:"minimumSpend,computed"`
	MinimumSpendAccountingProductID   types.String  `tfsdk:"minimum_spend_accounting_product_id" json:"minimumSpendAccountingProductId,computed"`
	MinimumSpendBillInAdvance         types.Bool    `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,computed"`
	MinimumSpendDescription           types.String  `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,computed"`
	Name                              types.String  `tfsdk:"name" json:"name,computed"`
	StandingCharge                    types.Float64 `tfsdk:"standing_charge" json:"standingCharge,computed"`
	StandingChargeAccountingProductID types.String  `tfsdk:"standing_charge_accounting_product_id" json:"standingChargeAccountingProductId,computed"`
	StandingChargeBillInAdvance       types.Bool    `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,computed"`
	StandingChargeDescription         types.String  `tfsdk:"standing_charge_description" json:"standingChargeDescription,computed"`
	Version                           types.Int64   `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	CustomFields                      types.Dynamic `tfsdk:"custom_fields" json:"customFields,computed"`
}

func (m *PlanGroupDataSourceModel) toReadParams(_ context.Context) (params m3ter.PlanGroupGetParams, diags diag.Diagnostics) {
	params = m3ter.PlanGroupGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
