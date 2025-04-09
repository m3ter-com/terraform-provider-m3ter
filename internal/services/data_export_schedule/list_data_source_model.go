// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type DataExportSchedulesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[DataExportSchedulesItemsDataSourceModel] `json:"data,computed"`
}

type DataExportSchedulesDataSourceModel struct {
	OrgID    types.String                                                          `tfsdk:"org_id" path:"orgId,optional"`
	IDs      *[]types.String                                                       `tfsdk:"ids" query:"ids,optional"`
	MaxItems types.Int64                                                           `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[DataExportSchedulesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *DataExportSchedulesDataSourceModel) toListParams(_ context.Context) (params m3ter.DataExportScheduleListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.DataExportScheduleListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type DataExportSchedulesItemsDataSourceModel struct {
	ID               types.String                   `tfsdk:"id" json:"id,computed"`
	Version          types.Int64                    `tfsdk:"version" json:"version,computed"`
	Code             types.String                   `tfsdk:"code" json:"code,computed"`
	CreatedBy        types.String                   `tfsdk:"created_by" json:"createdBy,computed"`
	DestinationIDs   customfield.List[types.String] `tfsdk:"destination_ids" json:"destinationIds,computed"`
	DtCreated        timetypes.RFC3339              `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified   timetypes.RFC3339              `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	ExportFileFormat types.String                   `tfsdk:"export_file_format" json:"exportFileFormat,computed"`
	LastModifiedBy   types.String                   `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name             types.String                   `tfsdk:"name" json:"name,computed"`
	Period           types.Int64                    `tfsdk:"period" json:"period,computed"`
	ScheduleType     types.String                   `tfsdk:"schedule_type" json:"scheduleType,computed"`
	SourceType       types.String                   `tfsdk:"source_type" json:"sourceType,computed"`
}
