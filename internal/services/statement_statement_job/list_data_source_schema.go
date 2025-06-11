// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*StatementStatementJobsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"active": schema.StringAttribute{
				Description: "Boolean filter on whether to only retrieve active *(i.e. not completed/cancelled)* StatementJobs.\n\n* TRUE - only active StatementJobs retrieved.\n* FALSE - all StatementJobs retrieved.",
				Optional:    true,
			},
			"bill_id": schema.StringAttribute{
				Description: "Filter Statement Jobs by billId",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Filter using the StatementJobs status. Possible values:\n\n* `PENDING`\n* `RUNNING`\n* `COMPLETE`\n* `CANCELLED`\n* `FAILED`",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[StatementStatementJobsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
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
						"org_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of your Organization. The Organization represents your company as a direct customer of our service.",
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
					},
				},
			},
		},
	}
}

func (d *StatementStatementJobsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *StatementStatementJobsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
