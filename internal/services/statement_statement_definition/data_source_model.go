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

type StatementStatementDefinitionDataSourceModel struct {
	ID                   types.String                                                                        `tfsdk:"id" path:"id,required"`
	OrgID                types.String                                                                        `tfsdk:"org_id" path:"orgId,required"`
	AggregationFrequency types.String                                                                        `tfsdk:"aggregation_frequency" json:"aggregationFrequency,computed"`
	CreatedBy            types.String                                                                        `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated            timetypes.RFC3339                                                                   `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified       timetypes.RFC3339                                                                   `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	IncludePricePerUnit  types.Bool                                                                          `tfsdk:"include_price_per_unit" json:"includePricePerUnit,computed"`
	LastModifiedBy       types.String                                                                        `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name                 types.String                                                                        `tfsdk:"name" json:"name,computed"`
	Version              types.Int64                                                                         `tfsdk:"version" json:"version,computed"`
	Dimensions           customfield.NestedObjectList[StatementStatementDefinitionDimensionsDataSourceModel] `tfsdk:"dimensions" json:"dimensions,computed"`
	Measures             customfield.NestedObjectList[StatementStatementDefinitionMeasuresDataSourceModel]   `tfsdk:"measures" json:"measures,computed"`
}

func (m *StatementStatementDefinitionDataSourceModel) toReadParams(_ context.Context) (params m3ter.StatementStatementDefinitionGetParams, diags diag.Diagnostics) {
	params = m3ter.StatementStatementDefinitionGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type StatementStatementDefinitionDimensionsDataSourceModel struct {
	DimensionAttributes customfield.List[types.String] `tfsdk:"dimension_attributes" json:"dimensionAttributes,computed"`
	DimensionName       types.String                   `tfsdk:"dimension_name" json:"dimensionName,computed"`
}

type StatementStatementDefinitionMeasuresDataSourceModel struct {
	Aggregations customfield.List[types.String] `tfsdk:"aggregations" json:"aggregations,computed"`
	MeterID      types.String                   `tfsdk:"meter_id" json:"meterId,computed"`
	Name         types.String                   `tfsdk:"name" json:"name,computed"`
}
