// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package usage_file_upload_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type UsageFileUploadJobDataSourceModel struct {
	ID            types.String `tfsdk:"id" path:"id,required"`
	OrgID         types.String `tfsdk:"org_id" path:"orgId,required"`
	ContentLength types.Int64  `tfsdk:"content_length" json:"contentLength,computed"`
	FailedRows    types.Int64  `tfsdk:"failed_rows" json:"failedRows,computed"`
	FileName      types.String `tfsdk:"file_name" json:"fileName,computed"`
	ProcessedRows types.Int64  `tfsdk:"processed_rows" json:"processedRows,computed"`
	Status        types.String `tfsdk:"status" json:"status,computed"`
	TotalRows     types.Int64  `tfsdk:"total_rows" json:"totalRows,computed"`
	UploadDate    types.String `tfsdk:"upload_date" json:"uploadDate,computed"`
	Version       types.Int64  `tfsdk:"version" json:"version,computed"`
}

func (m *UsageFileUploadJobDataSourceModel) toReadParams(_ context.Context) (params m3ter.UsageFileUploadJobGetParams, diags diag.Diagnostics) {
	params = m3ter.UsageFileUploadJobGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
