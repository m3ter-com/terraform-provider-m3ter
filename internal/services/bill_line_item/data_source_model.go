// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_line_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type BillLineItemDataSourceModel struct {
	BillID                 types.String                                                       `tfsdk:"bill_id" path:"billId,required"`
	ID                     types.String                                                       `tfsdk:"id" path:"id,required"`
	OrgID                  types.String                                                       `tfsdk:"org_id" path:"orgId,required"`
	AggregationID          types.String                                                       `tfsdk:"aggregation_id" json:"aggregationId,computed"`
	AverageUnitPrice       types.Float64                                                      `tfsdk:"average_unit_price" json:"averageUnitPrice,computed"`
	BalanceID              types.String                                                       `tfsdk:"balance_id" json:"balanceId,computed"`
	CommitmentID           types.String                                                       `tfsdk:"commitment_id" json:"commitmentId,computed"`
	CompoundAggregationID  types.String                                                       `tfsdk:"compound_aggregation_id" json:"compoundAggregationId,computed"`
	ContractID             types.String                                                       `tfsdk:"contract_id" json:"contractId,computed"`
	ConversionRate         types.Float64                                                      `tfsdk:"conversion_rate" json:"conversionRate,computed"`
	ConvertedSubtotal      types.Float64                                                      `tfsdk:"converted_subtotal" json:"convertedSubtotal,computed"`
	CounterID              types.String                                                       `tfsdk:"counter_id" json:"counterId,computed"`
	CreatedBy              types.String                                                       `tfsdk:"created_by" json:"createdBy,computed"`
	CreditTypeID           types.String                                                       `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
	Currency               types.String                                                       `tfsdk:"currency" json:"currency,computed"`
	Description            types.String                                                       `tfsdk:"description" json:"description,computed"`
	DtCreated              timetypes.RFC3339                                                  `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified         timetypes.RFC3339                                                  `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	JsonUsageGenerated     types.Bool                                                         `tfsdk:"json_usage_generated" json:"jsonUsageGenerated,computed"`
	LastModifiedBy         types.String                                                       `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	LineItemType           types.String                                                       `tfsdk:"line_item_type" json:"lineItemType,computed"`
	MeterID                types.String                                                       `tfsdk:"meter_id" json:"meterId,computed"`
	PlanGroupID            types.String                                                       `tfsdk:"plan_group_id" json:"planGroupId,computed"`
	PlanID                 types.String                                                       `tfsdk:"plan_id" json:"planId,computed"`
	PricingID              types.String                                                       `tfsdk:"pricing_id" json:"pricingId,computed"`
	ProductCode            types.String                                                       `tfsdk:"product_code" json:"productCode,computed"`
	ProductID              types.String                                                       `tfsdk:"product_id" json:"productId,computed"`
	ProductName            types.String                                                       `tfsdk:"product_name" json:"productName,computed"`
	Quantity               types.Float64                                                      `tfsdk:"quantity" json:"quantity,computed"`
	ReasonID               types.String                                                       `tfsdk:"reason_id" json:"reasonId,computed"`
	ReferencedBillID       types.String                                                       `tfsdk:"referenced_bill_id" json:"referencedBillId,computed"`
	ReferencedLineItemID   types.String                                                       `tfsdk:"referenced_line_item_id" json:"referencedLineItemId,computed"`
	SequenceNumber         types.Int64                                                        `tfsdk:"sequence_number" json:"sequenceNumber,computed"`
	ServicePeriodEndDate   timetypes.RFC3339                                                  `tfsdk:"service_period_end_date" json:"servicePeriodEndDate,computed" format:"date-time"`
	ServicePeriodStartDate timetypes.RFC3339                                                  `tfsdk:"service_period_start_date" json:"servicePeriodStartDate,computed" format:"date-time"`
	Subtotal               types.Float64                                                      `tfsdk:"subtotal" json:"subtotal,computed"`
	Unit                   types.String                                                       `tfsdk:"unit" json:"unit,computed"`
	Units                  types.Float64                                                      `tfsdk:"units" json:"units,computed"`
	Version                types.Int64                                                        `tfsdk:"version" json:"version,computed"`
	Group                  customfield.Map[types.String]                                      `tfsdk:"group" json:"group,computed"`
	Segment                customfield.Map[types.String]                                      `tfsdk:"segment" json:"segment,computed"`
	BandUsage              customfield.NestedObjectList[BillLineItemBandUsageDataSourceModel] `tfsdk:"band_usage" json:"bandUsage,computed"`
}

func (m *BillLineItemDataSourceModel) toReadParams(_ context.Context) (params m3ter.BillLineItemGetParams, diags diag.Diagnostics) {
	params = m3ter.BillLineItemGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type BillLineItemBandUsageDataSourceModel struct {
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
