// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*CustomFieldResource)(nil)

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
			"account": schema.DynamicAttribute{
				Description: "Updates to Account entity CustomFields.",
				Optional:    true,
			},
			"account_plan": schema.DynamicAttribute{
				Description: "Updates to AccountPlan entity CustomFields.",
				Optional:    true,
			},
			"aggregation": schema.DynamicAttribute{
				Description: "Updates to simple Aggregation entity CustomFields.",
				Optional:    true,
			},
			"compound_aggregation": schema.DynamicAttribute{
				Description: "Updates to Compound Aggregation entity CustomFields.",
				Optional:    true,
			},
			"contract": schema.DynamicAttribute{
				Description: "Updates to Contract entity CustomFields.",
				Optional:    true,
			},
			"meter": schema.DynamicAttribute{
				Description: "Updates to Meter entity CustomFields.",
				Optional:    true,
			},
			"organization": schema.DynamicAttribute{
				Description: "Updates to Organization CustomFields.",
				Optional:    true,
			},
			"plan": schema.DynamicAttribute{
				Description: "Updates to Plan entity CustomFields.",
				Optional:    true,
			},
			"plan_template": schema.DynamicAttribute{
				Description: "Updates to planTemplate entity CustomFields.",
				Optional:    true,
			},
			"product": schema.DynamicAttribute{
				Description: "Updates to Product entity CustomFields.",
				Optional:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this custom field.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when the Organization was created *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when a custom field was last modified - created, modified, or deleted - for the Organization *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified this custom field.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *CustomFieldResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CustomFieldResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
