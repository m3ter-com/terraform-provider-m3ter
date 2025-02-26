// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_destination

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*DataExportDestinationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the S3 bucket for the Export Destination.",
				Required:    true,
			},
			"code": schema.StringAttribute{
				Description: "The code of the Export Destination.",
				Required:    true,
			},
			"iam_role_arn": schema.StringAttribute{
				Description: "To enable m3ter to upload a Data Exports to your S3 bucket, the service has to assume an IAM role with PutObject permission for the specified `bucketName`. Create a suitable IAM role in your AWS account and enter ARN:\n\n**Formatting Constraints:**\n* The IAM role ARN must begin with \"arn:aws:iam\".\n* The general format required is: \"arn:aws:iam::<aws account id>:role/<role name>\". For example: \"arn:aws:iam::922609978421:role/IAMRole636\".\n* If the `iamRoleArn` used doesn't comply with this format, then an error message will be returned.\n\n**More Details:** For more details and examples of the Permission and Trust Policies you can use to create the required IAM Role ARN, see [Creating Data Export Destinations](https://www.m3ter.com/docs/guides/data-exports/creating-data-export-destinations) in our main User documentation.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Export Destination.",
				Required:    true,
			},
			"partition_order": schema.StringAttribute{
				Description: "Specify how you want the file path to be structured in your bucket destination - by Time first (Default) or Type first.\n\nType is dependent on whether the Export is for Usage data or Operational data:\n* **Usage.** Type is `measurements`.\n* **Operational.** Type is one of the entities for which operational data exports are available, such as `account`, `commitment`, `meter`, and so on.\n\nExample for Usage Data Export using .CSV format:\n* Time first:\n`{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`\n* Type first:\n`{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`\nAvailable values: \"TYPE_FIRST\", \"TIME_FIRST\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("TYPE_FIRST", "TIME_FIRST"),
				},
			},
			"prefix": schema.StringAttribute{
				Description: "Location in specified S3 bucket for the Export Destination. If you omit a `prefix`, then the root of the bucket will be used.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
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
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified the Export Destination.",
				Computed:    true,
			},
		},
	}
}

func (r *DataExportDestinationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *DataExportDestinationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
