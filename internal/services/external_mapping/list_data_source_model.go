// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package external_mapping

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ExternalMappingsDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[ExternalMappingsItemsDataSourceModel] `json:"data,computed"`
}

type ExternalMappingsDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
ExternalSystemID types.String `tfsdk:"external_system_id" query:"externalSystemId,optional"`
IntegrationConfigID types.String `tfsdk:"integration_config_id" query:"integrationConfigId,optional"`
M3terIDs *[]types.String `tfsdk:"m3ter_ids" query:"m3terIds,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[ExternalMappingsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ExternalMappingsDataSourceModel) toListParams(_ context.Context) (params m3ter.ExternalMappingListParams, diags diag.Diagnostics) {
  mM3terIDs := []string{}
  for _, item := range *m.M3terIDs {
    mM3terIDs = append(mM3terIDs, item.ValueString())
  }

  params = m3ter.ExternalMappingListParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
    M3terIDs: m3ter.F(mM3terIDs),
  }

  if !m.ExternalSystemID.IsNull() {
    params.ExternalSystemID = m3ter.F(m.ExternalSystemID.ValueString())
  }
  if !m.IntegrationConfigID.IsNull() {
    params.IntegrationConfigID = m3ter.F(m.IntegrationConfigID.ValueString())
  }

  return
}

type ExternalMappingsItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
ExternalID types.String `tfsdk:"external_id" json:"externalId,computed"`
ExternalSystem types.String `tfsdk:"external_system" json:"externalSystem,computed"`
ExternalTable types.String `tfsdk:"external_table" json:"externalTable,computed"`
M3terEntity types.String `tfsdk:"m3ter_entity" json:"m3terEntity,computed"`
M3terID types.String `tfsdk:"m3ter_id" json:"m3terId,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
IntegrationConfigID types.String `tfsdk:"integration_config_id" json:"integrationConfigId,computed"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}
