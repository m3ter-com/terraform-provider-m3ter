// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_pricing

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CounterPricingsDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[CounterPricingsItemsDataSourceModel] `json:"data,computed"`
}

type CounterPricingsDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
Date types.String `tfsdk:"date" query:"date,optional"`
PlanID types.String `tfsdk:"plan_id" query:"planId,optional"`
PlanTemplateID types.String `tfsdk:"plan_template_id" query:"planTemplateId,optional"`
IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[CounterPricingsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CounterPricingsDataSourceModel) toListParams(_ context.Context) (params m3ter.CounterPricingListParams, diags diag.Diagnostics) {
  mIDs := []string{}
  for _, item := range *m.IDs {
    mIDs = append(mIDs, item.ValueString())
  }

  params = m3ter.CounterPricingListParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
    IDs: m3ter.F(mIDs),
  }

  if !m.Date.IsNull() {
    params.Date = m3ter.F(m.Date.ValueString())
  }
  if !m.PlanID.IsNull() {
    params.PlanID = m3ter.F(m.PlanID.ValueString())
  }
  if !m.PlanTemplateID.IsNull() {
    params.PlanTemplateID = m3ter.F(m.PlanTemplateID.ValueString())
  }

  return
}

type CounterPricingsItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
AccountingProductID types.String `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
Code types.String `tfsdk:"code" json:"code,computed"`
CounterID types.String `tfsdk:"counter_id" json:"counterId,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
Cumulative types.Bool `tfsdk:"cumulative" json:"cumulative,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
EndDate timetypes.RFC3339 `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
PlanID types.String `tfsdk:"plan_id" json:"planId,computed"`
PlanTemplateID types.String `tfsdk:"plan_template_id" json:"planTemplateId,computed"`
PricingBands customfield.NestedObjectList[CounterPricingsPricingBandsDataSourceModel] `tfsdk:"pricing_bands" json:"pricingBands,computed"`
ProRateAdjustmentCredit types.Bool `tfsdk:"pro_rate_adjustment_credit" json:"proRateAdjustmentCredit,computed"`
ProRateAdjustmentDebit types.Bool `tfsdk:"pro_rate_adjustment_debit" json:"proRateAdjustmentDebit,computed"`
ProRateRunningTotal types.Bool `tfsdk:"pro_rate_running_total" json:"proRateRunningTotal,computed"`
RunningTotalBillInAdvance types.Bool `tfsdk:"running_total_bill_in_advance" json:"runningTotalBillInAdvance,computed"`
StartDate timetypes.RFC3339 `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
}

type CounterPricingsPricingBandsDataSourceModel struct {
FixedPrice types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,computed"`
LowerLimit types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,computed"`
UnitPrice types.Float64 `tfsdk:"unit_price" json:"unitPrice,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
CreditTypeID types.String `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
}
