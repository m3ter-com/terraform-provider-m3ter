// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*StatementStatementJobDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"bill_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the bill associated with the StatementJob.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this StatementJob.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the StatementJob was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the StatementJob was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"include_csv_format": schema.BoolAttribute{
				Description: "A Boolean value indicating whether the generated statement includes a CSV format.\n\n* TRUE - includes the statement in CSV format.\n* FALSE - no CSV format statement.",
				Computed:    true,
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this StatementJob.",
				Computed:    true,
			},
			"presigned_json_statement_url": schema.StringAttribute{
				Description: "The URL to access the generated statement in JSON format. This URL is temporary and has a limited lifetime.",
				Computed:    true,
			},
			"statement_job_status": schema.StringAttribute{
				Description: "The current status of the StatementJob. The status helps track the progress and outcome of a StatementJob.\nAvailable values: \"PENDING\", \"RUNNING\", \"COMPLETE\", \"CANCELLED\", \"FAILED\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PENDING",
						"RUNNING",
						"COMPLETE",
						"CANCELLED",
						"FAILED",
					),
				},
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (d *StatementStatementJobDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StatementStatementJobDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
