// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*PlanGroupResource)(nil)

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
			"currency": schema.StringAttribute{
				Description: "Currency code for the PlanGroup (For example, USD).",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the PlanGroup.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Optional. This PlanGroup is created as bespoke for the associated Account with this Account ID.",
				Optional:    true,
			},
			"code": schema.StringAttribute{
				Description: "The short code representing the PlanGroup.",
				Optional:    true,
			},
			"minimum_spend": schema.Float64Attribute{
				Description: "The minimum spend amount for the PlanGroup.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"minimum_spend_accounting_product_id": schema.StringAttribute{
				Description: "Optional. Product ID to attribute the PlanGroup's minimum spend for accounting purposes.",
				Optional:    true,
			},
			"minimum_spend_bill_in_advance": schema.BoolAttribute{
				Description: "A boolean flag that determines when the minimum spend is billed. This flag overrides the setting at Organizational level for minimum spend billing in arrears/in advance.\n\n* **TRUE** - minimum spend is billed at the start of each billing period. \n* **FALSE** - minimum spend is billed at the end of each billing period.",
				Optional:    true,
			},
			"minimum_spend_description": schema.StringAttribute{
				Description: "Description of the minimum spend, displayed on the bill line item.",
				Optional:    true,
			},
			"standing_charge": schema.Float64Attribute{
				Description: "Standing charge amount for the PlanGroup.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"standing_charge_accounting_product_id": schema.StringAttribute{
				Description: "Optional. Product ID to attribute the PlanGroup's standing charge for accounting purposes.",
				Optional:    true,
			},
			"standing_charge_bill_in_advance": schema.BoolAttribute{
				Description: "A boolean flag that determines when the standing charge is billed. This flag overrides the setting at Organizational level for standing charge billing in arrears/in advance.\n\n* **TRUE** - standing charge is billed at the start of each billing period. \n* **FALSE** - standing charge is billed at the end of each billing period.",
				Optional:    true,
			},
			"standing_charge_description": schema.StringAttribute{
				Description: "Description of the standing charge, displayed on the bill line item.",
				Optional:    true,
			},
			"custom_fields": schema.DynamicAttribute{
				Description:   "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:      true,
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *PlanGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PlanGroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
