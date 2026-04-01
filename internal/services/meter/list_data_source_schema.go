// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*MetersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Endpoints for listing, creating, updating, retrieving, or deleting Meters.\n\nUse Meters to submit usage data for the consumption of your products and services by end customers. This usage data then becomes the basis for setting up usage-based pricing for your products and services.\n\nExamples of usage data collected in Meters:\n* Number of logins.\n* Duration of session.\n* Amount of data downloaded.\n\nTo collect usage data and ingest it into the platform, you can define two types of fields for Meters:\n- `dataFields` Used to collect raw usage data measures - numeric quantitative data values or non-numeric point data values.\n- `derivedFields` Used to derive usage data measures that are the result of applying a calculation to `dataFields`, `customFields`, or system `Timestamp` fields.\n\nYou can also:\n- Create `customFields` for a Meter, which allows you to attach custom data to the Meter as name/value pairs.\n- Create Global Meters, which are not tied to a specific Product and allow you to collect usage data that will form the basis of usage-based pricing across multiple Products.\n\n**IMPORTANT! - use of PII:** The use of any of your end-customers' Personally Identifiable Information (PII) in m3ter is restricted to a few fields on the **Account** entity. Please ensure that any fields you configure for Meters, such as Data Fields or Derived Fields, do not contain any end-customer PII data. See the [Introduction section](https://www.m3ter.com/docs/api#section/Introduction) above for more details.\n\nSee also:\n- [Reviewing Meter Options](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/reviewing-meter-options).",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"codes": schema.ListAttribute{
				Description: "List of Meter codes to retrieve. These are the unique short codes that identify each Meter.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "List of Meter IDs to retrieve.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"product_id": schema.ListAttribute{
				Description: "The UUIDs of the Products to retrieve Meters for.",
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
				CustomType:  customfield.NewNestedObjectListType[MetersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"code": schema.StringAttribute{
							Description: "Code of the Meter - unique short code used to identify the Meter.",
							Computed:    true,
						},
						"custom_fields": schema.DynamicAttribute{
							Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
							Computed:    true,
							CustomType:  customfield.NormalizedDynamicType{},
						},
						"data_fields": schema.ListNestedAttribute{
							Description: "Used to submit categorized raw usage data values for ingest into the platform - either numeric quantitative values or non-numeric data values. At least one required per Meter; maximum 15 per Meter.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[MetersDataFieldsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"category": schema.StringAttribute{
										Description: "The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).\nAvailable values: \"WHO\", \"WHERE\", \"WHAT\", \"OTHER\", \"METADATA\", \"MEASURE\", \"INCOME\", \"COST\".",
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
							CustomType:  customfield.NewNestedObjectListType[MetersDerivedFieldsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"category": schema.StringAttribute{
										Description: "The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).\nAvailable values: \"WHO\", \"WHERE\", \"WHAT\", \"OTHER\", \"METADATA\", \"MEASURE\", \"INCOME\", \"COST\".",
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
									"calculation": schema.StringAttribute{
										Description: "The calculation used to transform the value of submitted `dataFields` in usage data. Calculation can reference `dataFields`, `customFields`, or system `Timestamp` fields. \n*(Example: datafieldms  datafieldgb)*",
										Computed:    true,
									},
								},
							},
						},
						"group_id": schema.StringAttribute{
							Description: "UUID of the MeterGroup the Meter belongs to. *(Optional)*.",
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
					},
				},
			},
		},
	}
}

func (d *MetersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MetersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
