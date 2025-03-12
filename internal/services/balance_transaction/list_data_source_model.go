// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance_transaction

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type BalanceTransactionsDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[BalanceTransactionsItemsDataSourceModel] `json:"data,computed"`
}

type BalanceTransactionsDataSourceModel struct {
BalanceID types.String `tfsdk:"balance_id" path:"balanceId,required"`
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
TransactionTypeID types.String `tfsdk:"transaction_type_id" query:"transactionTypeId,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[BalanceTransactionsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *BalanceTransactionsDataSourceModel) toListParams(_ context.Context) (params m3ter.BalanceTransactionListParams, diags diag.Diagnostics) {
  params = m3ter.BalanceTransactionListParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
  }

  if !m.TransactionTypeID.IsNull() {
    params.TransactionTypeID = m3ter.F(m.TransactionTypeID.ValueString())
  }

  return
}

type BalanceTransactionsItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
Amount types.Float64 `tfsdk:"amount" json:"amount,computed"`
AppliedDate timetypes.RFC3339 `tfsdk:"applied_date" json:"appliedDate,computed" format:"date-time"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
CurrencyPaid types.String `tfsdk:"currency_paid" json:"currencyPaid,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
EntityID types.String `tfsdk:"entity_id" json:"entityId,computed"`
EntityType types.String `tfsdk:"entity_type" json:"entityType,computed"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
Paid types.Float64 `tfsdk:"paid" json:"paid,computed"`
TransactionDate timetypes.RFC3339 `tfsdk:"transaction_date" json:"transactionDate,computed" format:"date-time"`
TransactionTypeID types.String `tfsdk:"transaction_type_id" json:"transactionTypeId,computed"`
}
