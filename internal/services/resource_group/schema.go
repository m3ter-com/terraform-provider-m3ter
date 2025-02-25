// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ResourceGroupResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique identifier (UUID) of the Resource Group.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"version": schema.Int64Attribute{
				Optional: true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this Resource Group.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the Resource Group was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the Resource Group was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this Resource Group.",
				Computed:    true,
			},
		},
	}
}

func (r *ResourceGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ResourceGroupResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
