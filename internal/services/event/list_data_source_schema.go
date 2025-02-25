// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*EventsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID associated with the Event to filter the results. Returns the Events that have been generated for the Account.",
				Optional:    true,
			},
			"event_name": schema.StringAttribute{
				Optional: true,
			},
			"event_type": schema.StringAttribute{
				Description: "The category of Events to filter the results by. Options:\n* Notification\n* IntegrationEvent\n* IngestValidationFailure\n* DataExportJobFailure",
				Optional:    true,
			},
			"include_actioned": schema.BoolAttribute{
				Description: "A Boolean flag indicating whether to return Events that have been actioned.\n\n* **TRUE** - include actioned Events.\n* **FALSE** - exclude actioned Events. ",
				Optional:    true,
			},
			"notification_code": schema.StringAttribute{
				Description: "Short code of the Notification to filter the results. Returns the Events that have triggered the Notification.",
				Optional:    true,
			},
			"notification_id": schema.StringAttribute{
				Description: "Notification UUID to filter the results. Returns the Events that have triggered the Notification.",
				Optional:    true,
			},
			"resource_id": schema.StringAttribute{
				Optional: true,
			},
			"ids": schema.ListAttribute{
				Description: "List of Event UUIDs to filter the results. \n\n**NOTE:** cannot be used with other filters.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[EventsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The uniqie identifier (UUID) of the Event.",
							Computed:    true,
						},
						"dt_actioned": schema.StringAttribute{
							Description: "When an Event was actioned. It follows the ISO 8601 date and time format.\n\nYou can action an Event to indicate that it has been followed up and resolved - this is useful when dealing with integration error Events or ingest failure Events.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"event_name": schema.StringAttribute{
							Description: "The name of the Event as it is registered in the system. This name is used to categorize and trigger associated actions.",
							Computed:    true,
						},
						"event_time": schema.StringAttribute{
							Description: "The time when the Event was triggered, using the ISO 8601 date and time format.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"m3ter_event": schema.StringAttribute{
							Description: "The Data Transfer Object (DTO) containing the details of the Event. ",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
					},
				},
			},
		},
	}
}

func (d *EventsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *EventsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
