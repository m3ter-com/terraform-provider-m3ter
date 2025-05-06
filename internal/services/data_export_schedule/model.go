// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type DataExportScheduleModel struct {
	ID                   types.String                                `tfsdk:"id" json:"id,computed"`
	OrgID                types.String                                `tfsdk:"org_id" path:"orgId,optional"`
	SourceType           types.String                                `tfsdk:"source_type" json:"sourceType,required,no_refresh"`
	TimePeriod           types.String                                `tfsdk:"time_period" json:"timePeriod,optional"`
	Version              types.Int64                                 `tfsdk:"version" json:"version,optional"`
	AccountIDs           *[]types.String                             `tfsdk:"account_ids" json:"accountIds,optional"`
	MeterIDs             *[]types.String                             `tfsdk:"meter_ids" json:"meterIds,optional"`
	OperationalDataTypes *[]types.String                             `tfsdk:"operational_data_types" json:"operationalDataTypes,optional"`
	Aggregations         *[]*DataExportScheduleAggregationsModel     `tfsdk:"aggregations" json:"aggregations,optional"`
	DimensionFilters     *[]*DataExportScheduleDimensionFiltersModel `tfsdk:"dimension_filters" json:"dimensionFilters,optional"`
	Groups               *[]*DataExportScheduleGroupsModel           `tfsdk:"groups" json:"groups,optional"`
}

func (m DataExportScheduleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DataExportScheduleModel) MarshalJSONForUpdate(state DataExportScheduleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type DataExportScheduleAggregationsModel struct {
	FieldCode types.String `tfsdk:"field_code" json:"fieldCode,required"`
	FieldType types.String `tfsdk:"field_type" json:"fieldType,required"`
	Function  types.String `tfsdk:"function" json:"function,required"`
	MeterID   types.String `tfsdk:"meter_id" json:"meterId,required"`
}

type DataExportScheduleDimensionFiltersModel struct {
	FieldCode types.String    `tfsdk:"field_code" json:"fieldCode,required"`
	MeterID   types.String    `tfsdk:"meter_id" json:"meterId,required"`
	Values    *[]types.String `tfsdk:"values" json:"values,required"`
}

type DataExportScheduleGroupsModel struct {
	GroupType types.String `tfsdk:"group_type" json:"groupType,optional"`
	FieldCode types.String `tfsdk:"field_code" json:"fieldCode,optional"`
	MeterID   types.String `tfsdk:"meter_id" json:"meterId,optional"`
	Frequency types.String `tfsdk:"frequency" json:"frequency,optional"`
}
