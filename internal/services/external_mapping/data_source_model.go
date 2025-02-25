// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package external_mapping

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type ExternalMappingDataSourceModel struct {
	ID                  types.String      `tfsdk:"id" path:"id,required"`
	OrgID               types.String      `tfsdk:"org_id" path:"orgId,required"`
	CreatedBy           types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated           timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	ExternalID          types.String      `tfsdk:"external_id" json:"externalId,computed"`
	ExternalSystem      types.String      `tfsdk:"external_system" json:"externalSystem,computed"`
	ExternalTable       types.String      `tfsdk:"external_table" json:"externalTable,computed"`
	IntegrationConfigID types.String      `tfsdk:"integration_config_id" json:"integrationConfigId,computed"`
	LastModifiedBy      types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	M3terEntity         types.String      `tfsdk:"m3ter_entity" json:"m3terEntity,computed"`
	M3terID             types.String      `tfsdk:"m3ter_id" json:"m3terId,computed"`
	Version             types.Int64       `tfsdk:"version" json:"version,computed"`
}

func (m *ExternalMappingDataSourceModel) toReadParams(_ context.Context) (params m3ter.ExternalMappingGetParams, diags diag.Diagnostics) {
	params = m3ter.ExternalMappingGetParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
	}

	return
}
