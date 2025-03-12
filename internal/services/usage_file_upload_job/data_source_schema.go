// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package usage_file_upload_job

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*UsageFileUploadJobDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Required: true,
      },
      "org_id": schema.StringAttribute{
        Required: true,
      },
      "content_length": schema.Int64Attribute{
        Description: "The size of the body in bytes. For example: `\"contentLength\": 485`, where 485 is the size in bytes of the file uploaded.",
        Computed: true,
      },
      "failed_rows": schema.Int64Attribute{
        Description: "The number of rows that failed processing during ingest.",
        Computed: true,
      },
      "file_name": schema.StringAttribute{
        Description: "The name of the measurements file for the upload job.",
        Computed: true,
      },
      "processed_rows": schema.Int64Attribute{
        Description: "The number of rows that were processed during ingest.",
        Computed: true,
      },
      "status": schema.StringAttribute{
        Description: "The status of the file upload job.\nAvailable values: \"notUploaded\", \"running\", \"failed\", \"succeeded\".",
        Computed: true,
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
        Computed: true,
      },
      "upload_date": schema.StringAttribute{
        Description: "The upload date for the upload job *(in ISO-8601 format)*.",
        Computed: true,
      },
      "version": schema.Int64Attribute{
        Description: "The version number. Default value when newly created is one.",
        Computed: true,
      },
    },
  }
}

func (d *UsageFileUploadJobDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *UsageFileUploadJobDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
