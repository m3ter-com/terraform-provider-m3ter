// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_config

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type OrganizationConfigModel struct {
	ID                              types.String                                   `tfsdk:"id" json:"id,computed"`
	OrgID                           types.String                                   `tfsdk:"org_id" path:"orgId,optional"`
	DaysBeforeBillDue               types.Int64                                    `tfsdk:"days_before_bill_due" json:"daysBeforeBillDue,required"`
	AutoApproveBillsGracePeriod     types.Int64                                    `tfsdk:"auto_approve_bills_grace_period" json:"autoApproveBillsGracePeriod,optional"`
	AutoApproveBillsGracePeriodUnit types.String                                   `tfsdk:"auto_approve_bills_grace_period_unit" json:"autoApproveBillsGracePeriodUnit,optional"`
	AutoGenerateStatementMode       types.String                                   `tfsdk:"auto_generate_statement_mode" json:"autoGenerateStatementMode,optional"`
	BillPrefix                      types.String                                   `tfsdk:"bill_prefix" json:"billPrefix,optional"`
	CommitmentFeeBillInAdvance      types.Bool                                     `tfsdk:"commitment_fee_bill_in_advance" json:"commitmentFeeBillInAdvance,optional"`
	ConsolidateBills                types.Bool                                     `tfsdk:"consolidate_bills" json:"consolidateBills,optional"`
	DefaultStatementDefinitionID    types.String                                   `tfsdk:"default_statement_definition_id" json:"defaultStatementDefinitionId,optional"`
	ExternalInvoiceDate             types.String                                   `tfsdk:"external_invoice_date" json:"externalInvoiceDate,optional"`
	MinimumSpendBillInAdvance       types.Bool                                     `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,optional"`
	ScheduledBillInterval           types.Float64                                  `tfsdk:"scheduled_bill_interval" json:"scheduledBillInterval,optional"`
	SequenceStartNumber             types.Int64                                    `tfsdk:"sequence_start_number" json:"sequenceStartNumber,optional"`
	StandingChargeBillInAdvance     types.Bool                                     `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,optional"`
	SuppressedEmptyBills            types.Bool                                     `tfsdk:"suppressed_empty_bills" json:"suppressedEmptyBills,optional"`
	Version                         types.Int64                                    `tfsdk:"version" json:"version,optional"`
	CreditApplicationOrder          *[]types.String                                `tfsdk:"credit_application_order" json:"creditApplicationOrder,optional"`
	CurrencyConversions             *[]*OrganizationConfigCurrencyConversionsModel `tfsdk:"currency_conversions" json:"currencyConversions,optional"`
	Currency                        types.String                                   `tfsdk:"currency" json:"currency,computed_optional"`
	DayEpoch                        types.String                                   `tfsdk:"day_epoch" json:"dayEpoch,computed_optional"`
	MonthEpoch                      types.String                                   `tfsdk:"month_epoch" json:"monthEpoch,computed_optional"`
	Timezone                        types.String                                   `tfsdk:"timezone" json:"timezone,computed_optional"`
	WeekEpoch                       types.String                                   `tfsdk:"week_epoch" json:"weekEpoch,computed_optional"`
	YearEpoch                       types.String                                   `tfsdk:"year_epoch" json:"yearEpoch,computed_optional"`
	CreatedBy                       types.String                                   `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                       timetypes.RFC3339                              `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified                  timetypes.RFC3339                              `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy                  types.String                                   `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m OrganizationConfigModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m OrganizationConfigModel) MarshalJSONForUpdate(state OrganizationConfigModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type OrganizationConfigCurrencyConversionsModel struct {
	From       types.String  `tfsdk:"from" json:"from,required"`
	To         types.String  `tfsdk:"to" json:"to,required"`
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,optional"`
}
