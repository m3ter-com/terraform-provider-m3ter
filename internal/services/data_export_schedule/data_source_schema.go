// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*DataExportScheduleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"time_period": schema.StringAttribute{
				Description: "Define a time period to control the range of usage data you want the data export to contain when it runs:\n\n* **TODAY**. Data collected for the current day up until the time the export runs.\n* **YESTERDAY**. Data collected for the day before the export runs - that is, the 24 hour period from midnight to midnight of the day before.\n* **WEEK_TO_DATE**. Data collected for the period covering the current week to the date and time the export runs, and weeks run Monday to Monday.\n* **CURRENT_MONTH**. Data collected for the current month in which the export is ran up to and including the date and time the export runs.\n* **LAST_30_DAYS**. Data collected for the 30 days prior to the date the export is ran.\n* **LAST_35_DAYS**. Data collected for the 35 days prior to the date the export is ran.\n* **PREVIOUS_WEEK**. Data collected for the previous full week period, and weeks run Monday to Monday.\n* **PREVIOUS_MONTH**. Data collected for the previous full month period.\n\nFor more details and examples, see the [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period) section in our main User Documentation.\nAvailable values: \"LAST_12_HOURS\", \"TODAY\", \"YESTERDAY\", \"WEEK_TO_DATE\", \"CURRENT_MONTH\", \"LAST_30_DAYS\", \"LAST_35_DAYS\", \"PREVIOUS_WEEK\", \"PREVIOUS_MONTH\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"LAST_12_HOURS",
						"TODAY",
						"YESTERDAY",
						"WEEK_TO_DATE",
						"CURRENT_MONTH",
						"LAST_30_DAYS",
						"LAST_35_DAYS",
						"PREVIOUS_WEEK",
						"PREVIOUS_MONTH",
					),
				},
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"account_ids": schema.ListAttribute{
				Description: "List of account IDs for which the usage data will be exported.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"meter_ids": schema.ListAttribute{
				Description: "List of meter IDs for which the usage data will be exported.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"operational_data_types": schema.ListAttribute{
				Description: "A list of the entities whose operational data is included in the data export.",
				Computed:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"BILLS",
							"COMMITMENTS",
							"ACCOUNTS",
							"BALANCES",
							"CONTRACTS",
							"ACCOUNT_PLANS",
							"AGGREGATIONS",
							"PLANS",
							"PRICING",
							"PRICING_BANDS",
							"BILL_LINE_ITEMS",
							"METERS",
							"PRODUCTS",
							"COMPOUND_AGGREGATIONS",
							"PLAN_GROUPS",
							"PLAN_GROUP_LINKS",
							"PLAN_TEMPLATES",
							"BALANCE_TRANSACTIONS",
						),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"aggregations": schema.ListNestedAttribute{
				Description: "List of aggregations to apply",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[DataExportScheduleAggregationsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"field_code": schema.StringAttribute{
							Description: "Field code",
							Computed:    true,
						},
						"field_type": schema.StringAttribute{
							Description: "Type of field\nAvailable values: \"DIMENSION\", \"MEASURE\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("DIMENSION", "MEASURE"),
							},
						},
						"function": schema.StringAttribute{
							Description: "Aggregation function\nAvailable values: \"SUM\", \"MIN\", \"MAX\", \"COUNT\", \"LATEST\", \"MEAN\", \"UNIQUE\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"SUM",
									"MIN",
									"MAX",
									"COUNT",
									"LATEST",
									"MEAN",
									"UNIQUE",
								),
							},
						},
						"meter_id": schema.StringAttribute{
							Description: "Meter ID",
							Computed:    true,
						},
					},
				},
			},
			"dimension_filters": schema.ListNestedAttribute{
				Description: "List of dimension filters to apply",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[DataExportScheduleDimensionFiltersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"field_code": schema.StringAttribute{
							Description: "Field code",
							Computed:    true,
						},
						"meter_id": schema.StringAttribute{
							Description: "Meter ID",
							Computed:    true,
						},
						"values": schema.ListAttribute{
							Description: "Values to filter by",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
			"groups": schema.ListNestedAttribute{
				Description: "List of groups to apply",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[DataExportScheduleGroupsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_type": schema.StringAttribute{
							Description: `Available values: "ACCOUNT", "DIMENSION", "TIME".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"ACCOUNT",
									"DIMENSION",
									"TIME",
								),
							},
						},
						"field_code": schema.StringAttribute{
							Description: "Field code to group by",
							Computed:    true,
						},
						"meter_id": schema.StringAttribute{
							Description: "Meter ID to group by",
							Computed:    true,
						},
						"frequency": schema.StringAttribute{
							Description: "Frequency of usage data\nAvailable values: \"DAY\", \"HOUR\", \"WEEK\", \"MONTH\", \"QUARTER\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"DAY",
									"HOUR",
									"WEEK",
									"MONTH",
									"QUARTER",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *DataExportScheduleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DataExportScheduleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
