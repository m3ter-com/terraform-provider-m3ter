// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pricing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type PricingsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[PricingsItemsDataSourceModel] `json:"data,computed"`
}

type PricingsDataSourceModel struct {
	OrgID          types.String                                               `tfsdk:"org_id" path:"orgId,optional"`
	AggregationID  types.String                                               `tfsdk:"aggregation_id" query:"aggregationId,optional"`
	Date           types.String                                               `tfsdk:"date" query:"date,optional"`
	PlanID         types.String                                               `tfsdk:"plan_id" query:"planId,optional"`
	PlanTemplateID types.String                                               `tfsdk:"plan_template_id" query:"planTemplateId,optional"`
	IDs            *[]types.String                                            `tfsdk:"ids" query:"ids,optional"`
	MaxItems       types.Int64                                                `tfsdk:"max_items"`
	Items          customfield.NestedObjectList[PricingsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *PricingsDataSourceModel) toListParams(_ context.Context) (params m3ter.PricingListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	if m.IDs != nil {
		for _, item := range *m.IDs {
			mIDs = append(mIDs, item.ValueString())
		}
	}

	params = m3ter.PricingListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.AggregationID.IsNull() {
		params.AggregationID = m3ter.F(m.AggregationID.ValueString())
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

type PricingsItemsDataSourceModel struct {
	ID                        types.String                                                             `tfsdk:"id" json:"id,computed"`
	AccountingProductID       types.String                                                             `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	AggregationID             types.String                                                             `tfsdk:"aggregation_id" json:"aggregationId,computed"`
	AggregationType           types.String                                                             `tfsdk:"aggregation_type" json:"aggregationType,computed"`
	Code                      types.String                                                             `tfsdk:"code" json:"code,computed"`
	CompoundAggregationID     types.String                                                             `tfsdk:"compound_aggregation_id" json:"compoundAggregationId,computed"`
	Cumulative                types.Bool                                                               `tfsdk:"cumulative" json:"cumulative,computed"`
	Description               types.String                                                             `tfsdk:"description" json:"description,computed"`
	EndDate                   timetypes.RFC3339                                                        `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
	MinimumSpend              types.Float64                                                            `tfsdk:"minimum_spend" json:"minimumSpend,computed"`
	MinimumSpendBillInAdvance types.Bool                                                               `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,computed"`
	MinimumSpendDescription   types.String                                                             `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,computed"`
	OveragePricingBands       customfield.NestedObjectList[PricingsOveragePricingBandsDataSourceModel] `tfsdk:"overage_pricing_bands" json:"overagePricingBands,computed"`
	PlanID                    types.String                                                             `tfsdk:"plan_id" json:"planId,computed"`
	PlanTemplateID            types.String                                                             `tfsdk:"plan_template_id" json:"planTemplateId,computed"`
	PricingBands              customfield.NestedObjectList[PricingsPricingBandsDataSourceModel]        `tfsdk:"pricing_bands" json:"pricingBands,computed"`
	Segment                   customfield.Map[types.String]                                            `tfsdk:"segment" json:"segment,computed"`
	SegmentString             types.String                                                             `tfsdk:"segment_string" json:"segmentString,computed"`
	StartDate                 timetypes.RFC3339                                                        `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
	TiersSpanPlan             types.Bool                                                               `tfsdk:"tiers_span_plan" json:"tiersSpanPlan,computed"`
	Type                      types.String                                                             `tfsdk:"type" json:"type,computed"`
	Version                   types.Int64                                                              `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

type PricingsOveragePricingBandsDataSourceModel struct {
	FixedPrice   types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,computed"`
	LowerLimit   types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unitPrice,computed"`
	ID           types.String  `tfsdk:"id" json:"id,computed"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
}

type PricingsPricingBandsDataSourceModel struct {
	FixedPrice   types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,computed"`
	LowerLimit   types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unitPrice,computed"`
	ID           types.String  `tfsdk:"id" json:"id,computed"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
}
