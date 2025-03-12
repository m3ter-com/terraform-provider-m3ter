// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package external_mapping

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type ExternalMappingModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
ExternalID types.String `tfsdk:"external_id" json:"externalId,required"`
ExternalSystem types.String `tfsdk:"external_system" json:"externalSystem,required"`
ExternalTable types.String `tfsdk:"external_table" json:"externalTable,required"`
M3terEntity types.String `tfsdk:"m3ter_entity" json:"m3terEntity,required"`
M3terID types.String `tfsdk:"m3ter_id" json:"m3terId,required"`
IntegrationConfigID types.String `tfsdk:"integration_config_id" json:"integrationConfigId,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m ExternalMappingModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ExternalMappingModel) MarshalJSONForUpdate(state ExternalMappingModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
