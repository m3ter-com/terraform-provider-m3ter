// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*DataExportJobsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"date_created_end": schema.StringAttribute{
				Description: "Include only Job entities created before this date. Format: yyyy-MM-dd'T'HH:mm:ss'Z'",
				Optional:    true,
			},
			"date_created_start": schema.StringAttribute{
				Description: "Include only Job entities created on or after this date. Format: yyyy-MM-dd'T'HH:mm:ss'Z'",
				Optional:    true,
			},
			"schedule_id": schema.StringAttribute{
				Description: "List Job entities for the schedule UUID",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "List Job entities for the status\navailable values: \"PENDING\", \"RUNNING\", \"SUCCEEDED\", \"FAILED\"",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PENDING",
						"RUNNING",
						"SUCCEEDED",
						"FAILED",
					),
				},
			},
			"ids": schema.ListAttribute{
				Description: "List Job entities for the given UUIDs",
				Optional:    true,
				ElementType: types.StringType,
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
				CustomType:  customfield.NewNestedObjectListType[DataExportJobsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the Export Job.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
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
					},
				},
			},
		},
	}
}

func (d *DataExportJobsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *DataExportJobsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
