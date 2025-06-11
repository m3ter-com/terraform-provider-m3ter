// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type StatementStatementDefinitionModel struct {
	ID                   types.String                                    `tfsdk:"id" json:"id,computed"`
	OrgID                types.String                                    `tfsdk:"org_id" path:"orgId,optional"`
	AggregationFrequency types.String                                    `tfsdk:"aggregation_frequency" json:"aggregationFrequency,required"`
	IncludePricePerUnit  types.Bool                                      `tfsdk:"include_price_per_unit" json:"includePricePerUnit,optional"`
	Name                 types.String                                    `tfsdk:"name" json:"name,optional"`
	Version              types.Int64                                     `tfsdk:"version" json:"version,optional"`
	Dimensions           *[]*StatementStatementDefinitionDimensionsModel `tfsdk:"dimensions" json:"dimensions,optional"`
	Measures             *[]*StatementStatementDefinitionMeasuresModel   `tfsdk:"measures" json:"measures,optional"`
	CreatedBy            types.String                                    `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated            timetypes.RFC3339                               `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified       timetypes.RFC3339                               `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy       types.String                                    `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m StatementStatementDefinitionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StatementStatementDefinitionModel) MarshalJSONForUpdate(state StatementStatementDefinitionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type StatementStatementDefinitionDimensionsModel struct {
	DimensionAttributes *[]types.String `tfsdk:"dimension_attributes" json:"dimensionAttributes,optional"`
	DimensionName       types.String    `tfsdk:"dimension_name" json:"dimensionName,optional"`
}

type StatementStatementDefinitionMeasuresModel struct {
	Aggregations *[]types.String `tfsdk:"aggregations" json:"aggregations,optional"`
	MeterID      types.String    `tfsdk:"meter_id" json:"meterId,optional"`
	Name         types.String    `tfsdk:"name" json:"name,optional"`
}
