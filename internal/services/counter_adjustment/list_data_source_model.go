// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CounterAdjustmentsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CounterAdjustmentsItemsDataSourceModel] `json:"data,computed"`
}

type CounterAdjustmentsDataSourceModel struct {
	OrgID        types.String                                                         `tfsdk:"org_id" path:"orgId,optional"`
	AccountID    types.String                                                         `tfsdk:"account_id" query:"accountId,optional"`
	CounterID    types.String                                                         `tfsdk:"counter_id" query:"counterId,optional"`
	Date         types.String                                                         `tfsdk:"date" query:"date,optional"`
	DateEnd      types.String                                                         `tfsdk:"date_end" query:"dateEnd,optional"`
	DateStart    types.String                                                         `tfsdk:"date_start" query:"dateStart,optional"`
	EndDateEnd   types.String                                                         `tfsdk:"end_date_end" query:"endDateEnd,optional"`
	EndDateStart types.String                                                         `tfsdk:"end_date_start" query:"endDateStart,optional"`
	SortOrder    types.String                                                         `tfsdk:"sort_order" query:"sortOrder,optional"`
	MaxItems     types.Int64                                                          `tfsdk:"max_items"`
	Items        customfield.NestedObjectList[CounterAdjustmentsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *CounterAdjustmentsDataSourceModel) toListParams(_ context.Context) (params m3ter.CounterAdjustmentListParams, diags diag.Diagnostics) {
	params = m3ter.CounterAdjustmentListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = m3ter.F(m.AccountID.ValueString())
	}
	if !m.CounterID.IsNull() {
		params.CounterID = m3ter.F(m.CounterID.ValueString())
	}
	if !m.Date.IsNull() {
		params.Date = m3ter.F(m.Date.ValueString())
	}
	if !m.DateEnd.IsNull() {
		params.DateEnd = m3ter.F(m.DateEnd.ValueString())
	}
	if !m.DateStart.IsNull() {
		params.DateStart = m3ter.F(m.DateStart.ValueString())
	}
	if !m.EndDateEnd.IsNull() {
		params.EndDateEnd = m3ter.F(m.EndDateEnd.ValueString())
	}
	if !m.EndDateStart.IsNull() {
		params.EndDateStart = m3ter.F(m.EndDateStart.ValueString())
	}
	if !m.SortOrder.IsNull() {
		params.SortOrder = m3ter.F(m.SortOrder.ValueString())
	}
	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type CounterAdjustmentsItemsDataSourceModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	Version             types.Int64       `tfsdk:"version" json:"version,computed"`
	AccountID           types.String      `tfsdk:"account_id" json:"accountId,computed"`
	CounterID           types.String      `tfsdk:"counter_id" json:"counterId,computed"`
	CreatedBy           types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	Date                timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
	DtCreated           timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy      types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	PurchaseOrderNumber types.String      `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	Value               types.Int64       `tfsdk:"value" json:"value,computed"`
}
