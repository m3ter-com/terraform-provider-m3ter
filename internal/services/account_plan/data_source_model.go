// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AccountPlanDataSourceModel struct {
	ID               types.String                         `tfsdk:"id" path:"id,computed_optional"`
	OrgID            types.String                         `tfsdk:"org_id" path:"orgId,required"`
	AccountID        types.String                         `tfsdk:"account_id" json:"accountId,computed"`
	BillEpoch        timetypes.RFC3339                    `tfsdk:"bill_epoch" json:"billEpoch,computed" format:"date"`
	ChildBillingMode types.String                         `tfsdk:"child_billing_mode" json:"childBillingMode,computed"`
	Code             types.String                         `tfsdk:"code" json:"code,computed"`
	ContractID       types.String                         `tfsdk:"contract_id" json:"contractId,computed"`
	EndDate          timetypes.RFC3339                    `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
	PlanGroupID      types.String                         `tfsdk:"plan_group_id" json:"planGroupId,computed"`
	PlanID           types.String                         `tfsdk:"plan_id" json:"planId,computed"`
	ProductID        types.String                         `tfsdk:"product_id" json:"productId,computed"`
	StartDate        timetypes.RFC3339                    `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
	Version          types.Int64                          `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	CustomFields     customfield.NormalizedDynamicValue   `tfsdk:"custom_fields" json:"customFields,computed"`
	FindOneBy        *AccountPlanFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *AccountPlanDataSourceModel) toReadParams(_ context.Context) (params m3ter.AccountPlanGetParams, diags diag.Diagnostics) {
	params = m3ter.AccountPlanGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *AccountPlanDataSourceModel) toListParams(_ context.Context) (params m3ter.AccountPlanListParams, diags diag.Diagnostics) {
	mFindOneByIDs := []string{}
	if m.FindOneBy.IDs != nil {
		for _, item := range *m.FindOneBy.IDs {
			mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
		}
	}

	params = m3ter.AccountPlanListParams{
		IDs: m3ter.F(mFindOneByIDs),
	}

	if !m.FindOneBy.Account.IsNull() {
		params.Account = m3ter.F(m.FindOneBy.Account.ValueString())
	}
	if !m.FindOneBy.Contract.IsNull() {
		params.Contract = m3ter.F(m.FindOneBy.Contract.ValueString())
	}
	if !m.FindOneBy.Date.IsNull() {
		params.Date = m3ter.F(m.FindOneBy.Date.ValueString())
	}
	if !m.FindOneBy.Includeall.IsNull() {
		params.Includeall = m3ter.F(m.FindOneBy.Includeall.ValueBool())
	}
	if !m.FindOneBy.Plan.IsNull() {
		params.Plan = m3ter.F(m.FindOneBy.Plan.ValueString())
	}
	if !m.FindOneBy.Product.IsNull() {
		params.Product = m3ter.F(m.FindOneBy.Product.ValueString())
	}

	return
}

type AccountPlanFindOneByDataSourceModel struct {
	Account    types.String    `tfsdk:"account" query:"account,optional"`
	Contract   types.String    `tfsdk:"contract" query:"contract,optional"`
	Date       types.String    `tfsdk:"date" query:"date,optional"`
	IDs        *[]types.String `tfsdk:"ids" query:"ids,optional"`
	Includeall types.Bool      `tfsdk:"includeall" query:"includeall,optional"`
	Plan       types.String    `tfsdk:"plan" query:"plan,optional"`
	Product    types.String    `tfsdk:"product" query:"product,optional"`
}
