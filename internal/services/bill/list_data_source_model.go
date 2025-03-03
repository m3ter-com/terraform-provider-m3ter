// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type BillsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[BillsItemsDataSourceModel] `json:"data,computed"`
}

type BillsDataSourceModel struct {
	OrgID                    types.String                                            `tfsdk:"org_id" path:"orgId,required"`
	AccountID                types.String                                            `tfsdk:"account_id" query:"accountId,optional"`
	BillDate                 types.String                                            `tfsdk:"bill_date" query:"billDate,optional"`
	BillDateEnd              types.String                                            `tfsdk:"bill_date_end" query:"billDateEnd,optional"`
	BillDateStart            types.String                                            `tfsdk:"bill_date_start" query:"billDateStart,optional"`
	BillingFrequency         types.String                                            `tfsdk:"billing_frequency" query:"billingFrequency,optional"`
	ExcludeLineItems         types.Bool                                              `tfsdk:"exclude_line_items" query:"excludeLineItems,optional"`
	ExternalInvoiceDateEnd   types.String                                            `tfsdk:"external_invoice_date_end" query:"externalInvoiceDateEnd,optional"`
	ExternalInvoiceDateStart types.String                                            `tfsdk:"external_invoice_date_start" query:"externalInvoiceDateStart,optional"`
	IncludeBillTotal         types.Bool                                              `tfsdk:"include_bill_total" query:"includeBillTotal,optional"`
	Locked                   types.Bool                                              `tfsdk:"locked" query:"locked,optional"`
	Status                   types.String                                            `tfsdk:"status" query:"status,optional"`
	IDs                      *[]types.String                                         `tfsdk:"ids" query:"ids,optional"`
	MaxItems                 types.Int64                                             `tfsdk:"max_items"`
	Items                    customfield.NestedObjectList[BillsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *BillsDataSourceModel) toListParams(_ context.Context) (params m3ter.BillListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.BillListParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
		IDs:   m3ter.F(mIDs),
	}

	if !m.AccountID.IsNull() {
		params.AccountID = m3ter.F(m.AccountID.ValueString())
	}
	if !m.BillDate.IsNull() {
		params.BillDate = m3ter.F(m.BillDate.ValueString())
	}
	if !m.BillDateEnd.IsNull() {
		params.BillDateEnd = m3ter.F(m.BillDateEnd.ValueString())
	}
	if !m.BillDateStart.IsNull() {
		params.BillDateStart = m3ter.F(m.BillDateStart.ValueString())
	}
	if !m.BillingFrequency.IsNull() {
		params.BillingFrequency = m3ter.F(m.BillingFrequency.ValueString())
	}
	if !m.ExcludeLineItems.IsNull() {
		params.ExcludeLineItems = m3ter.F(m.ExcludeLineItems.ValueBool())
	}
	if !m.ExternalInvoiceDateEnd.IsNull() {
		params.ExternalInvoiceDateEnd = m3ter.F(m.ExternalInvoiceDateEnd.ValueString())
	}
	if !m.ExternalInvoiceDateStart.IsNull() {
		params.ExternalInvoiceDateStart = m3ter.F(m.ExternalInvoiceDateStart.ValueString())
	}
	if !m.IncludeBillTotal.IsNull() {
		params.IncludeBillTotal = m3ter.F(m.IncludeBillTotal.ValueBool())
	}
	if !m.Locked.IsNull() {
		params.Locked = m3ter.F(m.Locked.ValueBool())
	}
	if !m.Status.IsNull() {
		params.Status = m3ter.F(m3ter.BillListParamsStatus(m.Status.ValueString()))
	}

	return
}

