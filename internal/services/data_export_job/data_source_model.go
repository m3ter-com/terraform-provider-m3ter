// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_job

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
)

type DataExportJobDataSourceModel struct {
ID types.String `tfsdk:"id" path:"id,required"`
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
DateCreated timetypes.RFC3339 `tfsdk:"date_created" json:"dateCreated,computed" format:"date-time"`
ScheduleID types.String `tfsdk:"schedule_id" json:"scheduleId,computed"`
SourceType types.String `tfsdk:"source_type" json:"sourceType,computed"`
StartedAt timetypes.RFC3339 `tfsdk:"started_at" json:"startedAt,computed" format:"date-time"`
Status types.String `tfsdk:"status" json:"status,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
}

func (m *DataExportJobDataSourceModel) toReadParams(_ context.Context) (params m3ter.DataExportJobGetParams, diags diag.Diagnostics) {
  params = m3ter.DataExportJobGetParams{

  }

  if !m.OrgID.IsNull() {
    params.OrgID = m3ter.F(m.OrgID.ValueString())
  }

  return
}
