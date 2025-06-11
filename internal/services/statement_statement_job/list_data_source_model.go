// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type StatementStatementJobsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[StatementStatementJobsItemsDataSourceModel] `json:"data,computed"`
}

type StatementStatementJobsDataSourceModel struct {
	OrgID    types.String                                                             `tfsdk:"org_id" path:"orgId,optional"`
	Active   types.String                                                             `tfsdk:"active" query:"active,optional"`
	BillID   types.String                                                             `tfsdk:"bill_id" query:"billId,optional"`
	Status   types.String                                                             `tfsdk:"status" query:"status,optional"`
	MaxItems types.Int64                                                              `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[StatementStatementJobsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *StatementStatementJobsDataSourceModel) toListParams(_ context.Context) (params m3ter.StatementStatementJobListParams, diags diag.Diagnostics) {
	params = m3ter.StatementStatementJobListParams{}

	if !m.Active.IsNull() {
		params.Active = m3ter.F(m.Active.ValueString())
	}
	if !m.BillID.IsNull() {
		params.BillID = m3ter.F(m.BillID.ValueString())
	}
	if !m.Status.IsNull() {
		params.Status = m3ter.F(m.Status.ValueString())
	}
	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type StatementStatementJobsItemsDataSourceModel struct {
	ID                        types.String      `tfsdk:"id" json:"id,computed"`
	Version                   types.Int64       `tfsdk:"version" json:"version,computed"`
	BillID                    types.String      `tfsdk:"bill_id" json:"billId,computed"`
	CreatedBy                 types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                 timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified            timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	IncludeCsvFormat          types.Bool        `tfsdk:"include_csv_format" json:"includeCsvFormat,computed"`
	LastModifiedBy            types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	OrgID                     types.String      `tfsdk:"org_id" json:"orgId,computed"`
	PresignedJsonStatementURL types.String      `tfsdk:"presigned_json_statement_url" json:"presignedJsonStatementUrl,computed"`
	StatementJobStatus        types.String      `tfsdk:"statement_job_status" json:"statementJobStatus,computed"`
}
