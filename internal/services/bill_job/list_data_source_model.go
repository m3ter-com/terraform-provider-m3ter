// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_job

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type BillJobsDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[BillJobsItemsDataSourceModel] `json:"data,computed"`
}

type BillJobsDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
Active types.String `tfsdk:"active" query:"active,optional"`
Status types.String `tfsdk:"status" query:"status,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[BillJobsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *BillJobsDataSourceModel) toListParams(_ context.Context) (params m3ter.BillJobListParams, diags diag.Diagnostics) {
  params = m3ter.BillJobListParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
  }

  if !m.Active.IsNull() {
    params.Active = m3ter.F(m.Active.ValueString())
  }
  if !m.Status.IsNull() {
    params.Status = m3ter.F(m.Status.ValueString())
  }

  return
}

type BillJobsItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
AccountIDs customfield.List[types.String] `tfsdk:"account_ids" json:"accountIds,computed"`
BillDate timetypes.RFC3339 `tfsdk:"bill_date" json:"billDate,computed" format:"date"`
BillFrequencyInterval types.Int64 `tfsdk:"bill_frequency_interval" json:"billFrequencyInterval,computed"`
BillIDs customfield.List[types.String] `tfsdk:"bill_ids" json:"billIds,computed"`
BillingFrequency types.String `tfsdk:"billing_frequency" json:"billingFrequency,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
CurrencyConversions customfield.NestedObjectList[BillJobsCurrencyConversionsDataSourceModel] `tfsdk:"currency_conversions" json:"currencyConversions,computed"`
DayEpoch timetypes.RFC3339 `tfsdk:"day_epoch" json:"dayEpoch,computed" format:"date"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
DueDate timetypes.RFC3339 `tfsdk:"due_date" json:"dueDate,computed" format:"date"`
ExternalInvoiceDate timetypes.RFC3339 `tfsdk:"external_invoice_date" json:"externalInvoiceDate,computed" format:"date"`
LastDateInBillingPeriod timetypes.RFC3339 `tfsdk:"last_date_in_billing_period" json:"lastDateInBillingPeriod,computed" format:"date"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
MonthEpoch timetypes.RFC3339 `tfsdk:"month_epoch" json:"monthEpoch,computed" format:"date"`
Pending types.Int64 `tfsdk:"pending" json:"pending,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
TargetCurrency types.String `tfsdk:"target_currency" json:"targetCurrency,computed"`
Timezone types.String `tfsdk:"timezone" json:"timezone,computed"`
Total types.Int64 `tfsdk:"total" json:"total,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
WeekEpoch timetypes.RFC3339 `tfsdk:"week_epoch" json:"weekEpoch,computed" format:"date"`
YearEpoch timetypes.RFC3339 `tfsdk:"year_epoch" json:"yearEpoch,computed" format:"date"`
}

type BillJobsCurrencyConversionsDataSourceModel struct {
From types.String `tfsdk:"from" json:"from,computed"`
To types.String `tfsdk:"to" json:"to,computed"`
Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
}