type BillsItemsDataSourceModel struct {
	ID                       types.String                                                          `tfsdk:"id" json:"id,computed"`
	Version                  types.Int64                                                           `tfsdk:"version" json:"version,computed"`
	AccountCode              types.String                                                          `tfsdk:"account_code" json:"accountCode,computed"`
	AccountID                types.String                                                          `tfsdk:"account_id" json:"accountId,computed"`
	BillDate                 timetypes.RFC3339                                                     `tfsdk:"bill_date" json:"billDate,computed" format:"date"`
	BillFrequencyInterval    types.Int64                                                           `tfsdk:"bill_frequency_interval" json:"billFrequencyInterval,computed"`
	BillingFrequency         types.String                                                          `tfsdk:"billing_frequency" json:"billingFrequency,computed"`
	BillJobID                types.String                                                          `tfsdk:"bill_job_id" json:"billJobId,computed"`
	BillTotal                types.Float64                                                         `tfsdk:"bill_total" json:"billTotal,computed"`
	CreatedBy                types.String                                                          `tfsdk:"created_by" json:"createdBy,computed"`
	CreatedDate              timetypes.RFC3339                                                     `tfsdk:"created_date" json:"createdDate,computed" format:"date-time"`
	CsvStatementGenerated    types.Bool                                                            `tfsdk:"csv_statement_generated" json:"csvStatementGenerated,computed"`
	Currency                 types.String                                                          `tfsdk:"currency" json:"currency,computed"`
	CurrencyConversions      customfield.NestedObjectList[BillsCurrencyConversionsDataSourceModel] `tfsdk:"currency_conversions" json:"currencyConversions,computed"`
	DtCreated                timetypes.RFC3339                                                     `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified           timetypes.RFC3339                                                     `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	DueDate                  timetypes.RFC3339                                                     `tfsdk:"due_date" json:"dueDate,computed" format:"date"`
	EndDate                  timetypes.RFC3339                                                     `tfsdk:"end_date" json:"endDate,computed" format:"date"`
	EndDateTimeUtc           timetypes.RFC3339                                                     `tfsdk:"end_date_time_utc" json:"endDateTimeUTC,computed" format:"date-time"`
	ExternalInvoiceDate      timetypes.RFC3339                                                     `tfsdk:"external_invoice_date" json:"externalInvoiceDate,computed" format:"date"`
	ExternalInvoiceReference types.String                                                          `tfsdk:"external_invoice_reference" json:"externalInvoiceReference,computed"`
	JsonStatementGenerated   types.Bool                                                            `tfsdk:"json_statement_generated" json:"jsonStatementGenerated,computed"`
	LastCalculatedDate       timetypes.RFC3339                                                     `tfsdk:"last_calculated_date" json:"lastCalculatedDate,computed" format:"date-time"`
	LastModifiedBy           types.String                                                          `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	LineItems                customfield.NestedObjectList[BillsLineItemsDataSourceModel]           `tfsdk:"line_items" json:"lineItems,computed"`
	Locked                   types.Bool                                                            `tfsdk:"locked" json:"locked,computed"`
	PurchaseOrderNumber      types.String                                                          `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	SequentialInvoiceNumber  types.String                                                          `tfsdk:"sequential_invoice_number" json:"sequentialInvoiceNumber,computed"`
	StartDate                timetypes.RFC3339                                                     `tfsdk:"start_date" json:"startDate,computed" format:"date"`
	StartDateTimeUtc         timetypes.RFC3339                                                     `tfsdk:"start_date_time_utc" json:"startDateTimeUTC,computed" format:"date-time"`
	Status                   types.String                                                          `tfsdk:"status" json:"status,computed"`
	Timezone                 types.String                                                          `tfsdk:"timezone" json:"timezone,computed"`
}

type BillsCurrencyConversionsDataSourceModel struct {
	From       types.String  `tfsdk:"from" json:"from,computed"`
	To         types.String  `tfsdk:"to" json:"to,computed"`
	Multiplier types.Float64 `tfsdk:"multiplier" json:"multiplier,computed"`
}

