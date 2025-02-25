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
			"aggregation": schema.StringAttribute{
				Description: "Specifies the aggregation method applied to usage data collected in the numeric Data Fields of Meters included for the Data Export Schedule - that is, Data Fields of type **MEASURE**, **INCOME**, or **COST**:\n\n* **SUM**. Adds the values.\n* **MIN**. Uses the minimum value.\n* **MAX**. Uses the maximum value.\n* **COUNT**. Counts the number of values.\n* **LATEST**. Uses the most recent value. Note: Based on the timestamp `ts` value of usage data measurement submissions. If using this method, please ensure *distinct* `ts` values are used for usage data measurement submissions.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"SUM",
						"MIN",
						"MAX",
						"COUNT",
						"LATEST",
						"MEAN",
					),
				},
			},
			"aggregation_frequency": schema.StringAttribute{
				Description: "Specifies the time period for the aggregation of usage data included each time the Data Export Schedule runs:\n* **ORIGINAL**. Usage data is *not aggregated*. If you select to not aggregate, then raw usage data measurements collected by all Data Field types and any Derived Fields on the selected Meters are included in the export. This is the *Default*.\n\nIf you want to aggregate usage data for the Export Schedule you must define an `aggregationFrequency`:\n\n* **HOUR**. Aggregated hourly.\n* **DAY**. Aggregated daily.\n* **WEEK**. Aggregated weekly.\n* **MONTH**. Aggregated monthly.\n\n* If you select to aggregate usage data for a Export Schedule, then only the aggregated usage data collected by numeric Data Fields of type **MEASURE**, **INCOME**, or **COST** on selected Meters are included in the export.\n\n**NOTE**: If you define an `aggregationFrequency` other than **ORIGINAL** and do not define an `aggregation` method, then you'll receive and error. ",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ORIGINAL",
						"HOUR",
						"DAY",
						"WEEK",
						"MONTH",
					),
				},
			},
			"time_period": schema.StringAttribute{
				Description: "Define a time period to control the range of usage data you want the data export to contain when it runs:\n\n* **TODAY**. Data collected for the current day up until the time the export runs.\n* **YESTERDAY**. Data collected for the day before the export runs - that is, the 24 hour period from midnight to midnight of the day before.\n* **WEEK_TO_DATE**. Data collected for the period covering the current week to the date and time the export runs, and weeks run Monday to Monday.\n* **CURRENT_MONTH**. Data collected for the current month in which the export is ran up to and including the date and time the export runs.\n* **LAST_30_DAYS**. Data collected for the 30 days prior to the date the export is ran.\n* **LAST_35_DAYS**. Data collected for the 35 days prior to the date the export is ran.\n* **PREVIOUS_WEEK**. Data collected for the previous full week period, and weeks run Monday to Monday.\n* **PREVIOUS_MONTH**. Data collected for the previous full month period.\n\nFor more details and examples, see the [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period) section in our main User Documentation.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
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
		},
	}
}

func (d *DataExportScheduleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DataExportScheduleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
