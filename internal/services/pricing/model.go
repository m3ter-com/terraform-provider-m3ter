// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pricing

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type PricingModel struct {
	ID                        types.String                        `tfsdk:"id" json:"id,computed"`
	OrgID                     types.String                        `tfsdk:"org_id" path:"orgId,optional"`
	StartDate                 timetypes.RFC3339                   `tfsdk:"start_date" json:"startDate,required" format:"date-time"`
	PricingBands              *[]*PricingPricingBandsModel        `tfsdk:"pricing_bands" json:"pricingBands,required"`
	AccountingProductID       types.String                        `tfsdk:"accounting_product_id" json:"accountingProductId,optional"`
	AggregationID             types.String                        `tfsdk:"aggregation_id" json:"aggregationId,optional"`
	Code                      types.String                        `tfsdk:"code" json:"code,optional"`
	CompoundAggregationID     types.String                        `tfsdk:"compound_aggregation_id" json:"compoundAggregationId,optional"`
	Cumulative                types.Bool                          `tfsdk:"cumulative" json:"cumulative,optional"`
	Description               types.String                        `tfsdk:"description" json:"description,optional"`
	EndDate                   timetypes.RFC3339                   `tfsdk:"end_date" json:"endDate,optional" format:"date-time"`
	MinimumSpend              types.Float64                       `tfsdk:"minimum_spend" json:"minimumSpend,optional"`
	MinimumSpendBillInAdvance types.Bool                          `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,optional"`
	MinimumSpendDescription   types.String                        `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,optional"`
	PlanID                    types.String                        `tfsdk:"plan_id" json:"planId,optional"`
	PlanTemplateID            types.String                        `tfsdk:"plan_template_id" json:"planTemplateId,optional"`
	TiersSpanPlan             types.Bool                          `tfsdk:"tiers_span_plan" json:"tiersSpanPlan,optional"`
	Type                      types.String                        `tfsdk:"type" json:"type,optional"`
	Segment                   *map[string]types.String            `tfsdk:"segment" json:"segment,optional"`
	OveragePricingBands       *[]*PricingOveragePricingBandsModel `tfsdk:"overage_pricing_bands" json:"overagePricingBands,optional"`
	Version                   types.Int64                         `tfsdk:"version" json:"version,computed_optional"`
	AggregationType           types.String                        `tfsdk:"aggregation_type" json:"aggregationType,computed"`
	CreatedBy                 types.String                        `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                 timetypes.RFC3339                   `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified            timetypes.RFC3339                   `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy            types.String                        `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	SegmentString             types.String                        `tfsdk:"segment_string" json:"segmentString,computed"`
}

func (m PricingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PricingModel) MarshalJSONForUpdate(state PricingModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type PricingPricingBandsModel struct {
	FixedPrice   types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,required"`
	LowerLimit   types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unitPrice,required"`
	ID           types.String  `tfsdk:"id" json:"id,optional"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"creditTypeId,optional"`
}

type PricingOveragePricingBandsModel struct {
	FixedPrice   types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,required"`
	LowerLimit   types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unitPrice,required"`
	ID           types.String  `tfsdk:"id" json:"id,optional"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"creditTypeId,optional"`
}
