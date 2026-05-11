// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type BalancesDataListDataSourceEnvelope struct {
	Data customfield.NormalizedDynamicValue `json:"data,computed"`
}

type BalancesDataSourceModel struct {
	OrgID        types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	AccountID    types.String                       `tfsdk:"account_id" query:"accountId,optional"`
	Contract     types.String                       `tfsdk:"contract" query:"contract,optional"`
	ContractID   types.String                       `tfsdk:"contract_id" query:"contractId,optional"`
	EndDateEnd   types.String                       `tfsdk:"end_date_end" query:"endDateEnd,optional"`
	EndDateStart types.String                       `tfsdk:"end_date_start" query:"endDateStart,optional"`
	IDs          *[]types.String                    `tfsdk:"ids" query:"ids,optional"`
	MaxItems     types.Int64                        `tfsdk:"max_items"`
	Items        customfield.NormalizedDynamicValue `tfsdk:"items"`
}

func (m *BalancesDataSourceModel) toListParams(_ context.Context) (params m3ter.BalanceListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	if m.IDs != nil {
		for _, item := range *m.IDs {
			mIDs = append(mIDs, item.ValueString())
		}
	}

	params = m3ter.BalanceListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.AccountID.IsNull() {
		params.AccountID = m3ter.F(m.AccountID.ValueString())
	}
	if !m.Contract.IsNull() {
		params.Contract = m3ter.F(m.Contract.ValueString())
	}
	if !m.ContractID.IsNull() {
		params.ContractID = m3ter.F(m.ContractID.ValueString())
	}
	if !m.EndDateEnd.IsNull() {
		params.EndDateEnd = m3ter.F(m.EndDateEnd.ValueString())
	}
	if !m.EndDateStart.IsNull() {
		params.EndDateStart = m3ter.F(m.EndDateStart.ValueString())
	}

	return
}
