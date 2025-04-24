// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type DataExportScheduleDataSourceModel struct {
	ID                   types.String                                                                    `tfsdk:"id" path:"id,required"`
	OrgID                types.String                                                                    `tfsdk:"org_id" path:"orgId,required"`
	TimePeriod           types.String                                                                    `tfsdk:"time_period" json:"timePeriod,computed"`
	Version              types.Int64                                                                     `tfsdk:"version" json:"version,computed"`
	AccountIDs           customfield.List[types.String]                                                  `tfsdk:"account_ids" json:"accountIds,computed"`
	MeterIDs             customfield.List[types.String]                                                  `tfsdk:"meter_ids" json:"meterIds,computed"`
	OperationalDataTypes customfield.List[types.String]                                                  `tfsdk:"operational_data_types" json:"operationalDataTypes,computed"`
	Aggregations         customfield.NestedObjectList[DataExportScheduleAggregationsDataSourceModel]     `tfsdk:"aggregations" json:"aggregations,computed"`
	DimensionFilters     customfield.NestedObjectList[DataExportScheduleDimensionFiltersDataSourceModel] `tfsdk:"dimension_filters" json:"dimensionFilters,computed"`
	Groups               customfield.NestedObjectList[DataExportScheduleGroupsDataSourceModel]           `tfsdk:"groups" json:"groups,computed"`
}

func (m *DataExportScheduleDataSourceModel) toReadParams(_ context.Context) (params m3ter.DataExportScheduleGetParams, diags diag.Diagnostics) {
	params = m3ter.DataExportScheduleGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type DataExportScheduleAggregationsDataSourceModel struct {
	FieldCode types.String `tfsdk:"field_code" json:"fieldCode,computed"`
	FieldType types.String `tfsdk:"field_type" json:"fieldType,computed"`
	Function  types.String `tfsdk:"function" json:"function,computed"`
	MeterID   types.String `tfsdk:"meter_id" json:"meterId,computed"`
}

type DataExportScheduleDimensionFiltersDataSourceModel struct {
	FieldCode types.String                   `tfsdk:"field_code" json:"fieldCode,computed"`
	MeterID   types.String                   `tfsdk:"meter_id" json:"meterId,computed"`
	Values    customfield.List[types.String] `tfsdk:"values" json:"values,computed"`
}

type DataExportScheduleGroupsDataSourceModel struct {
	GroupType types.String `tfsdk:"group_type" json:"groupType,computed"`
	FieldCode types.String `tfsdk:"field_code" json:"fieldCode,computed"`
	MeterID   types.String `tfsdk:"meter_id" json:"meterId,computed"`
	Frequency types.String `tfsdk:"frequency" json:"frequency,computed"`
}
