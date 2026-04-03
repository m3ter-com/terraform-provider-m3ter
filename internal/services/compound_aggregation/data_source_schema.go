// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CompoundAggregationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for Compound Aggregation related operations such as creation, update, list and delete.\n\nUse Compound Aggregations to create numerical measures from usage data by applying a calculation to one or more simple Aggregations or Custom Fields. These numerical measures can then be used as pricing metrics to price your Product Plans, enabling you to implement a wide range of usage-based pricing use cases.\n\nYou can create two types of Compound Aggregation:\n\n**Global**\n- Pricing: Not tied to any specific product and can be used to price Plans belonging to any Product.\n- Calculation: can reference all simple Aggregations - both Global simple Aggregations and any product-specific simple Aggregations.\n\n**Product-specific**\n- Pricing: belong to a specific Product and can only be used to price Plans belonging to the same Product.\n- Calculation: can reference any simple Aggregations belonging to the same Product and any Global simple Aggregations.\n\n**IMPORTANT!** If a simple Aggregation referenced by a Compound Aggregation has a **Quantity per unit** defined or a **Rounding** defined, these will not be factored into the value used by the calculation. For example, if the simple Aggregation referenced has a base value of 100 and has **Quantity per unit** set at 10, the Compound Aggregation calculation *will use the base value of 100 not 10*. \n\nTo better understand and use Compound Aggregations, refer to the example [Compound Aggregation Use Case](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/compound-aggregations#example-use-case) in the m3ter documentation.\n\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"accounting_product_id": schema.StringAttribute{
				Description: "Optional Product ID this Aggregation should be attributed to for accounting purposes.",
				Computed:    true,
			},
			"calculation": schema.StringAttribute{
				Description: "This field is a string that represents the formula for the calculation. This formula determines how the CompoundAggregation is calculated from the underlying usage data.",
				Computed:    true,
			},
			"code": schema.StringAttribute{
				Description: "Code of the Aggregation. A unique short code to identify the Aggregation.",
				Computed:    true,
			},
			"evaluate_null_aggregations": schema.BoolAttribute{
				Description: "This is a boolean True / False flag. \n\nIf set to TRUE, the calculation will be evaluated even if the referenced aggregation has no usage data.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Descriptive name for the Aggregation.",
				Computed:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "This field represents the unique identifier (UUID) of the Product that is associated with the CompoundAggregation.",
				Computed:    true,
			},
			"quantity_per_unit": schema.Float64Attribute{
				Description: "Defines how much of a quantity equates to 1 unit. Used when setting the price per unit for billing purposes - if charging for kilobytes per second (KiBy/s) at rate of $0.25 per 500 KiBy/s, then set quantityPerUnit to 500 and price Plan at $0.25 per unit.\n\nIf `quantityPerUnit` is set to a value other than one, rounding is typically set to UP.",
				Computed:    true,
			},
			"rounding": schema.StringAttribute{
				Description: "Specifies how you want to deal with non-integer, fractional number Aggregation values.\n\n**NOTES:**\n* **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is rounded to 4.\n* Also used in combination with `quantityPerUnit`. Rounds the number of units after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value other than one, you would typically set Rounding to **UP**. For example, suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` = 500, and set charge rate at $0.25 per unit used. If your customer used 48,900 KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up to 98 * 0.25 = $2.45.\n\nEnum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???\nAvailable values: \"UP\", \"DOWN\", \"NEAREST\", \"NONE\".",
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
			"unit": schema.StringAttribute{
				Description: "User defined or following the *Unified Code for Units of Measure* (UCUM). \n\nUsed as the label for billing, indicating to your customers what they are being charged for.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"segments": schema.ListAttribute{
				Description: "*(Optional)*. Used when creating a segmented Aggregation, which segments the usage data collected by a single Meter. Works together with `segmentedFields`.\n\nContains the values that are to be used as the segments, read from the fields in the meter pointed at by `segmentedFields`.",
				Computed:    true,
				CustomType:  customfield.NewListType[customfield.Map[types.String]](ctx),
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
			},
			"custom_fields": schema.DynamicAttribute{
				Computed:   true,
				CustomType: customfield.NormalizedDynamicType{},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"codes": schema.ListAttribute{
						Description: "An optional parameter to retrieve specific CompoundAggregations based on their short codes.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"ids": schema.ListAttribute{
						Description: "An optional parameter to retrieve specific CompoundAggregations based on their unique identifiers (UUIDs).",
						Optional:    true,
						ElementType: types.StringType,
					},
					"product_id": schema.ListAttribute{
						Description: "An optional parameter to filter the CompoundAggregations based on specific Product unique identifiers (UUIDs).",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *CompoundAggregationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CompoundAggregationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
