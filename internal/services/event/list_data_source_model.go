// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type EventsDataListDataSourceEnvelope struct {
Data customfield.NestedObjectList[EventsItemsDataSourceModel] `json:"data,computed"`
}

type EventsDataSourceModel struct {
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
AccountID types.String `tfsdk:"account_id" query:"accountId,optional"`
EventName types.String `tfsdk:"event_name" query:"eventName,optional"`
EventType types.String `tfsdk:"event_type" query:"eventType,optional"`
IncludeActioned types.Bool `tfsdk:"include_actioned" query:"includeActioned,optional"`
NotificationCode types.String `tfsdk:"notification_code" query:"notificationCode,optional"`
NotificationID types.String `tfsdk:"notification_id" query:"notificationId,optional"`
ResourceID types.String `tfsdk:"resource_id" query:"resourceId,optional"`
IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Items customfield.NestedObjectList[EventsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *EventsDataSourceModel) toListParams(_ context.Context) (params m3ter.EventListParams, diags diag.Diagnostics) {
  mIDs := []string{}
  for _, item := range *m.IDs {
    mIDs = append(mIDs, item.ValueString())
  }

  params = m3ter.EventListParams{
    IDs: m3ter.F(mIDs),
  }

  if !m.AccountID.IsNull() {
    params.AccountID = m3ter.F(m.AccountID.ValueString())
  }
  if !m.EventName.IsNull() {
    params.EventName = m3ter.F(m.EventName.ValueString())
  }
  if !m.EventType.IsNull() {
    params.EventType = m3ter.F(m.EventType.ValueString())
  }
  if !m.IncludeActioned.IsNull() {
    params.IncludeActioned = m3ter.F(m.IncludeActioned.ValueBool())
  }
  if !m.NotificationCode.IsNull() {
    params.NotificationCode = m3ter.F(m.NotificationCode.ValueString())
  }
  if !m.NotificationID.IsNull() {
    params.NotificationID = m3ter.F(m.NotificationID.ValueString())
  }
  if !m.ResourceID.IsNull() {
    params.ResourceID = m3ter.F(m.ResourceID.ValueString())
  }
  if !m.OrgID.IsNull() {
    params.OrgID = m3ter.F(m.OrgID.ValueString())
  }

  return
}

type EventsItemsDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
DtActioned timetypes.RFC3339 `tfsdk:"dt_actioned" json:"dtActioned,computed" format:"date-time"`
EventName types.String `tfsdk:"event_name" json:"eventName,computed"`
EventTime timetypes.RFC3339 `tfsdk:"event_time" json:"eventTime,computed" format:"date-time"`
M3terEvent jsontypes.Normalized `tfsdk:"m3ter_event" json:"m3terEvent,computed"`
}
