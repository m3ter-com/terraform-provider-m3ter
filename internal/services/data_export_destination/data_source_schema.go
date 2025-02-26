// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_destination

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*DataExportDestinationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the S3 bucket for the Export Destination.",
				Computed:    true,
			},
			"code": schema.StringAttribute{
				Description: "The code of the data Export Destination.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created the Export Destination.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when the Export Destination was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when the Export Destination was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"iam_role_arn": schema.StringAttribute{
				Description: "The specified IAM role ARN with PutObject permission for the specified `bucketName`, which allows the service to upload Data Exports to your S3 bucket.",
				Computed:    true,
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified the Export Destination.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the data Export Destination.",
				Computed:    true,
			},
			"partition_order": schema.StringAttribute{
				Description: "Specify how you want the file path to be structured in your bucket destination - by Time first (Default) or Type first.\n\nType is dependent on whether the Export is for Usage data or Operational data:\n* **Usage.** Type is `measurements`.\n* **Operational.** Type is one of the entities for which operational data exports are available, such as `account`, `commitment`, `meter`, and so on.\n\nExample for Usage Data Export using .CSV format:\n* Time first:\n`{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`\n* Type first:\n`{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`\nAvailable values: \"TYPE_FIRST\", \"TIME_FIRST\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("TYPE_FIRST", "TIME_FIRST"),
				},
			},
			"prefix": schema.StringAttribute{
				Description: "Location in specified S3 bucket for the Export Destination. If no `prefix` is specified, then the root of the bucket is used.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (d *DataExportDestinationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DataExportDestinationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
