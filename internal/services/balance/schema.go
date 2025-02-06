// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*BalanceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity. ",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the end customer Account.",
				Required:    true,
			},
			"currency": schema.StringAttribute{
				Description: "The currency code used for the Balance amount. For example: USD, GBP or EUR. ",
				Required:    true,
			},
			"end_date": schema.StringAttribute{
				Description: "The date *(in ISO 8601 format)* after which the Balance will no longer be active for the Account.\n\n**Note:** You can use the `rolloverEndDate` request parameter to define an extended grace period for continued draw-down against the Balance if any amount remains when the specified `endDate` is reached.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"start_date": schema.StringAttribute{
				Description: "The date *(in ISO 8601 format)* when the Balance becomes active.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"balance_draw_down_description": schema.StringAttribute{
				Description: "A description for the bill line items for draw-down charges against the Balance. *(Optional).*\n",
				Optional:    true,
			},
			"code": schema.StringAttribute{
				Description: "Unique short code for the Balance.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the Balance.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The official name for the Balance. ",
				Optional:    true,
			},
			"overage_description": schema.StringAttribute{
				Description: "A description for Bill line items overage charges.",
				Optional:    true,
			},
			"overage_surcharge_percent": schema.Float64Attribute{
				Description: "Define a surcharge level, as a percentage of regular usage rating, applied to overages *(usage charges that exceed the Balance amount)*. For example, if the regular usage rate is $10 per unit of usage consumed and `overageSurchargePercent` is set at 10%, then any usage charged above the original Balance amount is charged at $11 per unit of usage.",
				Optional:    true,
			},
			"rollover_amount": schema.Float64Attribute{
				Description: "The maximum amount that can be carried over past the Balance end date for draw-down at billing if there is any unused Balance amount when the end date is reached. Works with `rolloverEndDate` to define the amount and duration of a Balance \"grace period\". *(Optional)*\n\n**Notes:**\n- If you leave `rolloverAmount` empty and only enter a `rolloverEndDate`, any amount left over after the Balance end date is reached will be drawn-down against up to the specified `rolloverEndDate`.\n- You must enter a `rolloverEndDate`. If you only enter a `rolloverAmount` without entering a `rolloverEndDate`, you'll receive an error when trying to create or update the Balance.\n- If you don't want to grant any grace period for outstanding Balance amounts, then do not use `rolloverAmount` and `rolloverEndDate`. ",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"rollover_end_date": schema.StringAttribute{
				Description: "The end date *(in ISO 8601 format)* for the grace period during which unused Balance amounts can be carried over and drawn-down against at billing.\n\n**Note:** Use `rolloverAmount` if you want to specify a maximum amount that can be carried over and made available for draw-down.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"line_item_types": schema.ListAttribute{
				Description: "Specify the line item charge types that can draw-down at billing against the  Balance amount. Options are:\n- `\"MINIMUM_SPEND\"`\n- `\"STANDING_CHARGE\"`\n- `\"USAGE\"`\n- `\"COUNTER_RUNNING_TOTAL_CHARGE\"`\n- `\"COUNTER_ADJUSTMENT_DEBIT\"`\n\n**NOTE:** If no charge types are specified, by default *all types* can draw-down against the Balance amount at billing.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"STANDING_CHARGE",
							"USAGE",
							"MINIMUM_SPEND",
							"COUNTER_RUNNING_TOTAL_CHARGE",
							"COUNTER_ADJUSTMENT_DEBIT",
						),
					),
				},
				ElementType: types.StringType,
			},
			"product_ids": schema.ListAttribute{
				Description: "Specify the Products whose consumption charges due at billing can be drawn-down against the Balance amount.\n\n**Note:** If you don't specify any Products for Balance draw-down, by default the consumption charges for any Product the Account consumes will be drawn-down against the Balance amount.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"amount": schema.Float64Attribute{
				Description: "The financial value that the Balance holds.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who created the Balance.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the Balance was first created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the Balance was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who last modified the Balance.",
				Computed:    true,
			},
		},
	}
}

func (r *BalanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BalanceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
