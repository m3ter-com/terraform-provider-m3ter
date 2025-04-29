// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"custom_fields": schema.MapAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:    true,
				ElementType: types.DynamicType,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who created the PlanGroup.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the PlanGroup was first created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the PlanGroup was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who last modified the PlanGroup.",
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
