// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package commitment

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CommitmentDataSourceModel struct {
	ID                           types.String                                                    `tfsdk:"id" path:"id,required"`
	OrgID                        types.String                                                    `tfsdk:"org_id" path:"orgId,required"`
	AccountID                    types.String                                                    `tfsdk:"account_id" json:"accountId,computed"`
	AccountingProductID          types.String                                                    `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	Amount                       types.Float64                                                   `tfsdk:"amount" json:"amount,computed"`
	AmountFirstBill              types.Float64                                                   `tfsdk:"amount_first_bill" json:"amountFirstBill,computed"`
	AmountPrePaid                types.Float64                                                   `tfsdk:"amount_pre_paid" json:"amountPrePaid,computed"`
	AmountSpent                  types.Float64                                                   `tfsdk:"amount_spent" json:"amountSpent,computed"`
	BillEpoch                    timetypes.RFC3339                                               `tfsdk:"bill_epoch" json:"billEpoch,computed" format:"date"`
	BillingInterval              types.Int64                                                     `tfsdk:"billing_interval" json:"billingInterval,computed"`
	BillingOffset                types.Int64                                                     `tfsdk:"billing_offset" json:"billingOffset,computed"`
	BillingPlanID                types.String                                                    `tfsdk:"billing_plan_id" json:"billingPlanId,computed"`
	ChildBillingMode             types.String                                                    `tfsdk:"child_billing_mode" json:"childBillingMode,computed"`
	CommitmentFeeBillInAdvance   types.Bool                                                      `tfsdk:"commitment_fee_bill_in_advance" json:"commitmentFeeBillInAdvance,computed"`
	CommitmentFeeDescription     types.String                                                    `tfsdk:"commitment_fee_description" json:"commitmentFeeDescription,computed"`
	CommitmentUsageDescription   types.String                                                    `tfsdk:"commitment_usage_description" json:"commitmentUsageDescription,computed"`
	ContractID                   types.String                                                    `tfsdk:"contract_id" json:"contractId,computed"`
	CreatedBy                    types.String                                                    `tfsdk:"created_by" json:"createdBy,computed"`
	Currency                     types.String                                                    `tfsdk:"currency" json:"currency,computed"`
	DrawdownsAccountingProductID types.String                                                    `tfsdk:"drawdowns_accounting_product_id" json:"drawdownsAccountingProductId,computed"`
	DtCreated                    timetypes.RFC3339                                               `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified               timetypes.RFC3339                                               `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EndDate                      timetypes.RFC3339                                               `tfsdk:"end_date" json:"endDate,computed" format:"date"`
	FeesAccountingProductID      types.String                                                    `tfsdk:"fees_accounting_product_id" json:"feesAccountingProductId,computed"`
	LastModifiedBy               types.String                                                    `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	OverageDescription           types.String                                                    `tfsdk:"overage_description" json:"overageDescription,computed"`
	OverageSurchargePercent      types.Float64                                                   `tfsdk:"overage_surcharge_percent" json:"overageSurchargePercent,computed"`
	SeparateOverageUsage         types.Bool                                                      `tfsdk:"separate_overage_usage" json:"separateOverageUsage,computed"`
	StartDate                    timetypes.RFC3339                                               `tfsdk:"start_date" json:"startDate,computed" format:"date"`
	Version                      types.Int64                                                     `tfsdk:"version" json:"version,computed"`
	LineItemTypes                customfield.List[types.String]                                  `tfsdk:"line_item_types" json:"lineItemTypes,computed"`
	ProductIDs                   customfield.List[types.String]                                  `tfsdk:"product_ids" json:"productIds,computed"`
	FeeDates                     customfield.NestedObjectList[CommitmentFeeDatesDataSourceModel] `tfsdk:"fee_dates" json:"feeDates,computed"`
}

type CommitmentFeeDatesDataSourceModel struct {
	Amount                 types.Float64     `tfsdk:"amount" json:"amount,computed"`
	Date                   timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
	ServicePeriodEndDate   timetypes.RFC3339 `tfsdk:"service_period_end_date" json:"servicePeriodEndDate,computed" format:"date-time"`
	ServicePeriodStartDate timetypes.RFC3339 `tfsdk:"service_period_start_date" json:"servicePeriodStartDate,computed" format:"date-time"`
}
