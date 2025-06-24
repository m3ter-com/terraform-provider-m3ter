// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_pricing

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type CounterPricingModel struct {
	ID                        types.String                        `tfsdk:"id" json:"id,computed"`
	OrgID                     types.String                        `tfsdk:"org_id" path:"orgId,optional"`
	CounterID                 types.String                        `tfsdk:"counter_id" json:"counterId,required"`
	StartDate                 timetypes.RFC3339                   `tfsdk:"start_date" json:"startDate,required" format:"date-time"`
	PricingBands              *[]*CounterPricingPricingBandsModel `tfsdk:"pricing_bands" json:"pricingBands,required"`
	AccountingProductID       types.String                        `tfsdk:"accounting_product_id" json:"accountingProductId,optional"`
	Code                      types.String                        `tfsdk:"code" json:"code,optional"`
	Cumulative                types.Bool                          `tfsdk:"cumulative" json:"cumulative,optional"`
	Description               types.String                        `tfsdk:"description" json:"description,optional"`
	EndDate                   timetypes.RFC3339                   `tfsdk:"end_date" json:"endDate,optional" format:"date-time"`
	PlanID                    types.String                        `tfsdk:"plan_id" json:"planId,optional"`
	PlanTemplateID            types.String                        `tfsdk:"plan_template_id" json:"planTemplateId,optional"`
	ProRateAdjustmentCredit   types.Bool                          `tfsdk:"pro_rate_adjustment_credit" json:"proRateAdjustmentCredit,optional"`
	ProRateAdjustmentDebit    types.Bool                          `tfsdk:"pro_rate_adjustment_debit" json:"proRateAdjustmentDebit,optional"`
	ProRateRunningTotal       types.Bool                          `tfsdk:"pro_rate_running_total" json:"proRateRunningTotal,optional"`
	RunningTotalBillInAdvance types.Bool                          `tfsdk:"running_total_bill_in_advance" json:"runningTotalBillInAdvance,optional"`
	Version                   types.Int64                         `tfsdk:"version" json:"version,computed_optional"`
	CreatedBy                 types.String                        `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                 timetypes.RFC3339                   `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified            timetypes.RFC3339                   `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy            types.String                        `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m CounterPricingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CounterPricingModel) MarshalJSONForUpdate(state CounterPricingModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type CounterPricingPricingBandsModel struct {
	FixedPrice   types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,required"`
	LowerLimit   types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,required"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unitPrice,required"`
	ID           types.String  `tfsdk:"id" json:"id,optional"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"creditTypeId,optional"`
}
