// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_pricing

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*CounterPricingResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The UUID of the entity.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "org_id": schema.StringAttribute{
        Optional: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "counter_id": schema.StringAttribute{
        Description: "UUID of the Counter used to create the pricing.",
        Required: true,
      },
      "start_date": schema.StringAttribute{
        Description: "The start date *(in ISO-8601 format)* for when the Pricing starts to be active for the Plan of Plan Template.*(Required)*",
        Required: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "pricing_bands": schema.ListNestedAttribute{
        Required: true,
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "fixed_price": schema.Float64Attribute{
              Description: "Fixed price charged for the Pricing band.",
              Required: true,
            },
            "lower_limit": schema.Float64Attribute{
              Description: "Lower limit for the Pricing band.",
              Required: true,
              Validators: []validator.Float64{
              float64validator.AtLeast(0),
              },
            },
            "unit_price": schema.Float64Attribute{
              Description: "Unit price charged for the Pricing band.",
              Required: true,
            },
            "id": schema.StringAttribute{
              Description: "The ID for the Pricing band.",
              Optional: true,
            },
            "credit_type_id": schema.StringAttribute{
              Description: "**OBSOLETE - this is deprecated and no longer used.**",
              Optional: true,
            },
          },
        },
      },
      "accounting_product_id": schema.StringAttribute{
        Description: "Optional Product ID this Pricing should be attributed to for accounting purposes",
        Optional: true,
      },
      "code": schema.StringAttribute{
        Description: "Unique short code for the Pricing.",
        Optional: true,
      },
      "cumulative": schema.BoolAttribute{
        Description: "Controls whether or not charge rates under a set of pricing bands configured for a Pricing are applied according to each separate band or at the highest band reached.\n\n*(Optional)*. The default value is **FALSE**.\n\n* When TRUE, at billing charge rates are applied according to each separate band.\n\n* When FALSE, at billing charge rates are applied according to highest band reached.\n\n**NOTE:** Use the `cumulative` parameter to create the type of Pricing you require. For example, for Tiered Pricing set to **TRUE**; for Volume Pricing, set to **FALSE**.",
        Optional: true,
      },
      "description": schema.StringAttribute{
        Description: "Displayed on Bill line items.",
        Optional: true,
      },
      "end_date": schema.StringAttribute{
        Description: "The end date *(in ISO-8601 format)* for when the Pricing ceases to be active for the Plan or Plan Template.\n\n*(Optional)* If not specified, the Pricing remains active indefinitely.",
        Optional: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "plan_id": schema.StringAttribute{
        Description: "UUID of the Plan the Pricing is created for.",
        Optional: true,
      },
      "plan_template_id": schema.StringAttribute{
        Description: "UUID of the Plan Template the Pricing is created for.",
        Optional: true,
      },
      "pro_rate_adjustment_credit": schema.BoolAttribute{
        Description: "The default value is **TRUE**.\n\n* When TRUE, counter adjustment credits are prorated and are billed according to the number of days in billing period.\n\n* When FALSE, counter adjustment credits are not prorated and are billed for the entire billing period.\n\n*(Optional)*.",
        Optional: true,
      },
      "pro_rate_adjustment_debit": schema.BoolAttribute{
        Description: "The default value is **TRUE**.\n\n* When TRUE, counter adjustment debits are prorated and are billed according to the number of days in billing period.\n\n* When FALSE, counter adjustment debits are not prorated and are billed for the entire billing period.\n\n*(Optional)*.",
        Optional: true,
      },
      "pro_rate_running_total": schema.BoolAttribute{
        Description: "The default value is **TRUE**.\n\n* When TRUE, counter running total charges are prorated and are billed according to the number of days in billing period.\n\n* When FALSE, counter running total charges are not prorated and are billed for the entire billing period.\n\n*(Optional)*.",
        Optional: true,
      },
      "running_total_bill_in_advance": schema.BoolAttribute{
        Description: "The default value is **TRUE**.\n\n* When TRUE, running totals are billed at the start of each billing period.\n\n* When FALSE, running totals are billed at the end of each billing period.\n\n*(Optional)*.",
        Optional: true,
      },
      "version": schema.Int64Attribute{
        Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
        Optional: true,
      },
      "created_by": schema.StringAttribute{
        Description: "The ID of the user who created this item.",
        Computed: true,
      },
      "dt_created": schema.StringAttribute{
        Description: "The DateTime when this item was created *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "dt_last_modified": schema.StringAttribute{
        Description: "The DateTime when this item was last modified *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The ID of the user who last modified this item.",
        Computed: true,
      },
    },
  }
}

func (r *CounterPricingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *CounterPricingResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
