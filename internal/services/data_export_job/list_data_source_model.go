// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type DataExportJobsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[DataExportJobsItemsDataSourceModel] `json:"data,computed"`
}

type DataExportJobsDataSourceModel struct {
	OrgID            types.String                                                     `tfsdk:"org_id" path:"orgId,required"`
	DateCreatedEnd   types.String                                                     `tfsdk:"date_created_end" query:"dateCreatedEnd,optional"`
	DateCreatedStart types.String                                                     `tfsdk:"date_created_start" query:"dateCreatedStart,optional"`
	ScheduleID       types.String                                                     `tfsdk:"schedule_id" query:"scheduleId,optional"`
	Status           types.String                                                     `tfsdk:"status" query:"status,optional"`
	IDs              *[]types.String                                                  `tfsdk:"ids" query:"ids,optional"`
	MaxItems         types.Int64                                                      `tfsdk:"max_items"`
	Items            customfield.NestedObjectList[DataExportJobsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *DataExportJobsDataSourceModel) toListParams(_ context.Context) (params m3ter.DataExportJobListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.DataExportJobListParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
		IDs:   m3ter.F(mIDs),
	}

	if !m.DateCreatedEnd.IsNull() {
		params.DateCreatedEnd = m3ter.F(m.DateCreatedEnd.ValueString())
	}
	if !m.DateCreatedStart.IsNull() {
		params.DateCreatedStart = m3ter.F(m.DateCreatedStart.ValueString())
	}
	if !m.ScheduleID.IsNull() {
		params.ScheduleID = m3ter.F(m.ScheduleID.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = m3ter.F(m3ter.DataExportJobListParamsStatus(m.Status.ValueString()))
	}

	return
}

type DataExportJobsItemsDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	Version     types.Int64       `tfsdk:"version" json:"version,computed"`
	DateCreated timetypes.RFC3339 `tfsdk:"date_created" json:"dateCreated,computed" format:"date-time"`
	ScheduleID  types.String      `tfsdk:"schedule_id" json:"scheduleId,computed"`
	SourceType  types.String      `tfsdk:"source_type" json:"sourceType,computed"`
	StartedAt   timetypes.RFC3339 `tfsdk:"started_at" json:"startedAt,computed" format:"date-time"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
}
