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

type AccountPlansDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[AccountPlansItemsDataSourceModel] `json:"data,computed"`
}

type AccountPlansDataSourceModel struct {
	OrgID      types.String                                                   `tfsdk:"org_id" path:"orgId,optional"`
	Account    types.String                                                   `tfsdk:"account" query:"account,optional"`
	Contract   types.String                                                   `tfsdk:"contract" query:"contract,optional"`
	Date       types.String                                                   `tfsdk:"date" query:"date,optional"`
	Includeall types.Bool                                                     `tfsdk:"includeall" query:"includeall,optional"`
	Plan       types.String                                                   `tfsdk:"plan" query:"plan,optional"`
	Product    types.String                                                   `tfsdk:"product" query:"product,optional"`
	IDs        *[]types.String                                                `tfsdk:"ids" query:"ids,optional"`
	MaxItems   types.Int64                                                    `tfsdk:"max_items"`
	Items      customfield.NestedObjectList[AccountPlansItemsDataSourceModel] `tfsdk:"items"`
}

func (m *AccountPlansDataSourceModel) toListParams(_ context.Context) (params m3ter.AccountPlanListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.AccountPlanListParams{
		IDs: m3ter.F(mIDs),
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
	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type AccountPlansItemsDataSourceModel struct {
	ID               types.String      `tfsdk:"id" json:"id,computed"`
	AccountID        types.String      `tfsdk:"account_id" json:"accountId,computed"`
	BillEpoch        timetypes.RFC3339 `tfsdk:"bill_epoch" json:"billEpoch,computed" format:"date"`
	ChildBillingMode types.String      `tfsdk:"child_billing_mode" json:"childBillingMode,computed"`
	Code             types.String      `tfsdk:"code" json:"code,computed"`
	ContractID       types.String      `tfsdk:"contract_id" json:"contractId,computed"`
	CreatedBy        types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	CustomFields     types.Dynamic     `tfsdk:"custom_fields" json:"customFields,computed"`
	DtCreated        timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified   timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EndDate          timetypes.RFC3339 `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
	LastModifiedBy   types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	PlanGroupID      types.String      `tfsdk:"plan_group_id" json:"planGroupId,computed"`
	PlanID           types.String      `tfsdk:"plan_id" json:"planId,computed"`
	ProductID        types.String      `tfsdk:"product_id" json:"productId,computed"`
	StartDate        timetypes.RFC3339 `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
	Version          types.Int64       `tfsdk:"version" json:"version,computed"`
}
