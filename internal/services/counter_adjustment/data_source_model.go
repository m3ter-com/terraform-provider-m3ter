// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CounterAdjustmentDataSourceModel struct {
	OrgID               types.String      `tfsdk:"org_id" path:"orgId,required"`
	ID                  types.String      `tfsdk:"id" path:"id,computed"`
	AccountID           types.String      `tfsdk:"account_id" json:"accountId,computed"`
	CounterID           types.String      `tfsdk:"counter_id" json:"counterId,computed"`
	CreatedBy           types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	Date                timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
	DtCreated           timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy      types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	PurchaseOrderNumber types.String      `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	Value               types.Int64       `tfsdk:"value" json:"value,computed"`
	Version             types.Int64       `tfsdk:"version" json:"version,computed"`
}
