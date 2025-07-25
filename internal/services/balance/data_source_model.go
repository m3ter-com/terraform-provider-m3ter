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

type BalanceDataSourceModel struct {
	ID                              types.String                   `tfsdk:"id" path:"id,required"`
	OrgID                           types.String                   `tfsdk:"org_id" path:"orgId,required"`
	AccountID                       types.String                   `tfsdk:"account_id" json:"accountId,computed"`
	Amount                          types.Float64                  `tfsdk:"amount" json:"amount,computed"`
	BalanceDrawDownDescription      types.String                   `tfsdk:"balance_draw_down_description" json:"balanceDrawDownDescription,computed"`
	Code                            types.String                   `tfsdk:"code" json:"code,computed"`
	ConsumptionsAccountingProductID types.String                   `tfsdk:"consumptions_accounting_product_id" json:"consumptionsAccountingProductId,computed"`
	ContractID                      types.String                   `tfsdk:"contract_id" json:"contractId,computed"`
	CreatedBy                       types.String                   `tfsdk:"created_by" json:"createdBy,computed"`
	Currency                        types.String                   `tfsdk:"currency" json:"currency,computed"`
	Description                     types.String                   `tfsdk:"description" json:"description,computed"`
	DtCreated                       timetypes.RFC3339              `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified                  timetypes.RFC3339              `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EndDate                         timetypes.RFC3339              `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
	FeesAccountingProductID         types.String                   `tfsdk:"fees_accounting_product_id" json:"feesAccountingProductId,computed"`
	LastModifiedBy                  types.String                   `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name                            types.String                   `tfsdk:"name" json:"name,computed"`
	OverageDescription              types.String                   `tfsdk:"overage_description" json:"overageDescription,computed"`
	OverageSurchargePercent         types.Float64                  `tfsdk:"overage_surcharge_percent" json:"overageSurchargePercent,computed"`
	RolloverAmount                  types.Float64                  `tfsdk:"rollover_amount" json:"rolloverAmount,computed"`
	RolloverEndDate                 timetypes.RFC3339              `tfsdk:"rollover_end_date" json:"rolloverEndDate,computed" format:"date-time"`
	StartDate                       timetypes.RFC3339              `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
	Version                         types.Int64                    `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	LineItemTypes                   customfield.List[types.String] `tfsdk:"line_item_types" json:"lineItemTypes,computed"`
	ProductIDs                      customfield.List[types.String] `tfsdk:"product_ids" json:"productIds,computed"`
	CustomFields                    types.Dynamic                  `tfsdk:"custom_fields" json:"customFields,computed"`
}

func (m *BalanceDataSourceModel) toReadParams(_ context.Context) (params m3ter.BalanceGetParams, diags diag.Diagnostics) {
	params = m3ter.BalanceGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
