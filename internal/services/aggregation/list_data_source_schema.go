// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AggregationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"codes": schema.ListAttribute{
				Description: "List of Aggregation codes to retrieve. These are unique short codes to identify each Aggregation.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "List of Aggregation IDs to retrieve.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"product_id": schema.ListAttribute{
				Description: "The UUIDs of the Products to retrieve Aggregations for.",
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
				CustomType:  customfield.NewNestedObjectListType[AggregationsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity. ",
							Computed:    true,
						},
						"aggregation": schema.StringAttribute{
							Description: "Specifies the computation method applied to usage data collected in `targetField`. Aggregation unit value depends on the **Category** configured for the selected targetField.\n\nEnum: \n\n* **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **COUNT**. Counts the number of values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **LATEST**. Uses the most recent value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`) value of usage data measurement submissions. If using this method, please ensure *distinct* `ts` values are used for usage data measurment submissions.\n\n* **MEAN**. Uses the arithmetic mean of the values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **UNIQUE**. Uses unique values and returns a count of the number of unique values. Can be applied to a **Metadata** `targetField`.",
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
						"code": schema.StringAttribute{
							Description: "Code of the Aggregation. A unique short code to identify the Aggregation.",
							Computed:    true,
						},
						"created_by": schema.StringAttribute{
							Description: "The id of the user who created this aggregation.",
							Computed:    true,
						},
						"custom_fields": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
							ElementType: jsontypes.NormalizedType{},
						},
						"default_value": schema.Float64Attribute{
							Description: "Aggregation value used when no usage data is available to be aggregated. *(Optional)*.\n\n**Note:** Set to 0, if you expect to reference the Aggregation in a Compound Aggregation. This ensures that any null values are passed in correctly to the Compound Aggregation calculation with a value = 0.",
							Computed:    true,
						},
						"dt_created": schema.StringAttribute{
							Description: "The DateTime when the aggregation was created *(in ISO 8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_last_modified": schema.StringAttribute{
							Description: "The DateTime when the aggregation was last modified *(in ISO 8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The id of the user who last modified this aggregation.",
							Computed:    true,
						},
						"meter_id": schema.StringAttribute{
							Description: "The UUID of the Meter used as the source of usage data for the Aggregation.\n\nEach Aggregation is a child of a Meter, so the Meter must be selected. ",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Descriptive name for the Aggregation.",
							Computed:    true,
						},
						"quantity_per_unit": schema.Float64Attribute{
							Description: "Defines how much of a quantity equates to 1 unit. Used when setting the price per unit for billing purposes - if charging for kilobytes per second (KiBy/s) at rate of $0.25 per 500 KiBy/s, then set quantityPerUnit to 500 and price Plan at $0.25 per unit.\n\nIf `quantityPerUnit` is set to a value other than one, rounding is typically set to UP.",
							Computed:    true,
						},
						"rounding": schema.StringAttribute{
							Description: "Specifies how you want to deal with non-integer, fractional number Aggregation values.\n\n**NOTES:**\n* **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is rounded to 4.\n* Also used in combination with `quantityPerUnit`. Rounds the number of units after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value other than one, you would typically set Rounding to **UP**. For example, suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` = 500, and set charge rate at $0.25 per unit used. If your customer used 48,900 KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up to 98 * 0.25 = $2.45.\n\nEnum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???\n",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"UP",
									"DOWN",
									"NEAREST",
									"NONE",
								),
							},
						},
						"segmented_fields": schema.ListAttribute{
							Description: "*(Optional)*. Used when creating a segmented Aggregation, which segments the usage data collected by a single Meter. Works together with `segments`.\n\nThe `Codes` of the fields in the target Meter to use for segmentation purposes.\n\nString `dataFields` on the target Meter can be segmented. Any string `derivedFields` on the target Meter, such as one that concatenates two string `dataFields`, can also be segmented.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"segments": schema.ListAttribute{
							Description: "*(Optional)*. Used when creating a segmented Aggregation, which segments the usage data collected by a single Meter. Works together with `segmentedFields`.\n\nContains the values that are to be used as the segments, read from the fields in the meter pointed at by `segmentedFields`.  ",
							Computed:    true,
							CustomType:  customfield.NewListType[customfield.Map[types.String]](ctx),
							ElementType: types.MapType{
								ElemType: types.StringType,
							},
						},
						"target_field": schema.StringAttribute{
							Description: "`Code` of the target `dataField` or `derivedField` on the Meter used as the basis for the Aggregation.",
							Computed:    true,
						},
						"unit": schema.StringAttribute{
							Description: "User defined or following the *Unified Code for Units of Measure* (UCUM). \n\nUsed as the label for billing, indicating to your customers what they are being charged for.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *AggregationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AggregationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
