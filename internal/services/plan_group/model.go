// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type PlanGroupModel struct {
	ID                                types.String      `tfsdk:"id" json:"id,computed"`
	OrgID                             types.String      `tfsdk:"org_id" path:"orgId,optional"`
	Currency                          types.String      `tfsdk:"currency" json:"currency,required"`
	Name                              types.String      `tfsdk:"name" json:"name,required"`
	AccountID                         types.String      `tfsdk:"account_id" json:"accountId,optional"`
	Code                              types.String      `tfsdk:"code" json:"code,optional"`
	MinimumSpend                      types.Float64     `tfsdk:"minimum_spend" json:"minimumSpend,optional"`
	MinimumSpendAccountingProductID   types.String      `tfsdk:"minimum_spend_accounting_product_id" json:"minimumSpendAccountingProductId,optional"`
	MinimumSpendBillInAdvance         types.Bool        `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,optional"`
	MinimumSpendDescription           types.String      `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,optional"`
	StandingCharge                    types.Float64     `tfsdk:"standing_charge" json:"standingCharge,optional"`
	StandingChargeAccountingProductID types.String      `tfsdk:"standing_charge_accounting_product_id" json:"standingChargeAccountingProductId,optional"`
	StandingChargeBillInAdvance       types.Bool        `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,optional"`
	StandingChargeDescription         types.String      `tfsdk:"standing_charge_description" json:"standingChargeDescription,optional"`
	CustomFields                      types.Dynamic     `tfsdk:"custom_fields" json:"customFields,optional"`
	CreatedBy                         types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                         timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified                    timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy                    types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version                           types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m PlanGroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PlanGroupModel) MarshalJSONForUpdate(state PlanGroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
