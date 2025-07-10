// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_pricing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CounterPricingDataSourceModel struct {
	ID                        types.String                                                            `tfsdk:"id" path:"id,required"`
	OrgID                     types.String                                                            `tfsdk:"org_id" path:"orgId,required"`
	AccountingProductID       types.String                                                            `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	Code                      types.String                                                            `tfsdk:"code" json:"code,computed"`
	CounterID                 types.String                                                            `tfsdk:"counter_id" json:"counterId,computed"`
	CreatedBy                 types.String                                                            `tfsdk:"created_by" json:"createdBy,computed"`
	Cumulative                types.Bool                                                              `tfsdk:"cumulative" json:"cumulative,computed"`
	Description               types.String                                                            `tfsdk:"description" json:"description,computed"`
	DtCreated                 timetypes.RFC3339                                                       `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified            timetypes.RFC3339                                                       `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EndDate                   timetypes.RFC3339                                                       `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
	LastModifiedBy            types.String                                                            `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	PlanID                    types.String                                                            `tfsdk:"plan_id" json:"planId,computed"`
	PlanTemplateID            types.String                                                            `tfsdk:"plan_template_id" json:"planTemplateId,computed"`
	ProRateAdjustmentCredit   types.Bool                                                              `tfsdk:"pro_rate_adjustment_credit" json:"proRateAdjustmentCredit,computed"`
	ProRateAdjustmentDebit    types.Bool                                                              `tfsdk:"pro_rate_adjustment_debit" json:"proRateAdjustmentDebit,computed"`
	ProRateRunningTotal       types.Bool                                                              `tfsdk:"pro_rate_running_total" json:"proRateRunningTotal,computed"`
	RunningTotalBillInAdvance types.Bool                                                              `tfsdk:"running_total_bill_in_advance" json:"runningTotalBillInAdvance,computed"`
	StartDate                 timetypes.RFC3339                                                       `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
	Version                   types.Int64                                                             `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	PricingBands              customfield.NestedObjectList[CounterPricingPricingBandsDataSourceModel] `tfsdk:"pricing_bands" json:"pricingBands,computed"`
}

func (m *CounterPricingDataSourceModel) toReadParams(_ context.Context) (params m3ter.CounterPricingGetParams, diags diag.Diagnostics) {
	params = m3ter.CounterPricingGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type CounterPricingPricingBandsDataSourceModel struct {
	FixedPrice   types.Float64 `tfsdk:"fixed_price" json:"fixedPrice,computed"`
	LowerLimit   types.Float64 `tfsdk:"lower_limit" json:"lowerLimit,computed"`
	UnitPrice    types.Float64 `tfsdk:"unit_price" json:"unitPrice,computed"`
	ID           types.String  `tfsdk:"id" json:"id,computed"`
	CreditTypeID types.String  `tfsdk:"credit_type_id" json:"creditTypeId,computed"`
}
