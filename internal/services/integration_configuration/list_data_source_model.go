// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type IntegrationConfigurationsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[IntegrationConfigurationsItemsDataSourceModel] `json:"data,computed"`
}

type IntegrationConfigurationsDataSourceModel struct {
	OrgID         types.String                                                                `tfsdk:"org_id" path:"orgId,optional"`
	DestinationID types.String                                                                `tfsdk:"destination_id" query:"destinationId,optional"`
	MaxItems      types.Int64                                                                 `tfsdk:"max_items"`
	Items         customfield.NestedObjectList[IntegrationConfigurationsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *IntegrationConfigurationsDataSourceModel) toListParams(_ context.Context) (params m3ter.IntegrationConfigurationListParams, diags diag.Diagnostics) {
	params = m3ter.IntegrationConfigurationListParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.DestinationID.IsNull() {
		params.DestinationID = m3ter.F(m.DestinationID.ValueString())
	}

	return
}

type IntegrationConfigurationsItemsDataSourceModel struct {
	ID                       types.String                          `tfsdk:"id" json:"id,computed"`
	Destination              types.String                          `tfsdk:"destination" json:"destination,computed"`
	EntityType               types.String                          `tfsdk:"entity_type" json:"entityType,computed"`
	Authorized               types.Bool                            `tfsdk:"authorized" json:"authorized,computed"`
	ConfigData               customfield.Map[jsontypes.Normalized] `tfsdk:"config_data" json:"configData,computed"`
	DestinationID            types.String                          `tfsdk:"destination_id" json:"destinationId,computed"`
	Enabled                  types.Bool                            `tfsdk:"enabled" json:"enabled,computed"`
	EntityID                 types.String                          `tfsdk:"entity_id" json:"entityId,computed"`
	IntegrationCredentialsID types.String                          `tfsdk:"integration_credentials_id" json:"integrationCredentialsId,computed"`
	Name                     types.String                          `tfsdk:"name" json:"name,computed"`
	TriggerType              types.String                          `tfsdk:"trigger_type" json:"triggerType,computed"`
	Version                  types.Int64                           `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}
