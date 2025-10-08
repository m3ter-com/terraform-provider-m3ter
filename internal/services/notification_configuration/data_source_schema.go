// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*NotificationConfigurationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"active": schema.BoolAttribute{
				Description: "A Boolean flag indicating whether or not the Notification is active.\n\n* **TRUE** - active Notification. \n* **FALSE** - inactive Notification.",
				Computed:    true,
			},
			"always_fire_event": schema.BoolAttribute{
				Description: "A Boolean flag indicating whether the Notification is always triggered, regardless of other conditions and omitting reference to any calculation. This means the Notification will be triggered simply by the Event it is based on occurring and with no further conditions having to be met.\n \n*  **TRUE** - the Notification is always triggered and omits any reference to the calculation to check for other conditions being true before triggering the Notification.\n*  **FALSE** - the Notification is only triggered when the Event it is based on occurs and any calculation is checked and all conditions defined by the calculation are met.",
				Computed:    true,
			},
			"calculation": schema.StringAttribute{
				Description: "A logical expression that that is evaluated to a Boolean. If it evaluates as `True`, a Notification for the Event is created and sent to the configured destination. \nCalculations can reference numeric, string, and boolean Event fields.  \n\nSee [Creating Calculations](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications/key-concepts-and-relationships#creating-calculations) in the m3ter documentation for more information.",
				Computed:    true,
			},
			"code": schema.StringAttribute{
				Description: "The short code for the Notification.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description for the Notification providing a brief overview of its purpose and functionality.",
				Computed:    true,
			},
			"event_name": schema.StringAttribute{
				Description: "The name of the Event that the Notification is based on. When this Event occurs and the calculation evaluates to `True`, the Notification is triggered.\n\n**Note:** If the Notification is set to always fire, then the Notification will always be sent when the Event it is based on occurs, and without any other conditions defined by a calculation having to be met.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Notification.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Description: "A Boolean flag indicating whether to retrieve only active or only inactive Notifications.\n\n* **TRUE** - only active Notifications are returned. \n* **FALSE** - only inactive Notifications are returned.",
						Optional:    true,
					},
					"event_name": schema.StringAttribute{
						Description: "Use this to filter the Notifications returned - only those Notifications that are based on the *Event type* specified by `eventName` are returned.",
						Optional:    true,
					},
					"ids": schema.ListAttribute{
						Description: "A list of specific Notification UUIDs to retrieve.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *NotificationConfigurationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *NotificationConfigurationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
