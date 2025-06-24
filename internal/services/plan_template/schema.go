// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_template

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*PlanTemplateResource)(nil)

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
			"bill_frequency": schema.StringAttribute{
				Description: "Determines the frequency at which bills are generated.\n\n* **Daily**. Starting at midnight each day, covering the twenty-four hour period following.\n\n* **Weekly**. Starting at midnight on a Monday, covering the seven-day period following.\n\n* **Monthly**. Starting at midnight on the first day of each month, covering the entire calendar month following.\n\n* **Annually**. Starting at midnight on first day of each year covering the entire calendar year following.\nAvailable values: \"DAILY\", \"WEEKLY\", \"MONTHLY\", \"ANNUALLY\", \"AD_HOC\", \"MIXED\".",
				Required:    true,
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
			"currency": schema.StringAttribute{
				Description: "The ISO currency code for the currency used to charge end users - for example USD, GBP, EUR. This defines the *pricing currency* and is inherited by any Plans based on the Plan Template.\n\n**Notes:**\n* You can define a currency at Organization-level or Account-level to be used as the *billing currency*. This can be a different currency to that used for the Plan as the *pricing currency*.\n* If the billing currency for an Account is different to the pricing currency used by a Plan attached to the Account, you must ensure a *currency conversion rate* is defined for your Organization to convert the pricing currency into the billing currency at billing, otherwise Bills will fail for the Account.\n* To define any required currency conversion rates, use the `currencyConversions` request body parameter for the [Update OrganizationConfig](https://www.m3ter.com/docs/api#tag/OrganizationConfig/operation/UpdateOrganizationConfig) call.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Descriptive name for the PlanTemplate.",
				Required:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the Product associated with this PlanTemplate.",
				Required:    true,
			},
			"standing_charge": schema.Float64Attribute{
				Description: "The fixed charge *(standing charge)* applied to customer bills. This charge is prorated and must be a non-negative number.",
				Required:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"bill_frequency_interval": schema.Int64Attribute{
				Description: "How often bills are issued. \nFor example, if `billFrequency` is Monthly and `billFrequencyInterval` is 3, bills are issued every three months.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 365),
				},
			},
			"code": schema.StringAttribute{
				Description: "A unique, short code reference for the PlanTemplate. This code should not contain control characters or spaces.",
				Optional:    true,
			},
			"minimum_spend": schema.Float64Attribute{
				Description: "The Product minimum spend amount per billing cycle for end customer Accounts on a pricing Plan based on the PlanTemplate. This must be a non-negative number.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"minimum_spend_bill_in_advance": schema.BoolAttribute{
				Description: "A boolean that determines when the minimum spend is billed.\n\n* TRUE - minimum spend is billed at the start of each billing period.\n* FALSE - minimum spend is billed at the end of each billing period.\n\nOverrides the setting at Organizational level for minimum spend billing in arrears/in advance.",
				Optional:    true,
			},
			"minimum_spend_description": schema.StringAttribute{
				Description: "Minimum spend description *(displayed on the bill line item)*.",
				Optional:    true,
			},
			"ordinal": schema.Int64Attribute{
				Description: "The ranking of the PlanTemplate among your pricing plans. Lower numbers represent more basic plans, while higher numbers represent premium plans. This must be a non-negative integer.\n\n**NOTE: DEPRECATED** - do not use.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"standing_charge_bill_in_advance": schema.BoolAttribute{
				Description: "A boolean that determines when the standing charge is billed.\n\n* TRUE - standing charge is billed at the start of each billing period.\n* FALSE - standing charge is billed at the end of each billing period.\n\nOverrides the setting at Organizational level for standing charge billing in arrears/in advance.",
				Optional:    true,
			},
			"standing_charge_description": schema.StringAttribute{
				Description: "Standing charge description *(displayed on the bill line item)*.",
				Optional:    true,
			},
			"standing_charge_interval": schema.Int64Attribute{
				Description: "How often the standing charge is applied. \nFor example, if the bill is issued every three months and `standingChargeInterval` is 2, then the standing charge is applied every six months.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 365),
				},
			},
			"standing_charge_offset": schema.Int64Attribute{
				Description: "Defines an offset for when the standing charge is first applied. \nFor example, if the bill is issued every three months and the `standingChargeOfset` is 0, then the charge is applied to the first bill *(at three months)*; if 1, it would be applied to the second bill *(at six months)*, and so on.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 364),
				},
			},
			"custom_fields": schema.DynamicAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Computed:    true,
				Optional:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this PlanTemplate.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the PlanTemplate was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the PlanTemplate was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this PlanTemplate.",
				Computed:    true,
			},
		},
	}
}

func (r *PlanTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PlanTemplateResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
