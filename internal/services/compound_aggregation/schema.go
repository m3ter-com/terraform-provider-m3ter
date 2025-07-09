// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

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
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*CompoundAggregationResource)(nil)

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
			"calculation": schema.StringAttribute{
				Description: "String that represents the formula for the calculation. This formula determines how the CompoundAggregation value is calculated. The calculation can reference  simple Aggregations or Custom Fields. This field is required when creating or updating a CompoundAggregation.\n\n**NOTE:** If a simple Aggregation referenced by a Compound Aggregation has a **Quantity per unit** defined or a **Rounding** defined, these will not be factored into the value used by the calculation. For example, if the simple Aggregation referenced has a base value of 100 and has **Quantity per unit** set at 10, the Compound Aggregation calculation *will use the base value of 100 not 10*.",
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
			"evaluate_null_aggregations": schema.BoolAttribute{
				Description: "Boolean True / False flag:\n- **TRUE** - set to TRUE if you want to allow null values from the simple Aggregations referenced in the Compound Aggregation to be passed in. Simple Aggregations based on Meter Target Fields where no usage data is available will have null values.\n- **FALSE** Default.\n\n**Note:** If any of the simple Aggregations you reference in a Compound Aggregation calculation might have null values, you must set their Default Value to 0. This ensures that any null values passed into the Compound Aggregation are passed in correctly with value = 0.",
				Optional:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "Unique identifier (UUID) of the Product the CompoundAggregation belongs to.\n\n**Note:** Omit this parameter if you want to create a *Global* CompoundAggregation.",
				Optional:    true,
			},
			"custom_fields": schema.DynamicAttribute{
				Optional: true,
			},
			"aggregation": schema.StringAttribute{
				Description: "Specifies the computation method applied to usage data collected in `targetField`. Aggregation unit value depends on the **Category** configured for the selected targetField.\n\nEnum: \n\n* **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **COUNT**. Counts the number of values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **LATEST**. Uses the most recent value. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`) value of usage data measurement submissions. If using this method, please ensure *distinct* `ts` values are used for usage data measurment submissions.\n\n* **MEAN**. Uses the arithmetic mean of the values. Can be applied to a **Measure**, **Income**, or **Cost** `targetField`.\n\n* **UNIQUE**. Uses unique values and returns a count of the number of unique values. Can be applied to a **Metadata** `targetField`.\nAvailable values: \"SUM\", \"MIN\", \"MAX\", \"COUNT\", \"LATEST\", \"MEAN\", \"UNIQUE\", \"CUSTOM_SQL\".",
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
						"CUSTOM_SQL",
					),
				},
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this CompoundAggregation.",
				Computed:    true,
			},
			"custom_sql": schema.StringAttribute{
				Computed: true,
			},
			"default_value": schema.Float64Attribute{
				Description: "Aggregation value used when no usage data is available to be aggregated. *(Optional)*.\n\n**Note:** Set to 0, if you expect to reference the Aggregation in a Compound Aggregation. This ensures that any null values are passed in correctly to the Compound Aggregation calculation with a value = 0.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the CompoundAggregation was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the CompoundAggregation was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this CompoundAggregation.",
				Computed:    true,
			},
			"meter_id": schema.StringAttribute{
				Description: "The UUID of the Meter used as the source of usage data for the Aggregation.\n\nEach Aggregation is a child of a Meter, so the Meter must be selected.",
				Computed:    true,
			},
			"target_field": schema.StringAttribute{
				Description: "`Code` of the target `dataField` or `derivedField` on the Meter used as the basis for the Aggregation.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"segmented_fields": schema.ListAttribute{
				Description: "*(Optional)*. Used when creating a segmented Aggregation, which segments the usage data collected by a single Meter. Works together with `segments`.\n\nThe `Codes` of the fields in the target Meter to use for segmentation purposes.\n\nString `dataFields` on the target Meter can be segmented. Any string `derivedFields` on the target Meter, such as one that concatenates two string `dataFields`, can also be segmented.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"segments": schema.ListAttribute{
				Description: "*(Optional)*. Used when creating a segmented Aggregation, which segments the usage data collected by a single Meter. Works together with `segmentedFields`.\n\nContains the values that are to be used as the segments, read from the fields in the meter pointed at by `segmentedFields`.",
				Computed:    true,
				CustomType:  customfield.NewListType[customfield.Map[types.String]](ctx),
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
			},
		},
	}
}

func (r *CompoundAggregationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CompoundAggregationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
