// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_job

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type BillJobModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
BillDate timetypes.RFC3339 `tfsdk:"bill_date" json:"billDate,optional" format:"date"`
BillFrequencyInterval types.Int64 `tfsdk:"bill_frequency_interval" json:"billFrequencyInterval,optional"`
BillingFrequency types.String `tfsdk:"billing_frequency" json:"billingFrequency,optional"`
DayEpoch timetypes.RFC3339 `tfsdk:"day_epoch" json:"dayEpoch,optional" format:"date"`
DueDate timetypes.RFC3339 `tfsdk:"due_date" json:"dueDate,optional" format:"date"`
ExternalInvoiceDate timetypes.RFC3339 `tfsdk:"external_invoice_date" json:"externalInvoiceDate,optional" format:"date"`
LastDateInBillingPeriod timetypes.RFC3339 `tfsdk:"last_date_in_billing_period" json:"lastDateInBillingPeriod,optional" format:"date"`
MonthEpoch timetypes.RFC3339 `tfsdk:"month_epoch" json:"monthEpoch,optional" format:"date"`
TargetCurrency types.String `tfsdk:"target_currency" json:"targetCurrency,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
WeekEpoch timetypes.RFC3339 `tfsdk:"week_epoch" json:"weekEpoch,optional" format:"date"`
YearEpoch timetypes.RFC3339 `tfsdk:"year_epoch" json:"yearEpoch,optional" format:"date"`
AccountIDs *[]types.String `tfsdk:"account_ids" json:"accountIds,optional"`
Timezone types.String `tfsdk:"timezone" json:"timezone,computed_optional"`
CurrencyConversions customfield.NestedObjectList[BillJobCurrencyConversionsModel] `tfsdk:"currency_conversions" json:"currencyConversions,computed_optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
Pending types.Int64 `tfsdk:"pending" json:"pending,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
Total types.Int64 `tfsdk:"total" json:"total,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
BillIDs customfield.List[types.String] `tfsdk:"bill_ids" json:"billIds,computed"`
}

func (m BillJobModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m BillJobModel) MarshalJSONForUpdate(state BillJobModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type BillJobCurrencyConversionsModel struct {
From types.String `tfsdk:"from" json:"from,required"`
To types.String `tfsdk:"to" json:"to,required"`
Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,optional"`
}
