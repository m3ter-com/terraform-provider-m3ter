// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_debit_line_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type BillDebitLineItemDataSourceModel struct {
	BillID                 types.String      `tfsdk:"bill_id" path:"billId,required"`
	ID                     types.String      `tfsdk:"id" path:"id,required"`
	OrgID                  types.String      `tfsdk:"org_id" path:"orgId,required"`
	Amount                 types.Float64     `tfsdk:"amount" json:"amount,computed"`
	CreatedBy              types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DebitReasonID          types.String      `tfsdk:"debit_reason_id" json:"debitReasonId,computed"`
	Description            types.String      `tfsdk:"description" json:"description,computed"`
	DtCreated              timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified         timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy         types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	ProductID              types.String      `tfsdk:"product_id" json:"productId,computed"`
	ReferencedBillID       types.String      `tfsdk:"referenced_bill_id" json:"referencedBillId,computed"`
	ReferencedLineItemID   types.String      `tfsdk:"referenced_line_item_id" json:"referencedLineItemId,computed"`
	ServicePeriodEndDate   timetypes.RFC3339 `tfsdk:"service_period_end_date" json:"servicePeriodEndDate,computed" format:"date-time"`
	ServicePeriodStartDate timetypes.RFC3339 `tfsdk:"service_period_start_date" json:"servicePeriodStartDate,computed" format:"date-time"`
	Version                types.Int64       `tfsdk:"version" json:"version,computed"`
}

func (m *BillDebitLineItemDataSourceModel) toReadParams(_ context.Context) (params m3ter.BillDebitLineItemGetParams, diags diag.Diagnostics) {
	params = m3ter.BillDebitLineItemGetParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
	}

	return
}
