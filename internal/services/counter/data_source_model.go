// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type CounterDataSourceModel struct {
	ID             types.String      `tfsdk:"id" path:"id,required"`
	OrgID          types.String      `tfsdk:"org_id" path:"orgId,required"`
	Code           types.String      `tfsdk:"code" json:"code,computed"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name           types.String      `tfsdk:"name" json:"name,computed"`
	ProductID      types.String      `tfsdk:"product_id" json:"productId,computed"`
	Unit           types.String      `tfsdk:"unit" json:"unit,computed"`
	Version        types.Int64       `tfsdk:"version" json:"version,computed"`
}

func (m *CounterDataSourceModel) toReadParams(_ context.Context) (params m3ter.CounterGetParams, diags diag.Diagnostics) {
	params = m3ter.CounterGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
