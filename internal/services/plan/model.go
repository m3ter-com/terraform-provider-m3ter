// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type PlanModel struct {
	ID                                types.String  `tfsdk:"id" json:"id,computed"`
	OrgID                             types.String  `tfsdk:"org_id" path:"orgId,optional"`
	Code                              types.String  `tfsdk:"code" json:"code,required"`
	Name                              types.String  `tfsdk:"name" json:"name,required"`
	PlanTemplateID                    types.String  `tfsdk:"plan_template_id" json:"planTemplateId,required"`
	AccountID                         types.String  `tfsdk:"account_id" json:"accountId,optional"`
	Bespoke                           types.Bool    `tfsdk:"bespoke" json:"bespoke,optional"`
	MinimumSpend                      types.Float64 `tfsdk:"minimum_spend" json:"minimumSpend,optional"`
	MinimumSpendAccountingProductID   types.String  `tfsdk:"minimum_spend_accounting_product_id" json:"minimumSpendAccountingProductId,optional"`
	MinimumSpendBillInAdvance         types.Bool    `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,optional"`
	MinimumSpendDescription           types.String  `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,optional"`
	Ordinal                           types.Int64   `tfsdk:"ordinal" json:"ordinal,optional"`
	StandingCharge                    types.Float64 `tfsdk:"standing_charge" json:"standingCharge,optional"`
	StandingChargeAccountingProductID types.String  `tfsdk:"standing_charge_accounting_product_id" json:"standingChargeAccountingProductId,optional"`
	StandingChargeBillInAdvance       types.Bool    `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,optional"`
	StandingChargeDescription         types.String  `tfsdk:"standing_charge_description" json:"standingChargeDescription,optional"`
	CustomFields                      types.Dynamic `tfsdk:"custom_fields" json:"customFields,optional"`
	ProductID                         types.String  `tfsdk:"product_id" json:"productId,computed"`
	Version                           types.Int64   `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m PlanModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PlanModel) MarshalJSONForUpdate(state PlanModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
