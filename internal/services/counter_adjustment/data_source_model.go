// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type CounterAdjustmentDataSourceModel struct {
	ID                  types.String                               `tfsdk:"id" path:"id,computed_optional"`
	OrgID               types.String                               `tfsdk:"org_id" path:"orgId,required"`
	AccountID           types.String                               `tfsdk:"account_id" json:"accountId,computed"`
	CounterID           types.String                               `tfsdk:"counter_id" json:"counterId,computed"`
	Date                timetypes.RFC3339                          `tfsdk:"date" json:"date,computed" format:"date"`
	PurchaseOrderNumber types.String                               `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	Value               types.Int64                                `tfsdk:"value" json:"value,computed"`
	Version             types.Int64                                `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	FindOneBy           *CounterAdjustmentFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *CounterAdjustmentDataSourceModel) toReadParams(_ context.Context) (params m3ter.CounterAdjustmentGetParams, diags diag.Diagnostics) {
	params = m3ter.CounterAdjustmentGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *CounterAdjustmentDataSourceModel) toListParams(_ context.Context) (params m3ter.CounterAdjustmentListParams, diags diag.Diagnostics) {
	params = m3ter.CounterAdjustmentListParams{}

	if !m.FindOneBy.AccountID.IsNull() {
		params.AccountID = m3ter.F(m.FindOneBy.AccountID.ValueString())
	}
	if !m.FindOneBy.CounterID.IsNull() {
		params.CounterID = m3ter.F(m.FindOneBy.CounterID.ValueString())
	}
	if !m.FindOneBy.Date.IsNull() {
		params.Date = m3ter.F(m.FindOneBy.Date.ValueString())
	}
	if !m.FindOneBy.DateEnd.IsNull() {
		params.DateEnd = m3ter.F(m.FindOneBy.DateEnd.ValueString())
	}
	if !m.FindOneBy.DateStart.IsNull() {
		params.DateStart = m3ter.F(m.FindOneBy.DateStart.ValueString())
	}
	if !m.FindOneBy.EndDateEnd.IsNull() {
		params.EndDateEnd = m3ter.F(m.FindOneBy.EndDateEnd.ValueString())
	}
	if !m.FindOneBy.EndDateStart.IsNull() {
		params.EndDateStart = m3ter.F(m.FindOneBy.EndDateStart.ValueString())
	}
	if !m.FindOneBy.SortOrder.IsNull() {
		params.SortOrder = m3ter.F(m.FindOneBy.SortOrder.ValueString())
	}

	return
}

type CounterAdjustmentFindOneByDataSourceModel struct {
	AccountID    types.String `tfsdk:"account_id" query:"accountId,optional"`
	CounterID    types.String `tfsdk:"counter_id" query:"counterId,optional"`
	Date         types.String `tfsdk:"date" query:"date,optional"`
	DateEnd      types.String `tfsdk:"date_end" query:"dateEnd,optional"`
	DateStart    types.String `tfsdk:"date_start" query:"dateStart,optional"`
	EndDateEnd   types.String `tfsdk:"end_date_end" query:"endDateEnd,optional"`
	EndDateStart types.String `tfsdk:"end_date_start" query:"endDateStart,optional"`
	SortOrder    types.String `tfsdk:"sort_order" query:"sortOrder,optional"`
}
