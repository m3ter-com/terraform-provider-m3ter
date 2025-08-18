// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package commitment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CommitmentsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CommitmentsItemsDataSourceModel] `json:"data,computed"`
}

type CommitmentsDataSourceModel struct {
	OrgID        types.String                                                  `tfsdk:"org_id" path:"orgId,optional"`
	AccountID    types.String                                                  `tfsdk:"account_id" query:"accountId,optional"`
	ContractID   types.String                                                  `tfsdk:"contract_id" query:"contractId,optional"`
	Date         types.String                                                  `tfsdk:"date" query:"date,optional"`
	EndDateEnd   types.String                                                  `tfsdk:"end_date_end" query:"endDateEnd,optional"`
	EndDateStart types.String                                                  `tfsdk:"end_date_start" query:"endDateStart,optional"`
	ProductID    types.String                                                  `tfsdk:"product_id" query:"productId,optional"`
	IDs          *[]types.String                                               `tfsdk:"ids" query:"ids,optional"`
	MaxItems     types.Int64                                                   `tfsdk:"max_items"`
	Items        customfield.NestedObjectList[CommitmentsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CommitmentsDataSourceModel) toListParams(_ context.Context) (params m3ter.CommitmentListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	if m.IDs != nil {
		for _, item := range *m.IDs {
			mIDs = append(mIDs, item.ValueString())
		}
	}

	params = m3ter.CommitmentListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.AccountID.IsNull() {
		params.AccountID = m3ter.F(m.AccountID.ValueString())
	}
	if !m.ContractID.IsNull() {
		params.ContractID = m3ter.F(m.ContractID.ValueString())
	}
	if !m.Date.IsNull() {
		params.Date = m3ter.F(m.Date.ValueString())
	}
	if !m.EndDateEnd.IsNull() {
		params.EndDateEnd = m3ter.F(m.EndDateEnd.ValueString())
	}
	if !m.EndDateStart.IsNull() {
		params.EndDateStart = m3ter.F(m.EndDateStart.ValueString())
	}
	if !m.ProductID.IsNull() {
		params.ProductID = m3ter.F(m.ProductID.ValueString())
	}
	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type CommitmentsItemsDataSourceModel struct {
	ID                           types.String                                                     `tfsdk:"id" json:"id,computed"`
	AccountID                    types.String                                                     `tfsdk:"account_id" json:"accountId,computed"`
	AccountingProductID          types.String                                                     `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	Amount                       types.Float64                                                    `tfsdk:"amount" json:"amount,computed"`
	AmountFirstBill              types.Float64                                                    `tfsdk:"amount_first_bill" json:"amountFirstBill,computed"`
	AmountPrePaid                types.Float64                                                    `tfsdk:"amount_pre_paid" json:"amountPrePaid,computed"`
	AmountSpent                  types.Float64                                                    `tfsdk:"amount_spent" json:"amountSpent,computed"`
	BillEpoch                    timetypes.RFC3339                                                `tfsdk:"bill_epoch" json:"billEpoch,computed" format:"date"`
	BillingInterval              types.Int64                                                      `tfsdk:"billing_interval" json:"billingInterval,computed"`
	BillingOffset                types.Int64                                                      `tfsdk:"billing_offset" json:"billingOffset,computed"`
	BillingPlanID                types.String                                                     `tfsdk:"billing_plan_id" json:"billingPlanId,computed"`
	ChildBillingMode             types.String                                                     `tfsdk:"child_billing_mode" json:"childBillingMode,computed"`
	CommitmentFeeBillInAdvance   types.Bool                                                       `tfsdk:"commitment_fee_bill_in_advance" json:"commitmentFeeBillInAdvance,computed"`
	CommitmentFeeDescription     types.String                                                     `tfsdk:"commitment_fee_description" json:"commitmentFeeDescription,computed"`
	CommitmentUsageDescription   types.String                                                     `tfsdk:"commitment_usage_description" json:"commitmentUsageDescription,computed"`
	ContractID                   types.String                                                     `tfsdk:"contract_id" json:"contractId,computed"`
	Currency                     types.String                                                     `tfsdk:"currency" json:"currency,computed"`
	DrawdownsAccountingProductID types.String                                                     `tfsdk:"drawdowns_accounting_product_id" json:"drawdownsAccountingProductId,computed"`
	EndDate                      timetypes.RFC3339                                                `tfsdk:"end_date" json:"endDate,computed" format:"date"`
	FeeDates                     customfield.NestedObjectList[CommitmentsFeeDatesDataSourceModel] `tfsdk:"fee_dates" json:"feeDates,computed"`
	FeesAccountingProductID      types.String                                                     `tfsdk:"fees_accounting_product_id" json:"feesAccountingProductId,computed"`
	LineItemTypes                customfield.List[types.String]                                   `tfsdk:"line_item_types" json:"lineItemTypes,computed"`
	OverageDescription           types.String                                                     `tfsdk:"overage_description" json:"overageDescription,computed"`
	OverageSurchargePercent      types.Float64                                                    `tfsdk:"overage_surcharge_percent" json:"overageSurchargePercent,computed"`
	ProductIDs                   customfield.List[types.String]                                   `tfsdk:"product_ids" json:"productIds,computed"`
	SeparateOverageUsage         types.Bool                                                       `tfsdk:"separate_overage_usage" json:"separateOverageUsage,computed"`
	StartDate                    timetypes.RFC3339                                                `tfsdk:"start_date" json:"startDate,computed" format:"date"`
	Version                      types.Int64                                                      `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

type CommitmentsFeeDatesDataSourceModel struct {
	Amount                 types.Float64     `tfsdk:"amount" json:"amount,computed"`
	Date                   timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
	ServicePeriodEndDate   timetypes.RFC3339 `tfsdk:"service_period_end_date" json:"servicePeriodEndDate,computed" format:"date-time"`
	ServicePeriodStartDate timetypes.RFC3339 `tfsdk:"service_period_start_date" json:"servicePeriodStartDate,computed" format:"date-time"`
}
