// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AccountPlansDataListDataSourceEnvelope struct {
	Data customfield.NormalizedDynamicValue `json:"data,computed"`
}

type AccountPlansDataSourceModel struct {
	OrgID      types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	Account    types.String                       `tfsdk:"account" query:"account,optional"`
	Contract   types.String                       `tfsdk:"contract" query:"contract,optional"`
	Date       types.String                       `tfsdk:"date" query:"date,optional"`
	Includeall types.Bool                         `tfsdk:"includeall" query:"includeall,optional"`
	Plan       types.String                       `tfsdk:"plan" query:"plan,optional"`
	Product    types.String                       `tfsdk:"product" query:"product,optional"`
	IDs        *[]types.String                    `tfsdk:"ids" query:"ids,optional"`
	MaxItems   types.Int64                        `tfsdk:"max_items"`
	Items      customfield.NormalizedDynamicValue `tfsdk:"items"`
}

func (m *AccountPlansDataSourceModel) toListParams(_ context.Context) (params m3ter.AccountPlanListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	if m.IDs != nil {
		for _, item := range *m.IDs {
			mIDs = append(mIDs, item.ValueString())
		}
	}

	params = m3ter.AccountPlanListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.Account.IsNull() {
		params.Account = m3ter.F(m.Account.ValueString())
	}
	if !m.Contract.IsNull() {
		params.Contract = m3ter.F(m.Contract.ValueString())
	}
	if !m.Date.IsNull() {
		params.Date = m3ter.F(m.Date.ValueString())
	}
	if !m.Includeall.IsNull() {
		params.Includeall = m3ter.F(m.Includeall.ValueBool())
	}
	if !m.Plan.IsNull() {
		params.Plan = m3ter.F(m.Plan.ValueString())
	}
	if !m.Product.IsNull() {
		params.Product = m3ter.F(m.Product.ValueString())
	}

	return
}
