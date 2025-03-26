// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type IntegrationConfigurationModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
Destination types.String `tfsdk:"destination" json:"destination,required"`
DestinationID types.String `tfsdk:"destination_id" json:"destinationId,required"`
EntityID types.String `tfsdk:"entity_id" json:"entityId,required"`
EntityType types.String `tfsdk:"entity_type" json:"entityType,required"`
IntegrationCredentialsID types.String `tfsdk:"integration_credentials_id" json:"integrationCredentialsId,required"`
Name types.String `tfsdk:"name" json:"name,required"`
ConfigData *map[string]jsontypes.Normalized `tfsdk:"config_data" json:"configData,required"`
Credentials *IntegrationConfigurationCredentialsModel `tfsdk:"credentials" json:"credentials,required"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
Authorized types.Bool `tfsdk:"authorized" json:"authorized,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCompleted timetypes.RFC3339 `tfsdk:"dt_completed" json:"dtCompleted,computed" format:"date-time"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
DtStarted timetypes.RFC3339 `tfsdk:"dt_started" json:"dtStarted,computed" format:"date-time"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
Error types.String `tfsdk:"error" json:"error,computed"`
ExternalID types.String `tfsdk:"external_id" json:"externalId,computed"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
TriggerType types.String `tfsdk:"trigger_type" json:"triggerType,computed"`
URL types.String `tfsdk:"url" json:"url,computed"`
}

func (m IntegrationConfigurationModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m IntegrationConfigurationModel) MarshalJSONForUpdate(state IntegrationConfigurationModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type IntegrationConfigurationCredentialsModel struct {
Type types.String `tfsdk:"type" json:"type,required"`
Destination types.String `tfsdk:"destination" json:"destination,optional"`
Empty types.Bool `tfsdk:"empty" json:"empty,optional"`
Name types.String `tfsdk:"name" json:"name,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
}
