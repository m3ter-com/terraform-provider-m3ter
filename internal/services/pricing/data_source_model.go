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

type PricingDataSourceModel struct {
	ID                        types.String                                                            `tfsdk:"id" path:"id,computed_optional"`
	OrgID                     types.String                                                            `tfsdk:"org_id" path:"orgId,optional"`
	AccountingProductID       types.String                                                            `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	AggregationID             types.String                                                            `tfsdk:"aggregation_id" json:"aggregationId,computed"`
	AggregationType           types.String                                                            `tfsdk:"aggregation_type" json:"aggregationType,computed"`
	Code                      types.String                                                            `tfsdk:"code" json:"code,computed"`
	CompoundAggregationID     types.String                                                            `tfsdk:"compound_aggregation_id" json:"compoundAggregationId,computed"`
	Cumulative                types.Bool                                                              `tfsdk:"cumulative" json:"cumulative,computed"`
	Description               types.String                                                            `tfsdk:"description" json:"description,computed"`
	EndDate                   timetypes.RFC3339                                                       `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
	MinimumSpend              types.Float64                                                           `tfsdk:"minimum_spend" json:"minimumSpend,computed"`
	MinimumSpendBillInAdvance types.Bool                                                              `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,computed"`
	MinimumSpendDescription   types.String                                                            `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,computed"`
	PlanID                    types.String                                                            `tfsdk:"plan_id" json:"planId,computed"`
	PlanTemplateID            types.String                                                            `tfsdk:"plan_template_id" json:"planTemplateId,computed"`
	SegmentString             types.String                                                            `tfsdk:"segment_string" json:"segmentString,computed"`
	StartDate                 timetypes.RFC3339                                                       `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
	TiersSpanPlan             types.Bool                                                              `tfsdk:"tiers_span_plan" json:"tiersSpanPlan,computed"`
	Type                      types.String                                                            `tfsdk:"type" json:"type,computed"`
	Version                   types.Int64                                                             `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	Segment                   customfield.Map[types.String]                                           `tfsdk:"segment" json:"segment,computed"`
	OveragePricingBands       customfield.NestedObjectList[PricingOveragePricingBandsDataSourceModel] `tfsdk:"overage_pricing_bands" json:"overagePricingBands,computed"`
	PricingBands              customfield.NestedObjectList[PricingPricingBandsDataSourceModel]        `tfsdk:"pricing_bands" json:"pricingBands,computed"`
	FindOneBy                 *PricingFindOneByDataSourceModel                                        `tfsdk:"find_one_by"`
}

func (m *PricingDataSourceModel) toReadParams(_ context.Context) (params m3ter.PricingGetParams, diags diag.Diagnostics) {
	params = m3ter.PricingGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *PricingDataSourceModel) toListParams(_ context.Context) (params m3ter.PricingListParams, diags diag.Diagnostics) {
	mFindOneByIDs := []string{}
	if m.FindOneBy.IDs != nil {
		for _, item := range *m.FindOneBy.IDs {
			mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
		}
	}

	params = m3ter.PricingListParams{
		IDs: m3ter.F(mFindOneByIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.FindOneBy.AggregationID.IsNull() {
		params.AggregationID = m3ter.F(m.FindOneBy.AggregationID.ValueString())
	}
	if !m.FindOneBy.Date.IsNull() {
		params.Date = m3ter.F(m.FindOneBy.Date.ValueString())
	}
	if !m.FindOneBy.PlanID.IsNull() {
		params.PlanID = m3ter.F(m.FindOneBy.PlanID.ValueString())
	}
	if !m.FindOneBy.PlanTemplateID.IsNull() {
		params.PlanTemplateID = m3ter.F(m.FindOneBy.PlanTemplateID.ValueString())
	}

	return
}

type PricingOveragePricingBandsDataSourceModel struct {
	FixedPrice   types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,computed"`
	LowerLimit   types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unitPrice,computed"`
	ID           types.String  `tfsdk:"id" json:"id,computed"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
}

type PricingPricingBandsDataSourceModel struct {
	FixedPrice   types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,computed"`
	LowerLimit   types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unitPrice,computed"`
	ID           types.String  `tfsdk:"id" json:"id,computed"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
}

type PricingFindOneByDataSourceModel struct {
	AggregationID  types.String    `tfsdk:"aggregation_id" query:"aggregationId,optional"`
	Date           types.String    `tfsdk:"date" query:"date,optional"`
	IDs            *[]types.String `tfsdk:"ids" query:"ids,optional"`
	PlanID         types.String    `tfsdk:"plan_id" query:"planId,optional"`
	PlanTemplateID types.String    `tfsdk:"plan_template_id" query:"planTemplateId,optional"`
}
