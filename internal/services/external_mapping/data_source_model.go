// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package external_mapping

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type ExternalMappingDataSourceModel struct {
	ID                  types.String                             `tfsdk:"id" path:"id,computed_optional"`
	OrgID               types.String                             `tfsdk:"org_id" path:"orgId,optional"`
	ExternalID          types.String                             `tfsdk:"external_id" json:"externalId,computed"`
	ExternalSystem      types.String                             `tfsdk:"external_system" json:"externalSystem,computed"`
	ExternalTable       types.String                             `tfsdk:"external_table" json:"externalTable,computed"`
	IntegrationConfigID types.String                             `tfsdk:"integration_config_id" json:"integrationConfigId,computed"`
	M3terEntity         types.String                             `tfsdk:"m3ter_entity" json:"m3terEntity,computed"`
	M3terID             types.String                             `tfsdk:"m3ter_id" json:"m3terId,computed"`
	Version             types.Int64                              `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	FindOneBy           *ExternalMappingFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *ExternalMappingDataSourceModel) toReadParams(_ context.Context) (params m3ter.ExternalMappingGetParams, diags diag.Diagnostics) {
	params = m3ter.ExternalMappingGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *ExternalMappingDataSourceModel) toListParams(_ context.Context) (params m3ter.ExternalMappingListParams, diags diag.Diagnostics) {
	mFindOneByM3terIDs := []string{}
	if m.FindOneBy.M3terIDs != nil {
		for _, item := range *m.FindOneBy.M3terIDs {
			mFindOneByM3terIDs = append(mFindOneByM3terIDs, item.ValueString())
		}
	}

	params = m3ter.ExternalMappingListParams{
		M3terIDs: m3ter.F(mFindOneByM3terIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}
	if !m.FindOneBy.ExternalSystemID.IsNull() {
		params.ExternalSystemID = m3ter.F(m.FindOneBy.ExternalSystemID.ValueString())
	}
	if !m.FindOneBy.IntegrationConfigID.IsNull() {
		params.IntegrationConfigID = m3ter.F(m.FindOneBy.IntegrationConfigID.ValueString())
	}

	return
}

type ExternalMappingFindOneByDataSourceModel struct {
	ExternalSystemID    types.String    `tfsdk:"external_system_id" query:"externalSystemId,optional"`
	IntegrationConfigID types.String    `tfsdk:"integration_config_id" query:"integrationConfigId,optional"`
	M3terIDs            *[]types.String `tfsdk:"m3ter_ids" query:"m3terIds,optional"`
}
