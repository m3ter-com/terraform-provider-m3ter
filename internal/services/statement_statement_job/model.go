// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_job

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type StatementStatementJobModel struct {
	ID                        types.String      `tfsdk:"id" json:"id,computed"`
	OrgID                     types.String      `tfsdk:"org_id" path:"orgId,optional"`
	BillID                    types.String      `tfsdk:"bill_id" json:"billId,required"`
	IncludeCsvFormat          types.Bool        `tfsdk:"include_csv_format" json:"includeCsvFormat,optional"`
	Version                   types.Int64       `tfsdk:"version" json:"version,optional"`
	CreatedBy                 types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                 timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified            timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy            types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	PresignedJsonStatementURL types.String      `tfsdk:"presigned_json_statement_url" json:"presignedJsonStatementUrl,computed"`
	StatementID               types.String      `tfsdk:"statement_id" json:"statementId,computed"`
	StatementJobStatus        types.String      `tfsdk:"statement_job_status" json:"statementJobStatus,computed"`
}

func (m StatementStatementJobModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StatementStatementJobModel) MarshalJSONForUpdate(state StatementStatementJobModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
