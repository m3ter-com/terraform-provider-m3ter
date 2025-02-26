// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package usage_file_upload_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*UsageFileUploadJobsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"date_created_end": schema.StringAttribute{
				Description: "Include only File Upload jobs created before this date. Required format is ISO-8601: yyyy-MM-dd'T'HH:mm:ss'Z'",
				Optional:    true,
			},
			"date_created_start": schema.StringAttribute{
				Description: "Include only File Upload jobs created on or after this date. Required format is ISO-8601: yyyy-MM-dd'T'HH:mm:ss'Z'",
				Optional:    true,
			},
			"file_key": schema.StringAttribute{
				Description: "<<deprecated>>",
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
				CustomType:  customfield.NewNestedObjectListType[UsageFileUploadJobsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID of the file upload job.",
							Computed:    true,
						},
						"content_length": schema.Int64Attribute{
							Description: "The size of the body in bytes. For example: `\"contentLength\": 485`, where 485 is the size in bytes of the file uploaded.",
							Computed:    true,
						},
						"failed_rows": schema.Int64Attribute{
							Description: "The number of rows that failed processing during ingest.",
							Computed:    true,
						},
						"file_name": schema.StringAttribute{
							Description: "The name of the measurements file for the upload job.",
							Computed:    true,
						},
						"processed_rows": schema.Int64Attribute{
							Description: "The number of rows that were processed during ingest.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "The status of the file upload job.\navailable values: \"notUploaded\", \"running\", \"failed\", \"succeeded\"",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"notUploaded",
									"running",
									"failed",
									"succeeded",
								),
							},
						},
						"total_rows": schema.Int64Attribute{
							Description: "The total number of rows in the file.",
							Computed:    true,
						},
						"upload_date": schema.StringAttribute{
							Description: "The upload date for the upload job *(in ISO-8601 format)*.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number. Default value when newly created is one.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *UsageFileUploadJobsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *UsageFileUploadJobsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
