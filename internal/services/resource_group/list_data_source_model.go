// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ResourceGroupsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[ResourceGroupsItemsDataSourceModel] `json:"data,computed"`
}

type ResourceGroupsDataSourceModel struct {
	Type     types.String                                                     `tfsdk:"type" path:"type,required"`
	OrgID    types.String                                                     `tfsdk:"org_id" path:"orgId,optional"`
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[ResourceGroupsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ResourceGroupsDataSourceModel) toListParams(_ context.Context) (params m3ter.ResourceGroupListParams, diags diag.Diagnostics) {
	params = m3ter.ResourceGroupListParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type ResourceGroupsItemsDataSourceModel struct {
	ID             types.String      `tfsdk:"id" json:"id,computed"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name           types.String      `tfsdk:"name" json:"name,computed"`
	Version        types.Int64       `tfsdk:"version" json:"version,computed"`
}