type BillsLineItemsDataSourceModel struct {
	AverageUnitPrice       types.Float64                                                                  `tfsdk:"average_unit_price" json:"averageUnitPrice,computed"`
	ConversionRate         types.Float64                                                                  `tfsdk:"conversion_rate" json:"conversionRate,computed"`
	ConvertedSubtotal      types.Float64                                                                  `tfsdk:"converted_subtotal" json:"convertedSubtotal,computed"`
	Currency               types.String                                                                   `tfsdk:"currency" json:"currency,computed"`
	Description            types.String                                                                   `tfsdk:"description" json:"description,computed"`
	LineItemType           types.String                                                                   `tfsdk:"line_item_type" json:"lineItemType,computed"`
	Quantity               types.Float64                                                                  `tfsdk:"quantity" json:"quantity,computed"`
	Subtotal               types.Float64                                                                  `tfsdk:"subtotal" json:"subtotal,computed"`
	Unit                   types.String                                                                   `tfsdk:"unit" json:"unit,computed"`
	Units                  types.Float64                                                                  `tfsdk:"units" json:"units,computed"`
	ID                     types.String                                                                   `tfsdk:"id" json:"id,computed"`
	AggregationID          types.String                                                                   `tfsdk:"aggregation_id" json:"aggregationId,computed"`
	BalanceID              types.String                                                                   `tfsdk:"balance_id" json:"balanceId,computed"`
	ChargeID               types.String                                                                   `tfsdk:"charge_id" json:"chargeId,computed"`
	ChildAccountCode       types.String                                                                   `tfsdk:"child_account_code" json:"childAccountCode,computed"`
	ChildAccountID         types.String                                                                   `tfsdk:"child_account_id" json:"childAccountId,computed"`
	CommitmentID           types.String                                                                   `tfsdk:"commitment_id" json:"commitmentId,computed"`
	CompoundAggregationID  types.String                                                                   `tfsdk:"compound_aggregation_id" json:"compoundAggregationId,computed"`
	ContractID             types.String                                                                   `tfsdk:"contract_id" json:"contractId,computed"`
	CounterID              types.String                                                                   `tfsdk:"counter_id" json:"counterId,computed"`
	CreditTypeID           types.String                                                                   `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
	Group                  customfield.Map[types.String]                                                  `tfsdk:"group" json:"group,computed"`
	MeterID                types.String                                                                   `tfsdk:"meter_id" json:"meterId,computed"`
	PlanGroupID            types.String                                                                   `tfsdk:"plan_group_id" json:"planGroupId,computed"`
	PlanID                 types.String                                                                   `tfsdk:"plan_id" json:"planId,computed"`
	PricingID              types.String                                                                   `tfsdk:"pricing_id" json:"pricingId,computed"`
	ProductCode            types.String                                                                   `tfsdk:"product_code" json:"productCode,computed"`
	ProductID              types.String                                                                   `tfsdk:"product_id" json:"productId,computed"`
	ProductName            types.String                                                                   `tfsdk:"product_name" json:"productName,computed"`
	ReasonID               types.String                                                                   `tfsdk:"reason_id" json:"reasonId,computed"`
	ReferencedBillID       types.String                                                                   `tfsdk:"referenced_bill_id" json:"referencedBillId,computed"`
	ReferencedLineItemID   types.String                                                                   `tfsdk:"referenced_line_item_id" json:"referencedLineItemId,computed"`
	Segment                customfield.Map[types.String]                                                  `tfsdk:"segment" json:"segment,computed"`
	SequenceNumber         types.Int64                                                                    `tfsdk:"sequence_number" json:"sequenceNumber,computed"`
	ServicePeriodEndDate   timetypes.RFC3339                                                              `tfsdk:"service_period_end_date" json:"servicePeriodEndDate,computed" format:"date-time"`
	ServicePeriodStartDate timetypes.RFC3339                                                              `tfsdk:"service_period_start_date" json:"servicePeriodStartDate,computed" format:"date-time"`
	UsagePerPricingBand    customfield.NestedObjectList[BillsLineItemsUsagePerPricingBandDataSourceModel] `tfsdk:"usage_per_pricing_band" json:"usagePerPricingBand,computed"`
}

type BillsLineItemsUsagePerPricingBandDataSourceModel struct {
	BandQuantity  types.Float64 `tfsdk:"band_quantity" json:"bandQuantity,computed"`
	BandSubtotal  types.Float64 `tfsdk:"band_subtotal" json:"bandSubtotal,computed"`
	BandUnits     types.Float64 `tfsdk:"band_units" json:"bandUnits,computed"`
	CreditTypeID  types.String  `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
	FixedPrice    types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,computed"`
	LowerLimit    types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,computed"`
	PricingBandID types.String  `tfsdk:"pricing_band_id" json:"pricingBandId,computed"`
	UnitPrice     types.Float64 `tfsdk:"unit_price" json:"unitPrice,computed"`
	UnitSubtotal  types.Float64 `tfsdk:"unit_subtotal" json:"unitSubtotal,computed"`
}
