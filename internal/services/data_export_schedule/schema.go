// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*DataExportScheduleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The id of the schedule.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"source_type": schema.StringAttribute{
				Description: "The type of data to export. Possible values are: OPERATIONAL\nAvailable values: \"OPERATIONAL\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("OPERATIONAL", "USAGE"),
				},
			},
			"time_period": schema.StringAttribute{
				Description: "Define a time period to control the range of usage data you want the data export to contain when it runs:\n\n* **TODAY**. Data collected for the current day up until the time the export runs.\n* **YESTERDAY**. Data collected for the day before the export runs - that is, the 24 hour period from midnight to midnight of the day before.\n* **WEEK_TO_DATE**. Data collected for the period covering the current week to the date and time the export runs, and weeks run Monday to Monday.\n* **CURRENT_MONTH**. Data collected for the current month in which the export is ran up to and including the date and time the export runs.\n* **LAST_30_DAYS**. Data collected for the 30 days prior to the date the export is ran.\n* **LAST_35_DAYS**. Data collected for the 35 days prior to the date the export is ran.\n* **PREVIOUS_WEEK**. Data collected for the previous full week period, and weeks run Monday to Monday.\n* **PREVIOUS_MONTH**. Data collected for the previous full month period.\n\nFor more details and examples, see the [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period) section in our main User Documentation.\nAvailable values: \"LAST_12_HOURS\", \"TODAY\", \"YESTERDAY\", \"WEEK_TO_DATE\", \"CURRENT_MONTH\", \"LAST_30_DAYS\", \"LAST_35_DAYS\", \"PREVIOUS_WEEK\", \"PREVIOUS_MONTH\".",
				Optional:    true,
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
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"account_ids": schema.ListAttribute{
				Description: "List of account IDs to export",
				Optional:    true,
				ElementType: types.StringType,
			},
			"meter_ids": schema.ListAttribute{
				Description: "List of meter IDs to export",
				Optional:    true,
				ElementType: types.StringType,
			},
			"operational_data_types": schema.ListAttribute{
				Description: "A list of the entities whose operational data is included in the data export.",
				Optional:    true,
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
				ElementType: types.StringType,
			},
			"aggregations": schema.ListNestedAttribute{
				Description: "List of aggregations to apply",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[DataExportScheduleAggregationsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"field_code": schema.StringAttribute{
							Description: "Field code",
							Required:    true,
						},
						"field_type": schema.StringAttribute{
							Description: "Type of field\nAvailable values: \"DIMENSION\", \"MEASURE\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("DIMENSION", "MEASURE"),
							},
						},
						"function": schema.StringAttribute{
							Description: "Aggregation function\nAvailable values: \"SUM\", \"MIN\", \"MAX\", \"COUNT\", \"LATEST\", \"MEAN\", \"UNIQUE\".",
							Required:    true,
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
							Required:    true,
						},
					},
				},
			},
			"dimension_filters": schema.ListNestedAttribute{
				Description: "List of dimension filters to apply",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[DataExportScheduleDimensionFiltersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"field_code": schema.StringAttribute{
							Description: "Field code",
							Required:    true,
						},
						"meter_id": schema.StringAttribute{
							Description: "Meter ID",
							Required:    true,
						},
						"values": schema.ListAttribute{
							Description: "Values to filter by",
							Required:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"groups": schema.ListNestedAttribute{
				Description: "List of groups to apply",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[DataExportScheduleGroupsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_type": schema.StringAttribute{
							Description: `Available values: "ACCOUNT", "DIMENSION", "TIME".`,
							Optional:    true,
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
							Optional:    true,
						},
						"meter_id": schema.StringAttribute{
							Description: "Meter ID to group by",
							Optional:    true,
						},
						"frequency": schema.StringAttribute{
							Description: "Frequency of usage data\nAvailable values: \"DAY\", \"HOUR\", \"WEEK\", \"MONTH\", \"QUARTER\".",
							Optional:    true,
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

func (r *DataExportScheduleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *DataExportScheduleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
