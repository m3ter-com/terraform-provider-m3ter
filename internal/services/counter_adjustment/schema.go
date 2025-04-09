// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*CounterAdjustmentResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID the CounterAdjustment is created for.",
				Required:    true,
			},
			"counter_id": schema.StringAttribute{
				Description: "The ID of the Counter used for the CounterAdjustment on the Account.",
				Required:    true,
			},
			"value": schema.Int64Attribute{
				Description: "Integer Value of the Counter used for the CounterAdjustment.\n\n**Note:** Use the new absolute value for the Counter for the selected date - if it was 15 and has increased to 20, enter 20; if it was 15 and has decreased to 10, enter 10. *Do not enter* the plus or minus value relative to the previous Counter value on the Account.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"purchase_order_number": schema.StringAttribute{
				Description: "Purchase Order Number for the Counter Adjustment. *(Optional)*",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"date": schema.StringAttribute{
				Description: "The date the CounterAdjustment is created for the Account *(in ISO-8601 date format)*.\n\n**Note:** CounterAdjustments on Accounts are supported down to a *specific day* of granularity - you cannot create more than one CounterAdjustment for any given day using the same Counter and you'll receive an error if you try to do this.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("2022-01-04"),
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
		},
	}
}

func (r *CounterAdjustmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CounterAdjustmentResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
