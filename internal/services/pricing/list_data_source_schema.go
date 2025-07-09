// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pricing

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

var _ datasource.DataSourceWithConfigValidators = (*PricingsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"date": schema.StringAttribute{
				Description: "Date on which to retrieve active Pricings.",
				Optional:    true,
			},
			"plan_id": schema.StringAttribute{
				Description: "UUID of the Plan to retrieve Pricings for.",
				Optional:    true,
			},
			"plan_template_id": schema.StringAttribute{
				Description: "UUID of the PlanTemplate to retrieve Pricings for.",
				Optional:    true,
			},
			"ids": schema.ListAttribute{
				Description: "List of Pricing IDs to retrieve.",
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
				CustomType:  customfield.NewNestedObjectListType[PricingsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"accounting_product_id": schema.StringAttribute{
							Computed: true,
						},
						"aggregation_id": schema.StringAttribute{
							Description: "UUID of the Aggregation used to create the Pricing. Use this when creating a Pricing for a segmented aggregation.",
							Computed:    true,
						},
						"aggregation_type": schema.StringAttribute{
							Description: `Available values: "SIMPLE", "COMPOUND".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("SIMPLE", "COMPOUND"),
							},
						},
						"code": schema.StringAttribute{
							Description: "Unique short code for the Pricing.",
							Computed:    true,
						},
						"compound_aggregation_id": schema.StringAttribute{
							Description: "UUID of the Compound Aggregation used to create the Pricing.",
							Computed:    true,
						},
						"created_by": schema.StringAttribute{
							Description: "The ID of the user who created this item.",
							Computed:    true,
						},
						"cumulative": schema.BoolAttribute{
							Description: "Controls whether or not charge rates under a set of pricing bands configured for a Pricing are applied according to each separate band or at the highest band reached.\n\nThe default value is **TRUE**.\n\n* When TRUE, at billing charge rates are applied according to each separate band.\n\n* When FALSE, at billing charge rates are applied according to highest band reached.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Displayed on Bill line items.",
							Computed:    true,
						},
						"dt_created": schema.StringAttribute{
							Description: "The DateTime when this item was created *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_last_modified": schema.StringAttribute{
							Description: "The DateTime when this item was last modified *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"end_date": schema.StringAttribute{
							Description: "The end date *(in ISO-8601 format)* for when the Pricing ceases to be active for the Plan or Plan Template.\n\nIf not specified, the Pricing remains active indefinitely.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The ID of the user who last modified this item.",
							Computed:    true,
						},
						"minimum_spend": schema.Float64Attribute{
							Description: "The minimum spend amount per billing cycle for end customer Accounts on a Plan to which the Pricing is applied.",
							Computed:    true,
						},
						"minimum_spend_bill_in_advance": schema.BoolAttribute{
							Description: "The default value is **FALSE**.\n\n* When TRUE, minimum spend is billed at the start of each billing period.\n\n* When FALSE, minimum spend is billed at the end of each billing period.\n\n*(Optional)*. Overrides the setting at Organization level for minimum spend billing in arrears/in advance.",
							Computed:    true,
						},
						"minimum_spend_description": schema.StringAttribute{
							Description: "Minimum spend description *(displayed on the bill line item)*.",
							Computed:    true,
						},
						"overage_pricing_bands": schema.ListNestedAttribute{
							Description: "The Prepayment/Balance overage pricing in pricing bands for the case of a **Tiered** pricing structure.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[PricingsOveragePricingBandsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"fixed_price": schema.Float64Attribute{
										Description: "Fixed price charged for the Pricing band.",
										Computed:    true,
									},
									"lower_limit": schema.Float64Attribute{
										Description: "Lower limit for the Pricing band.",
										Computed:    true,
										Validators: []validator.Float64{
											float64validator.AtLeast(0),
										},
									},
									"unit_price": schema.Float64Attribute{
										Description: "Unit price charged for the Pricing band.",
										Computed:    true,
									},
									"id": schema.StringAttribute{
										Description: "The ID for the Pricing band.",
										Computed:    true,
									},
									"credit_type_id": schema.StringAttribute{
										Description: "**OBSOLETE - this is deprecated and no longer used.**",
										Computed:    true,
									},
								},
							},
						},
						"plan_id": schema.StringAttribute{
							Description: "UUID of the Plan the Pricing is created for.",
							Computed:    true,
						},
						"plan_template_id": schema.StringAttribute{
							Description: "UUID of the Plan Template the Pricing was created for.",
							Computed:    true,
						},
						"pricing_bands": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[PricingsPricingBandsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"fixed_price": schema.Float64Attribute{
										Description: "Fixed price charged for the Pricing band.",
										Computed:    true,
									},
									"lower_limit": schema.Float64Attribute{
										Description: "Lower limit for the Pricing band.",
										Computed:    true,
										Validators: []validator.Float64{
											float64validator.AtLeast(0),
										},
									},
									"unit_price": schema.Float64Attribute{
										Description: "Unit price charged for the Pricing band.",
										Computed:    true,
									},
									"id": schema.StringAttribute{
										Description: "The ID for the Pricing band.",
										Computed:    true,
									},
									"credit_type_id": schema.StringAttribute{
										Description: "**OBSOLETE - this is deprecated and no longer used.**",
										Computed:    true,
									},
								},
							},
						},
						"segment": schema.MapAttribute{
							Description: "Name of the segment for which you are defining a Pricing.\n\nFor each segment in a segmented aggregation, make a separate call using `aggregationId` parameter to update a Pricing.",
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"segment_string": schema.StringAttribute{
							Computed: true,
						},
						"start_date": schema.StringAttribute{
							Description: "The start date *(in ISO-8601 format)* for when the Pricing starts to be active for the Plan of Plan Template.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"tiers_span_plan": schema.BoolAttribute{
							Description: "The default value is **FALSE**.\n\n* If TRUE, usage accumulates over the entire period the priced Plan is active for the account, and is not reset for pricing band rates at the start of each billing period.\n\n* If FALSE, usage does not accumulate, and is reset for pricing bands at the start of each billing period.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "* **DEBIT**. Default setting. The amount calculated using the Pricing is added to the bill as a debit.\n\n* **PRODUCT_CREDIT**. The amount calculated using the Pricing is added to the bill as a credit *(negative amount)*. To prevent negative billing, the bill will be capped at the total of other line items for the *same* Product.\n\n* **GLOBAL_CREDIT**. The amount calculated using the Pricing is added to the bill as a credit *(negative amount)*. To prevent negative billing, the bill will be capped at the total of other line items for the entire bill, which might include other Products the Account consumes.\nAvailable values: \"DEBIT\", \"PRODUCT_CREDIT\", \"GLOBAL_CREDIT\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"DEBIT",
									"PRODUCT_CREDIT",
									"GLOBAL_CREDIT",
								),
							},
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

func (d *PricingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PricingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
