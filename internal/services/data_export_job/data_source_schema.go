// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*DataExportJobDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"date_created": schema.StringAttribute{
				Description: "When the data Export Job was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"schedule_id": schema.StringAttribute{
				Description: "The id of the data Export Schedule.",
				Computed:    true,
			},
			"source_type": schema.StringAttribute{
				Description: "available values: \"USAGE\", \"OPERATIONAL\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("USAGE", "OPERATIONAL"),
				},
			},
			"started_at": schema.StringAttribute{
				Description: "When the data Export Job started running",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "available values: \"PENDING\", \"RUNNING\", \"SUCCEEDED\", \"FAILED\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PENDING",
						"RUNNING",
						"SUCCEEDED",
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

func (d *DataExportJobDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DataExportJobDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
