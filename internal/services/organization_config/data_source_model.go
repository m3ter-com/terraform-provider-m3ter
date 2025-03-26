// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_config

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type OrganizationConfigDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
AutoApproveBillsGracePeriod types.Int64 `tfsdk:"auto_approve_bills_grace_period" json:"autoApproveBillsGracePeriod,computed"`
AutoApproveBillsGracePeriodUnit types.String `tfsdk:"auto_approve_bills_grace_period_unit" json:"autoApproveBillsGracePeriodUnit,computed"`
AutoGenerateStatementMode types.String `tfsdk:"auto_generate_statement_mode" json:"autoGenerateStatementMode,computed"`
BillPrefix types.String `tfsdk:"bill_prefix" json:"billPrefix,computed"`
CommitmentFeeBillInAdvance types.Bool `tfsdk:"commitment_fee_bill_in_advance" json:"commitmentFeeBillInAdvance,computed"`
ConsolidateBills types.Bool `tfsdk:"consolidate_bills" json:"consolidateBills,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
Currency types.String `tfsdk:"currency" json:"currency,computed"`
DayEpoch timetypes.RFC3339 `tfsdk:"day_epoch" json:"dayEpoch,computed" format:"date"`
DaysBeforeBillDue types.Int64 `tfsdk:"days_before_bill_due" json:"daysBeforeBillDue,computed"`
DefaultStatementDefinitionID types.String `tfsdk:"default_statement_definition_id" json:"defaultStatementDefinitionId,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
ExternalInvoiceDate types.String `tfsdk:"external_invoice_date" json:"externalInvoiceDate,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
MinimumSpendBillInAdvance types.Bool `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,computed"`
MonthEpoch timetypes.RFC3339 `tfsdk:"month_epoch" json:"monthEpoch,computed" format:"date"`
ScheduledBillInterval types.Float64 `tfsdk:"scheduled_bill_interval" json:"scheduledBillInterval,computed"`
SequenceStartNumber types.Int64 `tfsdk:"sequence_start_number" json:"sequenceStartNumber,computed"`
StandingChargeBillInAdvance types.Bool `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,computed"`
SuppressedEmptyBills types.Bool `tfsdk:"suppressed_empty_bills" json:"suppressedEmptyBills,computed"`
Timezone types.String `tfsdk:"timezone" json:"timezone,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
WeekEpoch timetypes.RFC3339 `tfsdk:"week_epoch" json:"weekEpoch,computed" format:"date"`
YearEpoch timetypes.RFC3339 `tfsdk:"year_epoch" json:"yearEpoch,computed" format:"date"`
CreditApplicationOrder customfield.List[types.String] `tfsdk:"credit_application_order" json:"creditApplicationOrder,computed"`
CurrencyConversions customfield.NestedObjectList[OrganizationConfigCurrencyConversionsDataSourceModel] `tfsdk:"currency_conversions" json:"currencyConversions,computed"`
}

func (m *OrganizationConfigDataSourceModel) toReadParams(_ context.Context) (params m3ter.OrganizationConfigGetParams, diags diag.Diagnostics) {
  params = m3ter.OrganizationConfigGetParams{

  }

  if !m.OrgID.IsNull() {
    params.OrgID = m3ter.F(m.OrgID.ValueString())
  }

  return
}

type OrganizationConfigCurrencyConversionsDataSourceModel struct {
From types.String `tfsdk:"from" json:"from,computed"`
To types.String `tfsdk:"to" json:"to,computed"`
Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
}
