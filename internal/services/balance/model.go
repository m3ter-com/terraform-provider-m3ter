// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type BalanceModel struct {
	ID                              types.String      `tfsdk:"id" json:"id,computed"`
	OrgID                           types.String      `tfsdk:"org_id" path:"orgId,optional"`
	AccountID                       types.String      `tfsdk:"account_id" json:"accountId,required"`
	Currency                        types.String      `tfsdk:"currency" json:"currency,required"`
	EndDate                         timetypes.RFC3339 `tfsdk:"end_date" json:"endDate,required" format:"date-time"`
	StartDate                       timetypes.RFC3339 `tfsdk:"start_date" json:"startDate,required" format:"date-time"`
	BalanceDrawDownDescription      types.String      `tfsdk:"balance_draw_down_description" json:"balanceDrawDownDescription,optional"`
	Code                            types.String      `tfsdk:"code" json:"code,optional"`
	ConsumptionsAccountingProductID types.String      `tfsdk:"consumptions_accounting_product_id" json:"consumptionsAccountingProductId,optional"`
	ContractID                      types.String      `tfsdk:"contract_id" json:"contractId,optional"`
	Description                     types.String      `tfsdk:"description" json:"description,optional"`
	FeesAccountingProductID         types.String      `tfsdk:"fees_accounting_product_id" json:"feesAccountingProductId,optional"`
	Name                            types.String      `tfsdk:"name" json:"name,optional"`
	OverageDescription              types.String      `tfsdk:"overage_description" json:"overageDescription,optional"`
	OverageSurchargePercent         types.Float64     `tfsdk:"overage_surcharge_percent" json:"overageSurchargePercent,optional"`
	RolloverAmount                  types.Float64     `tfsdk:"rollover_amount" json:"rolloverAmount,optional"`
	RolloverEndDate                 timetypes.RFC3339 `tfsdk:"rollover_end_date" json:"rolloverEndDate,optional" format:"date-time"`
	LineItemTypes                   *[]types.String   `tfsdk:"line_item_types" json:"lineItemTypes,optional"`
	ProductIDs                      *[]types.String   `tfsdk:"product_ids" json:"productIds,optional"`
	CustomFields                    types.Dynamic     `tfsdk:"custom_fields" json:"customFields,optional"`
	Version                         types.Int64       `tfsdk:"version" json:"version,computed_optional"`
	Amount                          types.Float64     `tfsdk:"amount" json:"amount,computed"`
	CreatedBy                       types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                       timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified                  timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy                  types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m BalanceModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m BalanceModel) MarshalJSONForUpdate(state BalanceModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
