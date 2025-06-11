// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type StatementStatementDefinitionsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[StatementStatementDefinitionsItemsDataSourceModel] `json:"data,computed"`
}

type StatementStatementDefinitionsDataSourceModel struct {
	OrgID    types.String                                                                    `tfsdk:"org_id" path:"orgId,optional"`
	MaxItems types.Int64                                                                     `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[StatementStatementDefinitionsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *StatementStatementDefinitionsDataSourceModel) toListParams(_ context.Context) (params m3ter.StatementStatementDefinitionListParams, diags diag.Diagnostics) {
	params = m3ter.StatementStatementDefinitionListParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type StatementStatementDefinitionsItemsDataSourceModel struct {
	ID                   types.String                                                                         `tfsdk:"id" json:"id,computed"`
	Version              types.Int64                                                                          `tfsdk:"version" json:"version,computed"`
	AggregationFrequency types.String                                                                         `tfsdk:"aggregation_frequency" json:"aggregationFrequency,computed"`
	CreatedBy            types.String                                                                         `tfsdk:"created_by" json:"createdBy,computed"`
	Dimensions           customfield.NestedObjectList[StatementStatementDefinitionsDimensionsDataSourceModel] `tfsdk:"dimensions" json:"dimensions,computed"`
	DtCreated            timetypes.RFC3339                                                                    `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified       timetypes.RFC3339                                                                    `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	IncludePricePerUnit  types.Bool                                                                           `tfsdk:"include_price_per_unit" json:"includePricePerUnit,computed"`
	LastModifiedBy       types.String                                                                         `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Measures             customfield.NestedObjectList[StatementStatementDefinitionsMeasuresDataSourceModel]   `tfsdk:"measures" json:"measures,computed"`
	Name                 types.String                                                                         `tfsdk:"name" json:"name,computed"`
}

type StatementStatementDefinitionsDimensionsDataSourceModel struct {
	DimensionAttributes customfield.List[types.String] `tfsdk:"dimension_attributes" json:"dimensionAttributes,computed"`
	DimensionName       types.String                   `tfsdk:"dimension_name" json:"dimensionName,computed"`
}

type StatementStatementDefinitionsMeasuresDataSourceModel struct {
	Aggregations customfield.List[types.String] `tfsdk:"aggregations" json:"aggregations,computed"`
	MeterID      types.String                   `tfsdk:"meter_id" json:"meterId,computed"`
	Name         types.String                   `tfsdk:"name" json:"name,computed"`
}
