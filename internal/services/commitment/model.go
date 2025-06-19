// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package commitment

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type CommitmentModel struct {
	ID                           types.String                `tfsdk:"id" json:"id,computed"`
	OrgID                        types.String                `tfsdk:"org_id" path:"orgId,optional"`
	AccountID                    types.String                `tfsdk:"account_id" json:"accountId,required"`
	Amount                       types.Float64               `tfsdk:"amount" json:"amount,required"`
	Currency                     types.String                `tfsdk:"currency" json:"currency,required"`
	EndDate                      timetypes.RFC3339           `tfsdk:"end_date" json:"endDate,required" format:"date"`
	StartDate                    timetypes.RFC3339           `tfsdk:"start_date" json:"startDate,required" format:"date"`
	AccountingProductID          types.String                `tfsdk:"accounting_product_id" json:"accountingProductId,optional"`
	AmountFirstBill              types.Float64               `tfsdk:"amount_first_bill" json:"amountFirstBill,optional"`
	AmountPrePaid                types.Float64               `tfsdk:"amount_pre_paid" json:"amountPrePaid,optional"`
	BillEpoch                    timetypes.RFC3339           `tfsdk:"bill_epoch" json:"billEpoch,optional" format:"date"`
	BillingInterval              types.Int64                 `tfsdk:"billing_interval" json:"billingInterval,optional"`
	BillingOffset                types.Int64                 `tfsdk:"billing_offset" json:"billingOffset,optional"`
	BillingPlanID                types.String                `tfsdk:"billing_plan_id" json:"billingPlanId,optional"`
	ChildBillingMode             types.String                `tfsdk:"child_billing_mode" json:"childBillingMode,optional"`
	CommitmentFeeBillInAdvance   types.Bool                  `tfsdk:"commitment_fee_bill_in_advance" json:"commitmentFeeBillInAdvance,optional"`
	CommitmentFeeDescription     types.String                `tfsdk:"commitment_fee_description" json:"commitmentFeeDescription,optional"`
	CommitmentUsageDescription   types.String                `tfsdk:"commitment_usage_description" json:"commitmentUsageDescription,optional"`
	ContractID                   types.String                `tfsdk:"contract_id" json:"contractId,optional"`
	DrawdownsAccountingProductID types.String                `tfsdk:"drawdowns_accounting_product_id" json:"drawdownsAccountingProductId,optional"`
	FeesAccountingProductID      types.String                `tfsdk:"fees_accounting_product_id" json:"feesAccountingProductId,optional"`
	OverageDescription           types.String                `tfsdk:"overage_description" json:"overageDescription,optional"`
	OverageSurchargePercent      types.Float64               `tfsdk:"overage_surcharge_percent" json:"overageSurchargePercent,optional"`
	SeparateOverageUsage         types.Bool                  `tfsdk:"separate_overage_usage" json:"separateOverageUsage,optional"`
	LineItemTypes                *[]types.String             `tfsdk:"line_item_types" json:"lineItemTypes,optional"`
	ProductIDs                   *[]types.String             `tfsdk:"product_ids" json:"productIds,optional"`
	FeeDates                     *[]*CommitmentFeeDatesModel `tfsdk:"fee_dates" json:"feeDates,optional"`
	AmountSpent                  types.Float64               `tfsdk:"amount_spent" json:"amountSpent,computed"`
	CreatedBy                    types.String                `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                    timetypes.RFC3339           `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified               timetypes.RFC3339           `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy               types.String                `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version                      types.Int64                 `tfsdk:"version" json:"version,computed"`
}

func (m CommitmentModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CommitmentModel) MarshalJSONForUpdate(state CommitmentModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type CommitmentFeeDatesModel struct {
	Amount                 types.Float64     `tfsdk:"amount" json:"amount,required"`
	Date                   timetypes.RFC3339 `tfsdk:"date" json:"date,required" format:"date"`
	ServicePeriodEndDate   timetypes.RFC3339 `tfsdk:"service_period_end_date" json:"servicePeriodEndDate,required" format:"date-time"`
	ServicePeriodStartDate timetypes.RFC3339 `tfsdk:"service_period_start_date" json:"servicePeriodStartDate,required" format:"date-time"`
}
