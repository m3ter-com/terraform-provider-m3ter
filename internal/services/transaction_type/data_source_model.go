// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package transaction_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type TransactionTypeDataSourceModel struct {
	ID       types.String `tfsdk:"id" path:"id,required"`
	OrgID    types.String `tfsdk:"org_id" path:"orgId,required"`
	Archived types.Bool   `tfsdk:"archived" json:"archived,computed"`
	Code     types.String `tfsdk:"code" json:"code,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Version  types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m *TransactionTypeDataSourceModel) toReadParams(_ context.Context) (params m3ter.TransactionTypeGetParams, diags diag.Diagnostics) {
	params = m3ter.TransactionTypeGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
