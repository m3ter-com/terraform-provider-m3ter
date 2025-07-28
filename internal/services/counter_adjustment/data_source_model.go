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
	ID                  types.String      `tfsdk:"id" path:"id,required"`
	OrgID               types.String      `tfsdk:"org_id" path:"orgId,required"`
	AccountID           types.String      `tfsdk:"account_id" json:"accountId,computed"`
	CounterID           types.String      `tfsdk:"counter_id" json:"counterId,computed"`
	Date                timetypes.RFC3339 `tfsdk:"date" json:"date,computed" format:"date"`
	PurchaseOrderNumber types.String      `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	Value               types.Int64       `tfsdk:"value" json:"value,computed"`
	Version             types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m *CounterAdjustmentDataSourceModel) toReadParams(_ context.Context) (params m3ter.CounterAdjustmentGetParams, diags diag.Diagnostics) {
	params = m3ter.CounterAdjustmentGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
