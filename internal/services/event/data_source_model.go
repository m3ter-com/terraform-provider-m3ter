// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
)

type EventDataSourceModel struct {
ID types.String `tfsdk:"id" path:"id,required"`
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
DtActioned timetypes.RFC3339 `tfsdk:"dt_actioned" json:"dtActioned,computed" format:"date-time"`
EventName types.String `tfsdk:"event_name" json:"eventName,computed"`
EventTime timetypes.RFC3339 `tfsdk:"event_time" json:"eventTime,computed" format:"date-time"`
M3terEvent jsontypes.Normalized `tfsdk:"m3ter_event" json:"m3terEvent,computed"`
}

func (m *EventDataSourceModel) toReadParams(_ context.Context) (params m3ter.EventGetParams, diags diag.Diagnostics) {
  params = m3ter.EventGetParams{
    OrgID: m3ter.F(m.OrgID.ValueString()),
  }

  return
}
