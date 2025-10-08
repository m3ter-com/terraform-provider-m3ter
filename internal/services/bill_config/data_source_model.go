// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type BillConfigDataSourceModel struct {
	ID           types.String      `tfsdk:"id" path:"orgId,computed"`
	OrgID        types.String      `tfsdk:"org_id" path:"orgId,optional"`
	BillLockDate timetypes.RFC3339 `tfsdk:"bill_lock_date" json:"billLockDate,computed" format:"date"`
	Version      types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m *BillConfigDataSourceModel) toReadParams(_ context.Context) (params m3ter.BillConfigGetParams, diags diag.Diagnostics) {
	params = m3ter.BillConfigGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
