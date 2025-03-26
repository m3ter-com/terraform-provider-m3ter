// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type PlansDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[PlansItemsDataSourceModel] `json:"data,computed"`
}

type PlansDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
ProductID types.String `tfsdk:"product_id" query:"productId,optional"`
AccountID *[]types.String `tfsdk:"account_id" query:"accountId,optional"`
IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[PlansItemsDataSourceModel] `tfsdk:"items"`
}

func (m *PlansDataSourceModel) toListParams(_ context.Context) (params m3ter.PlanListParams, diags diag.Diagnostics) {
  mAccountID := []string{}
  for _, item := range *m.AccountID {
    mAccountID = append(mAccountID, item.ValueString())
  }
  mIDs := []string{}
  for _, item := range *m.IDs {
    mIDs = append(mIDs, item.ValueString())
  }

  params = m3ter.PlanListParams{
    AccountID: m3ter.F(mAccountID),
    IDs: m3ter.F(mIDs),
  }

  if !m.ProductID.IsNull() {
    params.ProductID = m3ter.F(m.ProductID.ValueString())
  }
  if !m.OrgID.IsNull() {
    params.OrgID = m3ter.F(m.OrgID.ValueString())
  }

  return
}

type PlansItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
AccountID types.String `tfsdk:"account_id" json:"accountId,computed"`
Bespoke types.Bool `tfsdk:"bespoke" json:"bespoke,computed"`
Code types.String `tfsdk:"code" json:"code,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
CustomFields customfield.Map[types.Dynamic] `tfsdk:"custom_fields" json:"customFields,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
MinimumSpend types.Float64 `tfsdk:"minimum_spend" json:"minimumSpend,computed"`
MinimumSpendAccountingProductID types.String `tfsdk:"minimum_spend_accounting_product_id" json:"minimumSpendAccountingProductId,computed"`
MinimumSpendBillInAdvance types.Bool `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,computed"`
MinimumSpendDescription types.String `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Ordinal types.Int64 `tfsdk:"ordinal" json:"ordinal,computed"`
PlanTemplateID types.String `tfsdk:"plan_template_id" json:"planTemplateId,computed"`
ProductID types.String `tfsdk:"product_id" json:"productId,computed"`
StandingCharge types.Float64 `tfsdk:"standing_charge" json:"standingCharge,computed"`
StandingChargeAccountingProductID types.String `tfsdk:"standing_charge_accounting_product_id" json:"standingChargeAccountingProductId,computed"`
StandingChargeBillInAdvance types.Bool `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,computed"`
StandingChargeDescription types.String `tfsdk:"standing_charge_description" json:"standingChargeDescription,computed"`
}
