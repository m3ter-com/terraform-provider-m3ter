// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*PlanResource)(nil)

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
			"code": schema.StringAttribute{
				Description: "Unique short code reference for the Plan.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Descriptive name for the Plan.",
				Required:    true,
			},
			"plan_template_id": schema.StringAttribute{
				Description: "UUID of the PlanTemplate the Plan belongs to.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "*(Optional)*. Used to specify an Account for which the Plan will be a custom/bespoke Plan:\n* Use when first creating a Plan.\n* A custom/bespoke Plan can only be attached to the specified Account.\n* Once created, a custom/bespoke Plan cannot be updated to be made a custom/bespoke Plan for a different Account.",
				Optional:    true,
			},
			"bespoke": schema.BoolAttribute{
				Description: "TRUE/FALSE flag indicating whether the plan is a custom/bespoke Plan for a particular Account:\n* When creating a Plan, use the `accountId` request parameter to specify the Account for which the Plan will be custom/bespoke.\n* A custom/bespoke Plan can only be attached to the specified Account.",
				Optional:    true,
			},
			"minimum_spend": schema.Float64Attribute{
				Description: "The product minimum spend amount per billing cycle for end customer Accounts on a priced Plan.\n\n*(Optional)*. Overrides PlanTemplate value.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"minimum_spend_accounting_product_id": schema.StringAttribute{
				Description: "Optional Product ID this plan's minimum spend should be attributed to for accounting purposes",
				Optional:    true,
			},
			"minimum_spend_bill_in_advance": schema.BoolAttribute{
				Description: "When TRUE, minimum spend is billed at the start of each billing period.\n\nWhen FALSE, minimum spend is billed at the end of each billing period.\n\n*(Optional)*. Overrides the setting at PlanTemplate level for minimum spend billing in arrears/in advance.",
				Optional:    true,
			},
			"minimum_spend_description": schema.StringAttribute{
				Description: "Minimum spend description *(displayed on the bill line item)*.",
				Optional:    true,
			},
			"ordinal": schema.Int64Attribute{
				Description: "Assigns a rank or position to the Plan in your order of pricing plans - lower numbers represent more basic pricing plans; higher numbers represent more premium pricing plans.\n\n*(Optional)*. Overrides PlanTemplate value.\n\n**NOTE: DEPRECATED** - do not use.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"standing_charge": schema.Float64Attribute{
				Description: "The standing charge applied to bills for end customers. This is prorated.\n\n*(Optional)*. Overrides PlanTemplate value.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"standing_charge_accounting_product_id": schema.StringAttribute{
				Description: "Optional Product ID this plan's standing charge should be attributed to for accounting purposes",
				Optional:    true,
			},
			"standing_charge_bill_in_advance": schema.BoolAttribute{
				Description: "When TRUE, standing charge is billed at the start of each billing period.\n\nWhen FALSE, standing charge is billed at the end of each billing period.\n\n*(Optional)*. Overrides the setting at PlanTemplate level for standing charge billing in arrears/in advance.",
				Optional:    true,
			},
			"standing_charge_description": schema.StringAttribute{
				Description: "Standing charge description *(displayed on the bill line item)*.",
				Optional:    true,
			},
			"custom_fields": schema.DynamicAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "UUID of the Product the Plan belongs to.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *PlanResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PlanResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
