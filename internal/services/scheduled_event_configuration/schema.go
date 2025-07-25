// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduled_event_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ScheduledEventConfigurationResource)(nil)

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
			"entity": schema.StringAttribute{
				Description: "The referenced configuration or billing entity for which the desired scheduled Event will trigger.",
				Required:    true,
			},
			"field": schema.StringAttribute{
				Description: "A DateTime field for which the desired scheduled Event will trigger - this must be a DateTime field on the referenced billing or configuration entity.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the custom Scheduled Event Configuration.\n\nThis must be in the format: \n* scheduled.*name of entity*.*custom event name*\n\nFor example:\n* `scheduled.bill.endDateEvent`",
				Required:    true,
			},
			"offset": schema.Int64Attribute{
				Description: "The offset in days from the specified DateTime field on the referenced entity when the scheduled Event will trigger.",
				Required:    true,
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
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *ScheduledEventConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ScheduledEventConfigurationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
