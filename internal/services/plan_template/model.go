// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_template

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type PlanTemplateModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
BillFrequency types.String `tfsdk:"bill_frequency" json:"billFrequency,required"`
Currency types.String `tfsdk:"currency" json:"currency,required"`
Name types.String `tfsdk:"name" json:"name,required"`
ProductID types.String `tfsdk:"product_id" json:"productId,required"`
StandingCharge types.Float64 `tfsdk:"standing_charge" json:"standingCharge,required"`
BillFrequencyInterval types.Int64 `tfsdk:"bill_frequency_interval" json:"billFrequencyInterval,optional"`
Code types.String `tfsdk:"code" json:"code,optional"`
MinimumSpend types.Float64 `tfsdk:"minimum_spend" json:"minimumSpend,optional"`
MinimumSpendBillInAdvance types.Bool `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,optional"`
MinimumSpendDescription types.String `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,optional"`
Ordinal types.Int64 `tfsdk:"ordinal" json:"ordinal,optional"`
StandingChargeBillInAdvance types.Bool `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,optional"`
StandingChargeDescription types.String `tfsdk:"standing_charge_description" json:"standingChargeDescription,optional"`
StandingChargeInterval types.Int64 `tfsdk:"standing_charge_interval" json:"standingChargeInterval,optional"`
StandingChargeOffset types.Int64 `tfsdk:"standing_charge_offset" json:"standingChargeOffset,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
CustomFields *map[string]types.Dynamic `tfsdk:"custom_fields" json:"customFields,optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m PlanTemplateModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m PlanTemplateModel) MarshalJSONForUpdate(state PlanTemplateModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
