// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package usage_file_upload_job

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type UsageFileUploadJobsDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[UsageFileUploadJobsItemsDataSourceModel] `json:"data,computed"`
}

type UsageFileUploadJobsDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
DateCreatedEnd types.String `tfsdk:"date_created_end" query:"dateCreatedEnd,optional"`
DateCreatedStart types.String `tfsdk:"date_created_start" query:"dateCreatedStart,optional"`
FileKey types.String `tfsdk:"file_key" query:"fileKey,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[UsageFileUploadJobsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *UsageFileUploadJobsDataSourceModel) toListParams(_ context.Context) (params m3ter.UsageFileUploadJobListParams, diags diag.Diagnostics) {
  params = m3ter.UsageFileUploadJobListParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
  }

  if !m.DateCreatedEnd.IsNull() {
    params.DateCreatedEnd = m3ter.F(m.DateCreatedEnd.ValueString())
  }
  if !m.DateCreatedStart.IsNull() {
    params.DateCreatedStart = m3ter.F(m.DateCreatedStart.ValueString())
  }
  if !m.FileKey.IsNull() {
    params.FileKey = m3ter.F(m.FileKey.ValueString())
  }

  return
}

type UsageFileUploadJobsItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
ContentLength types.Int64 `tfsdk:"content_length" json:"contentLength,computed"`
FailedRows types.Int64 `tfsdk:"failed_rows" json:"failedRows,computed"`
FileName types.String `tfsdk:"file_name" json:"fileName,computed"`
ProcessedRows types.Int64 `tfsdk:"processed_rows" json:"processedRows,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
TotalRows types.Int64 `tfsdk:"total_rows" json:"totalRows,computed"`
UploadDate types.String `tfsdk:"upload_date" json:"uploadDate,computed"`
Version types.Int64 `tfsdk:"version" json:"version,computed"`
}
