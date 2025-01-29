// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*CounterResource)(nil)

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
			"name": schema.StringAttribute{
				Description: "Descriptive name for the Counter.",
				Required:    true,
			},
			"unit": schema.StringAttribute{
				Description: "User defined label for units shown on Bill line items, and indicating to your customers what they are being charged for.",
				Required:    true,
			},
			"code": schema.StringAttribute{
				Description: "Code for the Counter. A unique short code to identify the Counter.",
				Optional:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "UUID of the product the Counter belongs to. *(Optional)* - if left blank, the Counter is Global. A Global Counter can be used to price Plans or Plan Templates belonging to any Product.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
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

func (r *CounterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CounterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
