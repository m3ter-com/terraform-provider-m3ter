// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type IntegrationConfigurationDataSourceModel struct {
	ID             types.String      `tfsdk:"id" path:"id,required"`
	OrgID          types.String      `tfsdk:"org_id" path:"orgId,required"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	Destination    types.String      `tfsdk:"destination" json:"destination,computed"`
	DtCompleted    timetypes.RFC3339 `tfsdk:"dt_completed" json:"dtCompleted,computed" format:"date-time"`
	DtCreated      timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	DtStarted      timetypes.RFC3339 `tfsdk:"dt_started" json:"dtStarted,computed" format:"date-time"`
	EntityID       types.String      `tfsdk:"entity_id" json:"entityId,computed"`
	EntityType     types.String      `tfsdk:"entity_type" json:"entityType,computed"`
	Error          types.String      `tfsdk:"error" json:"error,computed"`
	ExternalID     types.String      `tfsdk:"external_id" json:"externalId,computed"`
	LastModifiedBy types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Status         types.String      `tfsdk:"status" json:"status,computed"`
	URL            types.String      `tfsdk:"url" json:"url,computed"`
	Version        types.Int64       `tfsdk:"version" json:"version,computed"`
}

func (m *IntegrationConfigurationDataSourceModel) toReadParams(_ context.Context) (params m3ter.IntegrationConfigurationGetParams, diags diag.Diagnostics) {
	params = m3ter.IntegrationConfigurationGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
