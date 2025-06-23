// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AggregationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
				PlanModifiers:      []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"aggregation": schema.StringAttribute{
				Description: "Specifies the computation method applied to usage data collected in `targetField`. Aggregation unit value depends on the **Category** configured for the selected targetField.\n\nEnum: \n\n* **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **COUNT**. Counts the number of values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **LATEST**. Uses the most recent value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`) value of usage data measurement submissions. If using this method, please ensure *distinct* `ts` values are used for usage data measurment submissions.\n\n* **MEAN**. Uses the arithmetic mean of the values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **UNIQUE**. Uses unique values and returns a count of the number of unique values. Can be applied to a **Metadata** `targetField`.\nAvailable values: \"SUM\", \"MIN\", \"MAX\", \"COUNT\", \"LATEST\", \"MEAN\", \"UNIQUE\".",
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
				Description: "The UUID of the Meter used as the source of usage data for the Aggregation.\n\nEach Aggregation is a child of a Meter, so the Meter must be selected.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Descriptive name for the Aggregation.",
				Required:    true,
			},
			"quantity_per_unit": schema.Float64Attribute{
				Description: "Defines how much of a quantity equates to 1 unit. Used when setting the price per unit for billing purposes - if charging for kilobytes per second (KiBy/s) at rate of $0.25 per 500 KiBy/s, then set quantityPerUnit to 500 and price Plan at $0.25 per unit.\n\n**Note:** If `quantityPerUnit` is set to a value other than one, `rounding` is typically set to `\"UP\"`.",
				Required:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"rounding": schema.StringAttribute{
				Description: "Specifies how you want to deal with non-integer, fractional number Aggregation values.\n\n**NOTES:**\n* **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is rounded to 4.\n* Also used in combination with `quantityPerUnit`. Rounds the number of units after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value other than one, you would typically set Rounding to **UP**. For example, suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` = 500, and set charge rate at $0.25 per unit used. If your customer used 48,900 KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up to 98 * 0.25 = $2.45.\n\nEnum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???\nAvailable values: \"UP\", \"DOWN\", \"NEAREST\", \"NONE\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"UP",
						"DOWN",
						"NEAREST",
						"NONE",
					),
				},
			},
			"target_field": schema.StringAttribute{
				Description: "`Code` of the target `dataField` or `derivedField` on the Meter used as the basis for the Aggregation.",
				Required:    true,
			},
			"unit": schema.StringAttribute{
				Description: "User defined label for units shown for Bill line items, indicating to your customers what they are being charged for.",
				Required:    true,
			},
			"accounting_product_id": schema.StringAttribute{
				Description: "Optional Product ID this Aggregation should be attributed to for accounting purposes",
				Optional:    true,
			},
			"code": schema.StringAttribute{
				Description: "Code of the new Aggregation. A unique short code to identify the Aggregation.",
				Optional:    true,
			},
			"custom_sql": schema.StringAttribute{
				Description: "**NOTE:** The `customSql` Aggregation type is currently only available in Beta release and on request. If you are interested in using this feature, please get in touch with m3ter Support or your m3ter contact.",
				Optional:    true,
			},
			"default_value": schema.Float64Attribute{
				Description: "Aggregation value used when no usage data is available to be aggregated. *(Optional)*.\n\n**Note:** Set to 0, if you expect to reference the Aggregation in a Compound Aggregation. This ensures that any null values are passed in correctly to the Compound Aggregation calculation with a value = 0.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"segmented_fields": schema.ListAttribute{
				Description: "*(Optional)*. Used when creating a segmented Aggregation, which segments the usage data collected by a single Meter. Works together with `segments`.\n\nEnter the `Codes` of the fields in the target Meter to use for segmentation purposes.\n\nString `dataFields` on the target Meter can be segmented. Any string `derivedFields` on the target Meter, such as one that concatenates two string `dataFields`, can also be segmented.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"segments": schema.ListAttribute{
				Description: "*(Optional)*. Used when creating a segmented Aggregation, which segments the usage data collected by a single Meter. Works together with `segmentedFields`.\n\nEnter the values that are to be used as the segments, read from the fields in the meter pointed at by `segmentedFields`.\n\nNote that you can use *wildcards* or *defaults* when setting up segment values. For more details on how to do this with an example, see [Using Wildcards - API Calls](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/segmented-aggregations#using-wildcards---api-calls) in our main User Docs.",
				Optional:    true,
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
			},
			"custom_fields": schema.DynamicAttribute{
				Optional: true,
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this aggregation.",
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
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *AggregationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AggregationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
