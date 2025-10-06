// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type ResourceGroupDataSourceModel struct {
	ID      types.String `tfsdk:"id" path:"id,computed_optional"`
	OrgID   types.String `tfsdk:"org_id" path:"orgId,required"`
	Type    types.String `tfsdk:"type" path:"type,required"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
	Version types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m *ResourceGroupDataSourceModel) toReadParams(_ context.Context) (params m3ter.ResourceGroupGetParams, diags diag.Diagnostics) {
	params = m3ter.ResourceGroupGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
