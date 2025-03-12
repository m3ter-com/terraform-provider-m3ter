// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type IntegrationConfigurationsDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[IntegrationConfigurationsItemsDataSourceModel] `json:"data,computed"`
}

type IntegrationConfigurationsDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[IntegrationConfigurationsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *IntegrationConfigurationsDataSourceModel) toListParams(_ context.Context) (params m3ter.IntegrationConfigurationListParams, diags diag.Diagnostics) {
  params = m3ter.IntegrationConfigurationListParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
  }

  return
}

type IntegrationConfigurationsItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Destination types.String `tfsdk:"destination" json:"destination,computed"`
EntityType types.String `tfsdk:"entity_type" json:"entityType,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
Authorized types.Bool `tfsdk:"authorized" json:"authorized,computed"`
ConfigData customfield.Map[jsontypes.Normalized] `tfsdk:"config_data" json:"configData,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DestinationID types.String `tfsdk:"destination_id" json:"destinationId,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
EntityID types.String `tfsdk:"entity_id" json:"entityId,computed"`
IntegrationCredentialsID types.String `tfsdk:"integration_credentials_id" json:"integrationCredentialsId,computed"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
TriggerType types.String `tfsdk:"trigger_type" json:"triggerType,computed"`
}
