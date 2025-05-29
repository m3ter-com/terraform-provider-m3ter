// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*StatementStatementDefinitionResource)(nil)

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
			"aggregation_frequency": schema.StringAttribute{
				Description: "This specifies how often the Statement should aggregate data.\nAvailable values: \"DAY\", \"WEEK\", \"MONTH\", \"QUARTER\", \"YEAR\", \"WHOLE_PERIOD\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"DAY",
						"WEEK",
						"MONTH",
						"QUARTER",
						"YEAR",
						"WHOLE_PERIOD",
					),
				},
			},
			"include_price_per_unit": schema.BoolAttribute{
				Description: "A Boolean indicating whether to include the price per unit in the Statement.\n\n* TRUE - includes the price per unit.\n* FALSE - excludes the price per unit.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Descriptive name for the StatementDefinition providing context and information.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"dimensions": schema.ListNestedAttribute{
				Description: "An array of objects, each representing a Dimension data field from a Meter *(for Meters that have Dimensions setup)*.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"filter": schema.ListAttribute{
							Description: `The value of a Dimension to use as a filter. Use "*" as a wildcard to filter on all Dimension values.`,
							Required:    true,
							ElementType: types.StringType,
						},
						"name": schema.StringAttribute{
							Description: "The name of the Dimension to target in the Meter.",
							Required:    true,
						},
						"attributes": schema.ListAttribute{
							Description: "The Dimension attribute to target.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"meter_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the Meter containing this Dimension.",
							Optional:    true,
						},
					},
				},
			},
			"measures": schema.ListNestedAttribute{
				Description: "An array of objects, each representing a Measure data field from a Meter.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"aggregations": schema.ListAttribute{
							Description: "A list of Aggregations to apply to the Measure.",
							Optional:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
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
								),
							},
							ElementType: types.StringType,
						},
						"meter_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the Meter containing this Measure.",
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of a Measure data field *(or blank to indicate a wildcard, i.e. all fields)*. Default value is blank.",
							Optional:    true,
						},
					},
				},
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this StatementDefinition.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the StatementDefinition was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the StatementDefinition was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this StatementDefinition.",
				Computed:    true,
			},
		},
	}
}

func (r *StatementStatementDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *StatementStatementDefinitionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
