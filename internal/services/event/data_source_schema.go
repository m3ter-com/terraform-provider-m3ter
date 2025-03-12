// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*EventDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Required: true,
      },
      "org_id": schema.StringAttribute{
        Required: true,
      },
      "dt_actioned": schema.StringAttribute{
        Description: "When an Event was actioned. It follows the ISO 8601 date and time format.\n\nYou can action an Event to indicate that it has been followed up and resolved - this is useful when dealing with integration error Events or ingest failure Events.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "event_name": schema.StringAttribute{
        Description: "The name of the Event as it is registered in the system. This name is used to categorize and trigger associated actions.",
        Computed: true,
      },
      "event_time": schema.StringAttribute{
        Description: "The time when the Event was triggered, using the ISO 8601 date and time format.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "m3ter_event": schema.StringAttribute{
        Description: "The Data Transfer Object (DTO) containing the details of the Event.",
        Computed: true,
        CustomType: jsontypes.NormalizedType{

        },
      },
    },
  }
}

func (d *EventDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *EventDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
