// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type CounterAdjustmentModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
AccountID types.String `tfsdk:"account_id" json:"accountId,required"`
CounterID types.String `tfsdk:"counter_id" json:"counterId,required"`
Value types.Int64 `tfsdk:"value" json:"value,required"`
PurchaseOrderNumber types.String `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
Date types.String `tfsdk:"date" json:"date,computed_optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m CounterAdjustmentModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m CounterAdjustmentModel) MarshalJSONForUpdate(state CounterAdjustmentModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
