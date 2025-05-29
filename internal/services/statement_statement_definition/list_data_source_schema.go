// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package statement_statement_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*StatementStatementDefinitionsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
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
				CustomType:  customfield.NewNestedObjectListType[StatementStatementDefinitionsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
						"aggregation_frequency": schema.StringAttribute{
							Description: "This specifies how often the Statement should aggregate data.\nAvailable values: \"DAY\", \"WEEK\", \"MONTH\", \"QUARTER\", \"YEAR\", \"WHOLE_PERIOD\".",
							Computed:    true,
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
						"created_by": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the user who created this StatementDefinition.",
							Computed:    true,
						},
						"dimensions": schema.ListNestedAttribute{
							Description: "An array of objects, each representing a Dimension data field from a Meter *(for Meters that have Dimensions setup)*.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[StatementStatementDefinitionsDimensionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"filter": schema.ListAttribute{
										Description: `The value of a Dimension to use as a filter. Use "*" as a wildcard to filter on all Dimension values.`,
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"name": schema.StringAttribute{
										Description: "The name of the Dimension to target in the Meter.",
										Computed:    true,
									},
									"attributes": schema.ListAttribute{
										Description: "The Dimension attribute to target.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"meter_id": schema.StringAttribute{
										Description: "The unique identifier (UUID) of the Meter containing this Dimension.",
										Computed:    true,
									},
								},
							},
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
						"include_price_per_unit": schema.BoolAttribute{
							Description: "A Boolean indicating whether to include the price per unit in the Statement.\n\n* TRUE - includes the price per unit.\n* FALSE - excludes the price per unit.",
							Computed:    true,
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the user who last modified this StatementDefinition.",
							Computed:    true,
						},
						"measures": schema.ListNestedAttribute{
							Description: "An array of objects, each representing a Measure data field from a Meter.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[StatementStatementDefinitionsMeasuresDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"aggregations": schema.ListAttribute{
										Description: "A list of Aggregations to apply to the Measure.",
										Computed:    true,
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
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"meter_id": schema.StringAttribute{
										Description: "The unique identifier (UUID) of the Meter containing this Measure.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "The name of a Measure data field *(or blank to indicate a wildcard, i.e. all fields)*. Default value is blank.",
										Computed:    true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "Descriptive name for the StatementDefinition providing context and information.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *StatementStatementDefinitionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *StatementStatementDefinitionsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
