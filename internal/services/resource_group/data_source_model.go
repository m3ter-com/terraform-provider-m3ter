// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type ResourceGroupDataSourceModel struct {
	ID             types.String      `tfsdk:"id" path:"id,required"`
	OrgID          types.String      `tfsdk:"org_id" path:"orgId,required"`
	Type           types.String      `tfsdk:"type" path:"type,required"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name           types.String      `tfsdk:"name" json:"name,computed"`
	Version        types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m *ResourceGroupDataSourceModel) toReadParams(_ context.Context) (params m3ter.ResourceGroupGetParams, diags diag.Diagnostics) {
	params = m3ter.ResourceGroupGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
