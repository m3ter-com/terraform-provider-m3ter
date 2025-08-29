// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type IntegrationConfigurationModel struct {
	ID                       types.String                                                       `tfsdk:"id" json:"id,computed"`
	OrgID                    types.String                                                       `tfsdk:"org_id" path:"orgId,optional"`
	Destination              types.String                                                       `tfsdk:"destination" json:"destination,required"`
	EntityType               types.String                                                       `tfsdk:"entity_type" json:"entityType,required"`
	DestinationID            types.String                                                       `tfsdk:"destination_id" json:"destinationId,optional,no_refresh"`
	EntityID                 types.String                                                       `tfsdk:"entity_id" json:"entityId,optional"`
	IntegrationCredentialsID types.String                                                       `tfsdk:"integration_credentials_id" json:"integrationCredentialsId,optional,no_refresh"`
	Name                     types.String                                                       `tfsdk:"name" json:"name,optional,no_refresh"`
	ConfigData               customfield.Map[jsontypes.Normalized]                              `tfsdk:"config_data" json:"configData,computed_optional,no_refresh"`
	Credentials              customfield.NestedObject[IntegrationConfigurationCredentialsModel] `tfsdk:"credentials" json:"credentials,computed_optional,no_refresh"`
	Authorized               types.Bool                                                         `tfsdk:"authorized" json:"authorized,computed,no_refresh"`
	DtCompleted              timetypes.RFC3339                                                  `tfsdk:"dt_completed" json:"dtCompleted,computed" format:"date-time"`
	DtStarted                timetypes.RFC3339                                                  `tfsdk:"dt_started" json:"dtStarted,computed" format:"date-time"`
	Enabled                  types.Bool                                                         `tfsdk:"enabled" json:"enabled,computed,no_refresh"`
	Error                    types.String                                                       `tfsdk:"error" json:"error,computed"`
	ExternalID               types.String                                                       `tfsdk:"external_id" json:"externalId,computed"`
	Status                   types.String                                                       `tfsdk:"status" json:"status,computed"`
	TriggerType              types.String                                                       `tfsdk:"trigger_type" json:"triggerType,computed,no_refresh"`
	URL                      types.String                                                       `tfsdk:"url" json:"url,computed"`
	Version                  types.Int64                                                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m IntegrationConfigurationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m IntegrationConfigurationModel) MarshalJSONForUpdate(state IntegrationConfigurationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type IntegrationConfigurationCredentialsModel struct {
	Type        types.String `tfsdk:"type" json:"type,required"`
	Destination types.String `tfsdk:"destination" json:"destination,optional"`
	Empty       types.Bool   `tfsdk:"empty" json:"empty,optional"`
	Name        types.String `tfsdk:"name" json:"name,optional"`
	Version     types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}
