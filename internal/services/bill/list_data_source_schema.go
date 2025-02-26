// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BillsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Description: "Optional filter. An Account ID - returns the Bills for the single specified Account.",
				Optional:    true,
			},
			"bill_date": schema.StringAttribute{
				Description: "The specific date in ISO 8601 format for which you want to retrieve Bills.",
				Optional:    true,
			},
			"bill_date_end": schema.StringAttribute{
				Description: "Only include Bills with bill dates earlier than this date.",
				Optional:    true,
			},
			"bill_date_start": schema.StringAttribute{
				Description: "Only include Bills with bill dates equal to or later than this date.",
				Optional:    true,
			},
			"billing_frequency": schema.StringAttribute{
				Optional: true,
			},
			"exclude_line_items": schema.BoolAttribute{
				Description: "Exclude Line Items",
				Optional:    true,
			},
			"external_invoice_date_end": schema.StringAttribute{
				Description: "Only include Bills with external invoice dates earlier than this date.",
				Optional:    true,
			},
			"external_invoice_date_start": schema.StringAttribute{
				Description: "Only include Bills with external invoice dates equal to or later than this date.",
				Optional:    true,
			},
			"include_bill_total": schema.BoolAttribute{
				Description: "Include Bill Total",
				Optional:    true,
			},
			"locked": schema.BoolAttribute{
				Description: "Boolean flag specifying whether to include Bills with \"locked\" status.\n\n* **TRUE** - the list inlcudes \"locked\" Bills.\n* **FALSE** - excludes \"locked\" Bills from the list.",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Only include Bills having the given status\nAvailable values: \"PENDING\", \"APPROVED\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("PENDING", "APPROVED"),
				},
			},
			"ids": schema.ListAttribute{
				Description: "Optional filter. The list of Bill IDs to retrieve.",
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
				CustomType:  customfield.NewNestedObjectListType[BillsItemsDataSourceModel](ctx),
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
						"account_code": schema.StringAttribute{
							Computed: true,
						},
						"account_id": schema.StringAttribute{
							Computed: true,
						},
						"bill_date": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"bill_frequency_interval": schema.Int64Attribute{
							Computed: true,
						},
						"billing_frequency": schema.StringAttribute{
							Description: "Available values: \"DAILY\", \"WEEKLY\", \"MONTHLY\", \"ANNUALLY\", \"AD_HOC\", \"MIXED\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"DAILY",
									"WEEKLY",
									"MONTHLY",
									"ANNUALLY",
									"AD_HOC",
									"MIXED",
								),
							},
						},
						"bill_job_id": schema.StringAttribute{
							Computed: true,
						},
						"bill_total": schema.Float64Attribute{
							Description: "The sum total for the Bill.",
							Computed:    true,
						},
						"created_by": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the user who created the Bill.",
							Computed:    true,
						},
						"created_date": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"csv_statement_generated": schema.BoolAttribute{
							Description: "Flag to indicate that the statement in CSV format has been generated for the Bill.\n\n* **TRUE** - CSV statement has been generated.\n* **FALSE** - no CSV statement generated.",
							Computed:    true,
						},
						"currency": schema.StringAttribute{
							Computed: true,
						},
						"currency_conversions": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[BillsCurrencyConversionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"from": schema.StringAttribute{
										Description: "Currency to convert from. For example: GBP.",
										Computed:    true,
									},
									"to": schema.StringAttribute{
										Description: "Currency to convert to. For example: USD.",
										Computed:    true,
									},
									"multiplier": schema.Float64Attribute{
										Description: "Conversion rate between currencies.",
										Computed:    true,
										Validators: []validator.Float64{
											float64validator.AtLeast(0),
										},
									},
								},
							},
						},
						"dt_created": schema.StringAttribute{
							Description: "The date and time *(in ISO 8601 format)* when the Bill was first created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_last_modified": schema.StringAttribute{
							Description: "The date and time *(in ISO 8601 format)* when the Bill was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"due_date": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"end_date": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"end_date_time_utc": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"external_invoice_date": schema.StringAttribute{
							Description: "For accounting purposes, the date set at Organization level to use for external invoicing with respect to billing periods - two options:\n* `FIRST_DAY_OF_NEXT_PERIOD` *(Default)*.\n* `LAST_DAY_OF_ARREARS`\n\nFor example, if the retrieved Bill was on a monthly billing frequency and the billing period for the Bill is September 2023 and the *External invoice date* is set at `FIRST_DAY_OF_NEXT_PERIOD`, then the `externalInvoiceDate` will be `\"2023-10-01\"`.\n\n**NOTE:** To change the `externalInvoiceDate` setting for your Organization, you can use the [Update OrganizationConfig](https://www.m3ter.com/docs/api#tag/OrganizationConfig/operation/GetOrganizationConfig) call.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"external_invoice_reference": schema.StringAttribute{
							Description: "The reference ID to use for external invoice.",
							Computed:    true,
						},
						"json_statement_generated": schema.BoolAttribute{
							Description: "Flag to indicate that the statement in JSON format has been generated for the Bill.\n\n* **TRUE** - JSON statement has been generated.\n* **FALSE** - no JSON statement generated.",
							Computed:    true,
						},
						"last_calculated_date": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the user who last modified this Bill.",
							Computed:    true,
						},
						"line_items": schema.ListNestedAttribute{
							Description: "An array of the Bill line items.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[BillsLineItemsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"average_unit_price": schema.Float64Attribute{
										Description: "The average unit price across all tiers / pricing bands.",
										Computed:    true,
									},
									"conversion_rate": schema.Float64Attribute{
										Description: "The currency conversion rate if currency conversion is required for the line item.",
										Computed:    true,
									},
									"converted_subtotal": schema.Float64Attribute{
										Description: "The converted subtotal amount if currency conversions have been used.",
										Computed:    true,
									},
									"currency": schema.StringAttribute{
										Description: "The currency code for the currency used in the line item. For example: USD, GBP, or EUR.",
										Computed:    true,
									},
									"description": schema.StringAttribute{
										Description: "Line item description.",
										Computed:    true,
									},
									"line_item_type": schema.StringAttribute{
										Description: "Available values: \"STANDING_CHARGE\", \"USAGE\", \"COUNTER_RUNNING_TOTAL_CHARGE\", \"COUNTER_ADJUSTMENT_DEBIT\", \"COUNTER_ADJUSTMENT_CREDIT\", \"USAGE_CREDIT\", \"MINIMUM_SPEND\", \"MINIMUM_SPEND_REFUND\", \"CREDIT_DEDUCTION\", \"MANUAL_ADJUSTMENT\", \"CREDIT_MEMO\", \"DEBIT_MEMO\", \"COMMITMENT_CONSUMED\", \"COMMITMENT_FEE\", \"OVERAGE_SURCHARGE\", \"OVERAGE_USAGE\", \"BALANCE_CONSUMED\", \"BALANCE_FEE\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"STANDING_CHARGE",
												"USAGE",
												"COUNTER_RUNNING_TOTAL_CHARGE",
												"COUNTER_ADJUSTMENT_DEBIT",
												"COUNTER_ADJUSTMENT_CREDIT",
												"USAGE_CREDIT",
												"MINIMUM_SPEND",
												"MINIMUM_SPEND_REFUND",
												"CREDIT_DEDUCTION",
												"MANUAL_ADJUSTMENT",
												"CREDIT_MEMO",
												"DEBIT_MEMO",
												"COMMITMENT_CONSUMED",
												"COMMITMENT_FEE",
												"OVERAGE_SURCHARGE",
												"OVERAGE_USAGE",
												"BALANCE_CONSUMED",
												"BALANCE_FEE",
											),
										},
									},
									"quantity": schema.Float64Attribute{
										Description: "The amount of usage for the line item.",
										Computed:    true,
									},
									"subtotal": schema.Float64Attribute{
										Description: "The subtotal amount for the line item, before any currency conversions.",
										Computed:    true,
									},
									"unit": schema.StringAttribute{
										Description: "The unit for the usage data in thie line item. For example: **GB** of disk storage space.",
										Computed:    true,
									},
									"units": schema.Float64Attribute{
										Description: "The number of units used for the line item.",
										Computed:    true,
									},
									"id": schema.StringAttribute{
										Description: "The UUID for the line item.",
										Computed:    true,
									},
									"aggregation_id": schema.StringAttribute{
										Description: "The Aggregation ID used for the line item.",
										Computed:    true,
									},
									"balance_id": schema.StringAttribute{
										Computed: true,
									},
									"child_account_code": schema.StringAttribute{
										Description: "If part of a Parent/Child account billing hierarchy, this is the code for the child Account.",
										Computed:    true,
									},
									"child_account_id": schema.StringAttribute{
										Description: "If part of a Parent/Child account billing hierarchy, this is the child Account UUID.",
										Computed:    true,
									},
									"commitment_id": schema.StringAttribute{
										Description: "If Commitments *(prepayments)* are used in the line item, this shows the Commitment UUID.",
										Computed:    true,
									},
									"compound_aggregation_id": schema.StringAttribute{
										Description: "The Compound Aggregation ID for the line item if a Compound Aggregation has been used.",
										Computed:    true,
									},
									"contract_id": schema.StringAttribute{
										Description: "The UUID for the Contract used in the line item.",
										Computed:    true,
									},
									"counter_id": schema.StringAttribute{
										Computed: true,
									},
									"credit_type_id": schema.StringAttribute{
										Computed: true,
									},
									"group": schema.MapAttribute{
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"meter_id": schema.StringAttribute{
										Description: "The UUID of the Meter used in the line item.",
										Computed:    true,
									},
									"plan_group_id": schema.StringAttribute{
										Description: "The UUID of the PlanGroup, provided the line item used a PlanGroup.",
										Computed:    true,
									},
									"plan_id": schema.StringAttribute{
										Description: "The ID of the Plan used for the line item.",
										Computed:    true,
									},
									"pricing_id": schema.StringAttribute{
										Description: "The UUID of the Pricing used on the line item.",
										Computed:    true,
									},
									"product_code": schema.StringAttribute{
										Computed: true,
									},
									"product_id": schema.StringAttribute{
										Description: "The UUID of the Product for the line item.",
										Computed:    true,
									},
									"product_name": schema.StringAttribute{
										Description: "The name of the Product for the line item.",
										Computed:    true,
									},
									"reason_id": schema.StringAttribute{
										Computed: true,
									},
									"referenced_bill_id": schema.StringAttribute{
										Computed: true,
									},
									"referenced_line_item_id": schema.StringAttribute{
										Computed: true,
									},
									"segment": schema.MapAttribute{
										Description: "Applies only when segmented Aggregations have been used. The Segment to which the usage data in this line item belongs.",
										Computed:    true,
										CustomType:  customfield.NewMapType[types.String](ctx),
										ElementType: types.StringType,
									},
									"sequence_number": schema.Int64Attribute{
										Description: "The number used for sequential invoices.",
										Computed:    true,
									},
									"service_period_end_date": schema.StringAttribute{
										Description: "The ending date *(exclusive)* for the service period *(in ISO 8601 format)*.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"service_period_start_date": schema.StringAttribute{
										Description: "The starting date *(inclusive)* for the service period *(in ISO 8601 format)*.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"usage_per_pricing_band": schema.ListNestedAttribute{
										Description: "Shows the usage by pricing band for tiered pricing structures.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[BillsLineItemsUsagePerPricingBandDataSourceModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"band_quantity": schema.Float64Attribute{
													Description: "Usage amount within the band.",
													Computed:    true,
												},
												"band_subtotal": schema.Float64Attribute{
													Description: "Subtotal amount for the band.",
													Computed:    true,
												},
												"band_units": schema.Float64Attribute{
													Description: "The number of units used within the band.",
													Computed:    true,
												},
												"credit_type_id": schema.StringAttribute{
													Description: "The UUID of the credit type.",
													Computed:    true,
												},
												"fixed_price": schema.Float64Attribute{
													Description: "Fixed price is a charge entered for certain pricing types such as Stairstep, Custom Tiered, and Custom Volume. It is a set price and not dependent on usage.",
													Computed:    true,
												},
												"lower_limit": schema.Float64Attribute{
													Description: "The lower limit *(start)* of the pricing band.",
													Computed:    true,
												},
												"pricing_band_id": schema.StringAttribute{
													Description: "The UUID for the pricing band.",
													Computed:    true,
												},
												"unit_price": schema.Float64Attribute{
													Description: "The price per unit in the band.",
													Computed:    true,
												},
												"unit_subtotal": schema.Float64Attribute{
													Description: "The subtotal of the unit usage.",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"locked": schema.BoolAttribute{
							Computed: true,
						},
						"purchase_order_number": schema.StringAttribute{
							Description: "Purchase Order number linked to the Account the Bill is for.",
							Computed:    true,
						},
						"sequential_invoice_number": schema.StringAttribute{
							Description: "The sequential invoice number of the Bill.\n\n**NOTE:** If you have not defined a `billPrefix` for your Organization, a `sequentialInvoiceNumber` is not returned in the response. See [Update OrganizationConfig](https://www.m3ter.com/docs/api#tag/OrganizationConfig/operation/UpdateOrganizationConfig)",
							Computed:    true,
						},
						"start_date": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"start_date_time_utc": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"status": schema.StringAttribute{
							Description: "Available values: \"PENDING\", \"APPROVED\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("PENDING", "APPROVED"),
							},
						},
						"timezone": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *BillsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *BillsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
