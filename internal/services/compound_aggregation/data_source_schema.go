// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CompoundAggregationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"accounting_product_id": schema.StringAttribute{
				Computed: true,
			},
			"calculation": schema.StringAttribute{
				Description: "This field is a string that represents the formula for the calculation. This formula determines how the CompoundAggregation is calculated from the underlying usage data.",
				Computed:    true,
			},
			"code": schema.StringAttribute{
				Description: "Code of the Aggregation. A unique short code to identify the Aggregation.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this CompoundAggregation.",
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
			"evaluate_null_aggregations": schema.BoolAttribute{
				Description: "This is a boolean True / False flag. \n\nIf set to TRUE, the calculation will be evaluated even if the referenced aggregation has no usage data.",
				Computed:    true,
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this CompoundAggregation.",
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
			"custom_fields": schema.MapAttribute{
				Computed:    true,
				CustomType:  customfield.NewMapType[types.Dynamic](ctx),
				ElementType: types.DynamicType,
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

func (d *CompoundAggregationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CompoundAggregationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
