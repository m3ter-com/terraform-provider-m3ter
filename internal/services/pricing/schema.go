// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pricing

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

var _ resource.ResourceWithConfigValidators = (*PricingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"start_date": schema.StringAttribute{
				Description: "The start date *(in ISO-8601 format)* for when the Pricing starts to be active for the Plan of Plan Template.*(Required)*",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"pricing_bands": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fixed_price": schema.Float64Attribute{
							Description: "Fixed price charged for the Pricing band.",
							Required:    true,
						},
						"lower_limit": schema.Float64Attribute{
							Description: "Lower limit for the Pricing band.",
							Required:    true,
							Validators: []validator.Float64{
								float64validator.AtLeast(0),
							},
						},
						"unit_price": schema.Float64Attribute{
							Description: "Unit price charged for the Pricing band.",
							Required:    true,
						},
						"id": schema.StringAttribute{
							Description: "The ID for the Pricing band.",
							Optional:    true,
						},
						"credit_type_id": schema.StringAttribute{
							Description: "**OBSOLETE - this is deprecated and no longer used.**",
							Optional:    true,
						},
					},
				},
			},
			"accounting_product_id": schema.StringAttribute{
				Description: "Optional Product ID this Pricing should be attributed to for accounting purposes",
				Optional:    true,
			},
			"aggregation_id": schema.StringAttribute{
				Description: "UUID of the Aggregation used to create the Pricing. Use this when creating a Pricing for a segmented aggregation.",
				Optional:    true,
			},
			"code": schema.StringAttribute{
				Description: "Unique short code for the Pricing.",
				Optional:    true,
			},
			"compound_aggregation_id": schema.StringAttribute{
				Description: "UUID of the Compound Aggregation used to create the Pricing.",
				Optional:    true,
			},
			"cumulative": schema.BoolAttribute{
				Description: "Controls whether or not charge rates under a set of pricing bands configured for a Pricing are applied according to each separate band or at the highest band reached.\n\n*(Optional)*. The default value is **FALSE**.\n\n* When TRUE, at billing charge rates are applied according to each separate band.\n\n* When FALSE, at billing charge rates are applied according to highest band reached.\n\n**NOTE:** Use the `cumulative` parameter to create the type of Pricing you require. For example, for Tiered Pricing set to **TRUE**; for Volume Pricing, set to **FALSE**.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Displayed on Bill line items.",
				Optional:    true,
			},
			"end_date": schema.StringAttribute{
				Description: "The end date *(in ISO-8601 format)* for when the Pricing ceases to be active for the Plan or Plan Template.\n\n*(Optional)* If not specified, the Pricing remains active indefinitely.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"minimum_spend": schema.Float64Attribute{
				Description: "The minimum spend amount per billing cycle for end customer Accounts on a Plan to which the Pricing is applied.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"minimum_spend_bill_in_advance": schema.BoolAttribute{
				Description: "The default value is **FALSE**.\n\n* When TRUE, minimum spend is billed at the start of each billing period.\n\n* When FALSE, minimum spend is billed at the end of each billing period.\n\n*(Optional)*. Overrides the setting at Organization level for minimum spend billing in arrears/in advance.",
				Optional:    true,
			},
			"minimum_spend_description": schema.StringAttribute{
				Description: "Minimum spend description *(displayed on the bill line item)*.",
				Optional:    true,
			},
			"plan_id": schema.StringAttribute{
				Description: "UUID of the Plan the Pricing is created for.",
				Optional:    true,
			},
			"plan_template_id": schema.StringAttribute{
				Description: "UUID of the Plan Template the Pricing is created for.",
				Optional:    true,
			},
			"tiers_span_plan": schema.BoolAttribute{
				Description: "The default value is **FALSE**.\n\n* If TRUE, usage accumulates over the entire period the priced Plan is active for the account, and is not reset for pricing band rates at the start of each billing period.\n\n* If FALSE, usage does not accumulate, and is reset for pricing bands at the start of each billing period.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "* **DEBIT**. Default setting. The amount calculated using the Pricing is added to the bill as a debit.\n\n* **PRODUCT_CREDIT**. The amount calculated using the Pricing is added to the bill as a credit *(negative amount)*. To prevent negative billing, the bill will be capped at the total of other line items for the same Product.\n\n* **GLOBAL_CREDIT**. The amount calculated using the Pricing is added to the bill as a credit *(negative amount)*. To prevent negative billing, the bill will be capped at the total of other line items for the entire bill, which might include other Products the Account consumes.\nAvailable values: \"DEBIT\", \"PRODUCT_CREDIT\", \"GLOBAL_CREDIT\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"DEBIT",
						"PRODUCT_CREDIT",
						"GLOBAL_CREDIT",
					),
				},
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"segment": schema.MapAttribute{
				Description: "Specifies the segment value which you are defining a Pricing for using this call:\n- For each segment value defined on a Segmented Aggregation you must create a separate Pricing and use the appropriate `aggregationId` parameter for the call.\n- If you specify a segment value that has not been defined for the Aggregation, you'll receive an error.\n- If you've defined segment values for the Aggregation using a single wildcard or multiple wildcards, you can create Pricing for these wildcard segment values also.\n\nFor more details on creating Pricings for segment values on a Segmented Aggregation using this call, together with some examples, see the [Using API Call to Create Segmented Pricings](https://www.m3ter.com/docs/guides/plans-and-pricing/pricing-plans/pricing-plans-using-segmented-aggregations#using-api-call-to-create-a-segmented-pricing) in our User Documentation.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"overage_pricing_bands": schema.ListNestedAttribute{
				Description: "Specify Prepayment/Balance overage pricing in pricing bands for the case of a **Tiered** pricing structure.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[PricingOveragePricingBandsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fixed_price": schema.Float64Attribute{
							Description: "Fixed price charged for the Pricing band.",
							Required:    true,
						},
						"lower_limit": schema.Float64Attribute{
							Description: "Lower limit for the Pricing band.",
							Required:    true,
							Validators: []validator.Float64{
								float64validator.AtLeast(0),
							},
						},
						"unit_price": schema.Float64Attribute{
							Description: "Unit price charged for the Pricing band.",
							Required:    true,
						},
						"id": schema.StringAttribute{
							Description: "The ID for the Pricing band.",
							Optional:    true,
						},
						"credit_type_id": schema.StringAttribute{
							Description: "**OBSOLETE - this is deprecated and no longer used.**",
							Optional:    true,
						},
					},
				},
			},
			"aggregation_type": schema.StringAttribute{
				Description: "Available values: \"SIMPLE\", \"COMPOUND\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("SIMPLE", "COMPOUND"),
				},
			},
			"created_by": schema.StringAttribute{
				Description: "The ID of the user who created this item.",
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
			"last_modified_by": schema.StringAttribute{
				Description: "The ID of the user who last modified this item.",
				Computed:    true,
			},
			"segment_string": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *PricingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PricingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
