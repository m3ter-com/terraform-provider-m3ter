// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance_transaction

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type BalanceTransactionModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
BalanceID types.String `tfsdk:"balance_id" path:"balanceId,required"`
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
Amount types.Float64 `tfsdk:"amount" json:"amount,required"`
AppliedDate timetypes.RFC3339 `tfsdk:"applied_date" json:"appliedDate,optional" format:"date-time"`
CurrencyPaid types.String `tfsdk:"currency_paid" json:"currencyPaid,optional"`
Description types.String `tfsdk:"description" json:"description,optional"`
Paid types.Float64 `tfsdk:"paid" json:"paid,optional"`
TransactionDate timetypes.RFC3339 `tfsdk:"transaction_date" json:"transactionDate,optional" format:"date-time"`
TransactionTypeID types.String `tfsdk:"transaction_type_id" json:"transactionTypeId,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
EntityID types.String `tfsdk:"entity_id" json:"entityId,computed"`
EntityType types.String `tfsdk:"entity_type" json:"entityType,computed"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m BalanceTransactionModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m BalanceTransactionModel) MarshalJSONForUpdate(state BalanceTransactionModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
