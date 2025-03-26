// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_configuration

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*NotificationConfigurationResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The UUID of the entity.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "org_id": schema.StringAttribute{
        Optional: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "code": schema.StringAttribute{
        Description: "The short code for the Notification.",
        Required: true,
      },
      "description": schema.StringAttribute{
        Description: "The description for the Notification providing a brief overview of its purpose and functionality.",
        Required: true,
      },
      "event_name": schema.StringAttribute{
        Description: "The name of the *Event type* that the Notification is based on. When an Event of this type occurs and any calculation built into the Notification evaluates to `True`, the Notification is triggered.\n\n**Note:** If the Notification is set to always fire, then the Notification will always be sent when the Event of the type it is based on occurs, and without any other conditions defined by a calculation having to be met.",
        Required: true,
      },
      "name": schema.StringAttribute{
        Description: "The name of the Notification.",
        Required: true,
      },
      "active": schema.BoolAttribute{
        Description: "Boolean flag that sets the Notification as active or inactive. Only active Notifications are sent when triggered by the Event they are based on:\n\n* **TRUE** - set Notification as active. \n* **FALSE** - set Notification as inactive.",
        Optional: true,
      },
      "always_fire_event": schema.BoolAttribute{
        Description: "A Boolean flag indicating whether the Notification is always triggered, regardless of other conditions and omitting reference to any calculation. This means the Notification will be triggered simply by the Event it is based on occurring and with no further conditions having to be met.\n \n*  **TRUE** - the Notification is always triggered and omits any reference to the calculation to check for other conditions being true before triggering the Notification.\n*  **FALSE** - the Notification is only triggered when the Event it is based on occurs and any calculation is checked and all conditions defined by the calculation are met.",
        Optional: true,
      },
      "calculation": schema.StringAttribute{
        Description: "A logical expression that that is evaluated to a Boolean. If it evaluates as `True`, a Notification for the Event is created and sent to the configured destination. \nCalculations can reference numeric, string, and boolean Event fields.  \n\nSee [Creating Calculations](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications/key-concepts-and-relationships#creating-calculations) in the m3ter documentation for more information.",
        Optional: true,
      },
      "version": schema.Int64Attribute{
        Description: "The version number for the Notification:\n\n- **Create:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update:** On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
        Optional: true,
      },
      "created_by": schema.StringAttribute{
        Description: "The ID of the user who created this item.",
        Computed: true,
      },
      "dt_created": schema.StringAttribute{
        Description: "The DateTime when this item was created *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "dt_last_modified": schema.StringAttribute{
        Description: "The DateTime when this item was last modified *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The ID of the user who last modified this item.",
        Computed: true,
      },
    },
  }
}

func (r *NotificationConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *NotificationConfigurationResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
