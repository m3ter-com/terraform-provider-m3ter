// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type BalancesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[BalancesItemsDataSourceModel] `json:"data,computed"`
}

type BalancesDataSourceModel struct {
	OrgID        types.String                                               `tfsdk:"org_id" path:"orgId,optional"`
	AccountID    types.String                                               `tfsdk:"account_id" query:"accountId,optional"`
	Contract     types.String                                               `tfsdk:"contract" query:"contract,optional"`
	EndDateEnd   types.String                                               `tfsdk:"end_date_end" query:"endDateEnd,optional"`
	EndDateStart types.String                                               `tfsdk:"end_date_start" query:"endDateStart,optional"`
	MaxItems     types.Int64                                                `tfsdk:"max_items"`
	Items        customfield.NestedObjectList[BalancesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *BalancesDataSourceModel) toListParams(_ context.Context) (params m3ter.BalanceListParams, diags diag.Diagnostics) {
	params = m3ter.BalanceListParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.AccountID.IsNull() {
		params.AccountID = m3ter.F(m.AccountID.ValueString())
	}
	if !m.Contract.IsNull() {
		params.Contract = m3ter.F(m.Contract.ValueString())
	}
	if !m.EndDateEnd.IsNull() {
		params.EndDateEnd = m3ter.F(m.EndDateEnd.ValueString())
	}
	if !m.EndDateStart.IsNull() {
		params.EndDateStart = m3ter.F(m.EndDateStart.ValueString())
	}

	return
}

type BalancesItemsDataSourceModel struct {
	ID                              types.String                       `tfsdk:"id" json:"id,computed"`
	AccountID                       types.String                       `tfsdk:"account_id" json:"accountId,computed"`
	Amount                          types.Float64                      `tfsdk:"amount" json:"amount,computed"`
	BalanceDrawDownDescription      types.String                       `tfsdk:"balance_draw_down_description" json:"balanceDrawDownDescription,computed"`
	Code                            types.String                       `tfsdk:"code" json:"code,computed"`
	ConsumptionsAccountingProductID types.String                       `tfsdk:"consumptions_accounting_product_id" json:"consumptionsAccountingProductId,computed"`
	ContractID                      types.String                       `tfsdk:"contract_id" json:"contractId,computed"`
	Currency                        types.String                       `tfsdk:"currency" json:"currency,computed"`
	CustomFields                    customfield.NormalizedDynamicValue `tfsdk:"custom_fields" json:"customFields,computed"`
	Description                     types.String                       `tfsdk:"description" json:"description,computed"`
	EndDate                         timetypes.RFC3339                  `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
	FeesAccountingProductID         types.String                       `tfsdk:"fees_accounting_product_id" json:"feesAccountingProductId,computed"`
	LineItemTypes                   customfield.List[types.String]     `tfsdk:"line_item_types" json:"lineItemTypes,computed"`
	Name                            types.String                       `tfsdk:"name" json:"name,computed"`
	OverageDescription              types.String                       `tfsdk:"overage_description" json:"overageDescription,computed"`
	OverageSurchargePercent         types.Float64                      `tfsdk:"overage_surcharge_percent" json:"overageSurchargePercent,computed"`
	ProductIDs                      customfield.List[types.String]     `tfsdk:"product_ids" json:"productIds,computed"`
	RolloverAmount                  types.Float64                      `tfsdk:"rollover_amount" json:"rolloverAmount,computed"`
	RolloverEndDate                 timetypes.RFC3339                  `tfsdk:"rollover_end_date" json:"rolloverEndDate,computed" format:"date-time"`
	StartDate                       timetypes.RFC3339                  `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
	Version                         types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}
