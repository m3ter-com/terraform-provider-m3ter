// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*MeterDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"code": schema.StringAttribute{
				Description: "Code of the Meter - unique short code used to identify the Meter.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this meter.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when the meter was created *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when the meter was last modified *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"group_id": schema.StringAttribute{
				Description: "UUID of the MeterGroup the Meter belongs to. *(Optional)*.",
				Computed:    true,
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified this meter.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Descriptive name for the Meter.",
				Computed:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "UUID of the Product the Meter belongs to. *(Optional)* - if blank, the Meter is global.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"custom_fields": schema.MapAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"data_fields": schema.ListNestedAttribute{
				Description: "Used to submit categorized raw usage data values for ingest into the platform - either numeric quantitative values or non-numeric data values. At least one required per Meter; maximum 15 per Meter.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MeterDataFieldsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"category": schema.StringAttribute{
							Description: "The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"WHO",
									"WHERE",
									"WHAT",
									"OTHER",
									"METADATA",
									"MEASURE",
									"INCOME",
									"COST",
								),
							},
						},
						"code": schema.StringAttribute{
							Description: "Short code to identify the field\n\n**NOTE:** Code has a maximum length of 80 characters and can only contain letters, numbers, underscore, and the dollar character, and must not start with a number.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Descriptive name of the field.",
							Computed:    true,
						},
						"unit": schema.StringAttribute{
							Description: "The units to measure the data with. Should conform to *Unified Code for Units of Measure* (UCUM). Required only for numeric field categories.",
							Computed:    true,
						},
					},
				},
			},
			"derived_fields": schema.ListNestedAttribute{
				Description: "Used to submit usage data values for ingest into the platform that are the result of a calculation performed on `dataFields`, `customFields`, or system `Timestamp` fields. Raw usage data is not submitted using `derivedFields`. Maximum 15 per Meter. *(Optional)*.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MeterDerivedFieldsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"category": schema.StringAttribute{
							Description: "The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"WHO",
									"WHERE",
									"WHAT",
									"OTHER",
									"METADATA",
									"MEASURE",
									"INCOME",
									"COST",
								),
							},
						},
						"code": schema.StringAttribute{
							Description: "Short code to identify the field\n\n**NOTE:** Code has a maximum length of 80 characters and can only contain letters, numbers, underscore, and the dollar character, and must not start with a number.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Descriptive name of the field.",
							Computed:    true,
						},
						"unit": schema.StringAttribute{
							Description: "The units to measure the data with. Should conform to *Unified Code for Units of Measure* (UCUM). Required only for numeric field categories.",
							Computed:    true,
						},
					},
				},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"codes": schema.ListAttribute{
						Description: "List of Meter codes to retrieve. These are the unique short codes that identify each Meter.   ",
						Optional:    true,
						ElementType: types.StringType,
					},
					"ids": schema.ListAttribute{
						Description: "List of Meter IDs to retrieve. ",
						Optional:    true,
						ElementType: types.StringType,
					},
					"product_id": schema.ListAttribute{
						Description: "The UUIDs of the Products to retrieve Meters for.",
						Optional:    true,
						ElementType: jsontypes.NormalizedType{},
					},
				},
			},
		},
	}
}

func (d *MeterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MeterDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
