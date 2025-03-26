// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule

import (
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type DataExportScheduleModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
SourceType types.String `tfsdk:"source_type" json:"sourceType,required"`
Aggregation types.String `tfsdk:"aggregation" json:"aggregation,optional"`
AggregationFrequency types.String `tfsdk:"aggregation_frequency" json:"aggregationFrequency,optional"`
TimePeriod types.String `tfsdk:"time_period" json:"timePeriod,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
AccountIDs *[]types.String `tfsdk:"account_ids" json:"accountIds,optional"`
MeterIDs *[]types.String `tfsdk:"meter_ids" json:"meterIds,optional"`
OperationalDataTypes *[]types.String `tfsdk:"operational_data_types" json:"operationalDataTypes,optional"`
}

func (m DataExportScheduleModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m DataExportScheduleModel) MarshalJSONForUpdate(state DataExportScheduleModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
