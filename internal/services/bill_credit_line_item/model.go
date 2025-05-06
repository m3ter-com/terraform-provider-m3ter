// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_credit_line_item

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type BillCreditLineItemModel struct {
	ID                     types.String      `tfsdk:"id" json:"id,computed"`
	BillID                 types.String      `tfsdk:"bill_id" path:"billId,required"`
	OrgID                  types.String      `tfsdk:"org_id" path:"orgId,optional"`
	Amount                 types.Float64     `tfsdk:"amount" json:"amount,required"`
	Description            types.String      `tfsdk:"description" json:"description,required"`
	ProductID              types.String      `tfsdk:"product_id" json:"productId,required"`
	ReferencedBillID       types.String      `tfsdk:"referenced_bill_id" json:"referencedBillId,required"`
	ReferencedLineItemID   types.String      `tfsdk:"referenced_line_item_id" json:"referencedLineItemId,required"`
	ServicePeriodEndDate   timetypes.RFC3339 `tfsdk:"service_period_end_date" json:"servicePeriodEndDate,required" format:"date-time"`
	ServicePeriodStartDate timetypes.RFC3339 `tfsdk:"service_period_start_date" json:"servicePeriodStartDate,required" format:"date-time"`
	CreditReasonID         types.String      `tfsdk:"credit_reason_id" json:"creditReasonId,optional"`
	LineItemType           types.String      `tfsdk:"line_item_type" json:"lineItemType,optional,no_refresh"`
	ReasonID               types.String      `tfsdk:"reason_id" json:"reasonId,optional,no_refresh"`
	Version                types.Int64       `tfsdk:"version" json:"version,optional"`
	CreatedBy              types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated              timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified         timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy         types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m BillCreditLineItemModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m BillCreditLineItemModel) MarshalJSONForUpdate(state BillCreditLineItemModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
