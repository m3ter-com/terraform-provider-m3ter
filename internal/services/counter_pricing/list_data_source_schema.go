// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_pricing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CounterPricingsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"date": schema.StringAttribute{
				Description: "Date on which to retrieve active CounterPricings.",
				Optional:    true,
			},
			"plan_id": schema.StringAttribute{
				Description: "UUID of the Plan to retrieve CounterPricings for.",
				Optional:    true,
			},
			"plan_template_id": schema.StringAttribute{
				Description: "UUID of the Plan Template to retrieve CounterPricings for.",
				Optional:    true,
			},
			"ids": schema.ListAttribute{
				Description: "List of CounterPricing IDs to retrieve.",
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
				CustomType:  customfield.NewNestedObjectListType[CounterPricingsItemsDataSourceModel](ctx),
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
						"accounting_product_id": schema.StringAttribute{
							Computed: true,
						},
						"code": schema.StringAttribute{
							Description: "Unique short code for the Pricing.",
							Computed:    true,
						},
						"counter_id": schema.StringAttribute{
							Description: "UUID of the Counter used to create the pricing.",
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
							CustomType: customfield.NewNestedObjectListType[CounterPricingsPricingBandsDataSourceModel](ctx),
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
						"pro_rate_adjustment_credit": schema.BoolAttribute{
							Description: "The default value is **TRUE**.\n\n* When TRUE, counter adjustment credits are prorated and are billed according to the number of days in billing period.\n\n* When FALSE, counter adjustment credits are not prorated and are billed for the entire billing period.",
							Computed:    true,
						},
						"pro_rate_adjustment_debit": schema.BoolAttribute{
							Description: "The default value is **TRUE**.\n\n* When TRUE, counter adjustment debits are prorated and are billed according to the number of days in billing period.\n\n* When FALSE, counter adjustment debits are not prorated and are billed for the entire billing period.",
							Computed:    true,
						},
						"pro_rate_running_total": schema.BoolAttribute{
							Description: "The default value is **TRUE**.\n\n* When TRUE, counter running total charges are prorated and are billed according to the number of days in billing period.\n\n* When FALSE, counter running total charges are not prorated and are billed for the entire billing period.",
							Computed:    true,
						},
						"running_total_bill_in_advance": schema.BoolAttribute{
							Description: "The default value is **TRUE**.\n\n* When TRUE, running totals are billed at the start of each billing period.\n\n* When FALSE, running totals are billed at the end of each billing period.",
							Computed:    true,
						},
						"start_date": schema.StringAttribute{
							Description: "The start date *(in ISO-8601 format)* for when the Pricing starts to be active for the Plan of Plan Template.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *CounterPricingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CounterPricingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
